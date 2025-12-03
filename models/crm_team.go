package models

import (
	"github.com/schollz/progressbar/v3"
)

// quest15data crm.team model
type CrmTeam_150 struct {
	Active bool `json:"active"` // Active
	// AliasBouncedContent   string `json:"alias_bounced_content"`   // Custom Bounced Message
	AliasContact any `json:"alias_contact"` // Alias Contact Security ⭐ required
	// AliasDefaults         string `json:"alias_defaults"`          // Default Values ⭐ required
	// AliasDomain           string `json:"alias_domain"`            // Alias domain
	// AliasForceThreadID    int    `json:"alias_force_thread_id"`   // Record Thread ID
	// AliasID               any    `json:"alias_id"`                // Alias 📦 relation: many2one mail.alias ⭐ required
	// AliasModelID          any    `json:"alias_model_id"`          // Aliased Model 📦 relation: many2one ir.model ⭐ required
	// AliasName             string `json:"alias_name"`              // Alias Name
	// AliasParentModelID    any    `json:"alias_parent_model_id"`   // Parent Model 📦 relation: many2one ir.model
	// AliasParentThreadID   int    `json:"alias_parent_thread_id"`  // Parent Record Thread ID
	// AliasUserID           any    `json:"alias_user_id"`           // Owner 📦 relation: many2one res.users
	// AssignmentAutoEnabled bool   `json:"assignment_auto_enabled"` // Auto Assignment
	// AssignmentDomain      string `json:"assignment_domain"`       // Assignment Domain
	// AssignmentEnabled     bool   `json:"assignment_enabled"`      // Lead Assign
	// AssignmentMax         int    `json:"assignment_max"`          // Lead Average Capacity
	// AssignmentOptout      bool   `json:"assignment_optout"`       // Skip auto assignment
	// Color                 int    `json:"color"`                   // Color Index
	// CompanyID             any    `json:"company_id"`              // Company 📦 relation: many2one res.company
	// CrmTeamMemberAllIDs any `json:"crm_team_member_all_ids"` // Sales Team Members (incl. inactive) 📦 relation: one2many crm.team.member
	CrmTeamMemberIDs any `json:"crm_team_member_ids"` // Sales Team Members 📦 relation: one2many crm.team.member
	// CurrencyID            any    `json:"currency_id"`             // Currency 📦 relation: many2one res.currency
	// DashboardButtonName   string `json:"dashboard_button_name"`   // Dashboard Button
	// DashboardGraphData    string `json:"dashboard_graph_data"`    // Dashboard Graph Data
	// DisplayName           string `json:"display_name"`            // Display Name
	// FavoriteUserIDs       any    `json:"favorite_user_ids"`       // Favorite Members 📦 relation: many2many res.users
	// HasMessage            bool   `json:"has_message"`             // Has Message
	ID int `json:"id"` // ID
	// Invoiced                   float64 `json:"invoiced"`                      // Invoiced This Month
	InvoicedTarget float64 `json:"invoiced_target"` // Invoicing Target
	// IsFavorite                 bool    `json:"is_favorite"`                   // Show on dashboard
	// IsMembershipMulti          bool    `json:"is_membership_multi"`           // Multiple Memberships Allowed
	// LeadAllAssignedMonthCount  int     `json:"lead_all_assigned_month_count"` // # Leads/Opps assigned this month
	// LeadUnassignedCount        int     `json:"lead_unassigned_count"`         // # Unassigned Leads
	// MemberCompanyIDs           any     `json:"member_company_ids"`            // Member Company 📦 relation: many2many res.company
	// MemberIDs                  any     `json:"member_ids"`                    // Salespersons 📦 relation: many2many res.users
	// MemberWarning              string  `json:"member_warning"`                // Membership Issue Warning
	// MessageAttachmentCount     int     `json:"message_attachment_count"`      // Attachment Count
	// MessageFollowerIDs         any     `json:"message_follower_ids"`          // Followers 📦 relation: one2many mail.followers
	// MessageHasError            bool    `json:"message_has_error"`             // Message Delivery error
	// MessageHasErrorCounter     int     `json:"message_has_error_counter"`     // Number of errors
	// MessageHasSmsError         bool    `json:"message_has_sms_error"`         // SMS Delivery error
	// MessageIDs                 any     `json:"message_ids"`                   // Messages 📦 relation: one2many mail.message
	// MessageIsFollower          bool    `json:"message_is_follower"`           // Is Follower
	// MessageMainAttachmentID    any     `json:"message_main_attachment_id"`    // Main Attachment 📦 relation: many2one ir.attachment
	// MessageNeedaction          bool    `json:"message_needaction"`            // Action Needed
	// MessageNeedactionCounter   int     `json:"message_needaction_counter"`    // Number of Actions
	// MessagePartnerIDs          any     `json:"message_partner_ids"`           // Followers (Partners) 📦 relation: many2many res.partner
	// MessageUnread              bool    `json:"message_unread"`                // Unread Messages
	// MessageUnreadCounter       int     `json:"message_unread_counter"`        // Unread Messages Counter
	Name string `json:"name"` // Sales Team ⭐ required
	// OpportunitiesAmount        float64 `json:"opportunities_amount"`          // Opportunities Revenues
	// OpportunitiesCount         int     `json:"opportunities_count"`           // # Opportunities
	// OpportunitiesOverdueAmount float64 `json:"opportunities_overdue_amount"`  // Overdue Opportunities Revenues
	// OpportunitiesOverdueCount  int     `json:"opportunities_overdue_count"`   // # Overdue Opportunities
	// QuotationsAmount           float64 `json:"quotations_amount"`             // Amount of quotations to invoice
	// QuotationsCount            int     `json:"quotations_count"`              // Number of quotations to invoice
	// SaleOrderCount             int     `json:"sale_order_count"`              // # Sale Orders
	// SalesToInvoiceCount        int     `json:"sales_to_invoice_count"`        // Number of sales to invoice
	// Sequence                   int     `json:"sequence"`                      // Sequence
	UseLeads         bool `json:"use_leads"`         // Leads
	UseOpportunities bool `json:"use_opportunities"` // Pipeline
	UseQuotations    bool `json:"use_quotations"`    // Quotations
	UserID           any  `json:"user_id"`           // Team Leader 📦 relation: many2one res.users
	// WebsiteMessageIDs          any     `json:"website_message_ids"`           // Website Messages 📦 relation: one2many mail.message
}

