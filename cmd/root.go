package cmd

import (
	"k8scli/cmd/action"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8scli",
	Short: "A cli to create and manage deployments on kubernetes",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
	}
}

func init() {
	// register commands
	rootCmd.AddCommand(action.CreateCmd)
	rootCmd.AddCommand(action.DeleteCmd)
}
