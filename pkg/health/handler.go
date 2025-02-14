package health

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Health(c *gin.Context)
}

type handler struct{}

func New() Handler {
	return &handler{}
}

// Health
// @Summary Método responsável por gerar um endpoint de helth (saúde da aplicação)
// @Description Método responsável por gerar um endpoint de helth (saúde da aplicação)
// @Tags health
// @Accept json
// @Produce json
// @Router /health [get]
func (a *handler) Health(c *gin.Context) {
	response := Entity{}
	response.IP = a.getLocalIP()
	response.Hostname, _ = os.Hostname()

	c.IndentedJSON(http.StatusOK, response)
}

func (a *handler) getLocalIP() string {
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			return ip.String()
		}
	}

	return ""
}
