package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data ncr.claim model
type NcrClaim_150 struct {
	Active                      bool      `json:"active"`                        // Active
	ActivityCalendarEventID     any       `json:"activity_calendar_event_id"`    // Next Activity Calendar Event 📦 relation: many2one calendar.event
	ActivityDateDeadline        time.Time `json:"activity_date_deadline"`        // Next Activity Deadline
	ActivityExceptionDecoration any       `json:"activity_exception_decoration"` // Activity Exception Decoration
	ActivityExceptionIcon       string    `json:"activity_exception_icon"`       // Icon
	ActivityIDs                 any       `json:"activity_ids"`                  // Activities 📦 relation: one2many mail.activity
	ActivityState               any       `json:"activity_state"`                // Activity State
	ActivitySummary             string    `json:"activity_summary"`              // Next Activity Summary
	ActivityTypeIcon            string    `json:"activity_type_icon"`            // Activity Type Icon
	ActivityTypeID              any       `json:"activity_type_id"`              // Next Activity Type 📦 relation: many2one mail.activity.type
	ActivityUserID              any       `json:"activity_user_id"`              // Responsible User 📦 relation: many2one res.users
	AmountTotal                 float64   `json:"amount_total"`                  // Total
	CategoryDepartmentID        any       `json:"category_department_id"`        // Department 📦 relation: many2one ncr.category
	CategoryIssueID             any       `json:"category_issue_id"`             // Issue 📦 relation: many2one ncr.category
	CategorySectionID           any       `json:"category_section_id"`           // Section 📦 relation: many2one ncr.category
	CategorySourceID            any       `json:"category_source_id"`            // Source 📦 relation: many2one ncr.category
	Classification              any       `json:"classification"`                // Classification
	ContactDetailsSource        any       `json:"contact_details_source"`        // Pull Details From
	CorrectiveActionsNotes      string    `json:"corrective_actions_notes"`      // Corrective Actions ⭐ required
	CostImpact                  any       `json:"cost_impact"`                   // Cost Impact
	CurrencyID                  any       `json:"currency_id"`                   // Currency 📦 relation: many2one res.currency ⭐ required
	DisplayName                 string    `json:"display_name"`                  // Display Name
	HasMessage                  bool      `json:"has_message"`                   // Has Message
	ID                          int       `json:"id"`                            // ID
	InvoiceID                   any       `json:"invoice_id"`                    // Invoice 📦 relation: many2one account.move
	IsManager                   bool      `json:"is_manager"`                    // Is Manager
	IssueType                   any       `json:"issue_type"`                    // Issue Type ⭐ required
	MessageAttachmentCount      int       `json:"message_attachment_count"`      // Attachment Count
	MessageFollowerIDs          any       `json:"message_follower_ids"`          // Followers 📦 relation: one2many mail.followers
	MessageHasError             bool      `json:"message_has_error"`             // Message Delivery error
	MessageHasErrorCounter      int       `json:"message_has_error_counter"`     // Number of errors
	MessageHasSmsError          bool      `json:"message_has_sms_error"`         // SMS Delivery error
	MessageIDs                  any       `json:"message_ids"`                   // Messages 📦 relation: one2many mail.message
	MessageIsFollower           bool      `json:"message_is_follower"`           // Is Follower
	MessageMainAttachmentID     any       `json:"message_main_attachment_id"`    // Main Attachment 📦 relation: many2one ir.attachment
	MessageNeedaction           bool      `json:"message_needaction"`            // Action Needed
	MessageNeedactionCounter    int       `json:"message_needaction_counter"`    // Number of Actions
	MessagePartnerIDs           any       `json:"message_partner_ids"`           // Followers (Partners) 📦 relation: many2many res.partner
	MessageUnread               bool      `json:"message_unread"`                // Unread Messages
	MessageUnreadCounter        int       `json:"message_unread_counter"`        // Unread Messages Counter
	MyActivityDateDeadline      time.Time `json:"my_activity_date_deadline"`     // My Activity Deadline
	Name                        string    `json:"name"`                          // Name ⭐ required
	NonConformanceNotes         string    `json:"non_conformance_notes"`         // Non-Conformance ⭐ required
	Notes                       string    `json:"notes"`                         // Notes
	Origin                      string    `json:"origin"`                        // Source Document
	PartnerContactID            any       `json:"partner_contact_id"`            // Contact 📦 relation: many2one res.partner
	PartnerID                   any       `json:"partner_id"`                    // Customer/Vendor 📦 relation: many2one res.partner
	PartnerShippingID           any       `json:"partner_shipping_id"`           // Delivery Address 📦 relation: many2one res.partner
	PickingID                   any       `json:"picking_id"`                    // Transfer 📦 relation: many2one stock.picking
	ProductLine                 any       `json:"product_line"`                  // Products 📦 relation: one2many ncr.product_line
	ProductLineSource           any       `json:"product_line_source"`           // Pull Products From
	PurchaseOrderID             any       `json:"purchase_order_id"`             // Purchase Order 📦 relation: many2one purchase.order
	ResolutionLeadTime          any       `json:"resolution_lead_time"`          // Resolution Lead Time
	ResonsibleForError          any       `json:"resonsible_for_error"`          // Responsible For Error 📦 relation: many2one hr.employee
	RootCauseAnalysisNotes      string    `json:"root_cause_analysis_notes"`     // Root Cause Analysis ⭐ required
	SaleOrderID                 any       `json:"sale_order_id"`                 // Sale Order 📦 relation: many2one sale.order
	SeverityLevel               any       `json:"severity_level"`                // Severity
	State                       any       `json:"state"`                         // Status
	UserID                      any       `json:"user_id"`                       // Resolved By 📦 relation: many2one res.users
	VerificationNotes           string    `json:"verification_notes"`            // Verification
	WebsiteMessageIDs           any       `json:"website_message_ids"`           // Website Messages 📦 relation: one2many mail.message
	XStudioDateTime             any       `json:"x_studio_date_time"`            // Date & Time ⚠️ manual: True --> Date
}

