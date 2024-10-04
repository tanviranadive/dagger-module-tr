// A generated module for DaggerModuleTr functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-module-tr/internal/dagger"
	"fmt"
	"log"
)

type DaggerModuleTr struct{}

// Returns a container that echoes whatever string argument is provided
func (m *DaggerModuleTr) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
//func (m *DaggerModuleTr) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
//	return dag.Container().
//		From("alpine:latest").
//		WithMountedDirectory("/mnt", directoryArg).
//		WithWorkdir("/mnt").
//		WithExec([]string{"grep", "-R", pattern, "."}).
//		Stdout(ctx)
//}

func (m *DaggerModuleTr) Setup(ctx context.Context, dir *dagger.Directory, filter string, test string) {
	// Initialize Dagger client
	client := dagger.Connect()
	ctx = context.Background()

	fmt.Println(test)
	_, err := ContainerEcho(client, filter, dir)
	if err != nil {
		log.Fatalf("Error in container echo: %v", err)
	}
}