package cmd

import (
	"log"

	"github.com/ToshihitoKon/m3utool/cmd/add"
	"github.com/ToshihitoKon/m3utool/cmd/check"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "m3utool",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err.Error())
			}
		},
	}

	rootCmd.AddCommand(add.GetCmd())
	rootCmd.AddCommand(check.GetCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
