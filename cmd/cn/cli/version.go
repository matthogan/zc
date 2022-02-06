package cli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/matthogan/zc/pkg/version"
)

const (
	NAME = "cn"
)

func Version() *cobra.Command {
	var outputJSON bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version",
		Long:  "Prints the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			v := version.GetVersionInfo()
			res := v.String()
			if outputJSON {
				j, err := v.JSONString()
				if err != nil {
					return errors.Wrap(err, "unable to generate JSON from version info")
				}
				res = j
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", res)
			return nil
		},
	}

	cmd.Flags().BoolVar(&outputJSON, "json", false,
		"print JSON instead of text")

	return cmd
}
