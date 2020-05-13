package entity

import "fmt"

type Bytes uint64

// @TODO: Move this to somewhere in the presentation layer
func (t Bytes) String() string {
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
	ContainerID   ContainerID
	CPUPercent    float64
	MemoryPercent float64
	MemoryUsage   Bytes
	MemoryTotal   Bytes
	NetworkInput  Bytes
	NetworkOutput Bytes
	BlockRead     Bytes
	BlockWrite    Bytes
}
