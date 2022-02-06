package resources

import (
	"embed"
)

var (
	//go:embed commands digest_get
	res embed.FS
)

type Resources struct{}

type ResourcesApi interface {
	GetResourceAsString(name string) string
}

// Returns a named embedded resource as a string. Panics
// on error retrieving the file
// name - filename
func (r Resources) GetResourceAsString(name string) string {
	data, err := res.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}
