package main

import (
	"context"

	"dagger/grayidea/internal/dagger"
)

type Grayidea struct {
	HelloClient *dagger.Hello
}

func (m *Grayidea) HelloWorld(ctx context.Context) (string, error) {
	return m.HelloClient.Hello(ctx, "", "", false, false, nil)
}

func (m *Grayidea) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

func (m *Grayidea) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
