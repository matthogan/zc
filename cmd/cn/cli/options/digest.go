package options

import (
	"github.com/spf13/cobra"
)

type DigestOptions struct {
	PrefixWithImage bool
	Registry        RegistryOptions
}

var _ Interface = (*DigestOptions)(nil)

// AddFlags implements Interface
func (o *DigestOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&o.PrefixWithImage, "prefix-with-image", false,
		"prefix the digest with the image ref")
	o.Registry.AddFlags(cmd)
}
