package main

import (
	"fmt"
)

type grupo = []host

//Coordinator defines methods to manage a group of hosts
type Coordinator interface {
	//Send a message to all hosts in a group
	sendToGroup(nodos grupo, msg []byte)
}

func escucharYResponerTodos(nodos grupo) {
	for _, nodo := range nodos {
		nodo.RecieveAndRespond()
	}
}

//Sends a message to the whole group
func (coord *host) sendToGroup(nodos grupo, msg []byte) {
	go escucharYResponerTodos(nodos)

	for i, nodo := range nodos {
		res := coord.Send2Way(&nodo, msg)
		fmt.Printf("Resultado del nodo %d %t\n", i, res)
	}
}

func main() {
	nodos := make(grupo, 2, 10)
	msg := []byte("un mensaje")
	coordinator := host{ipPort: "localhost:3100", protocol: "udp"}

	nodos[0] = host{ipPort: "localhost:3101", protocol: "udp"}
	nodos[1] = host{ipPort: "localhost:3102", protocol: "udp"}
	coordinator.sendToGroup(nodos, msg)
}
