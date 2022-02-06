package cli

import (
	"fmt"

	"github.com/matthogan/zc/cmd/cn/cli/options"
	container "github.com/matthogan/zc/pkg/container"
	"github.com/spf13/cobra"
)

var (
	digest container.DigestApi
)

func init() {
	digest = &container.Digest{}
}

func Digest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "digest",
		Short: "Operations related to artifact digests in a registry",
	}
	cmd.AddCommand(
		getDigest(),
	)
	return cmd
}

func getDigest() *cobra.Command {
	o := &options.DigestOptions{}
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "Get the digest of an image",
		Long:    r.GetResourceAsString("digest_get"),
		Example: "  cn digest get <image uri>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			digest, err := digest.GetDigestFromRegistry(cmd.Context(), &o.Registry, args[0])
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", digest.DigestStr())
			return nil
		},
	}
	return cmd
}
