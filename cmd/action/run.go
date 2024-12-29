package action

import (
	"fmt"
	"k8scli/pkg/resources/pod"
	"k8scli/utils"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run [podName]",
	Short: "Run a pod or image",
	Args:  cobra.ExactArgs(1),
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	podName := args[0]
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return fmt.Errorf("error getting image flag: %v", err)
	}
	port := utils.GetInt32ValueFromFlag(cmd, "port", 80)

	opt := pod.NewCreatePodOptions(podName, image, port)
	return opt.CreatePod()
}

func init() {
	RunCmd.Flags().String("image", "", "container image")
	RunCmd.Flags().String("port", "", "container port")

	RunCmd.MarkFlagRequired("image")
}
