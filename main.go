package main

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/cli"

	"os"

	log "github.com/sirupsen/logrus"
)

type Metadata struct {
	Status  string
	Version int
}

type Paging struct {
	Min     int
	Max     int
	Current int
}

type Response[T any] struct {
	Meta   Metadata
	Result T
	Pag    Paging
}

type UniversalResponse struct {
	A int
	B int
}

type SpecialResponse struct {
	Foo string
	Bar string
}

type Omnia struct{}

func universalCall[T any](o Omnia) Response[T] {
	return Response[T]{
		Meta: Metadata{Status: "ok", Version: 1},
	}
}

func main() {
	if err := cli.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
