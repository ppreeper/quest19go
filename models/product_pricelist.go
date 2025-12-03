package models

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data product.pricelist model
type ProductPricelist_150 struct {
	Active          bool   `json:"active"`            // Active
	CompanyID       any    `json:"company_id"`        // Company 📦 relation: many2one res.company
	CountryGroupIDs any    `json:"country_group_ids"` // Country Groups 📦 relation: many2many res.country.group
	CurrencyID      any    `json:"currency_id"`       // Currency 📦 relation: many2one res.currency ⭐ required
	DiscountPolicy  any    `json:"discount_policy"`   // Discount Policy ⭐ required
	DisplayName     string `json:"display_name"`      // Display Name
	ID              int    `json:"id"`                // ID
	ItemIDs         any    `json:"item_ids"`          // Pricelist Rules 📦 relation: one2many product.pricelist.item
	Name            string `json:"name"`              // Pricelist Name ⭐ required
	Sequence        int    `json:"sequence"`          // Sequence
	XStudioNotes    string `json:"x_studio_notes"`    // Notes: ⚠️ manual: True
	XStudioNotes1   string `json:"x_studio_notes_1"`  // Notes: ⚠️ manual: True
}

// quest19data product.pricelist model
type ProductPricelist_190 struct {
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
	CompanyID                   any       `json:"company_id"`                    // Company 📦 relation: many2one res.company
	CountryGroupIDs             any       `json:"country_group_ids"`             // Country Groups 📦 relation: many2many res.country.group
	CurrencyID                  any       `json:"currency_id"`                   // Currency 📦 relation: many2one res.currency ⭐ required
	DisplayName                 string    `json:"display_name"`                  // Display Name
	HasMessage                  bool      `json:"has_message"`                   // Has Message
	ID                          int       `json:"id"`                            // ID
	ItemIDs                     any       `json:"item_ids"`                      // Pricelist Rules 📦 relation: one2many product.pricelist.item
	MessageAttachmentCount      int       `json:"message_attachment_count"`      // Attachment Count
	MessageFollowerIDs          any       `json:"message_follower_ids"`          // Followers 📦 relation: one2many mail.followers
	MessageHasError             bool      `json:"message_has_error"`             // Message Delivery error
	MessageHasErrorCounter      int       `json:"message_has_error_counter"`     // Number of errors
	MessageHasSmsError          bool      `json:"message_has_sms_error"`         // SMS Delivery error
	MessageIDs                  any       `json:"message_ids"`                   // Messages 📦 relation: one2many mail.message
	MessageIsFollower           bool      `json:"message_is_follower"`           // Is Follower
	MessageNeedaction           bool      `json:"message_needaction"`            // Action Needed
	MessageNeedactionCounter    int       `json:"message_needaction_counter"`    // Number of Actions
	MessagePartnerIDs           any       `json:"message_partner_ids"`           // Followers (Partners) 📦 relation: many2many res.partner
	MyActivityDateDeadline      time.Time `json:"my_activity_date_deadline"`     // My Activity Deadline
	Name                        string    `json:"name"`                          // Pricelist Name ⭐ required
	RatingIDs                   any       `json:"rating_ids"`                    // Ratings 📦 relation: one2many rating.rating
	Sequence                    int       `json:"sequence"`                      // Sequence
	WebsiteMessageIDs           any       `json:"website_message_ids"`           // Website Messages 📦 relation: one2many mail.message
}

func (m *Model) ProductPricelist() {
	model := "product.pricelist"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	currency_map := m.currencyMap()

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductPricelist_150{}), []any{})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	m.writeRecord(model, map[string]any{"currency_id": currency_map["CAD"]}, 1)

	bar := progressbar.Default(int64(len(records)))
	for _, r := range records {
		var p ProductPricelist_150
		FillStruct(r, &p)
		// fmt.Println(prettyprint(p))

		ur := map[string]any{
			"name": p.Name,
		}

		// currency
		_, currency_name := ParseMany2One(p.CurrencyID)
		currency_id := currency_map[currency_name]
		if currency_id != -1 {
			ur["currency_id"] = currency_id
		}

		// xstudio_notes_1 needs to be inserted into model

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", p.Name},
		})

		m.writeRecord(model, ur, rid)

		bar.Add(1)
	}
	bar.Finish()
}
