package entity

type Volume struct {
	Name          string
	Type          string
	HostPath      string
	ContainerPath string
	Mode          string
	ReadWrite     bool
}
