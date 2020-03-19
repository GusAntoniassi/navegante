package entity

type Image struct {
	ID   string
	Name string
	Tag  string

	Architecture string
	Author       string
	Cmd          []string
	Digest       string
	Entrypoint   []string
	Env          []map[string]string
	ExposedPorts []string
	DomainName   string
	Hostname     string
	Labels       []map[string]string
	OS           string
	Size         uint64
	User         string
	Volumes      []string
	WorkingDir   string
}
