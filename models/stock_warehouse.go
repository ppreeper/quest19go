package models

import (
	"github.com/schollz/progressbar/v3"
)

type StockWarehouse_150 struct {
	Name                  string `json:"name"`                    // Warehouse ⭐ Required
	Active                bool   `json:"active"`                  // Active
	CompanyID             any    `json:"company_id"`              // Company 📦 Model Relation: res.company ⭐ Required
	PartnerID             any    `json:"partner_id"`              // Address 📦 Model Relation: res.partner
	ViewLocationID        any    `json:"view_location_id"`        // View Location 📦 Model Relation: stock.location ⭐ Required
	LotStockID            any    `json:"lot_stock_id"`            // Location Stock 📦 Model Relation: stock.location ⭐ Required
	Code                  string `json:"code"`                    // Short Name ⭐ Required
	RouteIDs              any    `json:"route_ids"`               // Routes 📦 Model Relation: stock.location.route
	ReceptionSteps        any    `json:"reception_steps"`         // Incoming Shipments ⭐ Required
	DeliverySteps         any    `json:"delivery_steps"`          // Outgoing Shipments ⭐ Required
	WhInputStockLocID     any    `json:"wh_input_stock_loc_id"`   // Input Location 📦 Model Relation: stock.location
	WhQcStockLocID        any    `json:"wh_qc_stock_loc_id"`      // Quality Control Location 📦 Model Relation: stock.location
	WhOutputStockLocID    any    `json:"wh_output_stock_loc_id"`  // Output Location 📦 Model Relation: stock.location
	WhPackStockLocID      any    `json:"wh_pack_stock_loc_id"`    // Packing Location 📦 Model Relation: stock.location
	MtoPullID             any    `json:"mto_pull_id"`             // MTO rule 📦 Model Relation: stock.rule
	PickTypeID            any    `json:"pick_type_id"`            // Pick Type 📦 Model Relation: stock.picking.type
	PackTypeID            any    `json:"pack_type_id"`            // Pack Type 📦 Model Relation: stock.picking.type
	OutTypeID             any    `json:"out_type_id"`             // Out Type 📦 Model Relation: stock.picking.type
	InTypeID              any    `json:"in_type_id"`              // In Type 📦 Model Relation: stock.picking.type
	IntTypeID             any    `json:"int_type_id"`             // Internal Type 📦 Model Relation: stock.picking.type
	ReturnTypeID          any    `json:"return_type_id"`          // Return Type 📦 Model Relation: stock.picking.type
	CrossdockRouteID      any    `json:"crossdock_route_id"`      // Crossdock Route 📦 Model Relation: stock.location.route
	ReceptionRouteID      any    `json:"reception_route_id"`      // Receipt Route 📦 Model Relation: stock.location.route
	DeliveryRouteID       any    `json:"delivery_route_id"`       // Delivery Route 📦 Model Relation: stock.location.route
	ResupplyWhIDs         any    `json:"resupply_wh_ids"`         // Resupply From 📦 Model Relation: stock.warehouse
	ResupplyRouteIDs      any    `json:"resupply_route_ids"`      // Resupply Routes 📦 Model Relation: stock.location.route
	Sequence              int    `json:"sequence"`                // Sequence
	ID                    int    `json:"id"`                      // ID
	LastUpdate            any    `json:"__last_update"`           // Last Modified on
	DisplayName           string `json:"display_name"`            // Display Name
	CreateUid             any    `json:"create_uid"`              // Created by 📦 Model Relation: res.users
	CreateDate            any    `json:"create_date"`             // Created on
	WriteUid              any    `json:"write_uid"`               // Last Updated by 📦 Model Relation: res.users
	WriteDate             any    `json:"write_date"`              // Last Updated on
	ManufactureToResupply bool   `json:"manufacture_to_resupply"` // Manufacture to Resupply
	ManufacturePullID     any    `json:"manufacture_pull_id"`     // Manufacture Rule 📦 Model Relation: stock.rule
	ManufactureMtoPullID  any    `json:"manufacture_mto_pull_id"` // Manufacture MTO Rule 📦 Model Relation: stock.rule
	PbmMtoPullID          any    `json:"pbm_mto_pull_id"`         // Picking Before Manufacturing MTO Rule 📦 Model Relation: stock.rule
	SamRuleID             any    `json:"sam_rule_id"`             // Stock After Manufacturing Rule 📦 Model Relation: stock.rule
	ManuTypeID            any    `json:"manu_type_id"`            // Manufacturing Operation Type 📦 Model Relation: stock.picking.type
	PbmTypeID             any    `json:"pbm_type_id"`             // Picking Before Manufacturing Operation Type 📦 Model Relation: stock.picking.type
	SamTypeID             any    `json:"sam_type_id"`             // Stock After Manufacturing Operation Type 📦 Model Relation: stock.picking.type
	ManufactureSteps      any    `json:"manufacture_steps"`       // Manufacture ⭐ Required
	PbmRouteID            any    `json:"pbm_route_id"`            // Picking Before Manufacturing Route 📦 Model Relation: stock.location.route
	PbmLocID              any    `json:"pbm_loc_id"`              // Picking before Manufacturing Location 📦 Model Relation: stock.location
	SamLocID              any    `json:"sam_loc_id"`              // Stock after Manufacturing Location 📦 Model Relation: stock.location
	BuyToResupply         bool   `json:"buy_to_resupply"`         // Buy to Resupply
	BuyPullID             any    `json:"buy_pull_id"`             // Buy rule 📦 Model Relation: stock.rule
	ReturnPartnerID       any    `json:"return_partner_id"`       // Return Address 📦 Model Relation: res.partner
}

func (m *Model) StockWarehouse() {
	model := "stock.warehouse"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(StockWarehouse_150{})
	domain := []any{}

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, domain)
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
		var w StockWarehouse_150
		FillStruct(r, &w)

		ur := map[string]any{
			"name":                    w.Name,
			"code":                    w.Code,
			"reception_steps":         w.ReceptionSteps,
			"delivery_steps":          w.DeliverySteps,
			"manufacture_steps":       w.ManufactureSteps,
			"buy_to_resupply":         w.BuyToResupply,
			"manufacture_to_resupply": w.ManufactureToResupply,
			"resupply_wh_ids":         w.ResupplyWhIDs,
		}

		if w.ID == 1 {
			ur["code"] = "EWH"
			m.writeRecord(model, ur, 1)
		} else if w.Name == "Calgary" {
			ur["code"] = "CWH"
			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", w.Name},
			})
			m.writeRecord(model, ur, rid)
		} else {
			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", w.Name},
			})
			m.writeRecord(model, ur, rid)
		}

		bar.Add(1)
	}
	bar.Finish()
}

// func (m *Model) GetWarehousesSource() ([]map[string]any, error) {
// 	model := "stock.warehouse"
// 	return m.Source.SearchRead(model, 0, 0, []string{"id", "name"}, []any{})
// }

// func (m *Model) GetWarehousesDest() ([]map[string]any, error) {
// 	model := "stock.warehouse"
// 	return m.Dest.SearchRead(model, 0, 0, []string{"id", "name"}, []any{})
// }
