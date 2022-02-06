package container

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/matthogan/zc/cmd/cn/cli/options"
	ociremote "github.com/matthogan/zc/pkg/oci/remote"
)

var (
	remote ociremote.RemoteApi
)

func init() {
	remote = &ociremote.Remote{}
}

type Digest struct{}

type DigestApi interface {
	GetDigestFromRegistry(ctx context.Context, regOpts *options.RegistryOptions, imageRef string) (*name.Digest, error)
}

func (d *Digest) GetDigestFromRegistry(ctx context.Context, regOpts *options.RegistryOptions, imageRef string) (*name.Digest, error) {
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return nil, err
	}
	ociremoteOpts, err := regOpts.ClientOpts(ctx)
	if err != nil {
		return nil, err
	}
	digest, err := remote.ResolveDigest(ref, ociremoteOpts...)
	if err != nil {
		return nil, err
	}
	return digest, nil
}
