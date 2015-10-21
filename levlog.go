package levlog

import (
	"log"
)

type debugLevel struct {
	name		string
	id			int
}

var (
    DEBUG  debugLevel = debugLevel{"DEBUG",1}
    ERROR debugLevel =  debugLevel{"ERROR",2}
    FATAL debugLevel =  debugLevel{"FATAL",3}
)

var writer RotateWriter

func SetOutput(filename string){
	RotateWriter := NewRotateWrite(filename)
	log.SetOutput(RotateWriter)
}

var DEBUG_LEVEL debugLevel = DEBUG 

func D(v ...interface{}){
	if DEBUG_LEVEL.id <= DEBUG.id{
		log.Printf("(DEBUG): %s\n",v)
	}
}

func E(v ...interface{}){
	if DEBUG_LEVEL.id <= ERROR.id{
		log.Printf("(ERROR): %s\n",v)
	}
}

func F(v ...interface{}){
	if DEBUG_LEVEL.id <= FATAL.id{
		log.Printf("(FATAL): %s\n",v)
	}
}

func DF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= DEBUG.id{
		log.Printf("(DEBUG): " + format+"\n",v)
	}
}

func EF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= ERROR.id{
		log.Printf("(ERROR): " + format+"\n",v)
	}
}

func FF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= FATAL.id{
		log.Fatalf("(FATAL): " + format+"\n",v)
	}
}