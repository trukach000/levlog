# levlog
simple level logger for go

Logging to console

```
import (
	log "github.com/trukach000/levlog"
)

...
D("Test debug")
...
x := 10
DF("Format = %d/%d",2,x)

```

To log to the file ouput with rotation use

```
shutdownChannel := make(chan bool) //make simple boolean shutdown channel

//then make log file correct name
logFileName := "." + string(filepath.Separator) + "logs" + string(filepath.Separator) + "main.log"
	
// use SetOutput function 
// second arg are used to set file rotating period
log.SetOutput(logFileName,time.Hour * 24, shutdownChannel)

//than use log function like D,E,F,DF,EF,FF as-is

```