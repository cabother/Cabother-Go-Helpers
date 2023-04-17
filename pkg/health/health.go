package health

type Entity struct {
	Hostname string `json:"hostname,omitempty"`
	IP       string `json:"ip,omitempty"`
}
