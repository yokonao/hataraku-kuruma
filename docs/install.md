# Install

hataraku-kuruma is a single binary for linux and macOS (amd64/arm64). Install it into a directory in your `PATH` in one of these ways:

1. [GitHub Releases](#github-releases)
1. [mise](#mise)
1. [go install](#go-install)

## GitHub Releases

Download the archive for your platform from [GitHub Releases](https://github.com/yokonao/hataraku-kuruma/releases), then install the binary into `PATH`:

```sh
asset=hataraku-kuruma_darwin_arm64.tar.gz # hataraku-kuruma_{linux,darwin}_{amd64,arm64}.tar.gz
gh release download -R yokonao/hataraku-kuruma -p "$asset" # the latest release; pass a tag for another
tar -xzf "$asset" hataraku-kuruma
mkdir -p ~/.local/bin
install -m 755 hataraku-kuruma ~/.local/bin/
```

The macOS binaries are not notarized. `gh` and `curl` do not mark downloads as quarantined, but a browser does, and macOS then refuses to run the binary; remove the mark with `xattr -d com.apple.quarantine hataraku-kuruma`.

### Verify the archive

Each archive has a [build provenance attestation](https://docs.github.com/en/actions/security-for-github-actions/using-artifact-attestations) from the release workflow. Verify it with the [GitHub CLI](https://cli.github.com/):

```sh
gh attestation verify "$asset" \
  -R yokonao/hataraku-kuruma \
  --signer-workflow yokonao/hataraku-kuruma/.github/workflows/release.yml
```

## mise

With [mise](https://mise.jdx.dev/):

```sh
mise use -g github:yokonao/hataraku-kuruma
```

mise verifies the [build provenance attestation](#verify-the-archive) automatically.

## go install

```sh
go install github.com/yokonao/hataraku-kuruma/cmd/hataraku-kuruma@latest
```

The binary lands in `$(go env GOBIN)`, or `~/go/bin` when that is unset.
