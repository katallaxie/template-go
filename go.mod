module example

go 1.24

tool (
	github.com/golangci/golangci-lint/cmd/golangci-lint
	github.com/goreleaser/goreleaser
	github.com/katallaxie/pkg/cmd/runproc
	github.com/vektra/mockery/v2
	gotest.tools/gotestsum
	mvdan.cc/gofumpt
)