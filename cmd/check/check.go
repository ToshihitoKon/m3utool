package check

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ToshihitoKon/m3utool/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	config    = ""
	basePath  = ""
	errorFlag = false
)

func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
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
			check(args[0])
		},
	}

	cmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
	cmd.Flags().StringVarP(&basePath, "basepath", "b", "", "m3u base file path")

	return cmd
}

func check(m3uPath string) {
	for _, v := range utils.ReadFile(m3uPath) {
		path := filepath.Join(basePath, v)
		if ok := utils.CheckFileExist(path); !ok {
			fmt.Println(path)
			errorFlag = true
			continue
		}
	}

	if errorFlag {
		os.Exit(1)
	}

	return
}
