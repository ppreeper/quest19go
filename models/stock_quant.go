package models

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

// > odooquery quest15data.odoopro.ca stock.quant -f
// ,,,,,,,,owner_id
// -d "[('location_id','not like','Virtual Locations'),('location_id','not like','Partner Locations')]" -c
// records: 10775

type StockQuant_150 struct {
	// LastUpdate                 time.Time `json:"__last_update"`                 // Last Modified on
	// AccountingDate             time.Time `json:"accounting_date"`               // Accounting Date
	// AvailableQuantity float64 `json:"available_quantity"` // Available Quantity
	CompanyID any `json:"company_id"` // Company 📦 relation: many2one res.company
	// CurrencyID  any    `json:"currency_id"`  // Currency 📦 relation: many2one res.currency
	// DisplayName string `json:"display_name"` // Display Name
	// DummyID                    string    `json:"dummy_id"`                      // Dummy
	ID int `json:"id"` // ID
	// InDate                     time.Time `json:"in_date"`                       // Incoming Date ⭐ required
	// InventoryDate              time.Time `json:"inventory_date"`                // Scheduled Date
	InventoryDiffQuantity float64 `json:"inventory_diff_quantity"` // Difference
	InventoryQuantity     float64 `json:"inventory_quantity"`      // Counted Quantity
	// InventoryQuantityAutoApply float64   `json:"inventory_quantity_auto_apply"` // Inventoried Quantity
	InventoryQuantitySet bool `json:"inventory_quantity_set"` // Inventory Quantity Set
	// IsAdjustmentManager        bool      `json:"is_adjustment_manager"`         // Is Adjustment Manager
	// IsOutdated                 bool      `json:"is_outdated"`                   // Quantity has been moved since last count
	// LargeDifference bool `json:"large_difference"` // Large Difference
	LocationID any `json:"location_id"` // Location 📦 relation: many2one stock.location ⭐ required
	LotID      any `json:"lot_id"`      // Lot/Serial Number 📦 relation: many2one stock.production.lot
	// OnHand           bool    `json:"on_hand"`           // On Hand
	OwnerID   any `json:"owner_id"`   // Owner 📦 relation: many2one res.partner
	PackageID any `json:"package_id"` // Package 📦 relation: many2one stock.quant.package
	// ProductCategID   any     `json:"product_categ_id"`  // Product Category 📦 relation: many2one product.category
	ProductID     any `json:"product_id"`      // Product 📦 relation: many2one product.product ⭐ required
	ProductTmplID any `json:"product_tmpl_id"` // Product Template 📦 relation: many2one product.template
	// ProductUomID     any     `json:"product_uom_id"`    // Unit of Measure 📦 relation: many2one uom.uom
	Quantity         float64 `json:"quantity"`          // Quantity
	ReservedQuantity float64 `json:"reserved_quantity"` // Reserved Quantity ⭐ required
	// Tracking         any     `json:"tracking"`          // Tracking
	// UserID           any     `json:"user_id"`           // Assigned To 📦 relation: many2one res.users
	// Value float64 `json:"value"` // Value
}

