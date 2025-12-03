package models

import "fmt"

type HrJob_150 struct {
	WebsiteID                any    `json:"website_id"`                 // Website 📦 Model Relation: website
	WebsitePublished         bool   `json:"website_published"`          // Visible on current website
	IsPublished              bool   `json:"is_published"`               // Is Published
	CanPublish               bool   `json:"can_publish"`                // Can Publish
	WebsiteUrl               string `json:"website_url"`                // Website URL
	IsSeoOptimized           bool   `json:"is_seo_optimized"`           // SEO optimized
	WebsiteMetaTitle         string `json:"website_meta_title"`         // Website meta title
	WebsiteMetaDescription   string `json:"website_meta_description"`   // Website meta description
	WebsiteMetaKeywords      string `json:"website_meta_keywords"`      // Website meta keywords
	WebsiteMetaOgImg         string `json:"website_meta_og_img"`        // Website opengraph image
	SeoName                  string `json:"seo_name"`                   // Seo name
	MessageIsFollower        bool   `json:"message_is_follower"`        // Is Follower
	MessageFollowerIDs       any    `json:"message_follower_ids"`       // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs        any    `json:"message_partner_ids"`        // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs               any    `json:"message_ids"`                // Messages 📦 Model Relation: mail.message
	HasMessage               bool   `json:"has_message"`                // Has Message
	MessageUnread            bool   `json:"message_unread"`             // Unread Messages
	MessageUnreadCounter     int    `json:"message_unread_counter"`     // Unread Messages Counter
	MessageNeedaction        bool   `json:"message_needaction"`         // Action Needed
	MessageNeedactionCounter int    `json:"message_needaction_counter"` // Number of Actions
	MessageHasError          bool   `json:"message_has_error"`          // Message Delivery error
	MessageHasErrorCounter   int    `json:"message_has_error_counter"`  // Number of errors
	MessageAttachmentCount   int    `json:"message_attachment_count"`   // Attachment Count
	MessageMainAttachmentID  any    `json:"message_main_attachment_id"` // Main Attachment 📦 Model Relation: ir.attachment
	WebsiteMessageIDs        any    `json:"website_message_ids"`        // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError       bool   `json:"message_has_sms_error"`      // SMS Delivery error
	Name                     string `json:"name"`                       // Job Position ⭐ Required
	Sequence                 int    `json:"sequence"`                   // Sequence
	ExpectedEmployees        int    `json:"expected_employees"`         // Total Forecasted Employees
	NoOfEmployee             int    `json:"no_of_employee"`             // Current Number of Employees
	NoOfRecruitment          int    `json:"no_of_recruitment"`          // Expected New Employees
	NoOfHiredEmployee        int    `json:"no_of_hired_employee"`       // Hired Employees
	EmployeeIDs              any    `json:"employee_ids"`               // Employees 📦 Model Relation: hr.employee
	Description              string `json:"description"`                // Job Description
	Requirements             string `json:"requirements"`               // Requirements
	DepartmentID             any    `json:"department_id"`              // Department 📦 Model Relation: hr.department
	CompanyID                any    `json:"company_id"`                 // Company 📦 Model Relation: res.company
	State                    any    `json:"state"`                      // Status ⭐ Required
	ID                       int    `json:"id"`                         // ID
	LastUpdate               any    `json:"__last_update"`              // Last Modified on
	DisplayName              string `json:"display_name"`               // Display Name
	CreateUid                any    `json:"create_uid"`                 // Created by 📦 Model Relation: res.users
	CreateDate               any    `json:"create_date"`                // Created on
	WriteUid                 any    `json:"write_uid"`                  // Last Updated by 📦 Model Relation: res.users
	WriteDate                any    `json:"write_date"`                 // Last Updated on
	AliasID                  any    `json:"alias_id"`                   // Alias 📦 Model Relation: mail.alias ⭐ Required
	AddressID                any    `json:"address_id"`                 // Job Location 📦 Model Relation: res.partner
	ApplicationIDs           any    `json:"application_ids"`            // Job Applications 📦 Model Relation: hr.applicant
	ApplicationCount         int    `json:"application_count"`          // Application Count
	AllApplicationCount      int    `json:"all_application_count"`      // All Application Count
	NewApplicationCount      int    `json:"new_application_count"`      // New Application
	OldApplicationCount      int    `json:"old_application_count"`      // Old Application
	ManagerID                any    `json:"manager_id"`                 // Department Manager 📦 Model Relation: hr.employee
	UserID                   any    `json:"user_id"`                    // Recruiter 📦 Model Relation: res.users
	HrResponsibleID          any    `json:"hr_responsible_id"`          // HR Responsible 📦 Model Relation: res.users
	DocumentIDs              any    `json:"document_ids"`               // Documents 📦 Model Relation: ir.attachment
	DocumentsCount           int    `json:"documents_count"`            // Document Count
	Color                    int    `json:"color"`                      // Color Index
	IsFavorite               bool   `json:"is_favorite"`                // Is Favorite
	FavoriteUserIDs          any    `json:"favorite_user_ids"`          // Favorite User 📦 Model Relation: res.users
	SurveyID                 any    `json:"survey_id"`                  // Interview Form 📦 Model Relation: survey.survey
	WebsiteDescription       string `json:"website_description"`        // Website description
	JobOpenDate              any    `json:"job_open_date"`              // Job Start Recruitment Date
	UtmCampaignID            any    `json:"utm_campaign_id"`            // Campaign 📦 Model Relation: utm.campaign
	MaxPoints                int    `json:"max_points"`                 // Max Points
	DirectClicks             int    `json:"direct_clicks"`              // Direct Clicks
	FacebookClicks           int    `json:"facebook_clicks"`            // Facebook Clicks
	TwitterClicks            int    `json:"twitter_clicks"`             // Twitter Clicks
	LinkedinClicks           int    `json:"linkedin_clicks"`            // Linkedin Clicks
	L10NBeScaleCategory      any    `json:"l10n_be_scale_category"`     // L10N Be Scale Category
	DisplayL10NBeScale       bool   `json:"display_l10n_be_scale"`      // Display L10N Be Scale
	AliasName                string `json:"alias_name"`                 // Alias Name
	AliasModelID             any    `json:"alias_model_id"`             // Aliased Model 📦 Model Relation: ir.model ⭐ Required
	AliasUserID              any    `json:"alias_user_id"`              // Owner 📦 Model Relation: res.users
	AliasDefaults            string `json:"alias_defaults"`             // Default Values ⭐ Required
	AliasForceThreadID       int    `json:"alias_force_thread_id"`      // Record Thread ID
	AliasDomain              string `json:"alias_domain"`               // Alias domain
	AliasParentModelID       any    `json:"alias_parent_model_id"`      // Parent Model 📦 Model Relation: ir.model
	AliasParentThreadID      int    `json:"alias_parent_thread_id"`     // Parent Record Thread ID
	AliasContact             any    `json:"alias_contact"`              // Alias Contact Security ⭐ Required
	AliasBouncedContent      string `json:"alias_bounced_content"`      // Custom Bounced Message
}

func (m *Model) HRJob() {
	model := "hr.job"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(HrJob_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"name", "!=", "Administrator"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}
	fmt.Println(model, "records found:", len(records))

	for _, record := range records {
		var job HrJob_150
		FillStruct(record, &job)
		// fmt.Println(prettyprint(job.Name), prettyprint(job.DepartmentID))

		ur := map[string]any{
			"name":        job.Name,
			"description": job.Description,
		}

		_, dep_name := ParseMany2One(job.DepartmentID)
		departmentID, _ := m.Dest.GetID("hr.department", []any{
			[]any{"name", "=", dep_name},
		})
		if departmentID != -1 {
			ur["department_id"] = departmentID
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", job.Name},
		})

		m.writeRecord(model, ur, rid)

	}
}
