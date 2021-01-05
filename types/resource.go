package types

type Resource struct {
	Name     string   `yaml:"name"`
	Labels   []string `yaml:"labels"`
	Metadata Metadata `yaml:"metadata"`
}
