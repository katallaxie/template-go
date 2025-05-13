package loader

import "github.com/katallaxie/go-template/pkg/spec"

// Loader is an interface that defines the methods for loading data.
type Loader interface {
	// Load loads data from a source and returns it as a byte slice.
	Load(tpl spec.Template, path string) error
}
