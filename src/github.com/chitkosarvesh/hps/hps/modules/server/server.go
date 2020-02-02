package server

import (
	"github.com/chitkosarvesh/hps/hps/modules/config"
	"fmt"
	"net"
	"os"
)
func Start(){
	address := config.Get("server.address")
	port := config.Get("server.port")
	var tcp_l net.Listener
	var udp_l net.Listener
	if(config.GetBool("server.tcp")){
		tcp_l,err := net.Listen("tcp",address+":"+port)
		if err != nil {
			fmt.Println("Error",err)
		}
		defer tcp_l.Close()
		fmt.Println("Listening on tcp://" + address + ":" + port)
	}
	if(config.GetBool("server.udp")){
		add,err :=net.ResolveUDPAddr("udp",address+":"+port)
		if err != nil {
			fmt.Println("Error",err)
		}
		udp_l,err := net.ListenUDP("udp",add)
		if err != nil {
			fmt.Println("Error",err)
		}
		defer udp_l.Close()
		fmt.Println("Listening on udp://" + address + ":" + port)

	}
	for {
		conn_tcp,err:=tcp_l.Accept()
		if err != nil {
			fmt.Println("Error in TCP Request",err)
			os.Exit(1)
		}
		go handleRequest(conn_tcp)
		conn_udp,err:=udp_l.Accept()
		if err != nil {
			fmt.Println("Error in UDP Request",err)
			os.Exit(1)
		}
		go handleRequest(conn_udp)
	}
}
// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
  }