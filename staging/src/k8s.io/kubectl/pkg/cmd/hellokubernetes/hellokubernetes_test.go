package hellokubernetes

import (
	"bytes"
	"os"
	"testing"
)

func TestHelloKubernetesCommand(t *testing.T) {
	// Fake manifest content
	yaml := `
apiVersion: v1
kind: Pod
metadata:
  name: mypod
`

	// Write the fake manifest to a temp file
	tmpfile, err := os.CreateTemp("", "pod-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(yaml)); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	tmpfile.Close()

	// Create the command
	cmd := NewHelloKubernetesCommand()

	// Capture the output
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	// Provide the -f flag with our temp file
	cmd.SetArgs([]string{"-f", tmpfile.Name()})

	// Run command
	if err := cmd.Execute(); err != nil {
		t.Fatalf("command failed: %v", err)
	}

	// Check output
	got := buf.String()
	want := "Hello Pod mypod\n"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
