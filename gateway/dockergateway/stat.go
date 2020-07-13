package dockergateway

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/gusantoniassi/navegante/core/entity"
)

type getStatResult struct {
	stat *entity.Stat
	err  error
}

// @TODO: test stats on Windows: https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L98
// OSType is in types.ContainerStats, along with the JSON response body

/**
 * Receives a Docker StatsJSON object and returns stats in a more human-friendly
 * format
 */
func hydrateStat(stat types.StatsJSON) entity.Stat {
	var cpuPercent, memPercent float64

	cpuPercent = getAverageCpuUsage(stat)

	memUsage := getMemUsage(stat.MemoryStats)
	memPercent = getMemPercent(memUsage, stat.MemoryStats.Limit)

	var netRx, netTx uint64
	for _, iface := range stat.Networks {
		netRx += iface.RxBytes
		netTx += iface.TxBytes
	}

	var blkRead, blkWrite uint64
	for _, svcBytes := range stat.BlkioStats.IoServiceBytesRecursive {
		switch svcBytes.Op {
		case "Read":
			blkRead += svcBytes.Value
		case "Write":
			blkWrite += svcBytes.Value
		}
	}

	entityStat := entity.Stat{
		ContainerID:   entity.ContainerID(stat.ID),
		CPUPercent:    cpuPercent,
		MemoryPercent: memPercent,
		MemoryUsage:   memUsage,
		MemoryTotal:   stat.MemoryStats.Limit,
		NetworkInput:  netRx,
		NetworkOutput: netTx,
		BlockRead:     blkRead,
		BlockWrite:    blkWrite,
	}

	return entityStat
}

/*
 * Calculates the average CPU usage based on the previous reading and the overall
 * system usage.
 * @see https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L166
 * @see https://stackoverflow.com/questions/35692667/in-docker-cpu-usage-calculation-what-are-totalusage-systemusage-percpuusage-a
 */
func getAverageCpuUsage(stat types.StatsJSON) float64 {
	previousContainerCpu := stat.PreCPUStats.CPUUsage.TotalUsage
	previousSystemCpu := stat.PreCPUStats.SystemUsage
	cpuPercent := 0.0

	// Container CPU usage changed from last reading
	containerCpuDelta := float64(stat.CPUStats.CPUUsage.TotalUsage) - float64(previousContainerCpu)
	// System CPU usage changed from last reading
	systemCpuDelta := float64(stat.CPUStats.SystemUsage) - float64(previousSystemCpu)
	// Number of system cores allocated to the container
	containerCPUCores := float64(stat.CPUStats.OnlineCPUs)

	// If the onlineCPU metric isn't present, use the number of PercpuUsage statistics returned
	if containerCPUCores == 0.0 {
		containerCPUCores = float64(len(stat.CPUStats.CPUUsage.PercpuUsage))
	}

	// If the system and the container CPU usage hasn't changed, we don't need to calc
	if systemCpuDelta > 0.0 && containerCpuDelta > 0.0 {
		// Calculate the average container CPU usage based on the total system usage
		averageCpuUsage := containerCpuDelta / systemCpuDelta
		// Multiply by containerCPUCores to consider the number of cores
		cpuPercent = averageCpuUsage * containerCPUCores * 100.0
	}

	return cpuPercent
}

/**
 * Extracts current container memory usage from MemoryStats type
 * @see https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L239
 */
func getMemUsage(memStat types.MemoryStats) uint64 {
	// Version 1 of the Linux cgroup API uses total_inactive_file
	if v, ok := memStat.Stats["total_inactive_file"]; ok && v < memStat.Usage {
		return memStat.Usage - v
	}

	// Version 2 of the Linux cgroup API uses inactive_file
	if v := memStat.Stats["inactive_file"]; v < memStat.Usage {
		return memStat.Usage - v
	}

	return memStat.Usage
}

/**
 * Extracts current container memory usage percentage, with a check to avoid
 * division by zero
 */
func getMemPercent(usage uint64, memLimit uint64) float64 {
	if memLimit == 0 {
		return 0
	}

	return float64(usage) / float64(memLimit) * 100
}

func getContainerStats(g *Gateway, id string, c chan getStatResult) {
	defer close(c)
	stat, err := g.ContainerStats(id)
	c <- getStatResult{stat: stat, err: err}
}

func getAllContainerStats(g *Gateway, containers []types.Container, c chan getStatResult) {
	defer close(c)
	var results []chan getStatResult

	for i, c := range containers {
		results = append(results, make(chan getStatResult))
		go getContainerStats(g, c.ID[:12], results[i])
	}

	for i := range results {
		for r1 := range results[i] {
			c <- r1
		}
	}
}

func (g *Gateway) ContainerStatsAll() ([]*entity.Stat, error) {
	ctx := context.Background()

	containers, err := g.Docker.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, nil
	}

	results := make(chan getStatResult, len(containers))
	go getAllContainerStats(g, containers, results)

	containerStats := make([]*entity.Stat, 0, len(containers))

	for r := range results {
		if r.err != nil {
			return nil, err
		}

		containerStats = append(containerStats, r.stat)
	}

	return containerStats, nil
}

func (g *Gateway) ContainerStats(cid string) (*entity.Stat, error) {
	ctx := context.Background()

	// @TODO: Implement streaming option
	response, err := g.Docker.ContainerStats(ctx, cid, false)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var result types.StatsJSON

	err = decoder.Decode(&result)

	if err != nil {
		return nil, err
	}

	stat := hydrateStat(result)

	return &stat, nil
}
