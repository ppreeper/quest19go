package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type StockLocation_150 struct {
	Name         string `json:"name"`          // Location Name ⭐ Required
	CompleteName string `json:"complete_name"` // Full Location Name
	Usage        any    `json:"usage"`         // Location Type ⭐ Required
	LocationID   any    `json:"location_id"`   // Parent Location 📦 Model Relation: stock.location
	// ChildIDs                 any    `json:"child_ids"`                   // Contains 📦 Model Relation: stock.location
	// ChildInternalLocationIDs any    `json:"child_internal_location_ids"` // Internal locations amoung descendants 📦 Model Relation: stock.location
	Comment string `json:"comment"` // Additional Information
	// Posx                     int    `json:"posx"`                        // Corridor (X)
	// Posy                     int    `json:"posy"`                        // Shelves (Y)
	// Posz                     int    `json:"posz"`                        // Height (Z)
	ParentPath     string `json:"parent_path"`     // Parent Path
	CompanyID      any    `json:"company_id"`      // Company 📦 Model Relation: res.company
	ScrapLocation  bool   `json:"scrap_location"`  // Is a Scrap Location?
	ReturnLocation bool   `json:"return_location"` // Is a Return Location?
	// RemovalStrategyID any    `json:"removal_strategy_id"` // Removal Strategy 📦 Model Relation: product.removal
	// PutawayRuleIDs    any    `json:"putaway_rule_ids"`    // Putaway Rules 📦 Model Relation: stock.putaway.rule
	Barcode string `json:"barcode"` // Barcode
	// QuantIDs                 any    `json:"quant_ids"`                  // Quant 📦 Model Relation: stock.quant
	CyclicInventoryFrequency int `json:"cyclic_inventory_frequency"` // Inventory Frequency (Days)
	// LastInventoryDate        any    `json:"last_inventory_date"`        // Last Effective Inventory
	// NextInventoryDate        any    `json:"next_inventory_date"`        // Next Expected Inventory
	WarehouseViewIDs  any `json:"warehouse_view_ids"`  // Warehouse View 📦 Model Relation: stock.warehouse
	WarehouseID       any `json:"warehouse_id"`        // Warehouse 📦 Model Relation: stock.warehouse
	StorageCategoryID any `json:"storage_category_id"` // Storage Category 📦 Model Relation: stock.storage.category
	// OutgoingMoveLineIDs      any    `json:"outgoing_move_line_ids"`     // Outgoing Move Line 📦 Model Relation: stock.move.line
	// IncomingMoveLineIDs      any    `json:"incoming_move_line_ids"`     // Incoming Move Line 📦 Model Relation: stock.move.line
	// NetWeight                any    `json:"net_weight"`                 // Net Weight
	// ForecastWeight           any    `json:"forecast_weight"`            // Forecasted Weight
	ID          int    `json:"id"`           // ID
	DisplayName string `json:"display_name"` // Display Name
	// ValuationInAccountID  any    `json:"valuation_in_account_id"`  // Stock Valuation Account (Incoming) 📦 Model Relation: account.account
	// ValuationOutAccountID any    `json:"valuation_out_account_id"` // Stock Valuation Account (Outgoing) 📦 Model Relation: account.account
	// IsInventoryAdmin      bool   `json:"is_inventory_admin"`       // Is Inventory Config Admin
}

