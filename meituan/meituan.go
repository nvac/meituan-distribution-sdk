package meituan

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"time"
)

type Options struct {
	BaseUrl   string
	Test      bool
	PartnerId int
	AccessKey string
	SecretKey string

	baseUrlV1 string
	baseUrlV2 string
}

type Client struct {
	options *Options
	client  *resty.Client
}

func NewMeituanClient(options *Options) Client {
	options.baseUrlV1 = options.BaseUrl
	options.baseUrlV2 = fmt.Sprintf("%s/v2", options.BaseUrl)
	client := resty.New()

	client.SetBaseURL(options.BaseUrl)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json; charset=utf-8")

	return Client{
		options: options,
		client:  client,
	}
}

type envelope struct {
	CustomerSessionId string
	Method            string
	Data              interface{}
}

type meituanRequestV1 struct {
	Accesskey string `json:"accesskey"`
	Data      string `json:"data"`
	Method    string `json:"method"`
	Nonce     int    `json:"nonce"`
	PartnerId int    `json:"partnerId"`
	Test      string `json:"test"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
	Signature string `json:"signature"`
}

type meituanRequestV2 struct {
	Accesskey         string `json:"accesskey"`
	CustomerSessionId string `json:"customerSessionId"`
	Data              string `json:"data"`
	Language          string `json:"language"`
	Method            string `json:"method"`
	Nonce             int32  `json:"nonce"`
	PartnerId         int    `json:"partnerId"`
	Test              string `json:"test,omitempty"`
	Timestamp         int64  `json:"timestamp"`
	Version           string `json:"version"`
	Signature         string `json:"signature"`
}

type meituanResponse struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	PartnerId int    `json:"partnerId,omitempty"`
	Result    any    `json:"result,omitempty"`
}

func (c Client) callV1(e envelope, resp any) (int, string, int, error) {
	marshal, err := json.Marshal(e.Data)
	if err != nil {
		return 0, "", 0, err
	}

	request := meituanRequestV1{
		Accesskey: c.options.AccessKey,
		Data:      string(marshal),
		Method:    e.Method,
		Nonce:     rand.Int(),
		PartnerId: c.options.PartnerId,
		Timestamp: time.Now().Unix(),
		Version:   "1.0",
	}

	plainText := fmt.Sprintf("accesskey=%s&data=%s&method=%s&nonce=%d&partnerId=%d&timestamp=%d&version=%s",
		request.Accesskey, request.Data, request.Method, request.Nonce, request.PartnerId, request.Timestamp, request.Version)

	if c.options.Test {
		request.Test = "test"
		plainText = fmt.Sprintf("accesskey=%s&data=%s&method=%s&nonce=%d&partnerId=%d&test=%s&timestamp=%d&version=%s",
			request.Accesskey, request.Data, request.Method, request.Nonce, request.PartnerId, request.Test,
			request.Timestamp, request.Version)
	} else {
		plainText = fmt.Sprintf("accesskey=%s&data==%s&method=%s&nonce=%d&partnerId=%d&timestamp=%d&version=%s",
			request.Accesskey, request.Data, request.Method, request.Nonce, request.PartnerId,
			request.Timestamp, request.Version)
	}

	mac := hmac.New(sha1.New, []byte(c.options.SecretKey))
	mac.Write([]byte(plainText))

	request.Signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	response := &meituanResponse{
		Result: resp,
	}

	_, err = c.client.R().
		SetBody(request).
		SetResult(response).
		Post(c.options.baseUrlV1)

	return response.Code, response.Message, response.PartnerId, err
}

func (c Client) callV2(e envelope, resp any) (int, string, int, error) {
	marshal, err := json.Marshal(e.Data)
	if err != nil {
		return 0, "", 0, err
	}

	request := meituanRequestV2{
		Accesskey:         c.options.AccessKey,
		CustomerSessionId: e.CustomerSessionId,
		Data:              string(marshal),
		Language:          "zh_CN",
		Method:            e.Method,
		Nonce:             rand.Int31(),
		PartnerId:         c.options.PartnerId,
		Timestamp:         time.Now().Unix(),
		Version:           "2.0",
	}

	var plainText string
	if c.options.Test {
		request.Test = "test"
		plainText = fmt.Sprintf("accesskey=%s&customerSessionId=%s&data=%s&language=%s&method=%s&nonce=%d&partnerId=%d&test=%s&timestamp=%d&version=%s",
			request.Accesskey, request.CustomerSessionId, request.Data, request.Language,
			request.Method, request.Nonce, request.PartnerId, request.Test,
			request.Timestamp, request.Version)
	} else {
		plainText = fmt.Sprintf("accesskey=%s&customerSessionId=%s&data=%s&language=%s&method=%s&nonce=%d&partnerId=%d&timestamp=%d&version=%s",
			request.Accesskey, request.CustomerSessionId, request.Data, request.Language,
			request.Method, request.Nonce, request.PartnerId,
			request.Timestamp, request.Version)
	}

	mac := hmac.New(sha1.New, []byte(c.options.SecretKey))
	mac.Write([]byte(plainText))

	request.Signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	response := &meituanResponse{
		Result: resp,
	}

	_, err = c.client.R().
		SetBody(request).
		SetResult(response).
		Post(c.options.baseUrlV2)

	return response.Code, response.Message, response.PartnerId, err
}
