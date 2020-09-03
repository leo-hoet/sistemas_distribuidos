package main

import (
	"fmt"
)

func send(sv *host, msg []byte) {
	client := host{ipPort: "localhost:3001", protocol: "udp"}
	client.Send(sv, msg)
}

func testSend() {
	sv := host{ipPort: "localhost:3100", protocol: "udp"}
	msg := []byte("hola mundo!")
	go send(&sv, msg)

	fmt.Printf("Envio %s de testSend\n", msg)

	buff := sv.Recieve()

	fmt.Printf("Recibio %s de testSend\n", buff)
}

func testRecieveAndResponse(sv *host) {
	msg := sv.RecieveAndRespond()
	fmt.Printf("El sv recibio: %s\n", msg)
}

func testSend2Way() {
	sv := host{ipPort: "localhost:3101", protocol: "udp"}
	go testRecieveAndResponse(&sv)

	client := host{ipPort: "localhost:3102", protocol: "udp"}
	res := client.Send2Way(&sv, []byte("Hola mundo!"))
	fmt.Printf("Resultado de testSend2Way %t\n", res)

}

func main(){
	testSend()
	testSend2Way()
}
