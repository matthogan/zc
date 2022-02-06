package options

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	ociremote "github.com/matthogan/zc/pkg/oci/remote"
	"github.com/spf13/cobra"

	"github.com/matthogan/zc/pkg/version"
)

// RegistryOptions is the wrapper for the registry options.
type RegistryOptions struct {
	AllowInsecure      bool
	KubernetesKeychain bool
}

var _ Interface = (*RegistryOptions)(nil)

// AddFlags implements Interface
func (o *RegistryOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVar(&o.AllowInsecure, "allow-insecure-registry", false,
		"whether to allow insecure connections to registries. Don't use this for anything but testing")
	cmd.Flags().BoolVar(&o.KubernetesKeychain, "k8s-keychain", false,
		"whether to use the kubernetes keychain instead of the default keychain (supports workload identity).")
}

func (o *RegistryOptions) ClientOpts(ctx context.Context) ([]ociremote.Option, error) {
	opts := []ociremote.Option{ociremote.WithRemoteOptions(o.GetRegistryClientOpts(ctx)...)}
	targetRepoOverride, err := ociremote.GetEnvTargetRepository()
	if err != nil {
		return nil, err
	}
	if (targetRepoOverride != name.Repository{}) {
		opts = append(opts, ociremote.WithTargetRepository(targetRepoOverride))
	}
	return opts, nil
}

func (o *RegistryOptions) GetRegistryClientOpts(ctx context.Context) []remote.Option {
	opts := []remote.Option{
		remote.WithContext(ctx),
		remote.WithUserAgent("cn/" + version.GetVersionInfo().GitVersion),
	}
	if o.KubernetesKeychain {
		kc, err := k8schain.NewNoClient(ctx)
		if err != nil {
			panic(err.Error())
		}
		opts = append(opts, remote.WithAuthFromKeychain(kc))
	} else {
		opts = append(opts, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	}
	if o != nil && o.AllowInsecure {
		opts = append(opts, remote.WithTransport(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}})) // #nosec G402
	}
	return opts
}
