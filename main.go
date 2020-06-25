package main

import (
	"log"

	"github.com/hpidcock/wgcf/cmd"
	"github.com/hpidcock/wgcf/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
