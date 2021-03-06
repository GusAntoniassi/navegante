package dockergateway

import (
	"context"
	"fmt"
	"github.com/gojuno/minimock/v3"
	"github.com/gusantoniassi/navegante/core/entity"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
)

func getMockContainerStatsLinux() types.ContainerStats {
	responseBody := `
{
  "id": "0123abcd456e",
  "read": "2015-01-08T22:57:31.547920715Z",
  "pids_stats": {
    "current": 3
  },
  "networks": {
    "eth0": {
      "rx_bytes": 5338,
      "rx_dropped": 0,
      "rx_errors": 0,
      "rx_packets": 36,
      "tx_bytes": 648,
      "tx_dropped": 0,
      "tx_errors": 0,
      "tx_packets": 8
    },
    "eth5": {
      "rx_bytes": 4641,
      "rx_dropped": 0,
      "rx_errors": 0,
      "rx_packets": 26,
      "tx_bytes": 690,
      "tx_dropped": 0,
      "tx_errors": 0,
      "tx_packets": 9
    }
  },
  "memory_stats": {
    "stats": {
      "total_pgmajfault": 0,
      "cache": 0,
      "mapped_file": 0,
      "total_inactive_file": 0,
      "pgpgout": 414,
      "rss": 6537216,
      "total_mapped_file": 0,
      "writeback": 0,
      "unevictable": 0,
      "pgpgin": 477,
      "total_unevictable": 0,
      "pgmajfault": 0,
      "total_rss": 6537216,
      "total_rss_huge": 6291456,
      "total_writeback": 0,
      "total_inactive_anon": 0,
      "rss_huge": 6291456,
      "hierarchical_memory_limit": 67108864,
      "total_pgfault": 964,
      "total_active_file": 0,
      "active_anon": 6537216,
      "total_active_anon": 6537216,
      "total_pgpgout": 414,
      "total_cache": 0,
      "inactive_anon": 0,
      "active_file": 0,
      "pgfault": 964,
      "inactive_file": 0,
      "total_pgpgin": 477
    },
    "max_usage": 6651904,
    "usage": 6537216,
    "failcnt": 0,
    "limit": 67108864
  },
  "blkio_stats": {
    "io_service_bytes_recursive": [
      {
        "major": 8,
        "minor": 16,
        "op": "Read",
        "value": 1048576
      },
      {
        "major": 8,
        "minor": 16,
        "op": "Write",
        "value": 1024
      }
    ],
    "io_serviced_recursive": [],
    "io_queue_recursive": [],
    "io_service_time_recursive": [],
    "io_wait_time_recursive": [],
    "io_merged_recursive": [],
    "io_time_recursive": [],
    "sectors_recursive": []
  },
  "cpu_stats": {
    "cpu_usage": {
      "percpu_usage": [
        8646879,
        24472255,
        36438778,
        30657443
      ],
      "usage_in_usermode": 50000000,
      "total_usage": 100215355,
      "usage_in_kernelmode": 30000000
    },
    "system_cpu_usage": 739306590000000,
    "online_cpus": 4,
    "throttling_data": {
      "periods": 0,
      "throttled_periods": 0,
      "throttled_time": 0
    }
  },
  "precpu_stats": {
    "cpu_usage": {
      "percpu_usage": [
        8646879,
        24350896,
        36438778,
        30657443
      ],
      "usage_in_usermode": 50000000,
      "total_usage": 100093996,
      "usage_in_kernelmode": 30000000
    },
    "system_cpu_usage": 9492140000000,
    "online_cpus": 4,
    "throttling_data": {
      "periods": 0,
      "throttled_periods": 0,
      "throttled_time": 0
    }
  }
}`

	r := ioutil.NopCloser(strings.NewReader(responseBody))

	return types.ContainerStats{
		Body:   r,
		OSType: "linux",
	}
}

func TestGateway_ContainerStats(t *testing.T) {
	mockStats := getMockContainerStatsLinux()

	mc := minimock.NewController(t)
	dockerMock := NewCommonAPIClientMock(mc).ContainerStatsMock.Set(func(context.Context, string, bool) (types.ContainerStats, error) {
		return mockStats, nil
	})

	gw := NewGateway(dockerMock)
	stats, err := gw.ContainerStats("abc123")

	assert.Nilf(t, err, "ContainerStats returns no error")
	assert.NotEmptyf(t, stats, "Should return stats")
	assert.Equal(t, stats.ContainerID, entity.ContainerID("0123abcd456e"), "Container ID should be equal to 0123abcd456e")
}

func TestGateway_ContainerStatsWithError(t *testing.T) {
	mc := minimock.NewController(t)
	dockerMock := NewCommonAPIClientMock(mc).ContainerStatsMock.Set(func(context.Context, string, bool) (types.ContainerStats, error) {
		return types.ContainerStats{}, fmt.Errorf("docker error")
	})

	gw := NewGateway(dockerMock)
	_, err := gw.ContainerStats("abc123")

	assert.NotNilf(t, err, "ContainerStats should return an error")
}

