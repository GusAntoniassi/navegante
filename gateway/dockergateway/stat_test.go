package dockergateway

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
)

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
