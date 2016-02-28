package healpay

import (
	"encoding/xml"
)

type RunTransactionRequest struct {
	XMLName xml.Name `xml:"ns1:runTransaction"`
	Parameters TransactionRequestObject
	HealpayRequest
}

type RunTransactionResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string
	HealpayResponse
}

func (res *RunTransactionResponse) GetBody() string {
	return res.Body
}

func (res *RunTransactionResponse) Decode(respBody []byte) ([]byte, error) {
	return respBody, nil
}
