package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	sut := New()
	require.NotNil(t, sut)
	require.NotNil(t, sut.Use)
	require.NotNil(t, sut.Short)
	require.NotNil(t, sut.Long)
	require.Nil(t, sut.RunE)
	require.False(t, sut.SilenceUsage)
	require.True(t, sut.DisableAutoGenTag)
	require.True(t, sut.SilenceErrors)
}
