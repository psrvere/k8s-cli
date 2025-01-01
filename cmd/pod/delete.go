package pod

import (
	"k8scli/pkg/resources/pod"

	"github.com/spf13/cobra"
)

var DeletePodCmd = &cobra.Command{
	Use:   "pod [podName]",
	Short: "Delete a pod",
	Args:  cobra.ExactArgs(1),
	RunE:  Delete,
}

func Delete(cmd *cobra.Command, args []string) error {
	name := args[0]
	opt := pod.NewDeletePodOption(name)
	return opt.DeletePod()
}
