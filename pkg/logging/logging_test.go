package logging

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogging_Init(t *testing.T) {
	actual := Init()
	require.NotEmpty(t, actual)
}

func TestLogging_Logger(t *testing.T) {
	actual := Logger("test")
	require.NotEmpty(t, actual)
}
