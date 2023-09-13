// The omnia package provides a way to interact with the API of the media
// management platform 3q nexx omnia. More information on the API can be
// obtained by reading the [official API documentation].
//
// [official API documentation]: https://api.docs.nexx.cloud/
package omnia

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alex-berlin-tv/nexx_omnia_go/enum"
	"github.com/alex-berlin-tv/nexx_omnia_go/params"
	"github.com/sirupsen/logrus"
)

// Internal distinction for the different API's. This is needed
// as the structure of a call differs depending on the API.
type apiType int

const (
	mediaApiType apiType = iota
	managementApiType
	uploadLinkManagementApiType
	systemApiType
	domainApiType
)

const omniaHeaderXRequestCid = "X-Request-CID"
const omniaHeaderXRequestToken = "X-Request-Token"

// Header for an Omnia API call.
type omniaHeader struct {
	xRequestCid   string
	xRequestToken string
}

func newOmniaHeader(operation, domainId, apiSecret, sessionId string) omniaHeader {
	logrus.Debugf("hash source: md5(%s+%s+API_SECRET)", operation, domainId)
	signature := md5.Sum([]byte(fmt.Sprintf("%s%s%s", operation, domainId, apiSecret)))
	return omniaHeader{
		xRequestCid:   sessionId,
		xRequestToken: hex.EncodeToString(signature[:]),
	}
}

// Encapsulates the calls to the nexxOmnia API. All implemented API methods are available
// as a method of this struct. In order to obtain the needed information in omnia go to
// Domains & Services > Domains and open the detail view (»Details anzeigen«) for the domain
// you want to wish to control with this package. You find the name of the variables in the
// documentation of the field.
type Client struct {
	// Named »ID« in the domain detail view.
	DomainId string `json:"domain_id"`
	// Named »API Secret« in the domain detail view.
	ApiSecret string `json:"api_secret"`
	// Named »Management API Session« in the domain detail view.
	SessionId string `json:"session_id"`
}

// Returns a new Omnia instance. For mor information on how to obtain the needed
// parameters please refer to the documentation of the [Client] type.
func NewClient(domainId string, apiSecret string, sessionId string) Client {
	return Client{
		DomainId:  domainId,
		ApiSecret: apiSecret,
		SessionId: sessionId,
	}
}

// Reads an Omnia instance from a json file.
func OmniaFromFile(path string) Client {
	file, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}
	rsl := Client{}
	if err := json.Unmarshal([]byte(file), &rsl); err != nil {
		logrus.Fatal(err)
	}
	return rsl
}

// Return a item of a given streamtype by it's id. Example, get the audio item with
// the id 72 with the publishing details:
//
//	id := 72
//	client := omnia.NewClient("23", "Secret", "42")
//	rsl, err := client.ById(enum.AudioStreamType, id, params.Custom{
//		"addPublishingDetails": 1,
//	})
func (o Client) ById(streamType enum.StreamType, id int, parameters params.QueryParameters) (*Response[MediaResultItem], error) {
	return Call(o, "get", streamType, "byid", []string{strconv.Itoa(id)}, parameters, 1, Response[MediaResultItem]{})
}

// Return a item of a given streamtype by it's global id.
func (o Client) ByGlobalId(streamType enum.StreamType, globalId int, parameters params.QueryParameters) (*Response[MediaResultItem], error) {
	return Call(o, "get", streamType, "byglobalid", []string{strconv.Itoa(globalId)}, parameters, 1, Response[MediaResultItem]{})
}

// Return a item of a given streamtype by it's hash.
func (o Client) ByHash(streamType enum.StreamType, hash string, parameters params.QueryParameters) (*Response[MediaResultItem], error) {
	return Call(o, "get", streamType, "byhash", []string{hash}, parameters, 1, Response[MediaResultItem]{})
}

// Return a item of a given streamtype by it's reference number.
func (o Client) ByRefNr(streamType enum.StreamType, reference string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byrefnr", []string{reference}, parameters, 1, Response[any]{})
}

// Return a item of a given streamtype by it's slug.
func (o Client) BySlug(streamType enum.StreamType, slug string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byslug", []string{slug}, parameters, 1, Response[any]{})
}

// Return a item of a given streamtype by it's remote reference number.
// This Call queries for an Item, that is (possibly) not hosted by nexxOMNIA. The API will
// call the given Remote Provider for Media Details and implicitly create the Item for
// future References within nexxOMNIA.
func (o Client) ByRemoteRef(streamType enum.StreamType, reference string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "byremotereference", []string{reference}, parameters, 1, Response[any]{})
}

