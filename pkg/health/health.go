package health

// Entity
// @Description Objeto de resposta para validação de saúde da aplicação
type Entity struct {
	Hostname string `json:"hostname,omitempty"`
	IP       string `json:"ip,omitempty"`
} // @name Entity
