package models

import "github.com/schollz/progressbar/v3"

// quest15data ncr.product_line model
type NcrProductLine_150 struct {
	CurrencyID          any     `json:"currency_id"`            // Currency 📦 relation: many2one res.currency
	DisplayName         string  `json:"display_name"`           // Display Name
	ID                  int     `json:"id"`                     // ID
	NcrClaimID          any     `json:"ncr_claim_id"`           // NCR 📦 relation: many2one ncr.claim
	ProductCost         float64 `json:"product_cost"`           // Cost
	ProductID           any     `json:"product_id"`             // Product 📦 relation: many2one product.product
	ProductLineDesc     string  `json:"product_line_desc"`      // Description
	ProductQty          float64 `json:"product_qty"`            // Quantity
	ProductSubtotal     float64 `json:"product_subtotal"`       // Subtotal
	PurchaseOrderID     any     `json:"purchase_order_id"`      // Purchase Order 📦 relation: many2one purchase.order
	PurchaseOrderLineID any     `json:"purchase_order_line_id"` // Purchase Order Line 📦 relation: many2one purchase.order.line
	SaleOrderID         any     `json:"sale_order_id"`          // Sale Order 📦 relation: many2one sale.order
	SaleOrderLineID     any     `json:"sale_order_line_id"`     // Sale Order Line 📦 relation: many2one sale.order.line
}

// quest19data ncr.product.line model
// type NcrProductLine_190 struct {
// 	CurrencyID          any     `json:"currency_id"`            // Currency 📦 relation: many2one res.currency
// 	DisplayName         string  `json:"display_name"`           // Display Name
// 	ID                  int     `json:"id"`                     // ID
// 	NcrClaimID          any     `json:"ncr_claim_id"`           // NCR 📦 relation: many2one ncr.claim
// 	ProductCost         float64 `json:"product_cost"`           // Cost
// 	ProductID           any     `json:"product_id"`             // Product 📦 relation: many2one product.product
// 	ProductLineDesc     string  `json:"product_line_desc"`      // Description
// 	ProductQty          float64 `json:"product_qty"`            // Quantity
// 	ProductSubtotal     float64 `json:"product_subtotal"`       // Subtotal
// 	ProductUomID        any     `json:"product_uom_id"`         // Unit 📦 relation: many2one uom.uom ***************** 19 only
// 	PurchaseOrderID     any     `json:"purchase_order_id"`      // Purchase Order 📦 relation: many2one purchase.order
// 	PurchaseOrderLineID any     `json:"purchase_order_line_id"` // Purchase Order Line 📦 relation: many2one purchase.order.line
// 	SaleOrderID         any     `json:"sale_order_id"`          // Sale Order 📦 relation: many2one sale.order
// 	SaleOrderLineID     any     `json:"sale_order_line_id"`     // Sale Order Line 📦 relation: many2one sale.order.line
// }

func (m *Model) NCRProductLine() {
	model := "ncr.product_line"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	currency := m.currencyMap()

	ncr_claims, err := m.Source.SearchRead("ncr.claim", 0, 0, ExtractJSONTags(NcrClaim_150{}), []any{
		// []any{"parent_category_id", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	// root
	// records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(NcrProductLine_150{}), []any{
	// 	// []any{"parent_category_id", "=", false},
	// })
	// if err != nil {
	// 	m.Log.Error(model, "func", trace(), "err", err)
	// 	return
	// }

	for _, r := range ncr_claims {
		var claim NcrClaim_150
		FillStruct(r, &claim)
		// fmt.Println(prettyprint(claim.ProductLine))
		ncr_claim_id, _ := m.Dest.GetID("ncr.claim", []any{[]any{"name", "=", claim.Name}})

		products, err := m.Source.SearchRead("ncr.product_line", 0, 0, ExtractJSONTags(NcrProductLine_150{}), []any{
			[]any{"id", "in", claim.ProductLine},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		bar := progressbar.Default(-1, claim.Name+" product")
		for _, r := range products {
			var product NcrProductLine_150
			FillStruct(r, &product)
			// fmt.Println(prettyprint(product))

			ur := map[string]any{
				"ncr_claim_id":      ncr_claim_id,
				"product_line_desc": product.ProductLineDesc,
				"product_qty":       product.ProductQty,
				"product_cost":      product.ProductCost,
				"product_subtotal":  product.ProductSubtotal,
				// 	"sale_order_id":          product.SaleOrderID,
				// 	"sale_order_line_id":     product.SaleOrderLineID,
				// 	"purchase_order_id":      product.PurchaseOrderID,
				// 	"purchase_order_line_id": product.PurchaseOrderLineID,
			}

			pid_source, _ := ParseMany2One(product.ProductID)
			ps, _ := m.Source.SearchRead("product.product", 0, 0, []string{"id", "name"}, []any{[]any{"id", "=", pid_source}})
			if len(ps) > 0 {
				pid := -1
				if pp_name_source, ok := ps[0]["name"].(string); ok {
					pid, _ = m.Dest.GetID("product.product", []any{[]any{"name", "=", pp_name_source}})
				}
				if pid != -1 {
					ur["product_id"] = pid
				}
			}
			// fmt.Println("prepared ncr.product.line record:", prettyprint(ur))

			// currency mapping
			cid := -1
			_, cur := ParseMany2One(product.CurrencyID)
			if mappedCur, ok := currency[cur]; ok {
				cid = mappedCur
			}
			if cid != -1 {
				ur["currency_id"] = cid
			}

			rid, err := m.Dest.GetID("ncr.product.line", []any{
				[]any{"ncr_claim_id", "=", ncr_claim_id},
				[]any{"product_id", "=", product.ProductID},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				return
			}
			// fmt.Println("mapped ncr.product.line id:", rid)

			m.writeRecord("ncr.product.line", ur, rid)

			bar.Add(1)
		}
		bar.Finish()
	}
}
