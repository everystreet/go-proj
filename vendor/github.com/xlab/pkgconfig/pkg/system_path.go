package pkg

import "os/exec"

// getSystemPath returns the installed pkg-config's default PKG_CONFIG_PATH, if
// it is available, otherwise it returns an empty string.
func getSystemPath() string {
	cmd := exec.Command("pkg-config", "--variable=pc_path", "pkg-config")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}
