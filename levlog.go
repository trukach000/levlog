package levlog

import (
	"log"
	"time"
	"os"
	"path/filepath"
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

func SetOutput(filename string, timeRotation time.Duration,shutdownChannel chan bool){
	var err error
	writer,err = NewRotateWrite(filename)
	if err != nil{
		panic(err)
	}
	log.SetOutput(writer)
	TimeRotating(timeRotation,shutdownChannel)
}

var DEBUG_LEVEL debugLevel = DEBUG 

func rotate()(error){
	if writer != nil{
		err := writer.Rotate()
		return err
	}
	return nil
}

func TimeRotating(dur time.Duration,shutdownChannel chan bool){
	go func (shutdown chan bool){
		for{
			select{
				case <-time.After(dur):
					err := rotate()
					if err != nil{
						panic(err)
					}
				case <-shutdown:
					return
			}
		}
	}(shutdownChannel)
}

func PanicLogInFile(v ...interface{}){
	panicFileName := "." + string(filepath.Separator) + "logs" + string(filepath.Separator) + "panic.log"
	f, err := os.OpenFile(panicFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    E("Error when open ",panicFileName,"for panic logging")
	    return
	}
	defer f.Close()
	
	log.SetOutput(f)
	log.Fatal(v)	
}

func D(v ...interface{}){
	nv := make([]interface{},0)
	nv = append(nv,"(DEBUG): ")
	nv = append(nv,v...)
	if DEBUG_LEVEL.id <= DEBUG.id{
		log.Println(nv...)
	}
}

func E(v ...interface{}){
	nv := make([]interface{},0)
	nv = append(nv,"(ERROR): ")
	nv = append(nv,v...)
	if DEBUG_LEVEL.id <= ERROR.id{
		log.Println(nv...)
	}
}

func F(v ...interface{}){
	nv := make([]interface{},0)
	nv = append(nv,"(FATAL): ")
	nv = append(nv,v...)
	if DEBUG_LEVEL.id <= FATAL.id{
		log.Println(nv...)
	}
}

func DF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= DEBUG.id{
		log.Printf("(DEBUG): " + format+"\n",v...)
	}
}

func EF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= ERROR.id{
		log.Printf("(ERROR): " + format+"\n",v...)
	}
}

func FF(format string,v ...interface{}){
	if DEBUG_LEVEL.id <= FATAL.id{
		log.Printf("(FATAL): " + format+"\n",v...)
	}
}
