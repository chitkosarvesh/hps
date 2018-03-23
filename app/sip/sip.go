//	Project: hps
//	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
//	Date: 22-March-2018 4:07 PM
/*

	Package -> sip

	Description: Main code which receives all the SIP messages, and then sends them to appropriate handler codes.
*/
package sip

import "net"

//Checks whether a request is a valid SIP request, checks what request is received and calls the handler function for those Requests
func findRequestType(m SIPMessage)	{

}

//	Method that receives the SIP messages from the main code
func Message(addr net.Addr ,body string){
	var m SIPMessage
	m.address=addr
	m.body=body
	findRequestType(m)
}
