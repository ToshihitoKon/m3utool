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
	checkCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
	checkCmd.Flags().StringVarP(&basePath, "basepath", "b", "", "m3u base file path")
	rootCmd.AddCommand(checkCmd)
}

var (
	config   = ""
	basePath = ""
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "check file exist",
		Long:  "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("require m3u file path")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("config path:", config)

			for i, v := range utils.ReadFile(args[0]) {
				path := filepath.Join(basePath, v)
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
