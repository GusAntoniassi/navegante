package entity

import "fmt"

type Byte uint64

// @TODO: Move this to somewhere in the presentation layer
func (t Byte) String() string {
	if t < 1024 {
		return fmt.Sprintf("%dB", t)
	}

	count := 0
	floatBytes := float64(t)

	for floatBytes > 1024 {
		floatBytes = floatBytes / 1024
		count++
	}

	return fmt.Sprintf("%.2f %ciB", floatBytes, "KMGTPE"[count-1])
}

type Stat struct {
	ContainerID   ContainerID `json:"containerId"`
	CPUPercent    float64     `json:"cpuPercent"`
	MemoryPercent float64     `json:"memoryPercent"`
	MemoryUsage   Byte        `json:"memoryUsage"`
	MemoryTotal   Byte        `json:"memoryTotal"`
	NetworkInput  Byte        `json:"networkInput"`
	NetworkOutput Byte        `json:"networkOutput"`
	BlockRead     Byte        `json:"blockRead"`
	BlockWrite    Byte        `json:"blockWrite"`
}
