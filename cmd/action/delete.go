package action

import (
	"fmt"
	"k8scli/cmd/deployments"
	"k8scli/cmd/pod"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [resource]",
	Short: "Delete a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must specify a resource type")
		}
		if args[0] != "deployment" {
			return fmt.Errorf("unknown resource type")
		}
		return nil
	},
}

func init() {
	DeleteCmd.AddCommand(deployments.DeleteDeploymentCmd)
	DeleteCmd.AddCommand(pod.DeletePodCmd)
}
