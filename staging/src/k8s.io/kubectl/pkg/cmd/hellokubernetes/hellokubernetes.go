package hellokubernetes

import (
    "fmt"
    "io/ioutil"

    "github.com/spf13/cobra"
    "sigs.k8s.io/yaml"
)

// minimal struct to grab what we need
type KubeMetadata struct {
    Kind     string `yaml:"kind"`
    Metadata struct {
        Name string `yaml:"name"`
    } `yaml:"metadata"`
}

func NewHelloKubernetesCommand() *cobra.Command {
    var filename string

    cmd := &cobra.Command{
        Use:   "hello-kubernetes",
        Short: "Print Hello <kind> <name> from a manifest",
        RunE: func(cmd *cobra.Command, args []string) error {
            if filename == "" {
                return fmt.Errorf("please provide a -f <file>")
            }

            // Read the YAML file
            data, err := ioutil.ReadFile(filename)
            if err != nil {
                return fmt.Errorf("failed to read file: %v", err)
            }

            // Decode YAML into our struct
            var meta KubeMetadata
            if err := yaml.Unmarshal(data, &meta); err != nil {
                return fmt.Errorf("failed to parse YAML: %v", err)
            }

            if meta.Kind == "" || meta.Metadata.Name == "" {
                return fmt.Errorf("file missing kind or metadata.name")
            }

            // Print hello message
            fmt.Fprintf(cmd.OutOrStdout(), "Hello %s %s\n", meta.Kind, meta.Metadata.Name)
            return nil
        },
    }

    cmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename of resource to read")
    return cmd
}
