package dockergateway

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/gusantoniassi/navegante/core/entity"
)

// @TODO: This is a very basic implementation, probably needs some work on the
// calculations. Need to understand more the docker stat source code and also
// take a look at how cadvisor calculates these values and compare the results
// @TODO: test stats on Windows: https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L98
func hydrateStat(stat types.StatsJSON) entity.Stat {
	var cpuPercent, memPercent float64

	if stat.CPUStats.SystemUsage > 0 {
		// @TODO: Check delta calc made in the docker stat command
		// https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L166
		cpuPercent = float64(stat.CPUStats.CPUUsage.TotalUsage) / float64(stat.CPUStats.SystemUsage) * 100
	}

	// @TODO: Check cgroup calc made in the docker stat command
	// https://github.com/docker/cli/blob/96e1d1d6421b725bdd5024f9a97af9bf97ad9619/cli/command/container/stats_helpers.go#L239
	memUsage := float64(stat.MemoryStats.Usage)

	if stat.MemoryStats.MaxUsage > 0 {
		memPercent = memUsage / float64(stat.MemoryStats.Limit) * 100
	}

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
		MemoryUsage:   entity.Bytes(memUsage),
		MemoryTotal:   entity.Bytes(stat.MemoryStats.Limit),
		NetworkInput:  entity.Bytes(netRx),
		NetworkOutput: entity.Bytes(netTx),
		BlockRead:     entity.Bytes(blkRead),
		BlockWrite:    entity.Bytes(blkWrite),
	}

	return entityStat
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

	containerStats := make([]*entity.Stat, 0, len(containers))

	for _, c := range containers {
		stat, err := g.ContainerStats(c.ID[:12])

		if err != nil {
			return nil, err
		}

		containerStats = append(containerStats, stat)
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
