package sysreport

type SystemReport struct {
	SPHardwareDataType     []*SPHardwareDataType     `json:"SPHardwareDataType"`
	SPApplicationsDataType []*SPApplicationsDataType `json:"SPApplicationsDataType"`
	SPMemoryDataType       []*SPMemoryDataType       `json:"SPMemoryDataType"`
}

type SPHardwareDataType struct {
	ActivationLockStatus string `json:"activation_lock_status"`
	BootRomVersion       string `json:"boot_rom_version"`
	ChipType             string `json:"chip_type"`
	MachineModel         string `json:"machine_model"`
	MachineName          string `json:"machine_name"`
	NumberProcessors     int    `json:"number_processors"`
	OsLoaderVersion      string `json:"os_loader_version"`
	PhysicalMemory       string `json:"physical_memory"`
	PlatformUUID         string `json:"platform_UUID"`
	ProvisioningUDID     string `json:"provisioning_UDID"`
	SerialNumber         string `json:"serial_number"`
}

type SPApplicationsDataType struct {
	Name         string   `json:"_name"`
	ArchKind     string   `json:"arch_kind"`
	Info         string   `json:"info"`
	LastModified string   `json:"lastModified"`
	ObtainedFrom string   `json:"obtained_from"`
	Path         string   `json:"path"`
	SignedBy     []string `json:"signed_by"`
	Version      string   `json:"version"`
}

type SPMemoryDataType struct {
	SPMemoryDataType string `json:"SPMemoryDataType"`
	DimmType         string `json:"dimm_type"`
}
