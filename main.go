package main

import (
	"fmt"

	"github.com/bonsai-oss/webbase"

	_ "github.com/fsrv-xyz/webbase-function-base/pkg/placeholder"
	"github.com/fsrv-xyz/webbase-function-base/pkg/registry"
)

func main() {
	router := webbase.NewRouter()
	for name, reg := range registry.GetRegisteredFunctions() {
		router.PathPrefix("/" + name).HandlerFunc(reg.HttpFunction)
	}
	fmt.Printf("%#v\n", router)
	webbase.ServeRouter("router", router)
}
