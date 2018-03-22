//	Project: hps
//	Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
//	Date: 22-March-2018 4:07 PM
/*

	Package -> sip

	Description: Main code which receives all the SIP messages, and then sends them to appropriate handler codes.
*/
package sip

import (
	"log"
	"net"
)
//	Method that receives the SIP messages from the main code
func Message(addr net.Addr ,message string){
	log.Println(addr,message)
}
