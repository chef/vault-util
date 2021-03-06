package commands

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/avast/retry-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	filesystem "github.com/chef/go-filesystem"
	"github.com/chef/vault-util/internal/pkg/install"
	"github.com/chef/vault-util/internal/pkg/secrets"
)

var (
	execCommand = exec.Command

	ciutils install.Install

	fs filesystem.FileSystem

	fslock filesystem.Locker

	version string

	rootCmd = &cobra.Command{
		Use:          "vault-util",
		Short:        "Utility to access secrets and account information stored in Hashicorp Vault from CI.",
		SilenceUsage: true,
		Version:      version,
	}

	secretsClient secrets.Client
)

// Execute handles the execution of child commands and flags.
func Execute() {
	var err error
	var retryAttempts uint = 10
	var retryDelay time.Duration = 100

	fs = filesystem.NewOsFs()
	ciutils = install.DefaultInstall()
	fslock = &filesystem.OsLock{
		RetryAttempts:  retryAttempts,
		RetryDelay:     retryDelay * time.Millisecond,
		RetryDelayType: retry.BackOffDelay,
	}

	secretsClient, err = secrets.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	if err = rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.SetOut(os.Stdout)

	rootCmd.SetVersionTemplate("{{.Version}}\n")

	cobra.OnInitialize(initConfig)

	viper.SetDefault("aws.region", "us-west-2")
	viper.SetDefault("vault.dynamic_mount", "account/dynamic")
	viper.SetDefault("vault.static_mount", "account/static")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Set environment variable prefix, eg: VAULT_UTIL_AWS_REGION
	viper.SetEnvPrefix("vault_util")

	// Override our config with any matching environment variables
	viper.AutomaticEnv()
}
