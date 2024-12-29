package deployments

import (
	"fmt"
	"k8scli/pkg/resources/deployments"
	"k8scli/utils"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create-deployment [name]",
	Short: "Create a deployment",
	Args:  cobra.ExactArgs(1),
	RunE:  Execute,
}

func init() {
	CreateCmd.Flags().String("image", "", "container image")
	CreateCmd.Flags().String("replicas", "", "number of replicas")
	CreateCmd.Flags().String("port", "", "port")

	CreateCmd.MarkFlagRequired("image")
}

func Execute(cmd *cobra.Command, args []string) error {
	// Get positional arguments
	name := args[0]

	// Get flags
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return fmt.Errorf("error getting image flag: %v", err)
	}

	replicas := utils.GetInt32ValueFromFlag(cmd, "replicas", 1)
	port := utils.GetInt32ValueFromFlag(cmd, "port", 80)

	// create deployment here
	opt := deployments.CreateDeploymentOptions{
		Name:          name,
		Image:         image,
		Replicas:      replicas,
		ContainerPort: port,
	}

	opt.CreateDeployment()
	return nil
}
