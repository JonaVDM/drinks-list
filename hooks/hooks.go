package hooks

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type DrinkHooks struct {
	PB     *pocketbase.PocketBase
	Config DrinkHooksConfig
}

type DrinkHooksConfig struct{}

func MustRegiser(app *pocketbase.PocketBase, config DrinkHooksConfig) DrinkHooks {
	hooks := DrinkHooks{
		PB:     app,
		Config: config,
	}

	hooks.RegisterCreateOrder()
	hooks.RegisterTotal()

	return hooks
}

func (h *DrinkHooks) RegisterTotal() {
	h.PB.OnRecordEnrich("orders").BindFunc(func(e *core.RecordEnrichEvent) error {
		e.Record.WithCustomData(true)

		rows := e.Record.ExpandedAll("rows")
		if len(rows) == 0 {
			fmt.Println(e.Record.GetStringSlice("rows"))
			rows, _ = h.PB.FindRecordsByIds("order_row", e.Record.GetStringSlice("rows"))
		}
		total := 0.0

		for _, row := range rows {
			price := row.GetFloat("price")
			total += price
			fmt.Println(total)
		}

		e.Record.Set("total", total)

		return e.Next()
	})
}
