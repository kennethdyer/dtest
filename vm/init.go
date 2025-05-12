package vm

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/viper"
)

func findRelevantVM() (string, string, bool) {
	// Set name
	base := viper.GetString("machine")
	if base == "" {
		lgr.Debug("machine configuration value is unset, defaulting to 'machine'")
		base = "machine"
	}
	name := fmt.Sprintf("dtest-%s", base)
	lgr.Debugf("Setting machine name for use with Podman: %s", name)

	// Find VM's
	lgr.Debug("Checking Podman for configured virtual machines")
	vms, ok := findVMs()
	if ok {
		nameMatch := regexp.MustCompile(fmt.Sprintf("^%s", name))
		for _, vm := range vms {
			if nameMatch.MatchString(vm) {
				return name, vm, true
			}
		}
		lgr.Debugf("No virtual machines for this project found, initializing %s", name)
		return name, "", false
	}
	lgr.Debugf("No virtual machines found, initializing %s", name)
	return name, "", false
}

// Init runs from the command-line and is used to initialize a
// virtual machine in Podman for use by DTest.
func Init() bool {
	if !CheckPlatform() {
		linuxError()
		return true
	}
	name, _, ready := findRelevantVM()

	if ready {
		lgr.Errorf("Podman machine %s already exists", name)
		return false
	}
	// Initialize Virtual Machine
	cmd := exec.Command("podman", "machine", "init", name)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		lgr.Errorf("Unable to initialize podman machine %q, due to error: %s", name, err)
		return false
	}
	lgr.Debugf("Podman machine %s initialized", name)

	return true
}
