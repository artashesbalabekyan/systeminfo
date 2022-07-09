package hwinfo

import (
	"system-info/pkg/macos/applications"
	"system-info/pkg/macos/baseboard"
	"system-info/pkg/macos/cpu"
	"system-info/pkg/macos/firewall"
	"system-info/pkg/macos/memory"
)

func Get() (*HWInfo, error) {
	board, err := baseboard.GetBaseBoardInfo()
	if err != nil {
		return nil, err
	}

	cpu, err := cpu.GetCPU()
	if err != nil {
		return nil, err
	}

	fw, err := firewall.Get()
	if err != nil {
		return nil, err
	}

	apps, err := applications.Get()
	if err != nil {
		return nil, err
	}

	mem, err := memory.Get()
	if err != nil {
		return nil, err
	}

	return &HWInfo{
		Baseboard: board,
		CPU:       cpu,
		Firewall:  fw,
		Apps:      apps,
		Memory:    mem,
	}, nil
}
