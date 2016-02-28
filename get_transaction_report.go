package healpay

import (
	"encoding/xml"
)

type GetTransactionReportResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>getTransactionReportResponse>getTransactionReportReturn"`
	HealpayResponse
}

func (res *GetTransactionReportResponse) GetBody() string {
	return res.Body
}

func (res *GetTransactionReportResponse) Decode(respBody []byte) ([]byte, error) {
	err := xml.Unmarshal(respBody, res)
	if err != nil { return nil, err }
	return Base64Decode(res.Body)
}
