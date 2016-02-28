package healpay

import (
	"encoding/xml"
)

type SearchTransactionsCustomResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>searchTransactionsCustomResponse>searchTransactionsCustomReturn"`
	HealpayResponse
}
func (res *SearchTransactionsCustomResponse) GetBody() string {
	return res.Body
}

func (res *SearchTransactionsCustomResponse) Decode(respBody []byte) ([]byte, error) {
	err := xml.Unmarshal(respBody, res)
	if err != nil { return nil, err }
	return Base64Decode(res.Body)
}
