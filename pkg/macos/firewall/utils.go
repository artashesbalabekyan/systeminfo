package firewall

func (app Application) GetRule() Rule {
	enabled := "Block incoming connections"
	if app.State == 0 {
		enabled = "Allow incoming connections"
	}
	rule := Rule{
		Name: app.Bundleid,
		Rule: enabled,
	}

	return rule
}

