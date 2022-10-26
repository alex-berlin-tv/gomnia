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

	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/params"
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
func (o Omnia) ById(streamType enums.StreamType, id int, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byid", []string{strconv.Itoa(id)}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's global id.
func (o Omnia) ByGlobalId(streamType enums.StreamType, globalId int, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byglobalid", []string{strconv.Itoa(globalId)}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's hash.
func (o Omnia) ByHash(streamType enums.StreamType, hash string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byhash", []string{hash}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's reference number.
func (o Omnia) ByRefNr(streamType enums.StreamType, reference string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byrefnr", []string{reference}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's slug.
func (o Omnia) BySlug(streamType enums.StreamType, slug string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byslug", []string{slug}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's remote reference number.
// This Call queries for an Item, that is (possibly) not hosted by nexxOMNIA. The API will
// call the given Remote Provider for Media Details and implicitely create the Item for
// future References within nexxOMNIA.
func (o Omnia) ByRemoteRef(streamType enums.StreamType, reference string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byremotereference", []string{reference}, parameters, Response[any]{})
}

// Return a item of a given streamtype by it's code name. Only available for container
// streamtypes.
func (o Omnia) ByCodeName(streamType enums.StreamType, codename string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "bycodename", []string{codename}, parameters, Response[any]{})
}

// Returns all media items of a given streamtype.
func (o Omnia) All(streamType enums.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "all", nil, parameters, Response[any]{})
}

// Returns all items, sorted by Creation Date (ignores the "order" Parameters).
func (o Omnia) Latest(streamType enums.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "latest", nil, parameters, Response[any]{})
}

// Returns all picked media items of a given streamtype. Ignores the order parameter.
func (o Omnia) Picked(streamType enums.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "picked", nil, parameters, Response[any]{})
}

// Returns all evergreen media items of a given streamtype.
func (o Omnia) Evergreens(streamType enums.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "evergreens", nil, parameters, Response[any]{})
}

// eturns all Items, marked as "created for Kids". This is NOT connected to
// any Age Restriction.
func (o Omnia) ForKids(streamType enums.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "forkids", nil, parameters, Response[any]{})
}

// Performs a regular Query on all Items. The "order" Parameters are ignored,
// if querymode is set to "fulltext".
func (o Omnia) ByQuery(streamType enums.StreamType, query string, parameters params.QueryParameters) (*Response[MediaResult], error) {
	rsl, err := Call(o, "get", streamType, "byquery", []string{query}, parameters, Response[MediaResult]{})
	if err != nil {
		return nil, err
	}
	return nil, errors.New(fmt.Sprintf("Wrong type, should be MediaResponse but is %T", rsl))
}

// Will update the general Metadata of a Media Item. Uses the Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#update
func (o Omnia) Update(
	streamType enums.StreamType,
	id int,
	parameters params.Custom,
) {
	log.Warn("result handling not implemented yet")
	ManagementCall(o, "put", streamType, "update", []string{strconv.Itoa(id)}, parameters, Response[any]{})
}

// Approves a media item of a given streamtype and item-id. Uses te Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#approve
func (o Omnia) Approve(
	streamType enums.StreamType,
	id int,
	parameters params.Approve,
) (*Response[any], error) {
	return ManagementCall(o, "post", streamType, "approve", []string{strconv.Itoa(id)}, parameters, Response[any]{})
}

// Rejects a media item of a given streamtype and item-id. Uses te Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#reject
func (o Omnia) Reject(
	streamType enums.StreamType,
	id int,
	parameters params.Reject,
) (*Response[any], error) {
	return ManagementCall(o, "post", streamType, "reject", []string{strconv.Itoa(id)}, parameters, Response[any]{})
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

// Lists all ediatble attributes for a given stream type.
func (o Omnia) EditableAttributes(streamType enums.StreamType) (*Response[EditableAttributesResponse], error) {
	rsl, err := SystemCall(o, "get", "editableattributesfor", []string{string(streamType)}, Response[EditableAttributesResponse]{})
	if err != nil {
		return nil, err
	}
	return nil, errors.New(fmt.Sprintf("Wrong type, should be EditableAttributesResponse but is %T", rsl))
}

// Generic call to the Omnia management API.
func ManagementCall[T any](
	o Omnia,
	method string,
	streamType enums.StreamType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, streamType, managementApiType, operation, args, parameters, response)
}

// Generic call to the Omnia Media API. Won't work with the management API's.
func Call[T any](
	o Omnia,
	method string,
	streamType enums.StreamType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, streamType, mediaApiType, operation, args, parameters, response)
}

// Generic call to the Omnia system API
func SystemCall[T any](
	o Omnia,
	method string,
	operation string,
	args []string,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, enums.VideoStreamType, systemApiType, operation, args, nil, response)
}

func universalCall[T any](
	o Omnia,
	method string,
	streamType enums.StreamType,
	aType apiType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	response Response[T],
) (*Response[T], error) {
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
			"https://api.nexx.cloud/v3.1/%s/manage%s/%s/%s",
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
		Timeout: time.Second * 10,
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
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	log.WithFields(response.Metadata.toMap()).Debug("Response Metadata")
	if response.Paging != nil {
		log.WithFields(response.Paging.toMap()).Debug("Response Paging")
	}
	if response.Metadata.Status != 200 {
		return &response, fmt.Errorf("call failed on server side, please see response object for more infromation")
	}
	log.Trace(response.Result)
	return &response, nil
}