// quest19data crm.team model
type CrmTeam_190 struct {
	Active                       bool    `json:"active"`                           // Active
	AliasBouncedContent          string  `json:"alias_bounced_content"`            // Custom Bounced Message
	AliasContact                 any     `json:"alias_contact"`                    // Alias Contact Security ⭐ required
	AliasDefaults                string  `json:"alias_defaults"`                   // Default Values ⭐ required
	AliasDomain                  string  `json:"alias_domain"`                     // Alias Domain Name
	AliasDomainID                any     `json:"alias_domain_id"`                  // Alias Domain 📦 relation: many2one mail.alias.domain
	AliasEmail                   string  `json:"alias_email"`                      // Email Alias
	AliasForceThreadID           int     `json:"alias_force_thread_id"`            // Record Thread ID
	AliasFullName                string  `json:"alias_full_name"`                  // Alias Email
	AliasID                      any     `json:"alias_id"`                         // Alias 📦 relation: many2one mail.alias ⭐ required
	AliasIncomingLocal           bool    `json:"alias_incoming_local"`             // Local-part based incoming detection
	AliasModelID                 any     `json:"alias_model_id"`                   // Aliased Model 📦 relation: many2one ir.model ⭐ required
	AliasName                    string  `json:"alias_name"`                       // Alias Name
	AliasParentModelID           any     `json:"alias_parent_model_id"`            // Parent Model 📦 relation: many2one ir.model
	AliasParentThreadID          int     `json:"alias_parent_thread_id"`           // Parent Record Thread ID
	AliasStatus                  any     `json:"alias_status"`                     // Alias Status
	AssignmentAutoEnabled        bool    `json:"assignment_auto_enabled"`          // Auto Assignment
	AssignmentDomain             string  `json:"assignment_domain"`                // Assignment Domain
	AssignmentEnabled            bool    `json:"assignment_enabled"`               // Lead Assign
	AssignmentMax                int     `json:"assignment_max"`                   // Lead Average Capacity
	AssignmentOptout             bool    `json:"assignment_optout"`                // Skip auto assignment
	Color                        int     `json:"color"`                            // Color Index
	CompanyID                    any     `json:"company_id"`                       // Company 📦 relation: many2one res.company
	CrmTeamMemberAllIDs          any     `json:"crm_team_member_all_ids"`          // Sales Team Members (incl. inactive) 📦 relation: one2many crm.team.member
	CrmTeamMemberIDs             any     `json:"crm_team_member_ids"`              // Sales Team Members 📦 relation: one2many crm.team.member
	CurrencyID                   any     `json:"currency_id"`                      // Currency 📦 relation: many2one res.currency
	DashboardButtonName          string  `json:"dashboard_button_name"`            // Dashboard Button
	DisplayName                  string  `json:"display_name"`                     // Display Name
	FavoriteUserIDs              any     `json:"favorite_user_ids"`                // Favorite Members 📦 relation: many2many res.users
	HasMessage                   bool    `json:"has_message"`                      // Has Message
	ID                           int     `json:"id"`                               // ID
	Invoiced                     float64 `json:"invoiced"`                         // Invoiced This Month
	InvoicedTarget               float64 `json:"invoiced_target"`                  // Invoicing Target
	IsFavorite                   bool    `json:"is_favorite"`                      // Show on dashboard
	IsMembershipMulti            bool    `json:"is_membership_multi"`              // Multiple Memberships Allowed
	LeadAllAssignedMonthCount    int     `json:"lead_all_assigned_month_count"`    // # Leads/Opps assigned this month
	LeadAllAssignedMonthExceeded bool    `json:"lead_all_assigned_month_exceeded"` // Exceed monthly lead assignement
	LeadPropertiesDefinition     any     `json:"lead_properties_definition"`       // Lead Properties
	LeadUnassignedCount          int     `json:"lead_unassigned_count"`            // # Unassigned Leads
	MemberCompanyIDs             any     `json:"member_company_ids"`               // Member Company 📦 relation: many2many res.company
	MemberIDs                    any     `json:"member_ids"`                       // Salespersons 📦 relation: many2many res.users
	MemberWarning                string  `json:"member_warning"`                   // Membership Issue Warning
	MessageAttachmentCount       int     `json:"message_attachment_count"`         // Attachment Count
	MessageFollowerIDs           any     `json:"message_follower_ids"`             // Followers 📦 relation: one2many mail.followers
	MessageHasError              bool    `json:"message_has_error"`                // Message Delivery error
	MessageHasErrorCounter       int     `json:"message_has_error_counter"`        // Number of errors
	MessageHasSmsError           bool    `json:"message_has_sms_error"`            // SMS Delivery error
	MessageIDs                   any     `json:"message_ids"`                      // Messages 📦 relation: one2many mail.message
	MessageIsFollower            bool    `json:"message_is_follower"`              // Is Follower
	MessageNeedaction            bool    `json:"message_needaction"`               // Action Needed
	MessageNeedactionCounter     int     `json:"message_needaction_counter"`       // Number of Actions
	MessagePartnerIDs            any     `json:"message_partner_ids"`              // Followers (Partners) 📦 relation: many2many res.partner
	Name                         string  `json:"name"`                             // Sales Team ⭐ required
	OriginSurveyIDs              any     `json:"origin_survey_ids"`                // Survey opportunities related to the sales team 📦 relation: one2many survey.survey
	RatingIDs                    any     `json:"rating_ids"`                       // Ratings 📦 relation: one2many rating.rating
	SaleOrderCount               int     `json:"sale_order_count"`                 // # Sale Orders
	Sequence                     int     `json:"sequence"`                         // Sequence
	UseLeads                     bool    `json:"use_leads"`                        // Leads
	UseOpportunities             bool    `json:"use_opportunities"`                // Pipeline
	UserID                       any     `json:"user_id"`                          // Team Leader 📦 relation: many2one res.users
	WebsiteMessageIDs            any     `json:"website_message_ids"`              // Website Messages 📦 relation: one2many mail.message
}

