package omnia

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

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

// Encapsulates the calls to the nexxOmnia API.
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

// Return a item of a given streamtype by it's id.
func (o Omnia) ById(streamType StreamType, id int, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byid", []string{strconv.Itoa(id)}, parameters)
}

// Return a item of a given streamtype by it's global id.
func (o Omnia) ByGlobalId(streamType StreamType, globalId int, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byglobalid", []string{strconv.Itoa(globalId)}, parameters)
}

// Return a item of a given streamtype by it's hash.
func (o Omnia) ByHash(streamType StreamType, hash string, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byhash", []string{hash}, parameters)
}

// Return a item of a given streamtype by it's reference number.
func (o Omnia) ByRefNr(streamType StreamType, reference string, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byrefnr", []string{reference}, parameters)
}

// Return a item of a given streamtype by it's slug.
func (o Omnia) BySlug(streamType StreamType, slug string, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byslug", []string{slug}, parameters)
}

// Return a item of a given streamtype by it's remote reference number.
// This Call queries for an Item, that is (possibly) not hosted by nexxOMNIA. The API will
// call the given Remote Provider for Media Details and implicitely create the Item for
// future References within nexxOMNIA.
func (o Omnia) ByRemoteRef(streamType StreamType, reference string, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "byremotereference", []string{reference}, parameters)
}

// Return a item of a given streamtype by it's code name. Only available for container
// streamtypes.
func (o Omnia) ByCodeName(streamType StreamType, codename string, parameters QueryParameters) (*Response, error) {
	return o.Call("get", streamType, "bycodename", []string{codename}, parameters)
}

// Generic call to the Omnia API. Won't work with the media management API.
func (o Omnia) Call(
	method string,
	streamType StreamType,
	operation string,
	args []string,
	parameters QueryParameters,
) (*Response, error) {
	return o.universalCall(method, streamType, false, operation, args, parameters)
}

// Generic call to the Omnia management API.
func (o Omnia) ManagementCall(
	method string,
	streamType StreamType,
	operation string,
	args []string,
	parameters QueryParameters,
) (*Response, error) {
	return o.universalCall(method, streamType, true, operation, args, parameters)
}

func (o Omnia) universalCall(
	method string,
	streamType StreamType,
	isManagement bool,
	operation string,
	args []string,
	parameters QueryParameters,
) (*Response, error) {
	method = strings.ToUpper(method)
	args_parts := ""
	if len(args) > 0 {
		args_parts = strings.Join(args, "/")
		args_parts = fmt.Sprintf("/%s", args_parts)
	}
	var reqUrl string
	if !isManagement {
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/%s/%s%s",
			o.DomainId, streamType, operation, args_parts,
		)
	} else {
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/manage/%s/%s/%s",
			o.DomainId, streamType, args_parts, operation,
		)
	}
	header := newOmniaHeader(operation, o.DomainId, o.ApiSecret, o.SessionId)
	paramUrl := ""
	if parameters != nil {
		var err error
		paramUrl, err = parameters.UrlEncode(nil)
		if err != nil {
			return nil, err
		}
	}
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
	if res.Paging != nil {
		log.WithFields(res.Paging.toMap()).Debug("Response Paging")
	}
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
