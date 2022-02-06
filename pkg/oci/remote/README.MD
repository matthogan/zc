# COSIGN

The `oci.remote` package has been copied from the
[cosign](https://github.com/sigstore/cosign) package. The
following changes have been made.

## Additional files

[api.go](pkg/../api.go)

Exposes a method for mocking in unit tests.

## Changed files

[digest.go](pkg/../digest.go)

Returns a pointer to the name.Digest value to enable nil references in unit tests.

## Other changes

Package renaming to enable the copy.

## Reference

<https://github.com/sigstore/cosign/tree/main/pkg/oci>
