package entity

type Stat struct {
	ContainerID   ContainerID `json:"containerId"`
	CPUPercent    float64     `json:"cpuPercent"`
	MemoryPercent float64     `json:"memoryPercent"`
	MemoryUsage   uint64      `json:"memoryUsage"`
	MemoryTotal   uint64      `json:"memoryTotal"`
	NetworkInput  uint64      `json:"networkInput"`
	NetworkOutput uint64      `json:"networkOutput"`
	BlockRead     uint64      `json:"blockRead"`
	BlockWrite    uint64      `json:"blockWrite"`
}
