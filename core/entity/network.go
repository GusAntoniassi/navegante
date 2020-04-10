package entity

type Network struct {
	ID        string
	Name      string
	Gateway   string
	IPAddress string

	Links   []string
	Aliases []string
}

func (n Network) String() string {
	return n.Name
}
