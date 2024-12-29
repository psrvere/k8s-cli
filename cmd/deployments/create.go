package deployments

import (
	"fmt"
	"k8scli/pkg/resources/deployments"
	"k8scli/utils"

	"github.com/spf13/cobra"
)

var CreateDeploymentCmd = &cobra.Command{
	Use:   "deployment [name]",
	Short: "Create a deployment",
	Args:  cobra.ExactArgs(1),
	RunE:  create,
}

func init() {
	CreateDeploymentCmd.Flags().String("image", "", "container image")
	CreateDeploymentCmd.Flags().String("replicas", "", "number of replicas")
	CreateDeploymentCmd.Flags().String("port", "", "port")

	CreateDeploymentCmd.MarkFlagRequired("image")
}

func create(cmd *cobra.Command, args []string) error {
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
	opt := deployments.NewCreateDeploymentOptions(name, image, replicas, port)

	opt.CreateDeployment()
	return nil
}
