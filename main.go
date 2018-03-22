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
	"log"
	"net"
)

func start(protocol string)	{
	log.Println("Starting "+protocol+" listener")
	if protocol == "udp" {
		address,err := net.ResolveUDPAddr("udp",":5060")
		if err != nil {
			log.Fatal("ERROR ",err)
		} else {
			ln,err := net.ListenUDP("udp", address)
			if err != nil	{
				log.Fatal("ERROR ",err)
			} else {
				log.Println("Listening on udp:5060")
				buf := make([]byte, 1024)
				for {
					n,addr,err := ln.ReadFromUDP(buf)
					log.Println("Received ",string(buf[0:n]), " from ",addr)

					if err != nil {
						log.Println("ERROR ",err)
					}
				}
			}
		}
	} else {
		ln,err := net.Listen(protocol, ":5060")
		if err != nil {
			log.Fatal("ERROR ",err,ln)
		} else	{
			log.Output(0,"Listening on "+protocol+":5060")
			buf := make([]byte, 1024)
			for {
				conn, connErr := ln.Accept()
				if connErr != nil {
					// handle error
					log.Println("ERROR ",connErr)
				} else {
					n,errRead := conn.Read(buf)
					if errRead != nil {
						log.Fatal("ERROR", errRead)
					} else {
						log.Println(string(buf[0:n]))
					}
				}

			}
		}
	}
}
func main(){
	log.Println("<<<==========HPS==========>>>")
	go start("tcp")
	go start("udp")
	i := 1
	for {
		if i<1000 {

		}
	}
}
