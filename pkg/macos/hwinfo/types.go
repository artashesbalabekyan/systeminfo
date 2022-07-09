package hwinfo

import (
	"system-info/pkg/macos/applications"
	"system-info/pkg/macos/baseboard"
	"system-info/pkg/macos/cpu"
	"system-info/pkg/macos/firewall"
	"system-info/pkg/macos/memory"
)

type HWInfo struct {
	Baseboard *baseboard.BaseBoard         `json:"baseboard"`
	CPU       *cpu.CPU                     `json:"cpu"`
	Firewall  *firewall.FirewallBeautified `json:"firewall"`
	Apps      []*applications.Application  `json:"applications"`
	Memory    []*memory.Memory             `json:"memory"`
}
