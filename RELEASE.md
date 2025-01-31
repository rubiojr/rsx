# Releasing RSX

* Make sure .goreleaser build tags match the go install tags in the README.
* script/release <version> (i.e. script/release v0.1.2)

A GitHub action will build and publish a new release using goreleaser.
