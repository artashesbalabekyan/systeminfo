package applications

import "system-info/pkg/macos/sysreport"

func Get() ([]*Application, error) {
	apps, err := sysreport.GetSPApplicationDataType()
	if err != nil {
		return nil, err
	}

	appsList := []*Application{}

	for _, app := range apps {
		appsList = append(appsList, &Application{
			Name:    app.Name,
			Path:    app.Path,
			Version: app.Version,
		})
	}

	return appsList, nil
}
