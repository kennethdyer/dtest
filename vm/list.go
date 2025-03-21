package vm

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var matchDtest = regexp.MustCompile("^dtest")

func List() {

	lgr.Debug("Checking for Podman machines")
	ret, size, dsize := listVms()

	if size == 0 {
		lgr.Errorf("podman machine ls returned no virtual machines")
		return
	}
	lgr.Infof("Found %d virtual machines", size)

	if dsize == 0 {
		lgr.Errorf("podman machine ls shows no dtest configured virutal machines")
		return
	}
	lgr.Infof("Found %d dtest configured virtual machines", size)
	fmt.Println(ret)

}

func CountVMs() int {
	_, _, dsize := listVms()
	return dsize
}

func listVms() (string, int, int) {
	cmd := exec.Command("podman", "machine", "list")
	out, err := cmd.Output()

	if err != nil {
		lgr.Fatalf("Unable to retrieve list of Podman machines: %s", err)
		return "", 0, 0
	}
	lines := strings.Split(string(out), "\n")
	size := len(lines) - 1
	if size == 0 {
		return "", size, 0
	}
	ret := []string{"", lines[0]}
	for i := 1; i < size; i++ {
		l := lines[i]
		if matchDtest.MatchString(l) {
			ret = append(ret, l)
		}
	}
	dsize := len(ret) - 2

	return strings.Join(ret, "\n"), size, dsize
}
