package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// OKMESSAGE used to send ok between hosts
var OKMESSAGE = [...]byte{1, 1, 1, 1, 1, 1, 1, 1}

// MSDELAY is Milliseconds  used to raise error when no message is recieved
const MSDELAY = 100

type host struct {
	ipPort   string
	protocol string
}

// SystemCommunication allows to send messages between two hosts
type SystemCommunication interface {
	// Send a message to the host and return true
	// if ok message has been recieved
	Send2Way(host *host, message []byte) bool

	// Send a message to one host
	Send(host *host, message []byte)

	//Waits for an upcoming message
	RecieveAndRespond() []byte

	Recieve() []byte

	//Returns the port of host
	Port() int
}

// Send a message to one host
func (h *host) Send(host *host, message []byte) {
	conn, err := net.Dial("udp", host.ipPort)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write(message)
}

func (h *host) Send2Way(sv *host, message []byte) bool {

	//Opens connection
	conn, err := net.Dial("udp", sv.ipPort)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//Send message
	conn.Write(message)

	response := make([]byte, 8)
	// if 1 sec has passed and no response has been recieved, raise error
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	//Wait for response
	_, err = conn.Read(response)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//TODO: implement ok message
	return true
}

// Receive waits for an upcomming connection
// and returns its message
func (h *host) Recieve() []byte {
	//Create local address where the server is going to listen
	port, _ := strconv.Atoi(strings.Split(h.ipPort, ":")[1])
	localaddr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(strings.Split(h.ipPort, ":")[0]),
	}

	//Listen for upcoming connections
	udpconn, err := net.ListenUDP("udp", &localaddr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer udpconn.Close()

	buffer := make([]byte, 160)

	//Read from udp connection
	_, _, err = udpconn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Printf(err.Error())
	}

	//clientHost := host{ipPort: remoteaddr.String(), protocol: "udp"}

	//go h.Send(&clientHost, []byte("Exito"))
	return buffer
}

func (h *host) RecieveAndRespond() []byte {
	//Create local address where the server is going to listen
	port, _ := strconv.Atoi(strings.Split(h.ipPort, ":")[1])
	localaddr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(strings.Split(h.ipPort, ":")[0]),
	}

	//Listen for upcoming connections
	udpconn, err := net.ListenUDP("udp", &localaddr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer udpconn.Close()

	buffer := make([]byte, 160)

	//Read from udp connection
	_, udpaddr, err := udpconn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Printf(err.Error())
	}
	udpconn.WriteToUDP([]byte("Exito"), udpaddr)
	return buffer
}
