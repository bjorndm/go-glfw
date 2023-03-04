//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/bjorndm/go-glfwff/deps/glad"
	_ "github.com/bjorndm/go-glfwff/deps/mingw"
	_ "github.com/bjorndm/go-glfwff/deps/vs2008"
)
