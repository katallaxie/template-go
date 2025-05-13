package spec

import (
	"io"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var validate = validator.New()

var _ io.Reader = (*Spec)(nil)

// Spec is the specification for the project.
type Spec struct {
	// Name is the name of the project.
	Name string `json:"name" yaml:"name"`
	// Version is the version of the project.
	Version int `json:"version" yaml:"version"`
	// Description is the description of the project.
	Description string `json:"description" yaml:"description"`
	// Author is the author of the project.
	Author string `json:"author" yaml:"author"`
	// License is the license of the project.
	License string `json:"license" yaml:"license"`
}

// Example returns an example specification.
func Example() *Spec {
	return &Spec{
		Name:        "example",
		Version:     DefaultVersion,
		Description: "This is an example project.",
	}
}

const (
	// DefaultVersion is the default version of the project.
	DefaultVersion = 1
	// DefaultFileName is the default file name of the specification.
	DefaultFileName = ".g.yml"
)

// UnmarshalYAML unmarshals the YAML data into the Spec struct.
func (s *Spec) UnmarshalYAML(data []byte) error {
	ss := &struct {
		Author      string `yaml:"author"`
		Description string `yaml:"description"`
		License     string `yaml:"license"`
		Name        string `yaml:"name"`
		Version     int    `yaml:"version"`
	}{}

	if err := yaml.Unmarshal(data, &ss); err != nil {
		return err
	}

	s.Name = ss.Name
	s.Version = ss.Version
	s.Description = ss.Description
	s.Author = ss.Author
	s.License = ss.License

	return nil
}

// Validate checks if the Spec is valid.
func (s *Spec) Validate() error {
	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return nil
}

// Write writes the Spec to the given writer in YAML format.
func (s *Spec) Read(p []byte) (n int, err error) {
	data, err := yaml.Marshal(s)
	if err != nil {
		return 0, err
	}

	return copy(p, data), io.EOF
}