func (m *Model) CrmTeam() {
	model := "crm.team"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(CrmTeam_150{}), []any{
		// []any{"name", "=", "EDMONTON"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var team CrmTeam_150
		FillStruct(r, &team)
		// fmt.Println(prettyprint(team))

		ur := make(map[string]any)
		ur["name"] = team.Name
		ur["use_opportunities"] = team.UseOpportunities
		ur["use_leads"] = team.UseLeads
		ur["invoiced_target"] = team.InvoicedTarget
		ur["active"] = team.Active
		ur["alias_contact"] = team.AliasContact

		// user_id -- team leader
		if team.UserID != false {
			_, user_name := ParseMany2One(team.UserID)
			user_id, err := m.Dest.GetID("res.users", []any{
				[]any{"name", "=", user_name},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				return
			}
			if user_id != -1 {
				ur["user_id"] = user_id
			}
		}

		// team member ids
		// if team.CrmTeamMemberIDs != false {
		// 	var member_ids []int
		// 	fmt.Println("team.CrmTeamMemberIDs", prettyprint(team.CrmTeamMemberIDs))
		// 	for _, mem := range team.CrmTeamMemberIDs.([]any) {
		// 		smid, _ := m.Source.SearchRead("res.users", 0, 0, []string{"name"}, []any{
		// 			[]any{"id", "=", mem.(float64)},
		// 		})
		// 		fmt.Println(" smid:", prettyprint(smid))
		// 		if len(smid) == 0 {
		// 			continue
		// 		}
		// 		fmt.Println(" smid:", prettyprint(smid))
		// 		mid, _ := m.Dest.GetID("res.users", []any{
		// 			[]any{"name", "=", smid[0]["name"]},
		// 		})
		// 		if mid != -1 {
		// 			member_ids = append(member_ids, mid)
		// 			fmt.Println("  member id:", mid)
		// 		}
		// 	}
		// 	if len(member_ids) > 0 {
		// 		ur["member_ids"] = member_ids
		// 	}
		// }

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", team.Name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
}
