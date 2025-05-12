package vm

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var matchDtest = regexp.MustCompile("^dtest-")

// List builds a list of virtual machines, singling out those
// configured for use by DTest, then prints its name to stdout.
func List() {
	if !CheckPlatform() {
		linuxError()
		return
	}
	ret, ok := findVMs()
	if ok {
		fmt.Println(strings.Join(ret, "\n"))
	}
}

func findVMs() ([]string, bool) {
	lgr.Debug("Checking for Podman machines")
	ret, size, dsize := listVms()

	if size == 0 {
		lgr.Errorf("podman machine ls returned no virtual machines")
		return ret, false
	}
	lgr.Infof("Found %d virtual machines", size)

	if dsize == 0 {
		lgr.Errorf("podman machine ls shows no dtest configured virtual machines")
		return ret, false
	}
	lgr.Infof("Found %d dtest configured virtual machines", dsize)
	return ret, true
}

// CountVMs returns a count of virtual machines configured for
// use by DTest. It's mainly used as a check, to determine
// whether Podman is properly configured with VM's.
func CountVMs() int {
	_, _, dsize := listVms()
	return dsize
}

func listVms() ([]string, int, int) {
	cmd := exec.Command("podman", "machine", "list")
	out, err := cmd.Output()
	empty := []string{}

	if err != nil {
		lgr.Fatalf("Unable to retrieve list of Podman machines: %s", err)
		return empty, 0, 0
	}
	lines := strings.Split(string(out), "\n")
	size := len(lines) - 1
	if size == 0 {
		return empty, size, 0
	}
	ret := []string{"", lines[0]}
	for i := 1; i < size; i++ {
		l := lines[i]
		if matchDtest.MatchString(l) {
			ret = append(ret, l)
		}
	}
	dsize := len(ret) - 2

	return ret, size, dsize
}
