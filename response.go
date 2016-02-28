package healpay

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"time"
)

type Response interface {
	Handle(req *http.Request, timeout int) (*http.Response, []byte, error)
	Decode([]byte) ([]byte, error)
	GetBody() string
	ToXML()
}

type HealpayResponse struct {
	Response
}

type RawResponse struct {
	XMLName xml.Name `xml:"Envelope"`
        Body string
        HealpayResponse
}

func (res *RawResponse) GetBody() string {
        return res.Body
}

func (res *RawResponse) Decode(respBody []byte) ([]byte, error) {
        return respBody, nil
}

func Base64Decode(body string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(body)
}

func (res *HealpayResponse) Handle(req *http.Request, timeout int) (*http.Response, []byte, error) {
	clientTimeout := time.Duration(timeout+1) * time.Second
	client := http.Client{
		Timeout: clientTimeout,
	}
	resp, err := client.Do(req)
	if err != nil { return nil, nil, err }
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return resp, b, err
	}
	if err != nil { return nil, nil, err }
	return resp, b, err
}
