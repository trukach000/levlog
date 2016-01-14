package levlog

import (
	"testing"
	"os"
	"bytes"
	"log"
	"regexp"
	. "github.com/smartystreets/goconvey/convey"
)

func ShouldBeRegexMatched(actual interface{}, expected ...interface{}) string {
    if match, err := regexp.MatchString(expected[0].(string), actual.(string)); err==nil && match {
        return ""   // empty string means the assertion passed
    } else {
        return "No matched to regexp '"+expected[0].(string)+
        	"' in: '"+ actual.(string)+"'"
    }
}

func TestLevlog(t *testing.T) {
    Convey("Keep backup of the real stdout and make fake stdout for test", t, func() {
        var buf bytes.Buffer
    	log.SetOutput(&buf)

        Convey("Print simple logging text into fake stdout", func() {
        	var out string	
            D("Test debug")
            out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(DEBUG\):\s.Test\sdebug`)
            buf.Reset()
			E("Test error")
			out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(ERROR\):\s.Test\serror`)
            buf.Reset()
			F("Test fatal")
			out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(FATAL\):\s.Test\sfatal`)
			
            Convey("Back to normal state", func() {
			    log.SetOutput(os.Stderr)
            })
        })
    })
}
func TestLevlogFormat(t *testing.T) {
	Convey("Keep backup of the real stdout and make fake stdout for test", t, func() {
        var buf bytes.Buffer
    	log.SetOutput(&buf)

        Convey("Print simple logging text into fake stdout", func() {
        	var out string	
            DF("Format = %s","Test debug")
            out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(DEBUG\):\sFormat\s=\sTest\sdebug`)
            buf.Reset()
			EF("Format = %s","Test error")
			out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(ERROR\):\sFormat\s=\sTest\serror`)
            buf.Reset()
			FF("Format = %s","Test fatal")
			out = buf.String()
            So(out, ShouldBeRegexMatched, 
                	`\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\(FATAL\):\sFormat\s=\sTest\sfatal`)
			
            Convey("Back to normal state", func() {
			    log.SetOutput(os.Stderr)
            })
        })
    })
}