/*
	Project: hps
	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
	Date: 23-March-2018 10:52 AM

	Package -> sip
	Description: Part of the SIP package which stores definitions of all the requests
*/
package sip

import "net"

// Definition of a normal aip incoming message
type SIPMessage struct {
	address net.Addr
	body    string
}

//Type definition for the REGISTER request
type Register struct {
	requestURI string
	to         string
	from       string
	callId     string
	cSeq       string
	contact    []string
	expires    int64
}
