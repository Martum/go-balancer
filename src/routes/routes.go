package routes

import (
	"regexp"

	"../config"
)

// Operations
const GiveMeAServer string = "givemeaserver"

type RouterRequest struct {
	Operation string
	Path      string // /foo/bar
	Method    string // GET, POST, PUT, PATCH, DELETE
	C         *chan RouterResponse
}

type RouterResponse struct {
	RouteRequest bool
	Server       string
}

func Router(c chan RouterRequest, reglas []config.Regla) {
	servers := initHandlers(reglas)

	for {
		msg := <-c

		switch msg.Operation {
		case GiveMeAServer:
			giveAServer(msg, servers)
		}
	}
}

// Private

// type serversQueuesMap map[string]serversLists

// func initHandlers(reglas []config.Regla) serversQueuesMap {
// 	serversQueues := make(serversQueuesMap)
//
// 	for _, rule := range reglas {
// 		cSer := make(chan string, 100)
// 		cUnSer := make(chan string, 100)
//
// 		for _, server := range rule.Servers {
// 			cSer <- server
// 		}
//
// 		serversQueues[rule.Ruta] = serversLists{servers: cSer, unavailableServers: cUnSer}
// 	}
//
// 	return serversQueues
// }

type serversHandler struct {
	pathRegEx *regexp.Regexp
	channel   chan string
}

func initHandlers(reglas []config.Regla) []serversHandler {
	handlersList := make([]serversHandler, len(reglas))

	var i = 0
	var defaultHandler serversHandler
	for _, rule := range reglas {
		if rule.Ruta == "*" {
			r, _ := regexp.Compile(".*")
			defaultHandler = serversHandler{pathRegEx: r, channel: make(chan string, 100)}
		} else {
			r, _ := regexp.Compile("^" + rule.Ruta + ".*")
			handlersList[i] = serversHandler{pathRegEx: r, channel: make(chan string, 100)}

			i = i + 1
		}
	}
	handlersList[i] = defaultHandler

	return handlersList
}

func giveAServer(msg RouterRequest, servers []serversHandler) {
	// Agregar de nuevo el server a la lista. Considerar el caso de que el channel
	// este vacio (usar select para esto)
	// Handlear el caso de que la ruta no exista (cae en el default *)
	// rsp := RouterResponse{routeRequest: true, server: <-servers[msg.path].servers}
	// *msg.c <- rsp
	for _, server :=
}
