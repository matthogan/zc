package options

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestDigestOptions(t *testing.T) {
	tests := []struct {
		name         string
		sut          DigestOptions
		defaultValue string
	}{{
		name:         "prefix-with-image",
		sut:          DigestOptions{},
		defaultValue: "false",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &tt.sut
			c := &cobra.Command{}
			o.AddFlags(c)
			f := c.Flags().Lookup(tt.name)
			if f == nil {
				t.Errorf("missing flag %s", tt.name)
				return
			}
			if f.Value.String() != tt.defaultValue {
				t.Errorf("%s unexpected flag value %s<>%s", tt.name, f.Value.String(), tt.defaultValue)
				return
			}
		})
	}
}
