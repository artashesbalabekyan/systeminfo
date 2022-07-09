package sysreport

import (
	"encoding/json"
	"fmt"
)

func parseSPHardwareDataType(s []byte) (*SPHardwareDataType, error) {
	hw := SystemReport{}
	err := json.Unmarshal(s, &hw)
	return hw.SPHardwareDataType[0], err
}

func parseSPApplicationDataType(s []byte) ([]*SPApplicationsDataType, error) {
	hw := SystemReport{}
	err := json.Unmarshal(s, &hw)
	return hw.SPApplicationsDataType, err
}

func parseSPMemoryDataType(s []byte) ([]*SPMemoryDataType, error) {
	hw := SystemReport{}
	err := json.Unmarshal(s, &hw)
	return hw.SPMemoryDataType, err
}

func GetSPHardwareDataType() (*SPHardwareDataType, error) {
	hw := &SPHardwareDataType{}
	s, err := execSPHardwareDataType()
	if err != nil {
		return hw, err
	}

	return parseSPHardwareDataType(s)
}

func GetSPApplicationDataType() ([]*SPApplicationsDataType, error) {
	hw := []*SPApplicationsDataType{}
	s, err := execSPApplicationsDataType()
	if err != nil {
		return hw, err
	}

	return parseSPApplicationDataType(s)
}

func GetSPMemoryDataType() ([]*SPMemoryDataType, error) {
	hw := []*SPMemoryDataType{}
	s, err := execSPMemoryDataType()
	if err != nil {
		return hw, err
	}

	js, _ := json.Marshal(hw)
	fmt.Println(string(js))

	return parseSPMemoryDataType(s)
}
