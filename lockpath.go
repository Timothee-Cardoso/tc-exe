// +build !windows

package tcexe

import "os/exec"

func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
