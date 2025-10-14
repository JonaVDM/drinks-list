package hooks

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func (h *DrinkHooks) RegisterCreateOrder() {
	h.PB.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/hello/{name}", func(e *core.RequestEvent) error {
			name := e.Request.PathValue("name")

			return e.String(http.StatusOK, "Hello "+name)
		})

		return se.Next()
	})
}