// quest19data ncr.claim model
type NcrClaim_190 struct {
	AmountTotal              float64   `json:"amount_total"`               // Total
	CategoryDepartmentID     any       `json:"category_department_id"`     // Department 📦 relation: many2one ncr.category
	CategoryIssueID          any       `json:"category_issue_id"`          // Issue 📦 relation: many2one ncr.category
	CategorySectionID        any       `json:"category_section_id"`        // Section 📦 relation: many2one ncr.category
	CategorySourceID         any       `json:"category_source_id"`         // Source 📦 relation: many2one ncr.category
	Classification           any       `json:"classification"`             // Classification ⭐ required
	CorrectiveActionsNotes   string    `json:"corrective_actions_notes"`   // Corrective Actions ⭐ required
	CostImpact               any       `json:"cost_impact"`                // Cost Impact
	CurrencyID               any       `json:"currency_id"`                // Currency 📦 relation: many2one res.currency ⭐ required
	Date                     time.Time `json:"date"`                       // Date
	DisplayName              string    `json:"display_name"`               // Display Name
	HasMessage               bool      `json:"has_message"`                // Has Message
	ID                       int       `json:"id"`                         // ID
	InvoiceID                any       `json:"invoice_id"`                 // Invoice 📦 relation: many2one account.move
	IsPullButtonClicked      bool      `json:"is_pull_button_clicked"`     // Is Pull Button Clicked
	IssueType                any       `json:"issue_type"`                 // Issue Type ⭐ required
	MessageAttachmentCount   int       `json:"message_attachment_count"`   // Attachment Count
	MessageFollowerIDs       any       `json:"message_follower_ids"`       // Followers 📦 relation: one2many mail.followers
	MessageHasError          bool      `json:"message_has_error"`          // Message Delivery error
	MessageHasErrorCounter   int       `json:"message_has_error_counter"`  // Number of errors
	MessageHasSmsError       bool      `json:"message_has_sms_error"`      // SMS Delivery error
	MessageIDs               any       `json:"message_ids"`                // Messages 📦 relation: one2many mail.message
	MessageIsFollower        bool      `json:"message_is_follower"`        // Is Follower
	MessageNeedaction        bool      `json:"message_needaction"`         // Action Needed
	MessageNeedactionCounter int       `json:"message_needaction_counter"` // Number of Actions
	MessagePartnerIDs        any       `json:"message_partner_ids"`        // Followers (Partners) 📦 relation: many2many res.partner
	Name                     string    `json:"name"`                       // Name ⭐ required
	NonConformanceNotes      string    `json:"non_conformance_notes"`      // Non-Conformance ⭐ required
	PartnerID                any       `json:"partner_id"`                 // Customer/Vendor 📦 relation: many2one res.partner
	PartnerShippingID        any       `json:"partner_shipping_id"`        // Delivery Address 📦 relation: many2one res.partner
	PartnerSource            any       `json:"partner_source"`             // Source
	PickingID                any       `json:"picking_id"`                 // Transfer 📦 relation: many2one stock.picking
	ProductLineIDs           any       `json:"product_line_ids"`           // Products 📦 relation: one2many ncr.product.line
	PurchaseOrderID          any       `json:"purchase_order_id"`          // Purchase Order 📦 relation: many2one purchase.order
	RatingIDs                any       `json:"rating_ids"`                 // Ratings 📦 relation: one2many rating.rating
	ResolutionLeadTime       any       `json:"resolution_lead_time"`       // Resolution Lead Time
	ResolvedBy               any       `json:"resolved_by"`                // Resolved By 📦 relation: many2one res.users
	ResponsibleID            any       `json:"responsible_id"`             // Responsible For Error 📦 relation: many2one hr.employee
	RootCauseAnalysisNotes   string    `json:"root_cause_analysis_notes"`  // Root Cause Analysis ⭐ required
	SaleOrderID              any       `json:"sale_order_id"`              // Sale Order 📦 relation: many2one sale.order
	Severity                 any       `json:"severity"`                   // Severity
	State                    any       `json:"state"`                      // State
	WebsiteMessageIDs        any       `json:"website_message_ids"`        // Website Messages 📦 relation: one2many mail.message
}

