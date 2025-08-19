package helloworld

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewHelloWorldCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "hello-world",
        Short: "Prints Hello World",
        Run: func(cmd *cobra.Command, args []string) {
             fmt.Fprintln(cmd.OutOrStdout(), "Hello World")
        },
    }
    return cmd
}
