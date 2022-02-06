package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	version "github.com/matthogan/zc/pkg/version"
	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	sut := Version()
	require.NotNil(t, sut)
	require.NotNil(t, sut.Use)
	require.NotNil(t, sut.Short)
	require.NotNil(t, sut.Example)
	require.NotNil(t, sut.RunE)
}

func TestVersion_RunE_Json(t *testing.T) {
	sut := Version()
	b := bytes.NewBufferString("")
	sut.SetOut(b)
	sut.SetArgs([]string{"--json", "true"})
	err := sut.Execute()
	if err != nil {
		t.Errorf("error %d", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	actual := version.Info{}
	if err := json.Unmarshal(out, &actual); err != nil {
		t.Fatalf("expected to be able to unmarshal \"%s\" as json. error %s", out, err.Error())
	}
}

func TestVersion_RunE(t *testing.T) {
	sut := Version()
	b := bytes.NewBufferString("")
	sut.SetOut(b)
	sut.SetArgs([]string{"--json", "false"})
	err := sut.Execute()
	if err != nil {
		t.Errorf("error %d", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatalf("expected a value but got \"%s\" instead", out)
	}
}
