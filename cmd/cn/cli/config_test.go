package cli

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestGetenv(t *testing.T) {
	sut := envVar.Getenv("_")
	require.NotEmpty(t, sut)
}

func TestGetenvSlice(t *testing.T) {
	sut := envVar.GetenvSlice("PATH", ":")
	require.NotNil(t, sut)
	require.Contains(t, sut, "/usr/local/bin")
}

func TestConfig(t *testing.T) {
	sut := Config()
	require.NotNil(t, sut)
	require.NotNil(t, sut.Use)
	require.NotNil(t, sut.Short)
	require.NotNil(t, sut.RunE)
	sut.RunE(nil, []string{})
	viper.AddConfigPath("config")
}

var getenvMock func(key string) string
var getenvSliceMock func(key string, delim string) []string

type envyVarMock struct{}

func (m envyVarMock) Getenv(key string) string {
	return getenvMock(key)
}
func (m envyVarMock) GetenvSlice(key string, delim string) []string {
	return getenvSliceMock(key, delim)
}

func TestLoad(t *testing.T) {
	sut := &Configuration{}
	envVar = envyVarMock{}
	getenvMock = func(key string) string {
		return ""
	}
	getenvSliceMock = func(key string, delim string) []string {
		return []string{}
	}
	sut.Load()
}

func TestInit(t *testing.T) {
	sut := &Configuration{}
	err := sut.Init([]string{"doesnotexist"}, "", "yaml")
	if err != nil {
		t.Errorf("error %d", err)
	}
}

func TestInit_FormatError(t *testing.T) {
	sut := &Configuration{}
	err := sut.Init([]string{}, "", "foobar")
	if err == nil {
		t.Error("is foobar a valid format?")
	}
}

func TestInit_Found(t *testing.T) {
	sut := &Configuration{}
	err := sut.Init([]string{"local"}, "../../../config", "yaml")
	if err != nil {
		t.Errorf("error %d", err)
	}
}