// Return a item of a given streamtype by it's code name. Only available for container
// streamtypes.
func (o Client) ByCodeName(streamType enum.StreamType, codename string, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "bycodename", []string{codename}, parameters, 1, Response[any]{})
}

// Returns all media items of a given streamtype. Please note that it's not possible
// to retrieve more than 100 items at once using the API. Thus if you have more than
// 100 items of a given streamtype you'll should use the [Client.AllPaged] method
// in order to get all items.
func (o Client) All(streamType enum.StreamType, parameters params.QueryParameters) (*Response[MediaResult], error) {
	return Call(o, "get", streamType, "all", nil, parameters, 1, Response[MediaResult]{})
}

// Joins results of multiple pages if there are more than 100 items and
// the API starts to use paging.
func (o Client) AllPaged(streamType enum.StreamType, parameters params.QueryParameters) (*Response[MediaResult], error) {
	rqs, err := Call(o, "get", streamType, "all", nil, parameters, 1, Response[MediaResult]{})
	if err != nil {
		return nil, err
	}
	if rqs.Paging.ResultCount <= 100 {
		return rqs, nil
	}
	for i := 100; i < rqs.Paging.ResultCount; i += 100 {
		tmp, err := Call(o, "get", streamType, "all", nil, parameters, i, Response[MediaResult]{})
		if err != nil {
			return nil, err
		}
		rqs.Result = append(rqs.Result, tmp.Result...)
	}
	return rqs, nil
}

// Returns all items, sorted by Creation Date (ignores the "order" Parameters).
func (o Client) Latest(streamType enum.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "latest", nil, parameters, 1, Response[any]{})
}

// Returns all picked media items of a given streamtype. Ignores the order parameter.
func (o Client) Picked(streamType enum.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "picked", nil, parameters, 1, Response[any]{})
}

// Returns all evergreen media items of a given streamtype.
func (o Client) Evergreens(streamType enum.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "evergreens", nil, parameters, 1, Response[any]{})
}

// Returns all Items, marked as "created for Kids". This is NOT connected to
// any Age Restriction.
func (o Client) ForKids(streamType enum.StreamType, parameters params.QueryParameters) (*Response[any], error) {
	return Call(o, "get", streamType, "forkids", nil, parameters, 1, Response[any]{})
}

// Performs a regular Query on all Items. The "order" Parameters are ignored,
// if query-mode is set to "fulltext".
func (o Client) ByQuery(streamType enum.StreamType, query string, parameters params.QueryParameters) (*Response[MediaResult], error) {
	rsl, err := Call(o, "get", streamType, "byquery", []string{query}, parameters, 1, Response[MediaResult]{})
	if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("wrong type, should be MediaResponse but is %T", rsl)
}

// Will update the general Metadata of a Media Item. Uses the Management API.
// Documentation can be found [here]. Example, change the title of the video
// item with the id 72:
//
//	client := omnia.NewClient("23", "Secret", "42")
//	client.Update(enums.VideoStreamType, 72, params.Custom{
//		"title": "My cool new title",
//	})
//
// As you see using a params.Custom map you can alter all the available metadata
// fields. You can use the [Client.EditableAttributes] method to get a list of all fields
// which are available and editable for an media item in omnia.
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#update
func (o Client) Update(
	streamType enum.StreamType,
	id int,
	parameters params.Custom,
) (*Response[any], error) {
	rsp, err := ManagementCall(o, "put", streamType, "update", []string{strconv.Itoa(id)}, parameters, Response[any]{})
	if err != nil {
		if rsp != nil {
			return rsp, fmtOmniaErr(*rsp)
		}
		return nil, err
	}
	return rsp, nil
}

// Approves a media item of a given streamtype and item-id. Uses te Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#approve
func (o Client) Approve(
	streamType enum.StreamType,
	id int,
	parameters params.Approve,
) (*Response[any], error) {
	return ManagementCall(o, "post", streamType, "approve", []string{strconv.Itoa(id)}, parameters, Response[any]{})
}

// Publish a media item of a given streamtype and item-id. Uses te Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#publish
func (o Client) Publish(
	streamType enum.StreamType,
	id int,
) (*Response[any], error) {
	return ManagementCall(o, "post", streamType, "publish", []string{strconv.Itoa(id)}, nil, Response[any]{})
}

// Rejects a media item of a given streamtype and item-id. Uses te Management API.
// Documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#reject
func (o Client) Reject(
	streamType enum.StreamType,
	id int,
	parameters params.Reject,
) (*Response[any], error) {
	return ManagementCall(o, "post", streamType, "reject", []string{strconv.Itoa(id)}, parameters, Response[any]{})
}

