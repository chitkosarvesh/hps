//	Project: hps
//	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
//	Date: 22-March-2018 4:07 PM
/*

	Package -> sip

	Description: Main code which receives all the SIP messages, and then sends them to appropriate handler codes.
*/
package sip

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Breaks down the request into manageable parts
func parseRequest(m SIPMessage) []string {
	return strings.Fields(m.body)
}

//Checks whether a request is a valid SIP request, checks what request is received and calls the handler function for those Requests
func findRequestType(m SIPMessage) {
	req := parseRequest(m)
	log.Print(m)
	if req[0] == "REGISTER" {
		log.Println("Register Request received")
		rReq := RegisterRequest(req)
		log.Println(m.address)
		conn, err := net.Dial("udp", m.address.String())
		if err != nil {
			log.Fatal("ERROR ", err)
		} else {
			log.Println(m.address.String())
			response := "SIP/2.0 200 OK\n" +
				"From: " + rReq.from + "\n" +
				"To: " + rReq.to + "\n" +
				"Call-ID:" + rReq.callId + "\n" +
				"CSeq: " + rReq.cSeq + "\n" +
				"Contact: " + rReq.contact[0] + "\n" +
				"Content-Length: 0" + "\n"
			log.Print(response)
			fmt.Fprintf(conn, response)
			status, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal("ERROR ", err)
			} else {
				log.Println(status)
			}
		}
	}

}

//	Method that receives the SIP messages from the main code
func Message(addr net.Addr, body string) {
	var m SIPMessage
	m.address = addr
	m.body = body
	findRequestType(m)
}
