package healpay

import (
	"testing"
)

func TestXML(t *testing.T) {
	token := NewToken("somekey", "somepin")
	t.Logf(token.XMLString())
}