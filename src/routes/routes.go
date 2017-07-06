package routes

import (
	"fmt"

	"../config"
)

// Operations
const GiveMeAServer string = "givemeaserver"

type RouterRequest struct {
	operation string
	path      string
	method    string
	c         *chan RouterResponse
}

type RouterResponse struct {
	zarasa string
}

func Router(c chan RouterRequest, reglas []config.Regla) {
	// for {
	// 	msg := <-c
	//
	// 	switch msg.operation {
	// 	case GiveMeAServer:
	// 		giveAServer(msg)
	// 	}
	// }

	fmt.Println(initColas(reglas))
}

// Private

type serversLists struct {
	servers            chan string
	unavailableServers chan string
}

func initColas(reglas []config.Regla) map[string]serversLists {
	serversQueues := make(map[string]serversLists)

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

func giveAServer(msg RouterRequest) {

}
