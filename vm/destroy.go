package vm

import (
	"os"
	"os/exec"
)

// Destroy removes the DTest configured virtual machine from
// Podman.
func Destroy() bool {
	name, _, ok := findRelevantVM()
	if !ok {
		lgr.Errorf("DTest VM (%s) does not exist", name)
		return false
	}

	cmd := exec.Command("podman", "machine", "rm", "-f", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		lgr.Errorf("Unable to remove DTest virtual machine %s from Podman due to error: %s", name, err)
		return false
	}
	lgr.Infof("DTest virtual machine %s removed", name)
	return true
}
