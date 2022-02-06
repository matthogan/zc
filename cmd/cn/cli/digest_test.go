package cli

import (
	"bytes"
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/matthogan/zc/cmd/cn/cli/options"
	"github.com/stretchr/testify/require"
)

func TestDigest(t *testing.T) {
	sut := Digest()
	require.NotNil(t, sut)
	require.NotNil(t, sut.Use)
	require.NotNil(t, sut.Short)
	require.NotNil(t, sut.Example)
	require.Nil(t, sut.RunE)
	require.NotNil(t, sut.Commands())
}

type DigestMock struct{}

var getDigestFromRegistryMock func(ctx context.Context, regOpts *options.RegistryOptions, imageRef string) (*name.Digest, error)

func (d *DigestMock) GetDigestFromRegistry(ctx context.Context, regOpts *options.RegistryOptions, imageRef string) (*name.Digest, error) {
	return getDigestFromRegistryMock(ctx, regOpts, imageRef)
}

func TestDigest_GetDigestFromRegistry(t *testing.T) {
	sut := Digest()
	b := bytes.NewBufferString("")
	sut.SetOut(b)
	sut.SetArgs([]string{"get", "abcd"})
	digest = &DigestMock{}
	d, err := name.NewDigest("image@sha256:0cae3cc26f4f6cf6d57d51239545bf46212bfed24a2d9df9a581a6b47ec6f532",
		name.WithDefaultRegistry("x.io"))
	if err != nil {
		t.Errorf("unexpected error %d", err)
	}
	getDigestFromRegistryMock = func(ctx context.Context, regOpts *options.RegistryOptions, imageRef string) (*name.Digest, error) {
		return &d, nil
	}
	sut.Execute()
	if err != nil {
		t.Errorf("unexpected error %d", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(out)) != d.DigestStr() {
		t.Fatalf("expected \"%s\" got \"%s\"", d.DigestStr(), strings.TrimSpace(string(out)))
	}
}
