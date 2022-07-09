package cpu

type CPU struct {
	Name    string `json:"name"`
	Threads int    `json:"threads"`
	Cores   int    `json:"cores"`
}