func (m *Model) NCRClaim() {
	model := "ncr.claim"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(NcrClaim_150{})

	// root
	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"parent_category_id", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var record NcrClaim_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record))
		ur := map[string]any{
			"name":                      record.Name,
			"corrective_actions_notes":  record.CorrectiveActionsNotes,
			"root_cause_analysis_notes": record.RootCauseAnalysisNotes,
			"state":                     record.State,
		}
		// add so and po reference to the non_conformance_notes field
		// so and po many2one fields are not directly mapped in quest19data ncr.claim model
		// we will just print them for now

		// ========
		// Column 1
		// ========

		// Classification values
		// [non_conform] - Non-Conformance | [nc] - Non-Conformance
		// [finding] - Finding             | [find] - Finding
		switch record.Classification {
		case "non_conform":
			ur["classification"] = "nc"
		case "finding":
			ur["classification"] = "find"
		}

		// Issue Type values
		// [vendor] - Vendor                       | [vendor] - Vendor
		// [customer] - Customer                   | [customer] - Customer
		// [audit] - Audit                         | [audit] - Audit
		// [opp_imp] - Opportunity for Improvement | [opp_imp] - Opportunity for Improvement
		// [objective] - Objective
		// [internal] - Internal                   | [internal] - Internal
		// [external] - External
		switch record.IssueType {
		case "vendor":
			ur["issue_type"] = "vendor"
		case "customer":
			ur["issue_type"] = "customer"
		case "audit":
			ur["issue_type"] = "audit"
		case "opp_imp":
			ur["issue_type"] = "opp_imp"
		case "internal":
			ur["issue_type"] = "internal"
		}

		_, sale_order_name := ParseMany2One(record.SaleOrderID)
		_, purchase_order_name := ParseMany2One(record.PurchaseOrderID)

		// fmt.Println("sale_order_name:", sale_order_name)
		// fmt.Println("purchase_order_name:", purchase_order_name)
		ur["non_conformance_notes"] = record.NonConformanceNotes +
			fmt.Sprintf("\n\nSale Order: %s\nPurchase Order: %s", sale_order_name, purchase_order_name)

		// ========
		// Column 2
		// ========

		// XStudioDateTime | Date
		if record.XStudioDateTime != nil {
			// fmt.Println("record.XStudioDateTime:", record.XStudioDateTime)
			ur["date"] = record.XStudioDateTime
		}

		// Resolved By
		// user_id | resolved_by
		if record.UserID != nil {
			_, uname := ParseMany2One(record.UserID)
			// fmt.Println("uname:", uname)
			userid, err := m.Dest.GetID("res.users", []any{[]any{"name", "=", uname}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if userid != -1 {
				ur["resolved_by"] = userid
			}
		}

		// Responsible For Error
		// resonsible_for_error | responsible_id
		if record.ResonsibleForError != nil {
			_, rname := ParseMany2One(record.ResonsibleForError)
			// fmt.Println("rname:", rname)
			responsibleid, err := m.Dest.GetID("hr.employee", []any{[]any{"name", "=", rname}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if responsibleid != -1 {
				ur["responsible_id"] = responsibleid
			}
		}

		// Severity
		// severity_level | severity
		// [minor] - Minor       | [minor] - Minor
		// [major] - Major       | [major] - Major
		// [critical] - Critical | [critical] - Critical
		switch record.SeverityLevel {
		case "minor":
			ur["severity"] = "minor"
		case "major":
			ur["severity"] = "major"
		case "critical":
			ur["severity"] = "critical"
		}

		// Cost Impact
		ur["cost_impact"] = record.CostImpact

		// Resolution Lead Time
		// resolution_lead_time
		// [same_day] - Same Day           | [same_day] - Same Day
		// [next_day] - Next Day           | [next_day] - Next Day
		// [over_two_days] - Over Two Days | [over_two_days] - Over Two Days
		// [cancelled] - Canceled          | [cancelled] - Canceled

		switch record.ResolutionLeadTime {
		case "same_day":
			ur["resolution_lead_time"] = "same_day"
		case "next_day":
			ur["resolution_lead_time"] = "next_day"
		case "over_two_days":
			ur["resolution_lead_time"] = "over_two_days"
		case "cancelled":
			ur["resolution_lead_time"] = "cancelled"
		}

		// Category - Source
		// category_source_id
		if record.CategorySourceID != nil {
			_, cname := ParseMany2One(record.CategorySourceID)

			// id, err := m.Source.SearchRead("ncr.category", 0, 0, []string{"id", "name", "parent_id"}, []any{[]any{"id", "=", cid}})
			// if err != nil {
			// 	m.Log.Error(model, "func", trace(), "err", err)
			// 	continue
			// }
			// parentid, err := m.Source.GetID("ncr.category", []any{[]any{"id", "=", id[0]["parent_id"].([]any)[0]}})
			// if err != nil {
			// 	m.Log.Error(model, "func", trace(), "err", err)
			// 	continue
			// }
			// id, err := m.Source.SearchRead("ncr.category", 0, 0, []string{"id", "name", "parent_id"}, []any{[]any{"id", "=", cid}})
			// if err != nil {
			// 	m.Log.Error(model, "func", trace(), "err", err)
			// 	continue
			// }

			categorysourceid, err := m.Dest.GetID("ncr.category", []any{[]any{"name", "=", strings.ToUpper(cname)}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			ur["category_source_id"] = categorysourceid
		}

		// Category - Department
		// category_department_id
		if record.CategoryDepartmentID != nil {
			_, dname := ParseMany2One(record.CategoryDepartmentID)

			categorydepartmentid, err := m.Dest.GetID("ncr.category", []any{[]any{"name", "=", strings.ToUpper(dname)}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			ur["category_department_id"] = categorydepartmentid
		}

		// Category - Section
		// category_section_id
		if record.CategorySectionID != nil {
			_, sname := ParseMany2One(record.CategorySectionID)

			categorysectionid, err := m.Dest.GetID("ncr.category", []any{[]any{"name", "=", strings.ToUpper(sname)}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			ur["category_section_id"] = categorysectionid
		}

		// Category - Issue
		// category_issue_id
		if record.CategoryIssueID != nil {
			_, iname := ParseMany2One(record.CategoryIssueID)

			categoryissueid, err := m.Dest.GetID("ncr.category", []any{[]any{"name", "=", strings.ToUpper(iname)}})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			ur["category_issue_id"] = categoryissueid
		}

		// get record id in destination

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", record.Name},
			// []any{"parent_category_id", "=", false},
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
