package deployments

import (
	"fmt"
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
	fmt.Println("In delete deployment")
	name := args[0]
	opt := deployments.NewDeleteDeploymentOptions(name)
	err := opt.DeleteDeployment()
	fmt.Println("err: ", err)
	return err
}
