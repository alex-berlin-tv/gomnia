package omnia

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Internal destinction for the different API's. This is needed
// as the structure of a call differs depending on the API.
type apiType int

const (
	mediaApiType apiType = iota
	managementApiType
	systemApiType
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
func (o Omnia) ById(streamType StreamType, id int, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byid", []string{strconv.Itoa(id)}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's global id.
func (o Omnia) ByGlobalId(streamType StreamType, globalId int, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byglobalid", []string{strconv.Itoa(globalId)}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's hash.
func (o Omnia) ByHash(streamType StreamType, hash string, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byhash", []string{hash}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's reference number.
func (o Omnia) ByRefNr(streamType StreamType, reference string, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byrefnr", []string{reference}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's slug.
func (o Omnia) BySlug(streamType StreamType, slug string, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byslug", []string{slug}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's remote reference number.
// This Call queries for an Item, that is (possibly) not hosted by nexxOMNIA. The API will
// call the given Remote Provider for Media Details and implicitely create the Item for
// future References within nexxOMNIA.
func (o Omnia) ByRemoteRef(streamType StreamType, reference string, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "byremotereference", []string{reference}, parameters, &UniversalResponse{})
}

// Return a item of a given streamtype by it's code name. Only available for container
// streamtypes.
func (o Omnia) ByCodeName(streamType StreamType, codename string, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "bycodename", []string{codename}, parameters, &UniversalResponse{})
}

// Returns all media items of a given streamtype.
func (o Omnia) All(streamType StreamType, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "all", nil, parameters, &UniversalResponse{})
}

// Returns all items, sorted by Creation Date (ignores the "order" Parameters).
func (o Omnia) Latest(streamType StreamType, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "latest", nil, parameters, &UniversalResponse{})
}

// Returns all picked media items of a given streamtype. Ignores the order parameter.
func (o Omnia) Picked(streamType StreamType, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "picked", nil, parameters, &UniversalResponse{})
}

// Returns all evergreen media items of a given streamtype.
func (o Omnia) Evergreens(streamType StreamType, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "evergreens", nil, parameters, &UniversalResponse{})
}

// eturns all Items, marked as "created for Kids". This is NOT connected to
// any Age Restriction.
func (o Omnia) ForKids(streamType StreamType, parameters QueryParameters) (Response, error) {
	return o.Call("get", streamType, "forkids", nil, parameters, &UniversalResponse{})
}

// Performs a regular Query on all Items. The "order" Parameters are ignored,
// if querymode is set to "fulltext".
func (o Omnia) ByQuery(streamType StreamType, query string, parameters QueryParameters) (*MediaResponse, error) {
	rsl, err := o.Call("get", streamType, "byquery", []string{query}, parameters, &MediaResponse{})
	if err != nil {
		return &MediaResponse{}, err
	}
	if rsl, ok := rsl.(*MediaResponse); ok {
		return rsl, nil
	}
	return &MediaResponse{}, errors.New(fmt.Sprintf("Wrong type, should be MediaResponse but is %T", rsl))
}

// Generic call to the Omnia Media API. Won't work with the management API's.
func (o Omnia) Call(
	method string,
	streamType StreamType,
	operation string,
	args []string,
	parameters QueryParameters,
	response Response,
) (Response, error) {
	return o.universalCall(method, streamType, mediaApiType, operation, args, parameters, response)
}

// Update a field in the metdata of a media item.
func (o Omnia) Update(
	streamType StreamType,
	id int,
	parameters CustomParameters,
) {
	o.ManagementCall("put", streamType, "update", []string{strconv.Itoa(id)}, parameters, UniversalResponse{})
}

// Generic call to the Omnia management API.
func (o Omnia) ManagementCall(
	method string,
	streamType StreamType,
	operation string,
	args []string,
	parameters QueryParameters,
	response Response,
) (Response, error) {
	return o.universalCall(method, streamType, managementApiType, operation, args, parameters, response)
}

// Lists all ediatble attributes for a given stream type.
func (o Omnia) EditableAttributes(streamType StreamType) (*EditableAttributesResponse, error) {
	rsl, err := o.SystemCall("get", "editableattributesfor", []string{string(streamType)}, &EditableAttributesResponse{})
	if err != nil {
		return &EditableAttributesResponse{}, err
	}
	if rsl, ok := rsl.(*EditableAttributesResponse); ok {
		return rsl, nil
	}
	return &EditableAttributesResponse{}, errors.New(fmt.Sprintf("Wrong type, should be EditableAttributesResponse but is %T", rsl))
}

// Generic call to the Omnia system API
func (o Omnia) SystemCall(
	method string,
	operation string,
	args []string,
	response Response,
) (Response, error) {
	return o.universalCall(method, VideoStreamType, systemApiType, operation, args, nil, response)
}

func (o Omnia) universalCall(
	method string,
	streamType StreamType,
	aType apiType,
	operation string,
	args []string,
	parameters QueryParameters,
	response Response,
) (Response, error) {
	method = strings.ToUpper(method)
	argsParts := ""
	if len(args) > 0 {
		argsParts = strings.Join(args, "/")
		argsParts = fmt.Sprintf("/%s", argsParts)
	}
	var reqUrl string
	switch aType {
	case mediaApiType:
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/%s/%s%s",
			o.DomainId, streamType, operation, argsParts,
		)
	case managementApiType:
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/manage/%s/%s/%s",
			o.DomainId, streamType, argsParts, operation,
		)
	case systemApiType:
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/system/%s%s",
			o.DomainId, operation, argsParts,
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

	reqUrl = fmt.Sprintf("%s?%s", reqUrl, paramUrl)

	req, err := http.NewRequest(method, reqUrl, nil)
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
	log.Trace(string(body))
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	log.WithFields(response.GetMetadata().toMap()).Debug("Response Metadata")
	if response.GetPaging() != nil {
		log.WithFields(response.GetPaging().toMap()).Debug("Response Paging")
	}
	log.Trace(response.GetResult())
	return response, nil
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
