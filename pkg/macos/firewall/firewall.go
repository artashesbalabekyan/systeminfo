package firewall

import (
	"encoding/json"
	"os"

	"github.com/groob/plist"
)

func GetFirewallInfo() (*Firewall, error) {
	f, err := os.Open("/Library/Preferences/com.apple.alf.plist")
	if err != nil {
		return nil, err
	}

	var tmp interface{}
	fv := Firewall{}

	d := plist.NewBinaryDecoder(f)
	err = d.Decode(&tmp)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(js, &fv)
	if err != nil {
		return nil, err
	}
	return &fv, nil
}

func Get() (*FirewallBeautified, error) {
	fw, err := GetFirewallInfo()
	if err != nil {
		return nil, err
	}

	rules := []Rule{}
	for _, app := range fw.Applications {
		if app.Bundleid != "" {
			rules = append(rules, app.GetRule())
		}
	}
	state := "disabled"
	if fw.Globalstate == 1 {
		state = "enabled"
	}
	fwBeautified := &FirewallBeautified{
		ActiveState: state,
		Rules:       rules,
	}
	return fwBeautified, nil
}
