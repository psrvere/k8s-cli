package crd

import (
	"k8scli/pkg/resources/crd"

	"github.com/spf13/cobra"
)

var CreateCustomResourceDefinitionCommand = &cobra.Command{
	Use:   "crd",
	Short: "create crd for custom pod",
	RunE:  create,
}

func create(cmd *cobra.Command, args []string) error {
	def := crd.NewCustomResourceDefinition()
	return def.CreateCustomSourceDefinition()
}
