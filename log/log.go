/*

Package log

Status: Stub

TODO:

* create drivers for centralized logging engines
* create drivers for local file writing
* establish standards that describe:
	clientId
	coralApp
	what's happening
	what happened
	any data
	metadata

	In a way that can be easily queried and reconstructed forensically

*/

package log

import (
	"log"
)

func Write(m ...interface{}) {

	log.Println(m)

}

func Fatal(m ...interface{}) {

	log.Fatal(m)

}