func (m *Model) StockLocation() {
	model := "stock.location"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(StockLocation_150{})

	// Root Locations
	fmt.Println("Fetching Root Locations...")
	root_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"name", "in", []string{"EWH", "CWH", "TDEWH", "CCWH", "PCLWH", "WWH"}},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println("root_records", len(root_records))

	// Root Locations
	bar := progressbar.Default(-1, "root locations")
	for _, r := range root_records {
		var root StockLocation_150
		FillStruct(r, &root)
		// fmt.Println(prettyprint(root))

		ur := map[string]any{
			"name":    root.Name,
			"barcode": odoo_barcode(root.Name),
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", root.Name},
			[]any{"complete_name", "=", odoo_pathjoin(root.Name)},
		})

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	// Layer 1 Locations
	fmt.Println("Fetching Layer 1 Locations...")
	layer1_records, _ := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"location_id", "in", ExtractIDs(root_records)},
	})
	fmt.Println("layer1_records", len(layer1_records))
	for _, r := range root_records {
		var root StockLocation_150
		FillStruct(r, &root)

		bar := progressbar.Default(-1, "root location "+root.Name+"\tprocessing layer 1")
		for _, r := range layer1_records {
			var layer1 StockLocation_150
			FillStruct(r, &layer1)
			// fmt.Println(prettyprint(layer1.CompleteName))

			_, layer1_location_name := ParseMany2One(layer1.LocationID)
			if root.CompleteName != layer1_location_name {
				continue
			}
			// m.Log.Info(model, "Root Location", root.Name, "Layer1", layer1.Name)

			dest_location_id, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", root.Name},
				[]any{"complete_name", "=", root.Name},
			})

			ur := map[string]any{
				"name":        layer1.Name,
				"location_id": dest_location_id,
				"usage":       "internal",
				"barcode":     odoo_barcode(root.Name, layer1.Name),
			}

			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", layer1.Name},
				[]any{"usage", "=", "internal"},
				[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name)},
			})

			// m.Log.Info(model, "rid", rid, "ur", ur)

			m.writeRecord(model, ur, rid)
			bar.Add(1)
		}
		bar.Finish()
	}

	// // Layer 2 Locations
	fmt.Println("Fetching Layer 2 Locations...")
	layer2_records, _ := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"location_id", "in", ExtractIDs(layer1_records)},
	})
	fmt.Println("layer2_records", len(layer2_records))
	for _, r := range root_records {
		var root StockLocation_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root location "+root.Name+"\tprocessing layer 2")

		for _, r := range layer1_records {
			var layer1 StockLocation_150
			FillStruct(r, &layer1)

			_, layer1_location_name := ParseMany2One(layer1.LocationID)
			if root.Name != layer1_location_name {
				continue
			}

			for _, r := range layer2_records {
				var layer2 StockLocation_150
				FillStruct(r, &layer2)
				// fmt.Println(prettyprint(layer2.CompleteName))

				_, layer2_location_name := ParseMany2One(layer2.LocationID)
				if layer1.CompleteName != layer2_location_name {
					continue
				}
				// m.Log.Info(model, "Root Location", root.Name, "Layer1", layer1.Name, "Layer2", layer2.Name)

				dest_location_id, _ := m.Dest.GetID(model, []any{
					[]any{"name", "=", layer1.Name},
					[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name)},
				})

				ur := map[string]any{
					"name":        layer2.Name,
					"location_id": dest_location_id,
					"usage":       "internal",
					"barcode":     odoo_barcode(root.Name, layer1.Name, layer2.Name),
				}

				rid, _ := m.Dest.GetID(model, []any{
					[]any{"name", "=", layer2.Name},
					[]any{"usage", "=", "internal"},
					[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name, layer2.Name)},
				})

				// m.Log.Info(model, "rid", rid, "ur", ur)

				m.writeRecord(model, ur, rid)
				bar.Add(1)
			}
		}
		bar.Finish()
	}

	// // Layer 3 Locations
	fmt.Println("Fetching Layer 3 Locations...")
	layer3_records, _ := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"location_id", "in", ExtractIDs(layer2_records)},
	})
	fmt.Println("layer3_records", len(layer3_records))
	for _, r := range root_records {
		var root StockLocation_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root location "+root.Name+"\tprocessing layer 3")

		for _, r := range layer1_records {
			var layer1 StockLocation_150
			FillStruct(r, &layer1)

			_, layer1_location_name := ParseMany2One(layer1.LocationID)
			if root.CompleteName != layer1_location_name {
				continue
			}

			for _, r := range layer2_records {
				var layer2 StockLocation_150
				FillStruct(r, &layer2)

				_, layer2_location_name := ParseMany2One(layer2.LocationID)
				if layer1.CompleteName != layer2_location_name {
					continue
				}

				for _, r := range layer3_records {
					var layer3 StockLocation_150
					FillStruct(r, &layer3)
					// fmt.Println(prettyprint(layer3.CompleteName))

					_, layer3_location_name := ParseMany2One(layer3.LocationID)
					if layer2.CompleteName != layer3_location_name {
						continue
					}
					// m.Log.Info(model, "Root Location", root.Name, "Layer1", layer1.Name, "Layer2", layer2.Name, "Layer3", layer3.Name)

					dest_location_id, _ := m.Dest.GetID(model, []any{
						[]any{"name", "=", layer2.Name},
						[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name, layer2.Name)},
					})

					ur := map[string]any{
						"name":        layer3.Name,
						"location_id": dest_location_id,
						"usage":       "internal",
						"barcode":     odoo_barcode(root.Name, layer1.Name, layer2.Name, layer3.Name),
					}

					rid, _ := m.Dest.GetID(model, []any{
						[]any{"name", "=", layer3.Name},
						[]any{"usage", "=", "internal"},
						[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name, layer2.Name, layer3.Name)},
					})

					// m.Log.Info(model, "rid", rid, "ur", ur)

					m.writeRecord(model, ur, rid)
					bar.Add(1)
				}
			}
		}
		bar.Finish()
	}

	// // Layer 4 Locations
	fmt.Println("Fetching Layer 4 Locations...")
	layer4_records, _ := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"location_id", "in", ExtractIDs(layer3_records)},
	})
	fmt.Println("layer4_records", len(layer4_records))
	for _, r := range root_records {
		var root StockLocation_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root location "+root.Name+"\tprocessing layer 4")

		for _, r := range layer1_records {
			var layer1 StockLocation_150
			FillStruct(r, &layer1)

			_, layer1_location_name := ParseMany2One(layer1.LocationID)
			if root.CompleteName != layer1_location_name {
				continue
			}

			for _, r := range layer2_records {
				var layer2 StockLocation_150
				FillStruct(r, &layer2)

				_, layer2_location_name := ParseMany2One(layer2.LocationID)
				if layer1.CompleteName != layer2_location_name {
					continue
				}

				for _, r := range layer3_records {
					var layer3 StockLocation_150
					FillStruct(r, &layer3)

					_, layer3_location_name := ParseMany2One(layer3.LocationID)
					if layer2.CompleteName != layer3_location_name {
						continue
					}

					for _, r := range layer4_records {
						var layer4 StockLocation_150
						FillStruct(r, &layer4)
						// fmt.Println(prettyprint(layer4.CompleteName))

						_, layer4_location_name := ParseMany2One(layer4.LocationID)
						if layer3.CompleteName != layer4_location_name {
							continue
						}
						// m.Log.Info(model, "Root Location", root.Name, "Layer1", layer1.Name, "Layer2", layer2.Name, "Layer3", layer3.Name, "Layer4", layer4.Name)

						dest_location_id, _ := m.Dest.GetID(model, []any{
							[]any{"name", "=", layer3.Name},
							[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name, layer2.Name, layer3.Name)},
						})

						ur := map[string]any{
							"name":        layer4.Name,
							"location_id": dest_location_id,
							"usage":       "internal",
							"barcode":     odoo_barcode(root.Name, layer1.Name, layer2.Name, layer3.Name, layer4.Name),
						}

						rid, _ := m.Dest.GetID(model, []any{
							[]any{"name", "=", layer4.Name},
							[]any{"usage", "=", "internal"},
							[]any{"complete_name", "=", odoo_pathjoin(root.Name, layer1.Name, layer2.Name, layer3.Name, layer4.Name)},
						})

						// m.Log.Info(model, "rid", rid, "ur", ur)

						m.writeRecord(model, ur, rid)
						bar.Add(1)
					}
				}
			}
		}
		bar.Finish()
	}
}
