package loader

import (
	"context"

	"github.com/katallaxie/go-template/pkg/spec"
)

// Loader is an interface that defines the methods for loading data.
type Loader interface {
	// Load loads data from a source and returns it as a byte slice.
	Load(ctx context.Context, tpl spec.Template, path string) error
}
