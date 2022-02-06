package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionText(t *testing.T) {
	sut := GetVersionInfo()
	require.NotEmpty(t, sut.String())
}

func TestVersionJSON(t *testing.T) {
	sut := GetVersionInfo()
	json, err := sut.JSONString()

	require.Nil(t, err)
	require.NotEmpty(t, json)
}
