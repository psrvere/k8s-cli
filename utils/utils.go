package utils

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func GetInt32ValueFromFlag(cmd *cobra.Command, flag string, defaultValue int32) int32 {
	var flagValue int32 = defaultValue
	flagStr, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalf("error getting replicas flag: %v", err)
	}
	var flagInt int
	if flagStr != "" {
		flagInt, err = strconv.Atoi(flagStr)
		if err != nil {
			log.Fatalf("error parsing replicas string value: %v", err)
		}
		flagValue = int32(flagInt)
	}
	return flagValue
}
