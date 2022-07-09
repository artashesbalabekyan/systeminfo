package baseboard

import (
	"system-info/pkg/macos/sysreport"
)

func GetBaseBoardInfo() (*BaseBoard, error) {
	hw, err := sysreport.GetSPHardwareDataType()
	if err != nil {
		return nil, err
	}

	b := BaseBoard(*hw)

	return &b, nil
}