func TestGateway_ContainerStatsAllLinux(t *testing.T) {
	mc := minimock.NewController(t)
	dockerMock := NewCommonAPIClientMock(mc).ContainerListMock.Set(func(context.Context, types.ContainerListOptions) ([]types.Container, error) {
		return getMockContainers(), nil
	}).ContainerStatsMock.Set(func(context.Context, string, bool) (types.ContainerStats, error) {
		return getMockContainerStatsLinux(), nil
	})

	gw := NewGateway(dockerMock)
	stats, err := gw.ContainerStatsAll()

	assert.Nilf(t, err, "ContainerStatsAll returns no error")
	assert.NotEmptyf(t, stats, "Should return stats")
	assert.Equal(t, stats[0].ContainerID, entity.ContainerID("0123abcd456e"), "Container ID should be equal to 0123abcd456e")
}

func TestGateway_ContainerStatsAllWithError(t *testing.T) {
	mc := minimock.NewController(t)
	dockerMock := NewCommonAPIClientMock(mc).ContainerListMock.Set(func(context.Context, types.ContainerListOptions) ([]types.Container, error) {
		return nil, fmt.Errorf("docker error")
	})

	gw := NewGateway(dockerMock)
	_, err := gw.ContainerStatsAll()

	assert.NotNilf(t, err, "ContainerStatsAll should return an error")
}

func TestGateway_getNetworkBytes(t *testing.T) {
	mockNetworkStats := map[string]types.NetworkStats{
		"eth0": {
			RxBytes: 1024,
			TxBytes: 2048,
		},
		"eth5": {
			RxBytes: 1024,
			TxBytes: 2048,
		},
		"eth999": {
			RxBytes: 1024,
			TxBytes: 2048,
		},
	}

	expectedRx, expectedTx := getNetworkBytes(mockNetworkStats)

	assert.Equal(t, expectedRx, uint64(1024+1024+1024), "getNetworkBytes should sum the RxBytes from all the network interfaces")
	assert.Equal(t, expectedTx, uint64(2048+2048+2048), "getNetworkBytes should sum the TxBytes from all the network interfaces")
}

func TestGateway_getDiskBytes(t *testing.T) {
	mockDiskStats := []types.BlkioStatEntry{
		// From my tests containerd will return only one of each Op, but I'll test the sum to be safe
		{Major: 8, Minor: 16, Op: "Read", Value: 1024},
		{Major: 8, Minor: 16, Op: "Read", Value: 1024},
		{Major: 8, Minor: 16, Op: "Write", Value: 2048},
		{Major: 8, Minor: 16, Op: "Write", Value: 2048},
	}

	expectedRead, expectedWrite := getDiskBytes(mockDiskStats)

	assert.Equal(t, expectedRead, uint64(1024+1024), "getDiskBytes should sum all the Read operations")
	assert.Equal(t, expectedWrite, uint64(2048+2048), "getDiskBytes should sum all the Write operations")
}

func getMockMemoryStats() types.MemoryStats {
	return types.MemoryStats{
		Usage:    100,
		MaxUsage: 1024,
		Stats:    map[string]uint64{},
	}
}

func TestGateway_getMemPercent(t *testing.T) {
	actual := getMemPercent(1024, 1024)
	expected := 100.0

	assert.Equal(t, expected, actual, "Memory percent calculation should return 100.0")
}

func TestGateway_getMemPercentWithZeroedMemLimit(t *testing.T) {
	actual := getMemPercent(1024, 0)
	expected := 0.0

	assert.Equal(t, expected, actual, "Memory percent calculation with a zero limit should return 0.0")
}

func TestGateway_getMemUsageDefault(t *testing.T) {
	memoryStats := getMockMemoryStats()

	actual := getMemUsage(memoryStats)
	expected := memoryStats.Usage

	assert.Equal(t, expected, actual, "Memory usage should return the value from 'Usage' if no Stats were specified")
}

func TestGateway_getMemUsageWithTotalInactiveFileLesserThanUsage(t *testing.T) {
	memoryStats := getMockMemoryStats()
	memoryStats.Stats["total_inactive_file"] = 10

	actual := getMemUsage(memoryStats)
	expected := memoryStats.Usage - memoryStats.Stats["total_inactive_file"]

	assert.Equal(t, expected, actual, "Memory usage should return (Usage - total_inactive_file) if total_inactive_file is lesser than Usage")
}

func TestGateway_getMemUsageWithTotalInactiveFileGreaterThanUsage(t *testing.T) {
	memoryStats := getMockMemoryStats()
	memoryStats.Stats["total_inactive_file"] = memoryStats.Usage

	actual := getMemUsage(memoryStats)
	expected := memoryStats.Usage

	assert.Equal(t, expected, actual, "Memory usage should return the value from 'Usage' if total_inactive_file is greater than Usage")
}

func TestGateway_getMemUsageWithInactiveFileLesserThanUsage(t *testing.T) {
	memoryStats := getMockMemoryStats()
	memoryStats.Stats["inactive_file"] = 10

	actual := getMemUsage(memoryStats)
	expected := memoryStats.Usage - memoryStats.Stats["inactive_file"]

	assert.Equal(t, expected, actual, "Memory usage should return (Usage - inactive_file) if inactive_file is lesser than Usage")
}

func TestGateway_getMemUsageWithInactiveFileGreaterThanUsage(t *testing.T) {
	memoryStats := getMockMemoryStats()
	memoryStats.Stats["inactive_file"] = memoryStats.Usage

	actual := getMemUsage(memoryStats)
	expected := memoryStats.Usage

	assert.Equal(t, expected, actual, "Memory usage should return the value from 'Usage' if inactive_file is greater than Usage")
}