// quest19data stock.quant model
type StockQuant_190 struct {
	AccountingDate             time.Time `json:"accounting_date"`               // Accounting Date
	AvailableQuantity          float64   `json:"available_quantity"`            // Available Quantity
	CompanyID                  any       `json:"company_id"`                    // Company 📦 relation: many2one res.company
	CostMethod                 any       `json:"cost_method"`                   // Cost Method
	CurrencyID                 any       `json:"currency_id"`                   // Currency 📦 relation: many2one res.currency
	CyclicInventoryFrequency   int       `json:"cyclic_inventory_frequency"`    // Inventory Frequency
	DisplayName                string    `json:"display_name"`                  // Display Name
	DummyID                    string    `json:"dummy_id"`                      // Dummy
	ID                         int       `json:"id"`                            // ID
	Image1920                  any       `json:"image_1920"`                    // Image
	InDate                     time.Time `json:"in_date"`                       // Incoming Date ⭐ required
	InventoryDate              time.Time `json:"inventory_date"`                // Scheduled
	InventoryDiffQuantity      float64   `json:"inventory_diff_quantity"`       // Difference
	InventoryQuantity          float64   `json:"inventory_quantity"`            // Counted
	InventoryQuantityAutoApply float64   `json:"inventory_quantity_auto_apply"` // Inventoried Quantity
	InventoryQuantitySet       bool      `json:"inventory_quantity_set"`        // Inventory Quantity Set
	IsFavorite                 bool      `json:"is_favorite"`                   // Favorite
	IsOutdated                 bool      `json:"is_outdated"`                   // Quantity has been moved since last count
	LastCountDate              time.Time `json:"last_count_date"`               // Last Count Date
	LocationID                 any       `json:"location_id"`                   // Location 📦 relation: many2one stock.location ⭐ required
	LotID                      any       `json:"lot_id"`                        // Lot/Serial Number 📦 relation: many2one stock.lot
	LotProperties              any       `json:"lot_properties"`                // Properties
	OnHand                     bool      `json:"on_hand"`                       // On Hand
	OwnerID                    any       `json:"owner_id"`                      // Owner 📦 relation: many2one res.partner
	PackageID                  any       `json:"package_id"`                    // Package 📦 relation: many2one stock.package
	ProductCategID             any       `json:"product_categ_id"`              // Product Category 📦 relation: many2one product.category
	ProductID                  any       `json:"product_id"`                    // Product 📦 relation: many2one product.product ⭐ required
	ProductReferenceCode       string    `json:"product_reference_code"`        // Product Reference Code
	ProductTmplID              any       `json:"product_tmpl_id"`               // Product Template 📦 relation: many2one product.template
	ProductUomID               any       `json:"product_uom_id"`                // Unit 📦 relation: many2one uom.uom
	Quantity                   float64   `json:"quantity"`                      // Quantity
	ReservedQuantity           float64   `json:"reserved_quantity"`             // Reserved Quantity ⭐ required
	SnDuplicated               bool      `json:"sn_duplicated"`                 // Duplicated Serial Number
	StorageCategoryID          any       `json:"storage_category_id"`           // Storage Category 📦 relation: many2one stock.storage.category
	Tracking                   any       `json:"tracking"`                      // Tracking
	UserID                     any       `json:"user_id"`                       // Assigned To 📦 relation: many2one res.users
	Value                      float64   `json:"value"`                         // Value
	WarehouseID                any       `json:"warehouse_id"`                  // Warehouse 📦 relation: many2one stock.warehouse
}

func (m *Model) StockQuant() {
	model := "stock.quant"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())
	records, err := m.Source.SearchRead(model, 0, 10, ExtractJSONTags(StockQuant_150{}), []any{
		[]any{"location_id", "not like", "Virtual Locations"},
		[]any{"location_id", "not like", "Partner Locations"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}
	bar := progressbar.Default(-1, "quants")
	for _, record := range records {
		var q StockQuant_150
		FillStruct(record, &q)
		// fmt.Println(prettyprint(q))

		ur := make(map[string]any)
		product_tmpl_id := m.GetDestProductTemplate(ParseMany2One(q.ProductTmplID))
		ur["product_tmpl_id"] = product_tmpl_id
		location_id := m.GetDestStockLocation(ParseMany2One(q.LocationID))
		ur["location_id"] = location_id
		ur["quantity"] = q.Quantity
		// ur["reserved_quantity"] = q.ReservedQuantity
		// ur["inventory_quantity_set"] = q.InventoryQuantitySet
		// ur["inventory_quantity"] = q.InventoryQuantity
		// ur["inventory_diff_quantity"] = q.InventoryDiffQuantity

		// fmt.Println(prettyprint(ur))

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"product_tmpl_id", "=", product_tmpl_id},
			[]any{"location_id", "=", location_id},
		})
		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
	fmt.Println("done", len(records))
}

// !Partner Locations 53517
// !Virtual Locations 44320

// EWH 9067
// CWH 1366
// CCWH 67
// PCLWH 139
// TDEWH 348
// WWH 68
// Physical Locations 135
