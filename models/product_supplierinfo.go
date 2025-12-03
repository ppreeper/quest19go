package models

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data product.supplierinfo model
type ProductSupplierinfo_150 struct {
	// CompanyID        any       `json:"company_id"`         // Company 📦 relation: many2one res.company
	CurrencyID  any       `json:"currency_id"`  // Currency 📦 relation: many2one res.currency ⭐ required
	DateEnd     time.Time `json:"date_end"`     // End Date
	DateStart   time.Time `json:"date_start"`   // Start Date
	Delay       int       `json:"delay"`        // Delivery Lead Time ⭐ required
	DisplayName string    `json:"display_name"` // Display Name
	ID          int       `json:"id"`           // ID
	// LastPurchaseDate time.Time `json:"last_purchase_date"` // Last Purchase
	MinQty        float64 `json:"min_qty"`         // Quantity ⭐ required
	Name          any     `json:"name"`            // Vendor 📦 relation: many2one res.partner ⭐ required
	Price         float64 `json:"price"`           // Price ⭐ required
	ProductCode   string  `json:"product_code"`    // Vendor Product Code
	ProductID     any     `json:"product_id"`      // Product Variant 📦 relation: many2one product.product
	ProductName   string  `json:"product_name"`    // Vendor Product Name
	ProductTmplID any     `json:"product_tmpl_id"` // Product Template 📦 relation: many2one product.template
	ProductUom    any     `json:"product_uom"`     // Unit of Measure 📦 relation: many2one uom.uom
	// ProductVariantCount       int       `json:"product_variant_count"`        // Variant Count
	// PurchaseRequisitionID     any       `json:"purchase_requisition_id"`      // Agreement 📦 relation: many2one purchase.requisition
	// PurchaseRequisitionLineID any       `json:"purchase_requisition_line_id"` // Purchase Requisition Line 📦 relation: many2one purchase.requisition.line
	Sequence int `json:"sequence"` // Sequence
	// ShowSetSupplierButton     bool      `json:"show_set_supplier_button"`     // Show Set Supplier Button
}

func (m *Model) ProductSupplierinfo() {
	model := "product.supplierinfo"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductSupplierinfo_150{}), []any{
		[]any{"product_tmpl_id.active", "=", true},
		[]any{"name.active", "=", true},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	currencies := m.currencyMap()
	uoms := m.uomMAP()

	bar := progressbar.Default(int64(len(records)), "product.supplierinfo")
	for _, r := range records {
		var p ProductSupplierinfo_150
		FillStruct(r, &p)
		// fmt.Println(prettyprint(p))

		ur := map[string]any{
			"price":        p.Price,
			"min_qty":      p.MinQty,
			"delay":        p.Delay,
			"product_name": p.ProductName,
			"product_code": p.ProductCode,
			"sequence":     p.Sequence,
		}

		// currency
		_, curr_name := ParseMany2One(p.CurrencyID)
		currency_id := currencies[curr_name]
		if currency_id == 0 {
			m.Log.Error(model, "func", trace(), "err", "currency not found", "currency", curr_name)
			continue
		}
		ur["currency_id"] = currency_id

		// product
		prod_id, product_name := ParseMany2One(p.ProductTmplID)
		prod, _ := m.Source.SearchRead("product.template", 0, 0, []string{"name"}, []any{
			[]any{"id", "=", prod_id},
		})
		if len(prod) == 0 {
			m.Log.Error(model, "func", trace(), "err", "product not found", "product_id", prod_id, "product_name", product_name)
			continue
		}
		prod_name := prod[0]["name"]
		product_id, err := m.Dest.GetID("product.template", []any{[]any{"name", "=", prod_name}})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if product_id != -1 {
			ur["product_tmpl_id"] = product_id
		}

		// product_uom
		_, uom_name := ParseMany2One(p.ProductUom)
		uom_id := 0
		switch uom_name {
		case "FT":
			uom_id = uoms["ft"]
		case "IN":
			uom_id = uoms["in"]
		case "LYD":
			uom_id = uoms["yd"]
		case "LFT":
			uom_id = uoms["ft"]
		case "Gram":
			uom_id = uoms["g"]
		default:
			uom_id = uoms[uom_name]
		}
		if uom_id == 0 {
			m.Log.Error(model, "func", trace(), "err", "uom not found", "uom", uom_name)
			continue
		}
		ur["product_uom_id"] = uom_id

		// vendor
		_, vendor_name := ParseMany2One(p.Name)
		vendor_id, err := m.Dest.GetID("res.partner", []any{[]any{"name", "=", vendor_name}})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if vendor_id == -1 {
			m.Log.Error(model, "func", trace(), "err", "vendor not found", "vendor", vendor_name)
			continue
		}
		ur["partner_id"] = vendor_id

		// date_start := FormatDate(p.DateStart)
		rid, err := m.Dest.GetID(model, []any{
			[]any{"partner_id", "=", vendor_id},
			[]any{"product_tmpl_id", "=", product_id},
			[]any{"product_uom_id", "=", uom_id},
			[]any{"min_qty", "=", p.MinQty},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		// if rid == -1 {
		// 	fmt.Println(prettyprint(p))
		// 	fmt.Println("ur", prettyprint(ur), "rid", rid)
		// }
		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
}
