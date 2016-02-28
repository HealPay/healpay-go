package healpay

import (
	"io"
	"crypto/sha1"
	"strconv"
	"time"
	"encoding/hex"
	"encoding/xml"
)

type Token struct {
	XMLName xml.Name `xml:"Token"`
	ClientIP string
	SourceKey string
	HashValue string `xml:"PinHash>HashValue"`
	Seed string `xml:"PinHash>Seed"`
	Type string `xml:"PinHash>Type"`
	Pin string `xml:"-"`
	unixNano string
}

func (t *Token) UnixNano() string {
	if t.unixNano == "" {
		t.unixNano = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return t.unixNano
}

func (t *Token) GetSeed() string {
	return t.UnixNano()
}

func (t *Token) GetHashValue(sourceKey, pin, seed string) string {
	h := sha1.New()
	io.WriteString(h, sourceKey)
	io.WriteString(h, seed)
	io.WriteString(h, pin)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Token) XMLString() string {
	m, _ := xml.MarshalIndent(t, "", "   ")
	return string(m)
}

func NewToken(sourceKey, pin string) *Token {
	token := &Token{
		SourceKey: sourceKey, 
		Pin: pin,
		ClientIP: "192.168.0.1", 
		Type: "sha1",
	}
	token.Seed = token.GetSeed()
	token.HashValue = token.GetHashValue(sourceKey, pin, token.Seed)
	return token
}