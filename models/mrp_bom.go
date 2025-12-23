package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type MrpBom_150 struct {
	MessageIsFollower                        bool   `json:"message_is_follower"`                           // Is Follower
	MessageFollowerIDs                       any    `json:"message_follower_ids"`                          // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs                        any    `json:"message_partner_ids"`                           // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs                               any    `json:"message_ids"`                                   // Messages 📦 Model Relation: mail.message
	HasMessage                               bool   `json:"has_message"`                                   // Has Message
	MessageUnread                            bool   `json:"message_unread"`                                // Unread Messages
	MessageUnreadCounter                     int    `json:"message_unread_counter"`                        // Unread Messages Counter
	MessageNeedaction                        bool   `json:"message_needaction"`                            // Action Needed
	MessageNeedactionCounter                 int    `json:"message_needaction_counter"`                    // Number of Actions
	MessageHasError                          bool   `json:"message_has_error"`                             // Message Delivery error
	MessageHasErrorCounter                   int    `json:"message_has_error_counter"`                     // Number of errors
	MessageAttachmentCount                   int    `json:"message_attachment_count"`                      // Attachment Count
	MessageMainAttachmentID                  any    `json:"message_main_attachment_id"`                    // Main Attachment 📦 Model Relation: ir.attachment
	WebsiteMessageIDs                        any    `json:"website_message_ids"`                           // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError                       bool   `json:"message_has_sms_error"`                         // SMS Delivery error
	Code                                     string `json:"code"`                                          // Reference
	Active                                   bool   `json:"active"`                                        // Production Ready
	Type                                     any    `json:"type"`                                          // BoM Type ⭐ Required
	ProductTmplID                            any    `json:"product_tmpl_id"`                               // Product 📦 Model Relation: product.template ⭐ Required
	ProductID                                any    `json:"product_id"`                                    // Product Variant 📦 Model Relation: product.product
	BomLineIDs                               any    `json:"bom_line_ids"`                                  // BoM Lines 📦 Model Relation: mrp.bom.line
	ByproductIDs                             any    `json:"byproduct_ids"`                                 // By-products 📦 Model Relation: mrp.bom.byproduct
	ProductQty                               any    `json:"product_qty"`                                   // Quantity ⭐ Required
	ProductUomID                             any    `json:"product_uom_id"`                                // Unit of Measure 📦 Model Relation: uom.uom ⭐ Required
	ProductUomCategoryID                     any    `json:"product_uom_category_id"`                       // Category 📦 Model Relation: uom.category
	Sequence                                 int    `json:"sequence"`                                      // Sequence
	OperationIDs                             any    `json:"operation_ids"`                                 // Operations 📦 Model Relation: mrp.routing.workcenter
	ReadyToProduce                           any    `json:"ready_to_produce"`                              // Manufacturing Readiness ⭐ Required
	PickingTypeID                            any    `json:"picking_type_id"`                               // Operation Type 📦 Model Relation: stock.picking.type
	CompanyID                                any    `json:"company_id"`                                    // Company 📦 Model Relation: res.company
	Consumption                              any    `json:"consumption"`                                   // Flexible Consumption ⭐ Required
	PossibleProductTemplateAttributeValueIDs any    `json:"possible_product_template_attribute_value_ids"` // Possible Product Template Attribute Value 📦 Model Relation: product.template.attribute.value
	ID                                       int    `json:"id"`                                            // ID
	LastUpdate                               any    `json:"__last_update"`                                 // Last Modified on
	DisplayName                              string `json:"display_name"`                                  // Display Name
	CreateUid                                any    `json:"create_uid"`                                    // Created by 📦 Model Relation: res.users
	CreateDate                               any    `json:"create_date"`                                   // Created on
	WriteUid                                 any    `json:"write_uid"`                                     // Last Updated by 📦 Model Relation: res.users
	WriteDate                                any    `json:"write_date"`                                    // Last Updated on
	AnalyticAccountID                        any    `json:"analytic_account_id"`                           // Analytic Account 📦 Model Relation: account.analytic.account
	Version                                  int    `json:"version"`                                       // Version
	PreviousBomID                            any    `json:"previous_bom_id"`                               // Previous BoM 📦 Model Relation: mrp.bom
	Image128                                 any    `json:"image_128"`                                     // Image 128
	EcoIDs                                   any    `json:"eco_ids"`                                       // ECO to be applied 📦 Model Relation: mrp.eco
	EcoCount                                 int    `json:"eco_count"`                                     // # ECOs
	EcoInprogressCount                       int    `json:"eco_inprogress_count"`                          // # ECOs in progress
	DepartmentID                             any    `json:"department_id"`                                 // Department 📦 Model Relation: quest_bom_ext.department
}

func (m *Model) MRPBom() {
	model := "mrp.bom"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MrpBom_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
		// []any{"product_tmpl_id", "=", "SNC:B7.1.1112"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	bar := progressbar.Default(int64(len(records)), model)
	for _, r := range records {
		var record MrpBom_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record))

		old_ptmpl_id, product_tmpl := ParseMany2One(record.ProductTmplID)
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

		ur := map[string]any{
			"product_tmpl_id":  product_tmpl_id,
			"code":             record.Code,
			"type":             record.Type,
			"product_qty":      record.ProductQty,
			"sequence":         record.Sequence,
			"ready_to_produce": record.ReadyToProduce,
			"consumption":      record.Consumption,
			// "version":          record.Version,
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"product_tmpl_id", "=", product_tmpl_id},
			[]any{"code", "=", record.Code},
			[]any{"type", "=", record.Type},
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
