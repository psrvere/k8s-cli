package action

import (
	"fmt"
	"k8scli/cmd/crd"
	"k8scli/cmd/customresource"
	"k8scli/cmd/deployments"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create [resource]",
	Short: "Create a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must specicy a resource type")
		}
		if args[0] != "deployment" {
			return fmt.Errorf("unknown resource type")
		}
		return nil
	},
}

func init() {
	CreateCmd.AddCommand(deployments.CreateDeploymentCmd)
	CreateCmd.AddCommand(customresource.CreateCustomResourceCmd)
	CreateCmd.AddCommand(crd.CreateCustomResourceDefinitionCommand)
}
