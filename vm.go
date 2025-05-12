package dtest

import "github.com/kennethdyer/dtest/vm"

// RunVM runs from the command-line and performs the default
// operation for the vm sub-command, which runs the list
// operation. The command prints an error to stderr if there's
// no dtest configured VM on the system.
func RunVM(args []string) {
	lgr.Debug("Called vm operation with sub-command qualifier. Defaulting to vm ls command.")
	RunVMList()
}

// RunVMList runs from the command-line and performs the list
// operation for the Podman VM. The ls sub-command prints an
// error to stderr if there isn't dtest configured VM on the
// system.
func RunVMList() {
	lgr.Info("Called vm list operation")
	vm.List()
}

// RunVMStatus runs from the command-line and performs the
// status operation for the Podman VM. The sub-command prints an
// error to stderr if there isn't a dtest configured VM on the
// system.
func RunVMStatus() {
	vm.Status()
}

// RunVMStart runs from the command-line and starts the dtest
// configured VM for Podman. The sub-command returns an error if
// there is no dtest configured VM on the system, or if there is
// and it's already running.
func RunVMStart() {
}

// RunVMStop runs from the command-line and stops the dtest
// configured VM for Podman. The sub-command returns an error if
// there is no dtest configured VM on the system, or if there is
// and it's already stopped.
func RunVMStop() {
}

// RunVMInit runs from the command-line and initializes the
// Podman virtual machine used by DTest.
func RunVMInit() {
	ok := vm.Init()
	if !ok {
		lgr.Fatal("VM initialization process exited with an error")
		return
	}
	lgr.Debug("VM initialization process complete.")
}

// RunVMDestroy runs from the command-line and destroys the
// Podman virtual machine used by DTest.
func RunVMDestroy() {
}
