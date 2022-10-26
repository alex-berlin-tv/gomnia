package cli

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

// Gets the id(s) from the cli.Context. The id(s) get determined by
// the following rules:
//
//  1. The id(s) can be set using one or more `--id` flag.
//  2. Each `--id` flag can itself contain one or more ids seperated by a
//     comma.
//  3. If _no_ `--id` flag is set and the global option `--id-from-clipboard`
//     _is_ set, the method will try to extract the id(s) from the clipbaord
//     according to the second rule.
func getIds(ctx *cli.Context) ([]int, error) {
	if !ctx.IsSet("id") && !ctx.Bool("id-from-clipboard") {
		return nil, fmt.Errorf("flag '--id' is not set")
	}
	var rsl []int
	for _, value := range ctx.StringSlice("id") {
		tmp, err := parseIdValue(value)
		if err != nil {
			return nil, err
		}
		rsl = append(rsl, tmp...)
	}
	if len(rsl) > 0 {
		return rsl, nil
	}
	if err := clipboard.Init(); err != nil {
		return nil, fmt.Errorf("usage of id-from-clipboard flag is not possible as clipboard cannot be read, %s", err)
	}
	data := string(clipboard.Read(clipboard.FmtText))
	values, err := parseIdValue(data)
	if err != nil {
		return nil, fmt.Errorf("content of clipbaord '%s' cannot be parsed as id(s)", data)
	}
	log.Infof("use id(s) %+v from clipboard", values)
	return values, nil
}

// Tries to parse the id value of a string according to the following rules:
//
//  1. The id's have to be integers.
//  2. Each value can contain one or more id(s).
//  3. Multiple id's are seperated by a comma.
func parseIdValue(raw string) ([]int, error) {
	var validId = regexp.MustCompile(`^(\d*[,]*)+$`)
	if !validId.MatchString(raw) {
		return nil, fmt.Errorf("'%s' cannot be parsed as id(s)", raw)
	}
	raws := strings.Split(raw, ",")
	var rsl []int
	for _, rawPart := range raws {
		value, err := strconv.Atoi(rawPart)
		if err != nil {
			return nil, fmt.Errorf("part '%s' of '%s' cannot be converted to a integer id", rawPart, raw)
		}
		rsl = append(rsl, value)
	}
	return rsl, nil
}

// Handles the errors returned by an API call. This is used
// to present the user with in depth information if the call
// fails on the server side (statuscode != 200).
func handleApiError(rsp *omnia.Response[any], err error, fatalOnError bool) {
	if rsp == nil {
		log.Error("client side error in the API library")
		if fatalOnError {
			log.Fatal(err)
		}
		log.Error(err)
		return
	}
	fields := log.Fields{
		"status":     rsp.Metadata.Status,
		"error_hint": *rsp.Metadata.ErrorHint,
		"called":     *rsp.Metadata.CalledWith,
		"action":     rsp.Metadata.Verb,
	}
	if fatalOnError {
		log.WithFields(fields).Fatal("server side error")
	}
	log.WithFields(fields).Error("server side error")
}
