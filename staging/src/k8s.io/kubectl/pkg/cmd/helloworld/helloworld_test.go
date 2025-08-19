package helloworld

import (
    "bytes"
    "strings"
    "testing"
)

func TestHelloWorldCommand(t *testing.T) {
    cmd := NewHelloWorldCommand()

    // Capture the output
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)

    // Execute the command (no args)
    if err := cmd.Execute(); err != nil {
        t.Fatalf("command execution failed: %v", err)
    }

    // Check that the output contains "Hello World"
    output := buf.String()
    if !strings.Contains(output, "Hello World") {
        t.Errorf("expected output to contain 'Hello World', got %q", output)
    }

    // Also check the command metadata
    if cmd.Use != "hello-world" {
        t.Errorf("expected command use to be hello-world, got %s", cmd.Use)
    }
}
