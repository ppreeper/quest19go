package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

// quest15data crm.team.member model
type CrmTeamMember_150 struct {
	Active bool `json:"active"` // Active
	// AssignmentDomain  string `json:"assignment_domain"`   // Assignment Domain
	// AssignmentEnabled bool   `json:"assignment_enabled"`  // Lead Assign
	// AssignmentMax     int    `json:"assignment_max"`      // Average Leads Capacity (on 30 days)
	// AssignmentOptout  bool   `json:"assignment_optout"`   // Skip auto assignment
	CompanyID   any    `json:"company_id"`   // Company 📦 relation: many2one res.company
	CrmTeamID   any    `json:"crm_team_id"`  // Sales Team 📦 relation: many2one crm.team ⭐ required
	DisplayName string `json:"display_name"` // Display Name
	Email       string `json:"email"`        // Email
	// HasMessage        bool   `json:"has_message"`         // Has Message
	ID                int    `json:"id"`                  // ID
	Image128          any    `json:"image_128"`           // Image (128)
	Image1920         any    `json:"image_1920"`          // Image
	IsMembershipMulti bool   `json:"is_membership_multi"` // Multiple Memberships Allowed
	LeadMonthCount    int    `json:"lead_month_count"`    // Leads (30 days)
	MemberWarning     string `json:"member_warning"`      // Member Warning
	// MessageAttachmentCount   int    `json:"message_attachment_count"`   // Attachment Count
	// MessageFollowerIDs       any    `json:"message_follower_ids"`       // Followers 📦 relation: one2many mail.followers
	// MessageHasError          bool   `json:"message_has_error"`          // Message Delivery error
	// MessageHasErrorCounter   int    `json:"message_has_error_counter"`  // Number of errors
	// MessageHasSmsError       bool   `json:"message_has_sms_error"`      // SMS Delivery error
	// MessageIDs               any    `json:"message_ids"`                // Messages 📦 relation: one2many mail.message
	// MessageIsFollower        bool   `json:"message_is_follower"`        // Is Follower
	// MessageMainAttachmentID  any    `json:"message_main_attachment_id"` // Main Attachment 📦 relation: many2one ir.attachment
	// MessageNeedaction        bool   `json:"message_needaction"`         // Action Needed
	// MessageNeedactionCounter int    `json:"message_needaction_counter"` // Number of Actions
	// MessagePartnerIDs        any    `json:"message_partner_ids"`        // Followers (Partners) 📦 relation: many2many res.partner
	// MessageUnread            bool   `json:"message_unread"`             // Unread Messages
	// MessageUnreadCounter     int    `json:"message_unread_counter"`     // Unread Messages Counter
	Mobile string `json:"mobile"` // Mobile
	Name   string `json:"name"`   // Name
	Phone  string `json:"phone"`  // Phone
	// UserCompanyIDs any    `json:"user_company_ids"`  // User Company 📦 relation: many2many res.company
	UserID any `json:"user_id"` // Salesperson 📦 relation: many2one res.users ⭐ required
	// UserInTeamsIDs any    `json:"user_in_teams_ids"` // User In Teams 📦 relation: many2many res.users
	// WebsiteMessageIDs any    `json:"website_message_ids"` // Website Messages 📦 relation: one2many mail.message
}

