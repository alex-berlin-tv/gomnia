package omnia

import (
	"crypto/md5"
	"encoding/json"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

// Header for an Omnia API call.
type omniaHeader struct {
	xRequestCid   string `json:"X-Request-CID"`
	xRequestToken string `json:"X-Request-Token"`
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
) Response {
	args_parts := ""
	if len(args) > 0 {
		args_parts = strings.Join(args, "/")
		args_parts = fmt.Sprintf("/%s", args_parts)
	}
	url := fmt.Sprintf(
		"https://api.nexx.cloud/v3.1/%s/%s/%s%s",
		o.DomainId, streamType, operation, args_parts,
	)
	header := newOmniaHeader(operation, o.DomainId, o.ApiSecret, o.SessionId)
	o.debugLog(method, url, header, parameters)

	return Response{}
}

// Logs parameters of API call.
func (o Omnia) debugLog(method string, url string, header omniaHeader, parameters *QueryParameters) {
	log.Debugf("METHOD:\t%s", method)
	log.Debugf("URL:\t\t%s", url)
	log.Debugf("HEADER:\t%+v", header)
	if parameters != nil {
		log.Debugf("URL PARAMS:\t%+v", parameters)
	} else {
		log.Debug("URL PARAMS:\t{}")
	}
}
