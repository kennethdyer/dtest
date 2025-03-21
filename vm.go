package dtest

import "github.com/kennethdyer/dtest/vm"

func RunVM(args []string) {
	lgr.Debug("Called vm operation with sub-command qualifier. Defaulting to vm ls command.")
	RunVMList()
}

func RunVMList() {
	lgr.Info("Called vm list operation")
	vm.List()
}
