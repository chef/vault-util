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

### Visual Studio Code Development Environment

This project contains the following Visual Studio Code (a.k.a. VS Code) configuration files to make it easy to start developing.

* [Development Container](https://code.visualstudio.com/docs/remote/containers) configuration files in [.devcontainer](.devcontainer)
* [Run/Debug (Launch)](https://code.visualstudio.com/docs/editor/debugging) configuration file in [.vscode/launch.json](.vscode/launch.json)

Please be sure that the following required software is installed on your workstation.

* [VS Code](https://code.visualstudio.com/download)
* [VS Code's Remote - Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
* [Docker](https://docs.docker.com/get-docker/)

Clone this repository to your workstation and open it in VS Code.

When you open the project folder in VS Code you should see a notification in the bottom, right saying `Folder contains a Development Container configuration file. Reopen folder to develop in a container`. Clicking the `Reopen in Container` button tells VS Code to build the development container image, start the container and connect to it for development. If you do not see the notification you can open the list of remote commands by clicking on the green section on the bottom, left side of the status bar and then selecting `Remote-Containers: Reopen in Container`.

It will take some time for VS Code to build, start and connect to the development container the first time you open the project but subsequent usage will be much faster.

VS Code should now be connected to development container with all the tools and settings necessary to build and test this project. If your workflow uses VS Code extensions that are not installed in the development container by default (e.g. eamodio.gitlens, github.vscode-pull-request-github) you can add them to the list of extensions in [.devcontainer/devcontainer.json](.devcontainer/devcontainer.json) and rebuild the development container.

The provided [.vscode/launch.json](.vscode/launch.json) configuration file contains a configuration named `configure-accounts` for running/debugging the `vault-util configure-accounts` command and a configuration named `fetch-secret-env` for running/debugging the `vault-util fetch-secret-env` command. These are configured to connect to our `vault.es.chef.co` server. If you want to use either of these you will need to copy the `.env.example` file to `.env` and replace `vault-token` with your vault token. You can view your vault token by running `vault print token`. Also, be sure to properly set `VAULT_ADDR` and `VAULT_NAMESPACE` in the [.vscode/launch.json](.vscode/launch.json) configuration file.

If you make changes to files in the `.devcontainer` folder you might need to rebuild the development container by selecting `Remote-Containers: Rebuild Container` from the list of remote commands.

When you are done using the development container you can stop it by simply closing the project's VS Code window or you can select `Remote-Containers: Reopen Locally` from the list of remote commands.
