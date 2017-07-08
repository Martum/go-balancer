package routes

func ServersHandler(c chan RouterRequest, servers []string) {
	availableServers := make(chan string, len(servers))
	unavailableServers := make(chan string, len(servers))

	for {
		msg := <-c

		switch msg.Operation {
		case GiveMeAServer:
			giveServerToClient(msg, availableServers)
		case ServerDown:
			removeServer(msg.Meta, availableServers, unavailableServers)
		case ServerUp:
			addServer(msg.Meta, availableServers, unavailableServers)
		}
	}
}

// Le mandamos un server al cliente y lo volvemos a encolar
// Si no hay servers para atender el pedido, indicamos que no se rutee el request
func giveServerToClient(msg RouterRequest, servers chan string) {
	select {
	case server := <-servers:
		*msg.C <- RouterResponse{RouteRequest: true, Server: server}
		servers <- server
	default:
		*msg.C <- RouterResponse{RouteRequest: false, Server: ""}
	}
}

func removeServer(serverDown string, availableServers chan string, unavailableServers chan string) {
	swapServer(serverDown, availableServers, unavailableServers)
}

func addServer(serverUp string, availableServers chan string, unavailableServers chan string) {
	swapServer(serverUp, unavailableServers, availableServers)
}

// Sacamos un server de from y lo pasamos a to, los demas servers los volvemos
// a agregar a from
func swapServer(swappingServer string, from chan string, to chan string) {
	var servers []string
	i := 0

	for {
		select {
		case server := <-from:
			if server == swappingServer {
				to <- server
				listToChan(servers, i, from)

				return
			} else {
				servers[i] = server
				i = i + 1
			}
		default:
			listToChan(servers, i, from)
			return
		}
	}
}

func listToChan(list []string, i int, c chan string) {
	for n := 0; n < i; n++ {
		c <- list[n]
	}
}
