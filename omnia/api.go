package omnia

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
)

const omniaHeaderXRequestCid = "X-Request-CID"
const omniaHeaderXRequestToken = "X-Request-Token"

// Header for an Omnia API call.
type omniaHeader struct {
	xRequestCid   string
	xRequestToken string
}

func newOmniaHeader(operation, domainId, apiSecret, sessionId string) omniaHeader {
	log.Debugf("HASH SRC:\tmd5(%s+%s+API_SECRET)", operation, domainId)
	signature := md5.Sum([]byte(fmt.Sprintf("%s%s%s", operation, domainId, apiSecret)))
	return omniaHeader{
		xRequestCid:   sessionId,
		xRequestToken: hex.EncodeToString(signature[:]),
	}
}

type Omnia struct {
	DomainId  string `json:"domain_id"`
	ApiSecret string `json:"api_secret"`
	SessionId string `json:"session_id"`
}

// Retuns a new Omnia instance.
func NewOmnia(domainId string, apiSecret string, sessionId string) Omnia {
	return Omnia{
		DomainId:  domainId,
		ApiSecret: apiSecret,
		SessionId: sessionId,
	}
}

// Reads an Omnia instance from a json file.
func OmniaFromFile(path string) Omnia {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	rsl := Omnia{}
	if err := json.Unmarshal([]byte(file), &rsl); err != nil {
		log.Fatal(err)
	}
	return rsl
}

// Generic call to the Omnia API. Won't work with the media management API.
func (o Omnia) Call(
	method string,
	streamType StreamType,
	operation string,
	args []string,
	parameters *QueryParameters,
) (*Response, error) {
	method = strings.ToUpper(method)
	args_parts := ""
	if len(args) > 0 {
		args_parts = strings.Join(args, "/")
		args_parts = fmt.Sprintf("/%s", args_parts)
	}
	reqUrl := fmt.Sprintf(
		"https://api.nexx.cloud/v3.1/%s/%s/%s%s",
		o.DomainId, streamType, operation, args_parts,
	)
	header := newOmniaHeader(operation, o.DomainId, o.ApiSecret, o.SessionId)
	paramQuery, err := query.Values(parameters)
	if err != nil {
		return nil, err
	}
	paramUrl := paramQuery.Encode()
	o.debugLog(method, reqUrl, header, paramUrl)

	req, err := http.NewRequest(method, reqUrl, strings.NewReader(paramUrl))
	if err != nil {
		return nil, err
	}
	req.Header.Add(omniaHeaderXRequestCid, header.xRequestCid)
	req.Header.Add(omniaHeaderXRequestToken, header.xRequestToken)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	res := &Response{}
	json.Unmarshal(body, res)
	log.WithFields(res.Metadata.toMap()).Debug("Response Metadata")
	log.WithFields(res.Paging.toMap()).Debug("Response Paging")
	log.Trace(res.Result)
	return res, nil
}

// Logs parameters of API call.
func (o Omnia) debugLog(method string, url string, header omniaHeader, parameters string) {
	log.Debugf("METHOD:\t%s", method)
	log.Debugf("URL:\t\t%s", url)
	log.Debugf("HEADER:\t%+v", header)
	if parameters != "" {
		log.Debugf("URL PARAMS:\t%+v", parameters)
	} else {
		log.Debug("URL PARAMS:\t{}")
	}
}
