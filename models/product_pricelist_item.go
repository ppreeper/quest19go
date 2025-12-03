package models

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data product.pricelist.item model
type ProductPricelistItem_150 struct {
	Active          bool      `json:"active"`            // Active
	AppliedOn       any       `json:"applied_on"`        // Apply On ⭐ required
	Base            any       `json:"base"`              // Based on ⭐ required
	BasePricelistID any       `json:"base_pricelist_id"` // Other Pricelist 📦 relation: many2one product.pricelist
	CategID         any       `json:"categ_id"`          // Product Category 📦 relation: many2one product.category
	CompanyID       any       `json:"company_id"`        // Company 📦 relation: many2one res.company
	ComputePrice    any       `json:"compute_price"`     // Compute Price ⭐ required
	CurrencyID      any       `json:"currency_id"`       // Currency 📦 relation: many2one res.currency
	DateEnd         time.Time `json:"date_end"`          // End Date
	DateStart       time.Time `json:"date_start"`        // Start Date
	DisplayName     string    `json:"display_name"`      // Display Name
	FixedPrice      float64   `json:"fixed_price"`       // Fixed Price
	ID              int       `json:"id"`                // ID
	MinQuantity     float64   `json:"min_quantity"`      // Min. Quantity
	Name            string    `json:"name"`              // Name
	PercentPrice    float64   `json:"percent_price"`     // Percentage Price
	Price           string    `json:"price"`             // Price
	PriceDiscount   float64   `json:"price_discount"`    // Price Discount
	PriceMaxMargin  float64   `json:"price_max_margin"`  // Max. Price Margin
	PriceMinMargin  float64   `json:"price_min_margin"`  // Min. Price Margin
	PriceRound      float64   `json:"price_round"`       // Price Rounding
	PriceSurcharge  float64   `json:"price_surcharge"`   // Price Surcharge
	PricelistID     any       `json:"pricelist_id"`      // Pricelist 📦 relation: many2one product.pricelist ⭐ required
	ProductID       any       `json:"product_id"`        // Product Variant 📦 relation: many2one product.product
	ProductTmplID   any       `json:"product_tmpl_id"`   // Product 📦 relation: many2one product.template
	RuleTip         string    `json:"rule_tip"`          // Rule Tip
}

// quest19data product.pricelist.item model
type ProductPricelistItem_190 struct {
	AppliedOn           any       `json:"applied_on"`            // Apply On ⭐ required
	Base                any       `json:"base"`                  // Based on ⭐ required
	BasePricelistID     any       `json:"base_pricelist_id"`     // Other Pricelist 📦 relation: many2one product.pricelist
	CategID             any       `json:"categ_id"`              // Category 📦 relation: many2one product.category
	CompanyID           any       `json:"company_id"`            // Company 📦 relation: many2one res.company
	ComputePrice        any       `json:"compute_price"`         // Compute Price ⭐ required
	CurrencyID          any       `json:"currency_id"`           // Currency 📦 relation: many2one res.currency
	DateEnd             time.Time `json:"date_end"`              // End Date
	DateStart           time.Time `json:"date_start"`            // Start Date
	DisplayAppliedOn    any       `json:"display_applied_on"`    // Display Applied On ⭐ required
	DisplayName         string    `json:"display_name"`          // Display Name
	FixedPrice          float64   `json:"fixed_price"`           // Fixed Price
	ID                  int       `json:"id"`                    // ID
	IsPricelistRequired bool      `json:"is_pricelist_required"` // Is Pricelist Required
	MinQuantity         float64   `json:"min_quantity"`          // Min. Quantity
	Name                string    `json:"name"`                  // Name
	PercentPrice        float64   `json:"percent_price"`         // Percentage Price
	Price               string    `json:"price"`                 // Price
	PriceDiscount       float64   `json:"price_discount"`        // Price Discount
	PriceMarkup         float64   `json:"price_markup"`          // Markup
	PriceMaxMargin      float64   `json:"price_max_margin"`      // Max. Price Margin
	PriceMinMargin      float64   `json:"price_min_margin"`      // Min. Price Margin
	PriceRound          float64   `json:"price_round"`           // Price Rounding
	PriceSurcharge      float64   `json:"price_surcharge"`       // Extra Fee
	PricelistID         any       `json:"pricelist_id"`          // Pricelist 📦 relation: many2one product.pricelist
	ProductID           any       `json:"product_id"`            // Variant 📦 relation: many2one product.product
	ProductTmplID       any       `json:"product_tmpl_id"`       // Product 📦 relation: many2one product.template
	ProductUomName      string    `json:"product_uom_name"`      // Unit Name
	ProductVariantCount int       `json:"product_variant_count"` // # Product Variants
	RuleTip             string    `json:"rule_tip"`              // Rule Tip
}

