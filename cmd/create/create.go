package create

import (
	"fmt"
	"k8scli/cmd/deployments"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create [resource]",
	Short: "create a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must specicy a resource argument")
		}
		if args[0] != "deployment" {
			return fmt.Errorf("unknown resource type")
		}
		return nil
	},
}

func init() {
	CreateCmd.AddCommand(deployments.DeploymentCmd)
}
