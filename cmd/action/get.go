package action

import (
	"fmt"
	"k8scli/cmd/pod"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get [resource]",
	Short: "get resource(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must specicy a resource type")
		}
		return nil
	},
}

func init() {
	GetCmd.AddCommand(pod.GetPodCmd)
}
