package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ToshihitoKon/m3utool/cmd/utils"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.Flags().StringVarP(&Config, "config", "c", "", "config file path")
	addCmd.Flags().StringVarP(&BasePath, "basepath", "b", "", "m3u base file path")
	rootCmd.AddCommand(addCmd)
}

var (
	Config   = ""
	BasePath = ""
	addCmd   = &cobra.Command{
		Use:   "add",
		Short: "add new song",
		Long:  "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("require add file path")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("config path:", Config)

			for i, v := range utils.ReadFile(args[0]) {
				path := filepath.Join(BasePath, v)
				fmt.Println(i, path)
				info, err := os.Stat(path)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				if info.IsDir() {
					log.Println(path, "is directory")
					continue
				}
			}
		},
	}
)
