package remote

import "github.com/google/go-containerregistry/pkg/name"

type Remote struct{}

type RemoteApi interface {
	ResolveDigest(ref name.Reference, opts ...Option) (*name.Digest, error)
}
