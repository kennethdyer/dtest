package vm

import (
	"os"
	"os/exec"
	"runtime"
)

// Start runs from the command-line and is used to start the
// DTest configured virtual machine in Podman.
func Start() bool {
	if !CheckPlatform() {
		linuxError()
		return true
	}
	name, vmLine, ok := findRelevantVM()
	if !ok {
		lgr.Errorf("DTest virtual machine %s does not exist", name)
		return
	}

	line := vmLineSplit.Split(vmLine, -1)
	status := line[3]
	switch status {
	case "Currently running":
		lgr.Warnf("DTest virtual machine %s is currently running", name)
		return false
	}
	cmd := exec.Command("podman", "machine", "start", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		lgr.Errorf("Unable to start DTest virtual machine %s due to error: %s", name, err)
		return false
	}
	lgr.Infof("DTest virtual machine %s started", name)
	return true
}

// Stop runs from the command-line and is used to stop the
// DTest configured virtual machine in Podman.
func Stop() bool {
	if !CheckPlatform() {
		linuxError()
		return true
	}
	name, vmLine, ok := findRelevantVM()
	if !ok {
		lgr.Errorf("DTest virtual machine %s does not exist", name)
	}
	line := vmLineSplit.Split(vmLine, -1)
	status := line[3]
	switch status {
	case "Currently running":
		lgr.Debugf("Stopping DTest virtual machine %s", name)
	default:
		lgr.Debugf("DTest virtual machine %s is not running", name)
		return false
	}
	cmd := exec.Command("podman", "machine", "stop", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		lgr.Errorf("Unable to stop DTest virtual machine %s due to error: %s", name, err)
		return false
	}
	lgr.Infof("DTest virtual machine %s stopped", name)
	return true
}

// CheckPlatform uses runtime.GOOS to identify the platform
// DTest is running on. If the platform is macOS, it means
// Podman requires VM management. If the platform is Linux, it
// means DTest should ignore VM commands. If Windows, it should
// error out, as DTest doesn't support Windows.
func CheckPlatform() bool {
	name := runtime.GOOS
	switch name {
	case "linux":
		return false
	case "darwin":
		return true
	default:
		lgr.Fatalf("DTest does not support the %s platform", name)
		os.Exit(1)
	}
	return false
}

func linuxError() {
	lgr.Errorf("DTest does not require virtual machine management when run on Linux")
}
