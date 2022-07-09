package baseboard

type BaseBoard struct {
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
