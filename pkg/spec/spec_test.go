package spec_test

import (
	"io"
	"testing"

	"github.com/katallaxie/go-template/pkg/spec"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	// Create a new example specification
	s := spec.Example()
	require.NotNil(t, s)
	require.Equal(t, "example", s.Name)
	require.Equal(t, spec.DefaultVersion, s.Version)
	require.Equal(t, "This is an example project.", s.Description)
}

func Test_UnmarshalYAML(t *testing.T) {
	// Create a new specification
	s := &spec.Spec{}

	// Example YAML data
	data := []byte(`
name: example
version: 1
description: This is an example project.
author: John Doe
license: MIT
`)
	// Unmarshal the YAML data into the specification
	err := s.UnmarshalYAML(data)
	require.NoError(t, err)
	require.Equal(t, "example", s.Name)
	require.Equal(t, 1, s.Version)
	require.Equal(t, "This is an example project.", s.Description)
	require.Equal(t, "John Doe", s.Author)
	require.Equal(t, "MIT", s.License)
}

func Test_Read(t *testing.T) {
	s := spec.Example()

	b, err := io.ReadAll(s)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.NotEmpty(t, b)
}
