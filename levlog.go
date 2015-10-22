package levlog

import (
	"log"
	"time"
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

var writer *RotateWriter

func SetOutput(filename string){
	writer = NewRotateWrite(filename)
	log.SetOutput(writer)
	TimeRotating(time.After(time.Minute * 1))
}

var DEBUG_LEVEL debugLevel = DEBUG 

func rotate()(error){
	if writer != nil{
		err := writer.Rotate()
		return err
	}
	return nil
}

func TimeRotating(t <-chan time.Time){
	go func (t <-chan time.Time){
		for{
			select{
				case <-t:
					rotate()
			}
		}
	}(t)
}

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