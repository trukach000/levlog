package levlog

import (
	"testing"
)

func TestLevlog(t *testing.T) {
	D("Test debug");
	E("Test error");
	F("Test fatal");
}
func TestLevlogFormat(t *testing.T) {
	D("Format = %s","Test debug");
	E("Format = %s","Test error");
	F("Format = %s","Test fatal");
}