package hooks

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func (h *DrinkHooks) RegisterCreateOrder() {
	h.PB.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/create-order", func(e *core.RequestEvent) error {
			data := struct {
				UserId   string `json:"user" form:"user"`
				Products []struct {
					Id     string `json:"id" form:"id"`
					Amount int    `json:"amount" form:"amount"`
				} `json:"products" form:"products"`
			}{}
			if err := e.BindBody(&data); err != nil {
				return e.BadRequestError("Failed to read request data", err)
			}

			if !h.ValidateId("users", data.UserId) {
				return e.BadRequestError("Invalid user", nil)
			}

			ids := make([]string, len(data.Products))
			for i, product := range data.Products {
				ids[i] = product.Id
			}
			productsList, err := h.PB.FindRecordsByIds("products", ids)
			if err != nil {
				return e.BadRequestError("Invalid product", err)
			}

			productPrice := make(map[string]float64)
			for _, p := range productsList {
				price := p.GetFloat("price")
				id := p.Id
				productPrice[id] = price
			}

			orderCollection, err := h.PB.FindCollectionByNameOrId("orders")
			if err != nil {
				return e.BadRequestError("Collection not found", err)
			}
			orderRowCollection, err := h.PB.FindCollectionByNameOrId("order_row")
			if err != nil {
				return e.BadRequestError("Collection not found", err)
			}

			err = h.PB.RunInTransaction(func(txApp core.App) error {
				orderRecord := core.NewRecord(orderCollection)
				orderRecord.Set("user", data.UserId)
				for _, row := range data.Products {
					rowRecord := core.NewRecord(orderRowCollection)
					rowRecord.Set("product", row.Id)
					rowRecord.Set("amount", row.Amount)
					rowRecord.Set("price", float64(row.Amount)*productPrice[row.Id])

					if err := txApp.Save(rowRecord); err != nil {
						return err
					}
					orderRecord.Set("rows+", rowRecord.Id)
				}
				return txApp.Save(orderRecord)
			})
			if err != nil {
				return e.BadRequestError("Couldn't insert order", err)
			}

			return e.String(http.StatusOK, "Ok")
		})

		return se.Next()
	})
}

func (h *DrinkHooks) ValidateId(collection, id string) bool {
	_, err := h.PB.FindRecordById(collection, id)
	return err == nil
}
