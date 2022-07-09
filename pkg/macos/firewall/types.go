package firewall

type Firewall struct {
	Allowdownloadsignedenabled int           `json:"allowdownloadsignedenabled"`
	Allowsignedenabled         int           `json:"allowsignedenabled"`
	Applications               []Application `json:"applications"`
	Exceptions                 []Exception   `json:"exceptions"`
	Explicitauths              []ID          `json:"explicitauths"`
	Firewall                   FW            `json:"firewall"`
	Firewallunload             int           `json:"firewallunload"`
	Globalstate                int           `json:"globalstate"`
	Loggingenabled             int           `json:"loggingenabled"`
	Loggingoption              int           `json:"loggingoption"`
	Previousonstate            int           `json:"previousonstate"`
	Stealthenabled             int           `json:"stealthenabled"`
	Version                    string        `json:"version"`
}

type Application struct {
	Alias    string `json:"alias"`
	Bundleid string `json:"bundleid"`
	Reqdata  string `json:"reqdata"`
	State    int    `json:"state"`
}

type ProcState struct {
	Proc  string `json:"proc"`
	State int    `json:"state"`
}

type ODSAgent struct {
	Proc            string `json:"proc"`
	Servicebundleid string `json:"servicebundleid"`
	State           int    `json:"state"`
}

type FW struct {
	Apple_Remote_Desktop  ProcState `json:"Apple Remote Desktop"`
	FTP_Access            ProcState `json:"FTP Access"`
	ODSAgent              ODSAgent  `json:"ODSAgent"`
	Personal_File_Sharing ProcState `json:"Personal File Sharing"`
	Personal_Web_Sharing  ProcState `json:"Personal Web Sharing"`
	Printer_Sharing       ProcState `json:"Printer Sharing"`
	Remote_Apple_Events   ProcState `json:"Remote Apple Events"`
	Remote_Login___SSH    ProcState `json:"Remote Login - SSH"`
	Samba_Sharing         ProcState `json:"Samba Sharing"`
}

type Exception struct {
	Bundleid string `json:"bundleid"`
	Path     string `json:"path"`
	State    int    `json:"state"`
}

type ID struct {
	ID string `json:"id"`
}

type FirewallBeautified struct {
	ActiveState string `json:"active_state"`
	Rules       []Rule `json:"fw_rules"`
}

type Rule struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
}
