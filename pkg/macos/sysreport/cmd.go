package sysreport

import (
	"os/exec"
)

func execSPHardwareDataType() ([]byte, error) {
	return execCmd("SPHardwareDataType")
}

func execSPApplicationsDataType() ([]byte, error) {
	return execCmd("SPApplicationsDataType")
}

func execSPMemoryDataType() ([]byte, error) {
	return execCmd("SPMemoryDataType")
}

func execCmd(cmd string) ([]byte, error) {
	return exec.Command("/usr/sbin/system_profiler", cmd, "-json", "-detailLevel", "full").Output()
}
