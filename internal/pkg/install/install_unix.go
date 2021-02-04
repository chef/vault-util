// +build !windows

package install

// DefaultRootParentDir - default parent dir.
const DefaultRootParentDir string = "/opt"

// DefaultSettingsDir - default dir for settings.
const DefaultSettingsDir string = "/var/opt/ci-settings"

// binName returns the name of binary based on the OS.
func binName(binary string) string {
	return binary
}
