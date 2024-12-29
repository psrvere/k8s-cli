package deployments

import (
	"fmt"
	"k8scli/pkg/resources/deployments"
	"strconv"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create-deployment [name]",
	Short: "Create a deployment",
	Args:  cobra.ExactArgs(1),
	RunE:  Execute,
}

func Execute(cmd *cobra.Command, args []string) error {
	// Get positional arguments
	name := args[0]

	// Get flags
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return fmt.Errorf("error getting image flag: %v", err)
	}

	var replicas int32 = 1
	rep, err := cmd.Flags().GetString("replicas")
	if err != nil {
		return fmt.Errorf("error getting replicas flag: %v", err)
	}
	repInt, err := strconv.Atoi(rep)
	if err != nil {
		return fmt.Errorf("error parsing replicas string value: %v", err)
	}
	replicas = int32(repInt)

	var port int32 = 80
	p, err := cmd.Flags().GetString("port")
	if err != nil {
		return fmt.Errorf("error getting port flag: %v", err)
	}
	pInt, err := strconv.Atoi(p)
	if err != nil {
		return fmt.Errorf("error parsing port string value: %v", err)
	}
	port = int32(pInt)

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

func init() {
	CreateCmd.Flags().String("image", "", "container image")
	CreateCmd.Flags().String("replicas", "", "number of replicas")
	CreateCmd.Flags().String("port", "", "port")

	CreateCmd.MarkFlagRequired("image")
}
