package levlog

import (
	"testing"
)

func TestLevlog(t *testing.T) {
	D("Test debug")
	E("Test error")
	F("Test fatal")
}
func TestLevlogFormat(t *testing.T) {
	DF("Format = %s","Test debug")
	EF("Format = %s","Test error")
	FF("Format = %s","Test fatal")
	x := 10
	DF("Format = %d/%d",2,x)
}