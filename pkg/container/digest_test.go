package container

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/matthogan/zc/cmd/cn/cli/options"
	ociremote "github.com/matthogan/zc/pkg/oci/remote"
)

type RemoteMock struct{}

var resolveDigestMock func(ref name.Reference, opts ...ociremote.Option) (*name.Digest, error)

func (d *RemoteMock) ResolveDigest(ref name.Reference, opts ...ociremote.Option) (*name.Digest, error) {
	return resolveDigestMock(ref, opts...)
}

func TestDigest_GetDigestFromRegistry(t *testing.T) {
	remote = &RemoteMock{}
	d, err := name.NewDigest("image@sha256:0cae3cc26f4f6cf6d57d51239545bf46212bfed24a2d9df9a581a6b47ec6f532",
		name.WithDefaultRegistry("x.io"))
	if err != nil {
		t.Errorf("unexpected error %d", err)
	}
	resolveDigestMock = func(ref name.Reference, opts ...ociremote.Option) (*name.Digest, error) {
		return &d, nil
	}
	sut := Digest{}
	imageRef := "image:latest"
	regOpts := &options.RegistryOptions{}
	actual, err := sut.GetDigestFromRegistry(context.Background(), regOpts, imageRef)
	if err != nil {
		t.Errorf("unexpected error %d", err)
	}
	if actual.DigestStr() != d.DigestStr() {
		t.Fatalf("expected \"%s\" got \"%s\"", d.DigestStr(), actual.DigestStr())
	}
}
func TestDigest_GetDigestFromRegistry_Error(t *testing.T) {
	remote = &RemoteMock{}
	resolveDigestMock = func(ref name.Reference, opts ...ociremote.Option) (*name.Digest, error) {
		return nil, errors.New("error")
	}
	sut := Digest{}
	imageRef := "image:latest"
	regOpts := &options.RegistryOptions{}
	if _, err := sut.GetDigestFromRegistry(context.Background(), regOpts, imageRef); err == nil {
		t.Errorf("expected error %s", "error")
	}
}
