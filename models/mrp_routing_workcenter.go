package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type MrpRoutingWorkcenter_150 struct {
	Name string `json:"name"` // Operation ⭐ Required
	// Active                                      bool   `json:"active"`                                            // Active
	WorkcenterID any `json:"workcenter_id"` // Work Center 📦 Model Relation: mrp.workcenter ⭐ Required
	// Sequence                                    int    `json:"sequence"`                                          // Sequence
	BomID any `json:"bom_id"` // Bill of Material 📦 Model Relation: mrp.bom ⭐ Required
	// CompanyID                                   any    `json:"company_id"`                                        // Company 📦 Model Relation: res.company
	// WorksheetType                               any    `json:"worksheet_type"`                                    // Work Sheet
	// Note                                        string `json:"note"`                                              // Description
	// Worksheet                                   any    `json:"worksheet"`                                         // PDF
	// WorksheetGoogleSlide                        string `json:"worksheet_google_slide"`                            // Google Slide
	TimeMode        any    `json:"time_mode"`         // Duration Computation
	TimeModeBatch   int    `json:"time_mode_batch"`   // Based on
	TimeComputedOn  string `json:"time_computed_on"`  // Computed on last
	TimeCycleManual any    `json:"time_cycle_manual"` // Manual Duration
	TimeCycle       any    `json:"time_cycle"`        // Duration
	// WorkorderCount                              int    `json:"workorder_count"`                                   // # Work Orders
	// WorkorderIDs                                any    `json:"workorder_ids"`                                     // Work Orders 📦 Model Relation: mrp.workorder
	// PossibleBomProductTemplateAttributeValueIDs any    `json:"possible_bom_product_template_attribute_value_ids"` // Possible Product Template Attribute Value 📦 Model Relation: product.template.attribute.value
	// BomProductTemplateAttributeValueIDs         any    `json:"bom_product_template_attribute_value_ids"`          // Apply on Variants 📦 Model Relation: product.template.attribute.value
	ID                int    `json:"id"`                  // ID
	DisplayName       string `json:"display_name"`        // Display Name
	QualityPointIDs   any    `json:"quality_point_ids"`   // Quality Point 📦 Model Relation: quality.point
	QualityPointCount int    `json:"quality_point_count"` // Steps
}

func (m *Model) MRPRoutingWorkcenter() {
	model := "mrp.routing.workcenter"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MrpRoutingWorkcenter_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"bom_id.active", "=", true},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	bar := progressbar.Default(int64(len(records)), model)
	for _, r := range records {
		var record MrpRoutingWorkcenter_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record.BomID))

		//
		// BILL OF MATERIAL
		//

		old_bom_id, _ := ParseMany2One(record.BomID)
		old_bom_record, err := m.Source.SearchRead("mrp.bom", 0, 1, ExtractJSONTags(MrpBom_150{}), []any{
			[]any{"id", "=", old_bom_id},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if len(old_bom_record) != 1 {
			continue
		}
		var old_bom MrpBom_150
		FillStruct(old_bom_record[0], &old_bom)

		old_ptmpl_id, product_tmpl := ParseMany2One(old_bom.ProductTmplID)
		old_ptmpl, _ := m.Source.SearchRead("product.template", 0, 0, []string{"name"}, []any{
			[]any{"id", "=", old_ptmpl_id},
		})
		if len(old_ptmpl) != 1 {
			continue
		}
		product_tmpl_id, _ := m.Dest.GetID("product.template", []any{
			[]any{"name", "=", old_ptmpl[0]["name"]},
		})
		if product_tmpl_id == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("product.template not found for %s", product_tmpl))
			continue
		}

		bom, err := m.Dest.GetID("mrp.bom", []any{
			[]any{"product_tmpl_id", "=", product_tmpl_id},
			[]any{"code", "=", old_bom.Code},
			[]any{"type", "=", old_bom.Type},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		//
		// workcenter
		//

		_, workcenter_name := ParseMany2One(record.WorkcenterID)
		dest_workcenter_id, err := m.Dest.GetID("mrp.workcenter", []any{
			[]any{"name", "=", workcenter_name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		ur := map[string]any{
			"name":              record.Name,
			"bom_id":            bom,
			"workcenter_id":     dest_workcenter_id,
			"time_mode":         record.TimeMode,
			"time_cycle_manual": record.TimeCycleManual,
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", record.Name},
			[]any{"bom_id", "=", bom},
			[]any{"workcenter_id", "=", dest_workcenter_id},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}
