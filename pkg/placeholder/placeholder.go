package placeholder

import (
	"net/http"

	"github.com/fsrv-xyz/webbase-function-base/pkg/registry"
)

func HandlerBuilder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placeholder"))
	}
}

func init() {
	registry.RegisterHttpFunction("placeholder", HandlerBuilder())
}
