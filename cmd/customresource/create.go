package customresource

import (
	"fmt"
	"k8scli/pkg/client"
	"k8scli/pkg/resources/customresource"

	"github.com/spf13/cobra"
)

var CreateCustomResourceCmd = &cobra.Command{
	Use:   "cr [name]",
	Short: "create custom resource",
	Args:  cobra.ExactArgs(1),
	RunE:  create,
}

func init() {
	CreateCustomResourceCmd.Flags().String("image", "", "image name")
	CreateCustomResourceCmd.Flags().String("container-name", "my-container", "contaner name")
	CreateCustomResourceCmd.Flags().Int32("port", 80, "container port")

	CreateCustomResourceCmd.MarkFlagRequired("image")
}

func create(cmd *cobra.Command, args []string) error {
	appName := args[0]

	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return fmt.Errorf("failed to get container image: %v", err)
	}
	containerName, err := cmd.Flags().GetString("container-name")
	if err != nil {
		return fmt.Errorf("failed to get container name: %v", err)
	}
	port, err := cmd.Flags().GetInt32("port")
	if err != nil {
		return fmt.Errorf("failed to get port: %v", err)
	}

	opts := customresource.NewCreateCustomResouceOptions(appName, "CustomPod", containerName, image, port)
	clinet := client.GetKubeClient()
	return opts.CreateCustomResource(clinet.DynamicClient)
}
