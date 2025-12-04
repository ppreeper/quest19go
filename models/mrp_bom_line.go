package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type MrpBomLine_150 struct {
	ProductID                                   any `json:"product_id"`                                        // Component 📦 Model Relation: product.product ⭐ Required
	ProductTmplID                               any `json:"product_tmpl_id"`                                   // Product Template 📦 Model Relation: product.template
	CompanyID                                   any `json:"company_id"`                                        // Company 📦 Model Relation: res.company
	ProductQty                                  any `json:"product_qty"`                                       // Quantity ⭐ Required
	ProductUomID                                any `json:"product_uom_id"`                                    // Product Unit of Measure 📦 Model Relation: uom.uom ⭐ Required
	ProductUomCategoryID                        any `json:"product_uom_category_id"`                           // Category 📦 Model Relation: uom.category
	Sequence                                    int `json:"sequence"`                                          // Sequence
	BomID                                       any `json:"bom_id"`                                            // Parent BoM 📦 Model Relation: mrp.bom ⭐ Required
	ParentProductTmplID                         any `json:"parent_product_tmpl_id"`                            // Parent Product Template 📦 Model Relation: product.template
	PossibleBomProductTemplateAttributeValueIDs any `json:"possible_bom_product_template_attribute_value_ids"` // Possible Product Template Attribute Value 📦 Model Relation: product.template.attribute.value
	BomProductTemplateAttributeValueIDs         any `json:"bom_product_template_attribute_value_ids"`          // Apply on Variants 📦 Model Relation: product.template.attribute.value
	AllowedOperationIDs                         any `json:"allowed_operation_ids"`                             // Operations 📦 Model Relation: mrp.routing.workcenter
	OperationID                                 any `json:"operation_id"`                                      // Consumed in Operation 📦 Model Relation: mrp.routing.workcenter
	ChildBomID                                  any `json:"child_bom_id"`                                      // Sub BoM 📦 Model Relation: mrp.bom
	ChildLineIDs                                any `json:"child_line_ids"`                                    // BOM lines of the referred bom 📦 Model Relation: mrp.bom.line
	AttachmentsCount                            int `json:"attachments_count"`                                 // Attachments Count
	ID                                          int `json:"id"`                                                // ID
	// LastUpdate                                  any    `json:"__last_update"`                                     // Last Modified on
	DisplayName string `json:"display_name"` // Display Name
	// CreateUid                                   any    `json:"create_uid"`                                        // Created by 📦 Model Relation: res.users
	// CreateDate                                  any    `json:"create_date"`                                       // Created on
	// WriteUid                                    any    `json:"write_uid"`                                         // Last Updated by 📦 Model Relation: res.users
	// WriteDate                                   any    `json:"write_date"`                                        // Last Updated on
}

func (m *Model) MRPBomLine() {
	model := "mrp.bom.line"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MrpBomLine_150{})

	// boms, _ := m.Source.SearchRead("mrp.bom", 0, 0, []string{"id"}, []any{
	// 	[]any{"display_name", "like", "[MTO] F:STUD.L7.118.12"},
	// })

	records, err := m.Source.SearchRead(model, 0, 5, sourceFields, []any{
		[]any{"bom_id.active", "=", true},
		// []any{"bom_id", "in", []int{16308, 22410, 22411, 22412, 22414, 22415, 22443, 22482, 22557, 22594}},
		// []any{"bom_id", "in", ExtractIDs(boms)},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	bar := progressbar.Default(int64(len(records)), model)
	for _, r := range records {
		var record MrpBomLine_150
		FillStruct(r, &record)
		// fmt.Println("----")
		// fmt.Println(
		// 	prettyprint(record.BomID),
		// 	prettyprint(record.ParentProductTmplID),
		// 	prettyprint(record.ProductID),
		// )

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
		// fmt.Println("BOM:", prettyprint(old_bom.ProductTmplID))

		old_ptmpl_id := m.GetDestProductTemplate(ParseMany2One(old_bom.ProductTmplID))
		// fmt.Println("old_ptmpl_id:", old_ptmpl_id, "code", old_bom.Code, "type", old_bom.Type)
		if old_ptmpl_id == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("product.template not found for %s", old_bom.ProductTmplID))
			continue
		}

		bom, err := m.Dest.GetID("mrp.bom", []any{
			[]any{"product_tmpl_id", "=", old_ptmpl_id},
			[]any{"code", "=", old_bom.Code},
			[]any{"type", "=", old_bom.Type},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if bom == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("mrp.bom not found for %v", old_bom))
			continue
		}
		// fmt.Println("DEST BOM ID:", bom)

		//
		// products
		//
		// fmt.Println("Component ProductID:", prettyprint(record.ProductID))
		dest_product_product_id := m.GetDestProductProduct(ParseMany2One(record.ProductID))
		if dest_product_product_id == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("product.product not found for %s", record.ProductID))
			continue
		}

		//
		// uom
		//

		_, uom_name := ParseMany2One(record.ProductUomID)
		dest_uom_id, err := m.Dest.GetID("uom.uom", []any{
			[]any{"name", "=", uom_name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if dest_uom_id == -1 {
			m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("uom.uom not found for %s", record.ProductUomID))
			continue
		}

		ur := map[string]any{
			"product_id":  dest_product_product_id,
			"bom_id":      bom,
			"product_qty": record.ProductQty,
		}

		if dest_uom_id != -1 {
			ur["product_uom_id"] = dest_uom_id
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"bom_id", "=", bom},
			[]any{"product_id", "=", dest_product_product_id},
			[]any{"product_qty", "=", record.ProductQty},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		// fmt.Println(rid, prettyprint(ur))

		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
}

func (m *Model) GetDestProductTemplate(product_id int, product_display_name string) int {
	old_product, _ := m.Source.SearchRead("product.template", 0, 0, []string{"name"}, []any{
		[]any{"id", "=", product_id},
	})
	if len(old_product) != 1 {
		return -1
	}
	product_name := old_product[0]["name"].(string)
	product_tmpl_id, _ := m.Dest.GetID("product.template", []any{
		[]any{"name", "=", product_name},
	})
	if product_tmpl_id == -1 {
		return -1
	}
	return product_tmpl_id
}

func (m *Model) GetDestProductProduct(product_id int, product_display_name string) int {
	// fmt.Println("GetDestProductProduct:", product_id, product_display_name)
	old_product, _ := m.Source.SearchRead("product.product", 0, 0, []string{"name"}, []any{
		[]any{"id", "=", product_id},
	})
	if len(old_product) != 1 {
		return -1
	}
	product_name := old_product[0]["name"].(string)
	product_product_id, _ := m.Dest.GetID("product.product", []any{
		[]any{"name", "=", product_name},
	})
	if product_product_id == -1 {
		return -1
	}
	return product_product_id
}
