package routes

import "../config"

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
	servers := initColas(reglas)

	for {
		msg := <-c

		switch msg.Operation {
		case GiveMeAServer:
			giveAServer(msg, servers)
		}
	}
}

// Private

type serversLists struct {
	servers            chan string
	unavailableServers chan string
}

type serversQueuesMap map[string]serversLists

func initColas(reglas []config.Regla) serversQueuesMap {
	serversQueues := make(serversQueuesMap)

	for _, rule := range reglas {
		cSer := make(chan string, 100)
		cUnSer := make(chan string, 100)

		for _, server := range rule.Servers {
			cSer <- server
		}

		serversQueues[rule.Ruta] = serversLists{servers: cSer, unavailableServers: cUnSer}
	}

	return serversQueues
}

func giveAServer(msg RouterRequest, servers serversQueuesMap) {
	// Agregar de nuevo el server a la lista. Considerar el caso de que el channel
	// este vacio (poner un timer)
	rsp := RouterResponse{RouteRequest: true, Server: <-servers[msg.Path].servers}
	*msg.C <- rsp
}
