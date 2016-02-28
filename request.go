package healpay

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"bytes"
)

type Body struct {
	XMLName xml.Name `xml:"SOAP-ENV:Body"`
	Content string `xml:",innerxml"`
}

type Envelope struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	SoapEnv string `xml:"xmlns:SOAP-ENV,attr"`
	Ns1 string `xml:"xmlns:ns1,attr"`
	Xsd string `xml:"xmlns:xsd,attr"`
	Xsi string `xml:"xmlns:xsi,attr"`
	SoapEnc string `xml:"xmlns:SOAP-ENC,attr"`
	EncodingStyle string `xml:"SOAP-ENV:encodingStyle,attr"`
	Body Body
}

type Request interface {
	SetToken(*Token)
	PostProcess() error
}

type HealpayRequest struct {
	Token *Token
	Request
}

func (r *HealpayRequest) SetToken(t *Token) {
	r.Token = t
}

func NewEnvelope(body string) *Envelope {
	return &Envelope{
		Xsd: "http://www.w3.org/2001/XMLSchema",
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ns1: "urn:healpay",
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		SoapEnc: "http://schemas.xmlsoap.org/soap/encoding/",
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: Body{Content: body},
	}
}

func NewRequest(location, body string) (*http.Request, error) {
	envelope := NewEnvelope(body)
	m, err := xml.MarshalIndent(envelope, "", "   ")
	if err != nil { return nil, err }
	buffer := bytes.NewBufferString(string(m))
	return http.NewRequest("POST", location, buffer)
}

func JSONToXML(report Request, data []byte) ([]byte, error) {
	err := json.Unmarshal(data, report)
	if err != nil { return nil, err }
	return xml.MarshalIndent(report, "", "   ")
}
