package main

import (
	"fmt"

	"github.com/alex-berlin-tv/nexx_omnia_go/cli"
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia"

	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	rsl, _ := omnia.EnumByValue[omnia.Bool](omnia.NoBool, "1")
	fmt.Printf("%T", rsl)
	if err := cli.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
