# vault-util

vault-util is a command line go utility/binary to access secrets and account information stored in Hashicorp Vault from CI.

### Installation

You can install the vault-util utility for a specific OS at https://github.com/chef/vault-util/releases

Currently available for:
OS | ARCH
--- | ---
Darwin | x86-64 (amd64)
Linux | x86-64 (amd64)
Windows | x86-64 (amd64)

### Build

You can optionally build vault-util via the [goreleaser](https://goreleaser.com/) utility. Goreleaser uses a `.goreleaser.yml` to define how you would like to build and release your go binary(ies).

### `vault-util`
<!-- stdout "./build/linux/vault-util --help" -->
```
Utility to access secrets and account information stored in Hashicorp Vault from CI.

Usage:
  vault-util [command]

Available Commands:
  configure-accounts    Configure the accounts specified in the VAULT_UTIL_ACCOUNTS environment variable.
  fetch-secret-env      Fetch the secrets specified in the VAULT_UTIL_SECRETS environment variable from Vault.
  help                  Help about any command
  print-git-credentials Utility that will print credentials for a GitHub App from Vault in git-credential-helper format.

Flags:
  -h, --help   help for vault-util

Use "vault-util [command] --help" for more information about a command.
```
<!-- stdout -->