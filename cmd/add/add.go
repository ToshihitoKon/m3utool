package add

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ToshihitoKon/m3utool/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	Config   = ""
	BasePath = ""
)

func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
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
			add(args[0])
		},
	}
	cmd.Flags().StringVarP(&Config, "config", "c", "", "config file path")
	cmd.Flags().StringVarP(&BasePath, "basepath", "b", "", "m3u base file path")

	return cmd
}

func add(m3uPath string) {
	fmt.Println("config path:", Config)

	for i, v := range utils.ReadFile(m3uPath) {
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
}
