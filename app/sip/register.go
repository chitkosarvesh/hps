/*
Project: hps
Author: Sarvesh Chitko (chitkosarvesh@gmail.com) (sarvesh)
Date: 23-March-2018 4:00 PM
*/
package sip

import (
	"strconv"
)

func RegisterRequest(m []string) Register {
	expires, err := strconv.ParseInt(m[24], 10, 64)
	if err != nil {
		expires = 3600
	}
	r := Register{
		requestURI: m[1],
		to:         m[12] + " " + m[13],
		from:       m[9] + " " + m[10],
		callId:     m[4],
		cSeq:       m[6],
		contact:    []string{m[27]},
		expires:    expires,
	}
	return r
}
