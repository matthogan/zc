package cli

import (
	"github.com/spf13/cobra"

	"github.com/matthogan/zc/pkg/resources"
)

var (
	co = &Configuration{}
	r  resources.Resources
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               NAME,
		Short:             "container ops",
		Long:              r.GetResourceAsString("commands"),
		DisableAutoGenTag: true,
		SilenceUsage:      false, // Do show usage on errors
		SilenceErrors:     true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	// Add sub-commands.
	cmd.AddCommand(Digest())
	cmd.AddCommand(Version())
	co.Load()

	return cmd
}
