package applications

type Application struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}
