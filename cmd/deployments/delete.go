package deployments

import (
	"k8scli/pkg/resources/deployments"

	"github.com/spf13/cobra"
)

var DeleteDeploymentCmd = &cobra.Command{
	Use:   "deployment [name]",
	Short: "Delete a deployment",
	Args:  cobra.ExactArgs(1),
	RunE:  delete,
}

func delete(cmd *cobra.Command, args []string) error {
	name := args[0]
	opt := deployments.NewDeleteDeploymentOptions(name)
	err := opt.DeleteDeployment()
	return err
}