func (m *Model) ProductPricelistItem() {
	model := "product.pricelist.item"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	currency_map := m.currencyMap()

	// plist, err := m.Source.GetID("product.pricelist", []any{
	// 	[]any{"name", "=", "DISTRIBUTION 00 PARTNER"},
	// })
	// if err != nil {
	// 	m.Log.Error(model, "func", trace(), "err", err)
	// 	return
	// }

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductPricelistItem_150{}), []any{
		// []any{"pricelist_id", "=", plist},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	bar := progressbar.Default(int64(len(records)))

	for _, r := range records {
		var p ProductPricelistItem_150
		FillStruct(r, &p)
		// fmt.Println(prettyprint(p))

		filter := []any{}

		ur := map[string]any{
			"name":         p.Name,
			"min_quantity": p.MinQuantity,
		}
		filter = append(filter, []any{"min_quantity", "=", p.MinQuantity})

		//
		// currency
		//
		_, currency_name := ParseMany2One(p.CurrencyID)
		currency_id := currency_map[currency_name]
		if currency_id != -1 {
			ur["currency_id"] = currency_id
		}
		filter = append(filter, []any{"currency_id", "=", currency_id})

		//
		// parent pricelist
		//
		old_pricelist_id, old_pricelist_name := ParseMany2One(p.PricelistID)
		old_pricelist, err := m.Source.SearchRead("product.pricelist", 0, 0, []string{"name", "display_name"}, []any{
			[]any{"id", "=", old_pricelist_id},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		pricelist_id, err := m.Dest.GetID("product.pricelist", []any{
			[]any{"name", "=", old_pricelist[0]["name"]},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if pricelist_id == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("pricelist not found: %s", old_pricelist_name))
			continue
		}
		ur["pricelist_id"] = pricelist_id
		filter = append(filter, []any{"pricelist_id", "=", pricelist_id})

		//
		// computation
		//
		ur["compute_price"] = p.ComputePrice
		filter = append(filter, []any{"compute_price", "=", p.ComputePrice})

		switch p.ComputePrice {
		case "fixed": // Fixed Price
			ur["fixed_price"] = p.FixedPrice
		case "percentage": // Percentage (discount / surcharge)
			ur["percent_price"] = p.PercentPrice
			filter = append(filter, []any{"percent_price", "=", p.PercentPrice})

		case "formula": // Formula
			ur["price_discount"] = p.PriceDiscount
			ur["price_surcharge"] = p.PriceSurcharge
			ur["price_round"] = p.PriceRound
			ur["price_min_margin"] = p.PriceMinMargin
			ur["price_max_margin"] = p.PriceMaxMargin

			filter = append(filter, []any{"base", "=", p.Base})
			switch p.Base {
			case "list_price":
				ur["base"] = "list_price"
			case "standard_price":
				ur["base"] = "standard_price"
			case "pricelist":
				ur["base"] = "pricelist"
				//
				// base pricelist
				//
				old_base_pricelist_id, old_base_pricelist_name := ParseMany2One(p.BasePricelistID)
				old_base_pricelist, err := m.Source.SearchRead("product.pricelist", 0, 0, []string{"name", "display_name"}, []any{
					[]any{"id", "=", old_base_pricelist_id},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}
				base_pricelist_id, err := m.Dest.GetID("product.pricelist", []any{
					[]any{"name", "=", old_base_pricelist[0]["name"]},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}
				if base_pricelist_id == -1 {
					m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("base pricelist not found: %s", old_base_pricelist_name))
					continue
				}
				ur["base_pricelist_id"] = base_pricelist_id
				filter = append(filter, []any{"base_pricelist_id", "=", base_pricelist_id})
			}
		}

		//
		// conditions
		//
		switch p.AppliedOn {
		case "0_product_variant": // Product Variant
			ur["applied_on"] = "1_product"
			ur["display_applied_on"] = "1_product"
			filter = append(filter, []any{"applied_on", "=", "1_product"})
			filter = append(filter, []any{"display_applied_on", "=", "1_product"})
			_, prod_var_name := ParseMany2One(p.ProductID)
			product_id, err := m.Dest.GetID("product.template", []any{
				[]any{"name", "=", prod_var_name},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if product_id == -1 {
				m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("product not found: %s", prod_var_name))
				continue
			}
			ur["product_tmpl_id"] = product_id
			filter = append(filter, []any{"product_tmpl_id", "=", product_id})
		case "1_product": // Product
			ur["applied_on"] = "1_product"
			ur["display_applied_on"] = "1_product"
			filter = append(filter, []any{"applied_on", "=", "1_product"})
			filter = append(filter, []any{"display_applied_on", "=", "1_product"})
			_, prod_name := ParseMany2One(p.ProductTmplID)
			product_id, err := m.Dest.GetID("product.template", []any{
				[]any{"name", "=", prod_name},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if product_id == -1 {
				m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("product not found: %s", prod_name))
				continue
			}
			ur["product_tmpl_id"] = product_id
			filter = append(filter, []any{"product_tmpl_id", "=", product_id})
		case "2_product_category": // Product Category
			ur["applied_on"] = "2_product_category"
			ur["display_applied_on"] = "2_product_category"
			filter = append(filter, []any{"applied_on", "=", "2_product_category"})
			filter = append(filter, []any{"display_applied_on", "=", "2_product_category"})
			categ, _ := ParseMany2One(p.CategID)

			categs, err := m.Source.SearchRead("product.category", 0, 0, ExtractJSONTags(ProductCategory_150{}), []any{
				[]any{"id", "=", categ},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}

			category_id, err := m.Dest.GetID("product.category", []any{
				[]any{"complete_name", "=", categs[0]["complete_name"]},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}

			ur["categ_id"] = category_id
			filter = append(filter, []any{"categ_id", "=", category_id})
		case "3_global": // All Products
			ur["applied_on"] = "1_product"
			ur["display_applied_on"] = "1_product"
			filter = append(filter, []any{"applied_on", "=", "1_product"})
			filter = append(filter, []any{"display_applied_on", "=", "1_product"})
		}

		rid, _ := m.Dest.GetID(model, filter)

		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
}

// quest19data product.pricelist.item model
// type ProductPricelistItem_190 struct {
// 	AppliedOn           any       `json:"applied_on"`            // Apply On ⭐ required
// 	Base                any       `json:"base"`                  // Based on ⭐ required
// 	BasePricelistID     any       `json:"base_pricelist_id"`     // Other Pricelist 📦 relation: many2one product.pricelist
// 	CategID             any       `json:"categ_id"`              // Category 📦 relation: many2one product.category
// 	CompanyID           any       `json:"company_id"`            // Company 📦 relation: many2one res.company
// 	ComputePrice        any       `json:"compute_price"`         // Compute Price ⭐ required
// 	CurrencyID          any       `json:"currency_id"`           // Currency 📦 relation: many2one res.currency
// 	DateEnd             time.Time `json:"date_end"`              // End Date
// 	DateStart           time.Time `json:"date_start"`            // Start Date
// 	DisplayAppliedOn    any       `json:"display_applied_on"`    // Display Applied On ⭐ required
// 	DisplayName         string    `json:"display_name"`          // Display Name
// 	FixedPrice          float64   `json:"fixed_price"`           // Fixed Price
// 	ID                  int       `json:"id"`                    // ID
// 	IsPricelistRequired bool      `json:"is_pricelist_required"` // Is Pricelist Required
// 	MinQuantity         float64   `json:"min_quantity"`          // Min. Quantity
// 	Name                string    `json:"name"`                  // Name
// 	PercentPrice        float64   `json:"percent_price"`         // Percentage Price
// 	Price               string    `json:"price"`                 // Price
// 	PriceDiscount       float64   `json:"price_discount"`        // Price Discount
// 	PriceMarkup         float64   `json:"price_markup"`          // Markup
// 	PriceMaxMargin      float64   `json:"price_max_margin"`      // Max. Price Margin
// 	PriceMinMargin      float64   `json:"price_min_margin"`      // Min. Price Margin
// 	PriceRound          float64   `json:"price_round"`           // Price Rounding
// 	PriceSurcharge      float64   `json:"price_surcharge"`       // Extra Fee
// 	PricelistID         any       `json:"pricelist_id"`          // Pricelist 📦 relation: many2one product.pricelist
// 	ProductID           any       `json:"product_id"`            // Variant 📦 relation: many2one product.product
// 	ProductTmplID       any       `json:"product_tmpl_id"`       // Product 📦 relation: many2one product.template
// 	ProductUomName      string    `json:"product_uom_name"`      // Unit Name
// 	ProductVariantCount int       `json:"product_variant_count"` // # Product Variants
// 	RuleTip             string    `json:"rule_tip"`              // Rule Tip
// }
