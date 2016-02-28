package healpay

import (
	"encoding/xml"
//	"time"
	"testing"
	"log"
)

func TestNewEnvelope(t *testing.T) {
	e := NewEnvelope("test")
	m, _ := xml.MarshalIndent(e, "", "   ")
	log.Println(string(m))
}
