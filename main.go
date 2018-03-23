//	Project: hps
//	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
//	Date: 20-March-2018 2:24 PM

/*

	Package -> main

	Description: Main Execution container for HPS. Starts listners and forwards incoming TCP/UDP messages to SIP processing packages in HPS

	Functions

	The following is a list of functions in this package:
		start : listener to start a listener based on parameters, called through goroutine to enable threading
			protocol (string): Specifies the type of the listener to start. Accepted values are "tcp" and "udp"
*/
package main

import (
	"hps/app/config"
	"hps/app/gonfig"
	"hps/app/sip"
	"log"
	"net"
)

//Variable that stores the configuration
var Cnf gonfig.Gonfig

func start(protocol string) {
	log.Println("Starting " + protocol + " listener")
	if protocol == "udp" {
		udpPort, _ := Cnf.GetString("server/udpPort", "5060")
		address, err := net.ResolveUDPAddr("udp", ":"+udpPort)
		if err != nil {
			log.Fatal("ERROR ", err)
		} else {
			ln, err := net.ListenUDP("udp", address)
			if err != nil {
				log.Fatal("ERROR ", err)
			} else {
				log.Println("Listening on udp:" + udpPort)
				buf := make([]byte, 1024)
				for {
					n, addr, err := ln.ReadFromUDP(buf)
					sip.Message(addr, string(buf[0:n])) //send the message to sip package
					if err != nil {
						log.Println("ERROR ", err)
					}
				}
			}
		}
	} else {
		tcpPort, _ := Cnf.GetString("server/tcpPort", "5060")
		address, err := net.ResolveTCPAddr("tcp", ":"+tcpPort)
		if err != nil {
			log.Println("ERROR ", err)
		} else {
			ln, err := net.ListenTCP("tcp", address)
			if err != nil {
				log.Fatal("ERROR ", err, ln)
			} else {
				log.Output(0, "Listening on "+protocol+":"+tcpPort)
				buf := make([]byte, 1024)
				for {
					conn, connErr := ln.Accept()
					if connErr != nil {
						// handle error
						log.Println("ERROR ", connErr)
					} else {
						n, errRead := conn.Read(buf)
						if errRead != nil {
							log.Fatal("ERROR", errRead)
						} else {
							sip.Message(conn.RemoteAddr(), string(buf[0:n])) //send the package to sip package
						}
					}

				}
			}
		}
	}
}
func main() {
	log.Println("<<<==========HPS==========>>>")
	Cnf = config.Read()
	go start("tcp")
	go start("udp")
	i := 1
	for {
		if i < 1000 {

		}
	}
}
