package models

import (
	"fmt"

	"github.com/ppreeper/str"
	"github.com/schollz/progressbar/v3"
)

type ProductCategory_150 struct {
	Name                                        string `json:"name"`                                             // Name ⭐ Required
	CompleteName                                string `json:"complete_name"`                                    // Complete Name
	ParentID                                    any    `json:"parent_id"`                                        // Parent Category 📦 Model Relation: product.category
	ParentPath                                  string `json:"parent_path"`                                      // Parent Path
	ChildID                                     any    `json:"child_id"`                                         // Child Categories 📦 Model Relation: product.category
	ProductCount                                int    `json:"product_count"`                                    // # Products
	ID                                          int    `json:"id"`                                               // ID
	LastUpdate                                  any    `json:"__last_update"`                                    // Last Modified on
	DisplayName                                 string `json:"display_name"`                                     // Display Name
	CreateUid                                   any    `json:"create_uid"`                                       // Created by 📦 Model Relation: res.users
	CreateDate                                  any    `json:"create_date"`                                      // Created on
	WriteUid                                    any    `json:"write_uid"`                                        // Last Updated by 📦 Model Relation: res.users
	WriteDate                                   any    `json:"write_date"`                                       // Last Updated on
	PropertyAccountIncomeCategID                any    `json:"property_account_income_categ_id"`                 // Income Account 📦 Model Relation: account.account
	PropertyAccountExpenseCategID               any    `json:"property_account_expense_categ_id"`                // Expense Account 📦 Model Relation: account.account
	RouteIDs                                    any    `json:"route_ids"`                                        // Routes 📦 Model Relation: stock.location.route
	RemovalStrategyID                           any    `json:"removal_strategy_id"`                              // Force Removal Strategy 📦 Model Relation: product.removal
	TotalRouteIDs                               any    `json:"total_route_ids"`                                  // Total routes 📦 Model Relation: stock.location.route
	PutawayRuleIDs                              any    `json:"putaway_rule_ids"`                                 // Putaway Rules 📦 Model Relation: stock.putaway.rule
	PackagingReserveMethod                      any    `json:"packaging_reserve_method"`                         // Reserve Packagings
	PropertyAccountCreditorPriceDifferenceCateg any    `json:"property_account_creditor_price_difference_categ"` // Price Difference Account 📦 Model Relation: account.account
	PropertyValuation                           any    `json:"property_valuation"`                               // Inventory Valuation ⭐ Required
	PropertyCostMethod                          any    `json:"property_cost_method"`                             // Costing Method ⭐ Required
	PropertyStockJournal                        any    `json:"property_stock_journal"`                           // Stock Journal 📦 Model Relation: account.journal
	PropertyStockAccountInputCategID            any    `json:"property_stock_account_input_categ_id"`            // Stock Input Account 📦 Model Relation: account.account
	PropertyStockAccountOutputCategID           any    `json:"property_stock_account_output_categ_id"`           // Stock Output Account 📦 Model Relation: account.account
	PropertyStockValuationAccountID             any    `json:"property_stock_valuation_account_id"`              // Stock Valuation Account 📦 Model Relation: account.account
}

