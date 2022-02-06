package resources

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResources_GetResourceAsString(t *testing.T) {
	sut := Resources{}
	actual := sut.GetResourceAsString("commands")
	require.NotEmpty(t, actual)
	require.True(t, strings.HasPrefix(actual, "~~~~## cn ##~~~~"))
}

func TestResources_GetResourceAsString_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("doesnotexist does exist")
		}
	}()
	sut := Resources{}
	_ = sut.GetResourceAsString("doesnotexist")
}
