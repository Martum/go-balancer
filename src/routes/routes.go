package routes

import (
	"regexp"
	"time"

	"../config"
)

// Operations
const GiveMeAServer string = "givemeaserver"
const ServerDown string = "serverdown"
const ServerUp string = "serverup"

type RouterRequest struct {
	Operation string
	Path      string // /foo/bar
	Method    string // GET, POST, PUT, PATCH, DELETE
	C         *chan RouterResponse
	Meta      string
}

type RouterResponse struct {
	RouteRequest bool
	Server       string
}

func Router(c chan RouterRequest, reglas []config.Regla, secondsToUp int) {
	servers := initHandlers(reglas)

	for {
		msg := <-c

		switch msg.Operation {
		case GiveMeAServer:
			giveAServer(msg, servers)
		case ServerDown:
			notifyServers(msg, servers)
			go notifyServerUp(secondsToUp, c, msg.Meta)
		case ServerUp:
			notifyServers(msg, servers)
		}
	}
}

// Private

type serversHandler struct {
	pathRegEx *regexp.Regexp
	channel   chan RouterRequest
}

func initHandlers(reglas []config.Regla) []serversHandler {
	handlersList := make([]serversHandler, len(reglas))

	var i = 0
	var defaultHandler serversHandler
	for _, rule := range reglas {
		if rule.Ruta == "*" {
			r, _ := regexp.Compile(".*")
			defaultHandler = serversHandler{pathRegEx: r, channel: make(chan RouterRequest, 1000)}
		} else {
			r, _ := regexp.Compile("^" + rule.Ruta + ".*")
			handlersList[i] = serversHandler{pathRegEx: r, channel: make(chan RouterRequest, 1000)}

			i = i + 1
		}
	}
	handlersList[i] = defaultHandler

	return handlersList
}

func giveAServer(msg RouterRequest, servers []serversHandler) {
	// Se despacha el pedido al primer serverHandler que responda al path solicitado
	for _, server := range servers {
		if server.pathRegEx.MatchString(msg.Path) {
			server.channel <- msg
			break
		}
	}
}

// Envia un mensaje a todos los serversHandlers
func notifyServers(msg RouterRequest, servers []serversHandler) {
	for _, server := range servers {
		server.channel <- msg
	}
}

// Me envia un mensaje para restaurar un server que habia sido dado de baja
func notifyServerUp(seconds int, c chan RouterRequest, server string) {
	time.Sleep(100 * time.Second)
	c <- RouterRequest{Operation: ServerUp, Meta: server}
}
