package cpu

import (
	"syscall"
)

func GetCPU() (*CPU, error) {
	name, err := getName()
	if err != nil {
		return nil, err
	}

	threads, err := getThreads()
	if err != nil {
		return nil, err
	}

	cores, err := getCores()
	if err != nil {
		return nil, err
	}

	return &CPU{
		Name:    name,
		Threads: int(threads),
		Cores:   int(cores),
	}, nil
}

func getName() (string, error) {
	return syscall.Sysctl("machdep.cpu.brand_string")
}

func getThreads() (uint32, error) {
	return syscall.SysctlUint32("machdep.cpu.thread_count")
}

func getCores() (uint32, error) {
	return syscall.SysctlUint32("machdep.cpu.core_count")
}
