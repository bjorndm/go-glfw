//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/bjorndm/go-glfwff-src/glfw/deps/glad"
	_ "github.com/bjorndm/go-glfwff-src/glfw/deps/mingw"
	_ "github.com/bjorndm/go-glfwff-src/glfw/deps/vs2008"
)