// Add a new channel. Documentation can be found [here].
//
// [here]: https://api.nexx.cloud/v3.1/manage/channels/add
func (o Client) AddChannel(parameters params.Channel) (*Response[any], error) {
	return ManagementCall(o, "post", "channels", "add", nil, parameters, Response[any]{})
}

// Adds a new UploadLink. UploadsLinks are dynamic URLs, that allow external Users to
// upload Files to a specific nexxOMNIA Account. Uses the Management API. Documentation
// can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/domain-management#uploadlinks
func (o Client) AddUploadLink(parameters params.UploadLink) (*Response[any], error) {
	if err := parameters.Validate(); err != nil {
		return nil, fmt.Errorf("invalid parameters given for AddUploadLink, %s", err)
	}
	return universalCall(o, http.MethodPost, enum.VideoStreamType, uploadLinkManagementApiType, "add", nil, parameters, 1, Response[any]{})
}

// Lists all editable attributes for a given stream type.
//
// This method is needed as there is no other documentation of all the available metadata
// fields in omnia. Especially useful if you want to know which metadata attributes
// you can alter using the [Client.Update] method.
func (o Client) EditableAttributes(streamType enum.StreamType) (*Response[EditableAttributesResponse], error) {
	rsl, err := SystemCall(o, "get", "editableattributesfor", []string{string(streamType)}, Response[EditableAttributesResponse]{})
	if err != nil {
		return nil, err
	}
	// return nil, fmt.Errorf("wrong type, should be EditableAttributesResponse but is %T", rsl)
	return rsl, nil
}

// Generic call to the Omnia Media API. Won't work with the management API's.
func Call[T any](
	o Client,
	method string,
	streamType enum.StreamType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	pagingStart int,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, streamType, mediaApiType, operation, args, parameters, pagingStart, response)
}

// Generic call to the Omnia management API.
func ManagementCall[T any](
	o Client,
	method string,
	streamType enum.StreamType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, streamType, managementApiType, operation, args, parameters, 1, response)
}

// Generic call to the Omnia system API
func SystemCall[T any](
	o Client,
	method string,
	operation string,
	args []string,
	response Response[T],
) (*Response[T], error) {
	return universalCall(o, method, enum.VideoStreamType, systemApiType, operation, args, nil, 1, response)
}

func universalCall[T any](
	o Client,
	method string,
	streamType enum.StreamType,
	aType apiType,
	operation string,
	args []string,
	parameters params.QueryParameters,
	pagingStart int,
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
			"https://api.nexx.cloud/v3.1/%s/manage/%s%s/%s",
			o.DomainId, streamType, argsParts, operation,
		)
	case uploadLinkManagementApiType:
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/manage/uploadlinks/%s",
			o.DomainId, operation,
		)
	case systemApiType:
		reqUrl = fmt.Sprintf(
			"https://api.nexx.cloud/v3.1/%s/system/%s%s",
			o.DomainId, operation, argsParts,
		)
	}
	header := newOmniaHeader(operation, o.DomainId, o.ApiSecret, o.SessionId)

	limitParam, err := params.Basic{
		Limit: 100,
		Start: pagingStart,
	}.UrlEncode()
	if err != nil {
		return nil, err
	}
	var paramUrl string
	if parameters != nil {
		var err error
		paramUrl, err = parameters.UrlEncode()
		if err != nil {
			return nil, err
		}
		paramUrl = fmt.Sprintf("%s&%s", paramUrl, limitParam)
	} else {
		paramUrl = limitParam
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
	logrus.Trace(string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	logrus.WithFields(response.Metadata.toMap()).Debug("Response Metadata")
	if response.Paging != nil {
		logrus.WithFields(response.Paging.toMap()).Debug("Response Paging")
	}
	if response.Metadata.Status != 200 && response.Metadata.Status != 201 {
		return &response, fmt.Errorf("call failed on server side with status code %d, %s", response.Metadata.Status, *response.Metadata.ErrorHint)
	}
	logrus.Trace(response.Result)
	return &response, nil
}

// Logs parameters of API call.
func (o Client) debugLog(method string, url string, header omniaHeader, parameters string) {
	var paramStr string
	if parameters != "" {
		paramStr = fmt.Sprintf("%+v", parameters)
	}
	logrus.WithFields(logrus.Fields{
		"method": method,
		"url":    url,
		"header": fmt.Sprintf("%+v", header),
		"params": paramStr,
	}).Debug("send request to Omnia")
}

// Formats an error message from an Omnia response.
func fmtOmniaErr[T any](rsp Response[T]) error {
	return fmt.Errorf("call to omnia failed with %d, %s", rsp.Metadata.Status, *rsp.Metadata.ErrorHint)
}
