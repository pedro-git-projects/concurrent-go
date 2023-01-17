package pkg

import (
	"io"
	"os"
)

type VoidFn func()

// Helper for testing where printed values must be temporarely
// piped to an alternative output
func PipeStdout(f VoidFn) string {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	return output
}
