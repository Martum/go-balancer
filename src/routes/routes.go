package routes

import "../config"

// Operations
const GiveMeAServer string = "givemeaserver"

type RouterRequest struct {
	operation string
	path      string // /foo/bar
	method    string // GET, POST, PUT, PATCH, DELETE
	c         *chan RouterResponse
}

type RouterResponse struct {
	routeRequest bool
	server       string
}

func Router(c chan RouterRequest, reglas []config.Regla) {
	servers := initColas(reglas)

	for {
		msg := <-c

		switch msg.operation {
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
	rsp := RouterResponse{routeRequest: true, server: <-servers[msg.path].servers}
	*msg.c <- rsp
}
