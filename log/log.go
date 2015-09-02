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