// quest19data crm.team.member model
type CrmTeamMember_190 struct {
	Active                    bool   `json:"active"`                      // Active
	AssignmentDomain          string `json:"assignment_domain"`           // Assignment Domain
	AssignmentDomainPreferred string `json:"assignment_domain_preferred"` // Preference assignment Domain
	AssignmentEnabled         bool   `json:"assignment_enabled"`          // Lead Assign
	AssignmentMax             int    `json:"assignment_max"`              // Average Leads Capacity (on 30 days)
	AssignmentOptout          bool   `json:"assignment_optout"`           // Pause assignment
	CompanyID                 any    `json:"company_id"`                  // Company 📦 relation: many2one res.company
	CrmTeamID                 any    `json:"crm_team_id"`                 // Sales Team 📦 relation: many2one crm.team ⭐ required
	DisplayName               string `json:"display_name"`                // Display Name
	Email                     string `json:"email"`                       // Email
	HasMessage                bool   `json:"has_message"`                 // Has Message
	ID                        int    `json:"id"`                          // ID
	Image128                  any    `json:"image_128"`                   // Image (128)
	Image1920                 any    `json:"image_1920"`                  // Image
	IsMembershipMulti         bool   `json:"is_membership_multi"`         // Multiple Memberships Allowed
	LeadDayCount              int    `json:"lead_day_count"`              // Leads (last 24h)
	LeadMonthCount            int    `json:"lead_month_count"`            // Leads (30 days)
	MemberWarning             string `json:"member_warning"`              // Member Warning
	MessageAttachmentCount    int    `json:"message_attachment_count"`    // Attachment Count
	MessageFollowerIDs        any    `json:"message_follower_ids"`        // Followers 📦 relation: one2many mail.followers
	MessageHasError           bool   `json:"message_has_error"`           // Message Delivery error
	MessageHasErrorCounter    int    `json:"message_has_error_counter"`   // Number of errors
	MessageHasSmsError        bool   `json:"message_has_sms_error"`       // SMS Delivery error
	MessageIDs                any    `json:"message_ids"`                 // Messages 📦 relation: one2many mail.message
	MessageIsFollower         bool   `json:"message_is_follower"`         // Is Follower
	MessageNeedaction         bool   `json:"message_needaction"`          // Action Needed
	MessageNeedactionCounter  int    `json:"message_needaction_counter"`  // Number of Actions
	MessagePartnerIDs         any    `json:"message_partner_ids"`         // Followers (Partners) 📦 relation: many2many res.partner
	Name                      string `json:"name"`                        // Name
	Phone                     string `json:"phone"`                       // Phone
	RatingIDs                 any    `json:"rating_ids"`                  // Ratings 📦 relation: one2many rating.rating
	UserCompanyIDs            any    `json:"user_company_ids"`            // User Company 📦 relation: many2many res.company
	UserID                    any    `json:"user_id"`                     // Salesperson 📦 relation: many2one res.users ⭐ required
	UserInTeamsIDs            any    `json:"user_in_teams_ids"`           // User In Teams 📦 relation: many2many res.users
	WebsiteMessageIDs         any    `json:"website_message_ids"`         // Website Messages 📦 relation: one2many mail.message
}

func (m *Model) CrmTeamMember() {
	model := "crm.team.member"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	teams, err := m.Source.SearchRead("crm.team", 0, 0, ExtractJSONTags(CrmTeam_150{}), []any{
		// []any{"name", "=", "EDMONTON"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	// bar := progressbar.Default(-1, model)
	for _, t := range teams {
		var team CrmTeam_150
		FillStruct(t, &team)
		// fmt.Println(prettyprint(team.Name))
		team_id, err := m.Dest.GetID("crm.team", []any{
			[]any{"name", "=", team.Name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		members, err := m.Source.SearchRead("crm.team.member", 0, 0, ExtractJSONTags(CrmTeamMember_150{}), []any{
			[]any{"crm_team_id", "=", team.ID},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}

		bar := progressbar.Default(-1, fmt.Sprintf("%s: %s", model, team.Name))
		for _, mbr := range members {
			var member CrmTeamMember_150
			FillStruct(mbr, &member)
			// fmt.Println("member", prettyprint(member.Name))

			ur := make(map[string]any)
			ur["crm_team_id"] = team_id

			oid, _ := m.Dest.GetID("res.users", []any{
				[]any{"name", "=", member.Name},
			})
			if oid == -1 {
				m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("no user found for member %s", member.Name))
				continue
			}
			ur["user_id"] = oid
			// fmt.Println(prettyprint(ur))

			rid, err := m.Dest.GetID("crm.team.member", []any{
				// []any{"name", "=", member.Name},
				[]any{"crm_team_id", "=", team_id},
				[]any{"user_id", "=", oid},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			// fmt.Println("rid", rid)
			m.writeRecord("crm.team.member", ur, rid)
			bar.Add(1)
		}
		bar.Finish()
	}
}
