/*
Project: hps
Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
Date: 23-March-2018 10:52 AM
*/
package sip

import "net"

type SIPMessage struct {
	address net.Addr
	body string
}

//Type Definition for Contact
type Contact struct {
	contacts []string
}

//Type definition for the REGISTER request
type Register struct {
	requestURI string
	to string
	from string
	callId string
	cSeq string
	contact Contact
	expires int
}