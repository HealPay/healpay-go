package healpay

import (
	"encoding/xml"
	"encoding/base64"
	"io/ioutil"
)

type CreateBatchUploadRequest struct {
	XMLName xml.Name `xml:"ns1:createBatchUpload"`
	FileName string
	AutoStart bool
	Format string
	Encoding string
	Fields []string `xml:"Fields>item"`
	Data string
	DataPath string `xml:"-"`
	OverrideDuplicates bool
	HealpayRequest
}

func (req *CreateBatchUploadRequest) PostProcess() error {
	data, err := ioutil.ReadFile(req.DataPath)
	req.Data = base64.StdEncoding.EncodeToString(data)
	req.Encoding = "base64"
	return err
}

type CreateBatchUploadResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string
	HealpayResponse
}

func (res *CreateBatchUploadResponse) GetBody() string {
	return res.Body
}

func (res *CreateBatchUploadResponse) Decode(respBody []byte) ([]byte, error) {
	return respBody, nil
}
