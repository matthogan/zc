package options

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
)

func TestRegistryOptions(t *testing.T) {
	tests := []struct {
		name    string
		sut     RegistryOptions
		wantErr bool
	}{{
		name:    "allow-insecure-registry",
		sut:     RegistryOptions{},
		wantErr: false,
	}, {
		name:    "k8s-keychain",
		sut:     RegistryOptions{},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &tt.sut
			c := &cobra.Command{}
			o.AddFlags(c)
			f := c.Flags().Lookup(tt.name)
			if (f == nil) != tt.wantErr {
				t.Errorf("missing flag %s", tt.name)
				return
			}
		})
	}
}

func TestClientOpts(t *testing.T) {
	tests := []struct {
		name     string
		sut      RegistryOptions
		ctx      context.Context
		expected int
	}{{
		name:     "basic",
		sut:      RegistryOptions{},
		ctx:      context.Background(),
		expected: 1,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &tt.sut
			opts, err := o.ClientOpts(tt.ctx)
			if err != nil {
				t.Errorf("%s unexpected err %s", tt.name, err.Error())
				return
			}
			if len(opts) != tt.expected {
				t.Errorf("%s unexpected options length %d<>%d", tt.name, len(opts), tt.expected)
				return
			}
		})
	}
}

func TestGetRegistryClientOpts(t *testing.T) {
	tests := []struct {
		name     string
		sut      RegistryOptions
		ctx      context.Context
		expected int
	}{{
		name:     "basic",
		sut:      RegistryOptions{},
		ctx:      context.Background(),
		expected: 3,
	}, {
		name: "all",
		sut: RegistryOptions{
			KubernetesKeychain: true,
			AllowInsecure:      true,
		},
		ctx:      context.Background(),
		expected: 4,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &tt.sut
			opts := o.GetRegistryClientOpts(tt.ctx) // private api
			if len(opts) != tt.expected {
				t.Errorf("%s unexpected options length %d<>%d",
					tt.name, len(opts), tt.expected)
				return
			}
		})
	}
}
