package memory

import "system-info/pkg/macos/sysreport"

func Get() ([]*Memory, error) {
	memory, err := sysreport.GetSPMemoryDataType()
	if err != nil {
		return nil, err
	}

	memList := []*Memory{}

	for _, m := range memory {
		mTmp := Memory(*m)
		memList = append(memList, &mTmp)
	}

	return memList, nil
}
