package hooks

import "github.com/pocketbase/pocketbase"

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
	//something else

	return hooks
}
