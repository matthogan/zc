# Usage

Operations provided by the built image. On the local build environment the
app image will be in the dist directory.

## Command line

The tool spits information, shortened here for brevity.

```shell
$ cn
~~~~## cn ##~~~~

Digests

Get digests from remote container registries.

        cn digest get <image-uri>

Usage:
  cn [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  digest      Operations related to artifact digests in a registry
  help        Help about any command
  version     Prints the version

Flags:
  -h, --help   help for cn

Use "cn [command] --help" for more information about a command.
```

## Download image digests

Get the image digest of an image in a remote registry.

```shell
$ dist/cn-darwin-amd64 digest get debian:latest
sha256:fb45fd4e25abe55a656ca69a7bef70e62099b8bb42a279a5e0ea4ae1ab410e0d
```
