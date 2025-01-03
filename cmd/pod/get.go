package pod

import (
	"fmt"
	"k8scli/pkg/resources/pod"

	"github.com/spf13/cobra"
)

var GetPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "get pod(s)",
	RunE:  get,
}

func get(cmd *cobra.Command, args []string) error {
	opt := pod.NewListPodOption()
	list, err := opt.ListPods()
	if err != nil {
		return err
	}
	for _, pod := range list.Items {
		fmt.Println(pod.Name)
	}
	return nil
}
