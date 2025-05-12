package vm

// Ready Initializes and/or starts the DTest configured virtual
// machine for use by podman.
func Ready() (bool, bool) {
	if !CheckPlatform() {
		return true, false
	}
	name, vmLine, ok := findRelevantVM()
	if !ok {
		ok = Init()
		if !ok {
			return false, false
		}
	}
	// Reset Variables
	name, vmLine, ok = findRelevantVM()
	line := vmLineSplit.Split(vmLine, -1)
	status := line[3]
	switch status {
	case "Currently running":
		lgr.Debugf("DTest virtual machine %s is currently running", name)
		return true, true
	default:
		ok = Start()
		if ok {
			return true, true
		}
	}
	return false, false
}