func (m *Model) ProductCategories() {
	model := "product.category"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())
	sourceFields := ExtractJSONTags(ProductCategory_150{})

	// Root Locations
	root_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"parent_id", "=", false},
		[]any{"name", "not ilike", "%DON'T USE%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println("root_records", len(root_records))
	root_records_ids := ExtractIDs(root_records)
	fmt.Println("root_records_ids", len(root_records_ids))

	layer1_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"parent_id", "in", root_records_ids},
		[]any{"name", "not ilike", "%DON'T USE%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println("layer1_records", len(layer1_records))
	layer1_records_ids := ExtractIDs(layer1_records)
	fmt.Println("layer1_records_ids", len(layer1_records_ids))

	layer2_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"parent_id", "in", layer1_records_ids},
		[]any{"name", "not ilike", "%DON'T USE%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println("layer2_records", len(layer2_records))
	layer2_records_ids := ExtractIDs(layer2_records)
	fmt.Println("layer2_records_ids", len(layer2_records_ids))

	layer3_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"parent_id", "in", layer2_records_ids},
		[]any{"name", "not ilike", "%DON'T USE%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println("layer3_records", len(layer3_records))
	layer3_records_ids := ExtractIDs(layer3_records)
	fmt.Println("layer3_records_ids", len(layer3_records_ids))

	// layer4_records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
	// 	[]any{"parent_id", "in", layer3_records_ids},
	// 	[]any{"name", "not ilike", "%DON'T USE%"},
	// })
	// if err != nil {
	// 	m.Log.Error(model, "func", trace(), "err", err)
	// 	return
	// }
	// fmt.Println("layer4_records", len(layer4_records))
	// layer4_records_ids := ExtractIDs(layer4_records)
	// fmt.Println("layer4_records_ids", len(layer4_records_ids))

	max_len := 0
	for _, r := range root_records {
		var root ProductCategory_150
		FillStruct(r, &root)
		namelen := len(root.Name)
		if namelen > max_len {
			max_len = namelen
		}
	}
	max_len += 2

	// Inventory Valuation v15
	// "Manual": "manual_periodic"
	// "Automated": "real_time"

	// Inventory Valuation v19
	// "Periodic (at closing)": "periodic"
	// "Perpetual (at invoicing)": "real_time"

	// Root Categories
	banner.Println(model, "root categories")
	bar := progressbar.Default(-1, "root categories")
	for _, r := range root_records {
		var root ProductCategory_150
		FillStruct(r, &root)
		// fmt.Println(prettyprint(root))
		// m.Log.Info(model, "root category", str.LJustLen(root.Name, max_len))

		property_valuation := "periodic"
		if root.PropertyValuation == "real_time" {
			property_valuation = "real_time"
		}
		_, property_account_income_categ := ParseMany2One(root.PropertyAccountIncomeCategID)
		_, property_account_expense_categ := ParseMany2One(root.PropertyAccountExpenseCategID)

		income_account_id, _ := m.Dest.GetID("account.account", []any{
			[]any{"display_name", "=", property_account_income_categ},
		})

		expense_account_id, _ := m.Dest.GetID("account.account", []any{
			[]any{"display_name", "=", property_account_expense_categ},
		})

		ur := map[string]any{
			"name":                              root.Name,
			"property_cost_method":              root.PropertyCostMethod,
			"property_valuation":                property_valuation,
			"property_account_income_categ_id":  income_account_id,
			"property_account_expense_categ_id": expense_account_id,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", root.Name},
			[]any{"complete_name", "=", root.Name},
		})

		// m.Log.Info(model, "rid", rid, "ur", ur)

		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()

	// Layer 1 Categories
	banner.Println(model, "layer 1 categories")
	for _, r := range root_records {
		var root ProductCategory_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root category "+str.LJustLen(root.Name, max_len)+" processing layer 1")
		for _, r := range layer1_records {
			var layer1 ProductCategory_150
			FillStruct(r, &layer1)
			_, layer1_parent_name := ParseMany2One(layer1.ParentID)
			if root.CompleteName != layer1_parent_name {
				continue
			}
			dest_parent_id, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", root.Name},
				[]any{"complete_name", "=", root.Name},
			})
			property_valuation := "periodic"
			if layer1.PropertyValuation == "real_time" {
				property_valuation = "real_time"
			}
			_, property_account_income_categ := ParseMany2One(layer1.PropertyAccountIncomeCategID)
			_, property_account_expense_categ := ParseMany2One(layer1.PropertyAccountExpenseCategID)

			income_account_id, _ := m.Dest.GetID("account.account", []any{
				[]any{"display_name", "=", property_account_income_categ},
			})

			expense_account_id, _ := m.Dest.GetID("account.account", []any{
				[]any{"display_name", "=", property_account_expense_categ},
			})
			ur := map[string]any{
				"name":                              layer1.Name,
				"parent_id":                         dest_parent_id,
				"property_cost_method":              layer1.PropertyCostMethod,
				"property_valuation":                property_valuation,
				"property_account_income_categ_id":  income_account_id,
				"property_account_expense_categ_id": expense_account_id,
			}
			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", layer1.Name},
				[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name)},
				[]any{"parent_id", "=", dest_parent_id},
			})
			// m.Log.Info(model, "rid", rid, "ur", ur)
			m.writeRecord(model, ur, rid)
			bar.Add(1)
		}
		bar.Finish()
	}
	// Layer 2 Categories
	banner.Println(model, "layer 2 categories")
	for _, r := range root_records {
		var root ProductCategory_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root category "+str.LJustLen(root.Name, max_len)+" processing layer 2")

		for _, r := range layer1_records {
			var layer1 ProductCategory_150
			FillStruct(r, &layer1)
			_, layer1_parent_name := ParseMany2One(layer1.ParentID)
			if root.CompleteName != layer1_parent_name {
				continue
			}
			for _, r := range layer2_records {
				var layer2 ProductCategory_150
				FillStruct(r, &layer2)
				_, layer2_parent_name := ParseMany2One(layer2.ParentID)
				if layer1.CompleteName != layer2_parent_name {
					continue
				}

				dest_parent_id, _ := m.Dest.GetID(model, []any{
					[]any{"name", "=", layer1.Name},
					[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name)},
				})
				property_valuation := "periodic"
				if layer2.PropertyValuation == "real_time" {
					property_valuation = "real_time"
				}
				_, property_account_income_categ := ParseMany2One(layer2.PropertyAccountIncomeCategID)
				_, property_account_expense_categ := ParseMany2One(layer2.PropertyAccountExpenseCategID)

				income_account_id, _ := m.Dest.GetID("account.account", []any{
					[]any{"display_name", "=", property_account_income_categ},
				})

				expense_account_id, _ := m.Dest.GetID("account.account", []any{
					[]any{"display_name", "=", property_account_expense_categ},
				})
				ur := map[string]any{
					"name":                              layer2.Name,
					"parent_id":                         dest_parent_id,
					"property_cost_method":              layer2.PropertyCostMethod,
					"property_valuation":                property_valuation,
					"property_account_income_categ_id":  income_account_id,
					"property_account_expense_categ_id": expense_account_id,
				}
				rid, _ := m.Dest.GetID(model, []any{
					[]any{"name", "=", layer2.Name},
					[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name, layer2.Name)},
					[]any{"parent_id", "=", dest_parent_id},
				})
				// m.Log.Info(model, "rid", rid, "ur", ur)

				m.writeRecord(model, ur, rid)
				bar.Add(1)
			}
		}
		bar.Finish()
	}
	// Layer 3 Categories
	banner.Println(model, "layer 3 categories")
	for _, r := range root_records {
		var root ProductCategory_150
		FillStruct(r, &root)
		bar := progressbar.Default(-1, "root category "+str.LJustLen(root.Name, max_len)+" processing layer 3")

		for _, r := range layer1_records {
			var layer1 ProductCategory_150
			FillStruct(r, &layer1)
			_, layer1_parent_name := ParseMany2One(layer1.ParentID)
			if root.CompleteName != layer1_parent_name {
				continue
			}
			for _, r := range layer2_records {
				var layer2 ProductCategory_150
				FillStruct(r, &layer2)
				_, layer2_parent_name := ParseMany2One(layer2.ParentID)
				if layer1.CompleteName != layer2_parent_name {
					continue
				}

				for _, r := range layer3_records {
					var layer3 ProductCategory_150
					FillStruct(r, &layer3)
					_, layer3_parent_name := ParseMany2One(layer3.ParentID)
					if layer2.CompleteName != layer3_parent_name {
						continue
					}
					_, property_account_income_categ := ParseMany2One(layer2.PropertyAccountIncomeCategID)
					_, property_account_expense_categ := ParseMany2One(layer2.PropertyAccountExpenseCategID)

					income_account_id, _ := m.Dest.GetID("account.account", []any{
						[]any{"display_name", "=", property_account_income_categ},
					})

					expense_account_id, _ := m.Dest.GetID("account.account", []any{
						[]any{"display_name", "=", property_account_expense_categ},
					})

					dest_parent_id, _ := m.Dest.GetID(model, []any{
						[]any{"name", "=", layer2.Name},
						[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name, layer2.Name)},
					})
					property_valuation := "periodic"
					if layer3.PropertyValuation == "real_time" {
						property_valuation = "real_time"
					}
					ur := map[string]any{
						"name":                              layer3.Name,
						"parent_id":                         dest_parent_id,
						"property_cost_method":              layer3.PropertyCostMethod,
						"property_valuation":                property_valuation,
						"property_account_income_categ_id":  income_account_id,
						"property_account_expense_categ_id": expense_account_id,
					}
					rid, _ := m.Dest.GetID(model, []any{
						[]any{"name", "=", layer3.Name},
						[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name, layer2.Name, layer3.Name)},
						[]any{"parent_id", "=", dest_parent_id},
					})
					// m.Log.Info(model, "rid", rid, "ur", ur)

					m.writeRecord(model, ur, rid)
					bar.Add(1)
				}
			}
		}
		bar.Finish()
	}
	// Layer 4 Categories
	// banner.Println(model, "layer 4 categories")
	// for _, r := range root_records {
	// 	var root ProductCategory
	// 	FillStruct(r, &root)
	// 	bar := progressbar.Default(-1, "root category "+str.LJustLen(root.Name, max_len)+" processing layer 4")

	// 	for _, r := range layer1_records {
	// 		var layer1 ProductCategory
	// 		FillStruct(r, &layer1)
	// 		_, layer1_parent_name := ParseMany2One(layer1.ParentID)
	// 		if root.CompleteName != layer1_parent_name {
	// 			continue
	// 		}
	// 		for _, r := range layer2_records {
	// 			var layer2 ProductCategory
	// 			FillStruct(r, &layer2)
	// 			_, layer2_parent_name := ParseMany2One(layer2.ParentID)
	// 			if layer1.CompleteName != layer2_parent_name {
	// 				continue
	// 			}

	// 			for _, r := range layer3_records {
	// 				var layer3 ProductCategory
	// 				FillStruct(r, &layer3)
	// 				_, layer3_parent_name := ParseMany2One(layer3.ParentID)
	// 				if layer2.CompleteName != layer3_parent_name {
	// 					continue
	// 				}

	// 				for _, r := range layer4_records {
	// 					var layer4 ProductCategory
	// 					FillStruct(r, &layer4)
	// 					_, layer4_parent_name := ParseMany2One(layer4.ParentID)
	// 					if layer3.CompleteName != layer4_parent_name {
	// 						continue
	// 					}

	// 					dest_parent_id, _ := m.Dest.GetID(model, []any{
	// 						[]any{"name", "=", layer3.Name},
	// 						[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name, layer2.Name, layer3.Name)},
	// 					})
	// 					ur := map[string]any{
	// 						"name":                 layer4.Name,
	// 						"parent_id":            dest_parent_id,
	// 						"property_cost_method": layer4.PropertyCostMethod,
	// 					}
	// 					rid, _ := m.Dest.GetID(model, []any{
	// 						[]any{"name", "=", layer4.Name},
	// 						[]any{"complete_name", "=", odoo_path(root.Name, layer1.Name, layer2.Name, layer3.Name, layer4.Name)},
	// 					})
	// 					m.Log.Info(model, "rid", rid, "ur", ur)

	// 					m.writeRecord(model, ur, rid)
	// 					bar.Add(1)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	bar.Finish()
	// }
}
