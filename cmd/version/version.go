package version

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Config    = ""
	BasePath  = ""
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
		RunE: run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("0.0.0")
	return nil
}
