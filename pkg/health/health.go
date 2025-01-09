package health

// Entity
// @Description Objeto de resposta para validação de saúde da aplicação
type Entity struct {
	Hostname string `json:"hostname,omitempty" example:"server-linux"`
	IP       string `json:"ip,omitempty" example:"192.168.0.1"`
} // @name Entity
