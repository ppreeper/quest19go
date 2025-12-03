package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

// quest15data res.users model
type ResUsers_150 struct {
	// AccessesCount int `json:"accesses_count"` // # Access Rights
	// AccountRepresentedCompanyIDs  any       `json:"account_represented_company_ids"`   // Account Represented Company 📦 relation: one2many res.company
	// ActionID                      any       `json:"action_id"`                         // Home Action 📦 relation: many2one ir.actions.actions
	// Active bool `json:"active"` // Active
	// ActiveLangCount               int       `json:"active_lang_count"`                 // Active Lang Count
	// ActivePartner                 bool      `json:"active_partner"`                    // Partner is Active
	// ActivityCalendarEventID       any       `json:"activity_calendar_event_id"`        // Next Activity Calendar Event 📦 relation: many2one calendar.event
	// ActivityDateDeadline          time.Time `json:"activity_date_deadline"`            // Next Activity Deadline
	// ActivityExceptionDecoration   any       `json:"activity_exception_decoration"`     // Activity Exception Decoration
	// ActivityExceptionIcon         string    `json:"activity_exception_icon"`           // Icon
	// ActivityIDs                   any       `json:"activity_ids"`                      // Activities 📦 relation: one2many mail.activity
	// ActivityState                 any       `json:"activity_state"`                    // Activity State
	// ActivitySummary               string    `json:"activity_summary"`                  // Next Activity Summary
	// ActivityTypeIcon              string    `json:"activity_type_icon"`                // Activity Type Icon
	// ActivityTypeID                any       `json:"activity_type_id"`                  // Next Activity Type 📦 relation: many2one mail.activity.type
	// ActivityUserID                any       `json:"activity_user_id"`                  // Responsible User 📦 relation: many2one res.users
	// AdditionalInfo                string    `json:"additional_info"`                   // Additional info
	// AdditionalNote                string    `json:"additional_note"`                   // Additional Note
	// AddressHomeID                 any       `json:"address_home_id"`                   // Address 📦 relation: many2one res.partner
	// AddressID                     any       `json:"address_id"`                        // Work Address 📦 relation: many2one res.partner
	// AllocationCount               float64   `json:"allocation_count"`                  // Total number of days allocated.
	// AllocationDisplay             string    `json:"allocation_display"`                // Allocation Display
	// AllocationUsedCount           float64   `json:"allocation_used_count"`             // Total number of days off used
	// AllocationUsedDisplay         string    `json:"allocation_used_display"`           // Allocation Used Display
	// ApiKeyIDs                     any       `json:"api_key_ids"`                       // API Keys 📦 relation: one2many res.users.apikeys
	// AttendanceState               any       `json:"attendance_state"`                  // Attendance Status
	Avatar1024 any `json:"avatar_1024"` // Avatar 1024
	Avatar128  any `json:"avatar_128"`  // Avatar 128
	Avatar1920 any `json:"avatar_1920"` // Avatar
	Avatar256  any `json:"avatar_256"`  // Avatar 256
	Avatar512  any `json:"avatar_512"`  // Avatar 512
	// BadgeIDs         any       `json:"badge_ids"`          // Badges 📦 relation: one2many gamification.badge.user
	// BankAccountCount int       `json:"bank_account_count"` // Bank
	// BankAccountID    any       `json:"bank_account_id"`    // Bank Account Number 📦 relation: many2one res.partner.bank
	// BankIDs          any       `json:"bank_ids"`           // Banks 📦 relation: one2many res.partner.bank
	// Barcode  string    `json:"barcode"`  // Badge ID
	// Birthday time.Time `json:"birthday"` // Date of Birth
	// BronzeBadge          int       `json:"bronze_badge"`            // Bronze badges count
	// CalendarLastNotifAck time.Time `json:"calendar_last_notif_ack"` // Last notification marked as read from base Calendar
	// CanAdjustCreditLimit bool      `json:"can_adjust_credit_limit"` // Can Adjust Credit Limit
	// CanEdit              bool      `json:"can_edit"`                // Can Edit
	// CanImpersonateUser   bool      `json:"can_impersonate_user"`    // Can Impersonate User
	// CanPublish           bool      `json:"can_publish"`             // Can Publish
	// CategoryID           any       `json:"category_id"`             // Tags 📦 relation: many2many res.partner.category
	// CategoryIDs          any       `json:"category_ids"`            // Employee Tags 📦 relation: many2many hr.employee.category
	// Certificate                   any       `json:"certificate"`                       // Certificate Level
	// CertificationsCompanyCount    int       `json:"certifications_company_count"`      // Company Certifications Count
	// CertificationsCount           int       `json:"certifications_count"`              // Certifications Count
	// ChannelIDs                    any       `json:"channel_ids"`                       // Channels 📦 relation: many2many mail.channel
	// ChildIDs any    `json:"child_ids"` // Contact 📦 relation: one2many res.partner
	// Children int    `json:"children"`  // Number of Children
	// City string `json:"city"` // City
	// CoachID  any    `json:"coach_id"`  // Coach 📦 relation: many2one hr.employee
	// Color    int    `json:"color"`     // Color Index
	// Comment string `json:"comment"` // Notes
	// CommercialCompanyName    string    `json:"commercial_company_name"`     // Company Name Entity
	// CommercialPartnerID      any       `json:"commercial_partner_id"`       // Commercial Entity 📦 relation: many2one res.partner
	// CompaniesCount           int       `json:"companies_count"`             // Number of Companies
	// CompanyID                any       `json:"company_id"`                  // Company 📦 relation: many2one res.company ⭐ required
	// CompanyIDs               any       `json:"company_ids"`                 // Companies 📦 relation: many2many res.company
	// CompanyName              string    `json:"company_name"`                // Company Name
	// CompanyType              any       `json:"company_type"`                // Company Type
	// ContactAddress           string    `json:"contact_address"`             // Complete Address
	// ContactAddressComplete   string    `json:"contact_address_complete"`    // Contact Address Complete
	// ContractIDs              any       `json:"contract_ids"`                // Partner Contracts 📦 relation: one2many account.analytic.account
	// CountryCode string `json:"country_code"` // Country Code
	// CountryID   any    `json:"country_id"`   // Country 📦 relation: many2one res.country
	// CountryOfBirth           any       `json:"country_of_birth"`            // Country of Birth 📦 relation: many2one res.country
	// Credit                   float64   `json:"credit"`                      // Total Receivable
	// CreditLimit              float64   `json:"credit_limit"`                // Credit Limit
	// CrmTeamIDs               any       `json:"crm_team_ids"`                // Sales Teams 📦 relation: many2many crm.team
	// CrmTeamMemberIDs         any       `json:"crm_team_member_ids"`         // Sales Team Members 📦 relation: one2many crm.team.member
	// CurrencyID               any       `json:"currency_id"`                 // Currency 📦 relation: many2one res.currency
	// CurrentLeaveState        any       `json:"current_leave_state"`         // Current Time Off Status
	// CustomerRank             int       `json:"customer_rank"`               // Customer Rank
	// Date                     time.Time `json:"date"`                        // Date
	// DateLocalization         time.Time `json:"date_localization"`           // Geolocation Date
	// Debit                    float64   `json:"debit"`                       // Total Payable
	// DebitLimit               float64   `json:"debit_limit"`                 // Payable Limit
	// DepartmentID             any       `json:"department_id"`               // Department 📦 relation: many2one hr.department
	// DependentChildren        int       `json:"dependent_children"`          // Considered number of dependent children
	// DependentJuniors         int       `json:"dependent_juniors"`           // Considered number of dependent juniors
	// DependentSeniors         int       `json:"dependent_seniors"`           // Considered number of dependent seniors
	// Disabled                 bool      `json:"disabled"`                    // Disabled
	// DisabledChildrenBool     bool      `json:"disabled_children_bool"`      // Disabled Children
	// DisabledChildrenNumber   int       `json:"disabled_children_number"`    // Number of disabled children
	// DisabledSpouseBool       bool      `json:"disabled_spouse_bool"`        // Disabled Spouse
	// DisplayName              string    `json:"display_name"`                // Display Name
	// DocumentCount            int       `json:"document_count"`              // Documents
	// DocumentIDs              any       `json:"document_ids"`                // Document 📦 relation: one2many documents.document
	Email string `json:"email"` // Email
	// EmailFormatted           string    `json:"email_formatted"`             // Formatted Email
	// EmailNormalized          string    `json:"email_normalized"`            // Normalized Email
	// EmergencyContact         string    `json:"emergency_contact"`           // Emergency Contact
	// EmergencyPhone           string    `json:"emergency_phone"`             // Emergency Phone
	Employee bool `json:"employee"` // Employee
	// EmployeeBankAccountID    any       `json:"employee_bank_account_id"`    // Employee's Bank Account Number 📦 relation: many2one res.partner.bank
	// EmployeeCarsCount        int       `json:"employee_cars_count"`         // Cars
	// EmployeeCount            int       `json:"employee_count"`              // Employee Count
	// EmployeeCountryID        any       `json:"employee_country_id"`         // Employee's Country 📦 relation: many2one res.country
	// EmployeeID               any       `json:"employee_id"`                 // Company employee 📦 relation: many2one hr.employee
	// EmployeeIDs              any       `json:"employee_ids"`                // Related employee 📦 relation: one2many hr.employee
	// EmployeeParentID         any       `json:"employee_parent_id"`          // Manager 📦 relation: many2one hr.employee
	// EmployeePhone            string    `json:"employee_phone"`              // Private Phone
	// EmployeeSkillIDs         any       `json:"employee_skill_ids"`          // Skills 📦 relation: one2many hr.employee.skill
	// EmployeeType             any       `json:"employee_type"`               // Employee Type
	// EmployeesCount           int       `json:"employees_count"`             // Employees Count
	// EquipmentCount           int       `json:"equipment_count"`             // Assigned Equipments
	// EquipmentIDs             any       `json:"equipment_ids"`               // Managed Equipments 📦 relation: one2many maintenance.equipment
	// ExpenseManagerID         any       `json:"expense_manager_id"`          // Expense 📦 relation: many2one res.users
	// FavoriteLunchProductIDs  any       `json:"favorite_lunch_product_ids"`  // Favorite Lunch Product 📦 relation: many2many lunch.product
	// FollowupLevel            any       `json:"followup_level"`              // Follow-up Level 📦 relation: many2one account_followup.followup.line
	// FollowupStatus           any       `json:"followup_status"`             // Follow-up Status
	// ForumWaitingPostsCount   int       `json:"forum_waiting_posts_count"`   // Waiting post
	Function string `json:"function"` // Job Position
	Gender   any    `json:"gender"`   // Gender
	// GoalIDs                  any       `json:"goal_ids"`                    // Goal 📦 relation: one2many gamification.goal
	// GoldBadge                int       `json:"gold_badge"`                  // Gold badges count
	// GroupsCount              int       `json:"groups_count"`                // # Groups
	// GroupsID                 any       `json:"groups_id"`                   // Groups 📦 relation: many2many res.groups
	// GstRegNo                 string    `json:"gst_reg_no"`                  // GST Registration No
	// HasMessage               bool      `json:"has_message"`                 // Has Message
	// HasUnreconciledEntries   bool      `json:"has_unreconciled_entries"`    // Has Unreconciled Entries
	// HelpdeskTargetClosed     float64   `json:"helpdesk_target_closed"`      // Target Tickets to Close
	// HelpdeskTargetRating     float64   `json:"helpdesk_target_rating"`      // Target Customer Rating
	// HelpdeskTargetSuccess    float64   `json:"helpdesk_target_success"`     // Target Success Rate
	// HoursLastMonth           float64   `json:"hours_last_month"`            // Hours Last Month
	// HoursLastMonthDisplay    string    `json:"hours_last_month_display"`    // Hours Last Month Display
	// HrIconDisplay            any       `json:"hr_icon_display"`             // Hr Icon Display
	// HrPresenceState          any       `json:"hr_presence_state"`           // Hr Presence State
	// HrReferralLevelID        any       `json:"hr_referral_level_id"`        // Hr Referral Level 📦 relation: many2one hr.referral.level
	// HrReferralOnboardingPage bool      `json:"hr_referral_onboarding_page"` // Hr Referral Onboarding Page
	ID int `json:"id"` // ID
	// IdentificationID      string `json:"identification_id"`         // Identification No
	// ImStatus              string `json:"im_status"`                 // IM Status
	Image1024   any `json:"image_1024"`   // Image 1024
	Image128    any `json:"image_128"`    // Image 128
	Image1920   any `json:"image_1920"`   // Image
	Image256    any `json:"image_256"`    // Image 256
	Image512    any `json:"image_512"`    // Image 512
	ImageMedium any `json:"image_medium"` // Medium-sized image
	// IndustryID            any    `json:"industry_id"`               // Industry 📦 relation: many2one res.partner.industry
	// InvoiceIDs            any    `json:"invoice_ids"`               // Invoices 📦 relation: one2many account.move
	// InvoiceWarn           any    `json:"invoice_warn"`              // Invoice
	// InvoiceWarnMsg        string `json:"invoice_warn_msg"`          // Message for Invoice
	// IsAbsent              bool   `json:"is_absent"`                 // Absent Today
	// IsAddressHomeACompany bool   `json:"is_address_home_a_company"` // The employee address has a company linked
	// IsBlacklisted         bool   `json:"is_blacklisted"`            // Blacklist
	// IsCompany             bool   `json:"is_company"`                // Is a Company
	// IsPublished           bool   `json:"is_published"`              // Is Published
	// IsSeoOptimized        bool   `json:"is_seo_optimized"`          // SEO optimized
	// IsSystem              bool   `json:"is_system"`                 // Is System
	JobTitle string `json:"job_title"` // Job Title
	// JournalItemCount      int    `json:"journal_item_count"`        // Journal Items
	// Karma                 int    `json:"karma"`                     // Karma
	// KarmaTrackingIDs      any    `json:"karma_tracking_ids"`        // Karma Changes 📦 relation: one2many gamification.karma.tracking
	// KmHomeWork            int    `json:"km_home_work"`              // Home-Work Distance
	// L10NBeScaleSeniority  int    `json:"l10n_be_scale_seniority"`   // Seniority at Hiring
	// L10NCaPst             string `json:"l10n_ca_pst"`               // PST
	Lang any `json:"lang"` // Language
	// LastActivity             time.Time `json:"last_activity"`              // Last Activity
	// LastActivityTime         string    `json:"last_activity_time"`         // Last Activity Time
	// LastAppraisalDate        time.Time `json:"last_appraisal_date"`        // Last Appraisal Date
	// LastAppraisalID          any       `json:"last_appraisal_id"`          // Last Appraisal 📦 relation: many2one hr.appraisal
	// LastCheckIn              time.Time `json:"last_check_in"`              // Check In
	// LastCheckOut             time.Time `json:"last_check_out"`             // Check Out
	// LastLunchLocationID      any       `json:"last_lunch_location_id"`     // Last Lunch Location 📦 relation: many2one lunch.location
	// LastTimeEntriesChecked   time.Time `json:"last_time_entries_checked"`  // Latest Invoices & Payments Matching Date
	// LeaveDateTo              time.Time `json:"leave_date_to"`              // To Date
	// LeaveManagerID           any       `json:"leave_manager_id"`           // Time Off 📦 relation: many2one res.users
	// LivechatUsername         string    `json:"livechat_username"`          // Livechat Username
	// LogIDs                   any       `json:"log_ids"`                    // User log entries 📦 relation: one2many res.users.log
	Login string `json:"login"` // Login ⭐ required
	// LoginDate                time.Time `json:"login_date"`                 // Latest authentication
	// Marital any `json:"marital"` // Marital Status
	// MeetingCount             int       `json:"meeting_count"`              // # Meetings
	// MeetingIDs               any       `json:"meeting_ids"`                // Meetings 📦 relation: many2many calendar.event
	// MessageAttachmentCount   int       `json:"message_attachment_count"`   // Attachment Count
	// MessageBounce            int       `json:"message_bounce"`             // Bounce
	// MessageFollowerIDs       any       `json:"message_follower_ids"`       // Followers 📦 relation: one2many mail.followers
	// MessageHasError          bool      `json:"message_has_error"`          // Message Delivery error
	// MessageHasErrorCounter   int       `json:"message_has_error_counter"`  // Number of errors
	// MessageHasSmsError       bool      `json:"message_has_sms_error"`      // SMS Delivery error
	// MessageIDs               any       `json:"message_ids"`                // Messages 📦 relation: one2many mail.message
	// MessageIsFollower        bool      `json:"message_is_follower"`        // Is Follower
	// MessageMainAttachmentID  any       `json:"message_main_attachment_id"` // Main Attachment 📦 relation: many2one ir.attachment
	// MessageNeedaction        bool      `json:"message_needaction"`         // Action Needed
	// MessageNeedactionCounter int       `json:"message_needaction_counter"` // Number of Actions
	// MessagePartnerIDs        any       `json:"message_partner_ids"`        // Followers (Partners) 📦 relation: many2many res.partner
	// MessageUnread            bool      `json:"message_unread"`             // Unread Messages
	// MessageUnreadCounter     int       `json:"message_unread_counter"`     // Unread Messages Counter
	Mobile string `json:"mobile"` // Mobile
	// MobileBlacklisted        bool      `json:"mobile_blacklisted"`         // Blacklisted Phone Is Mobile
	// MobilePhone              string    `json:"mobile_phone"`               // Work Mobile
	// MyActivityDateDeadline   time.Time `json:"my_activity_date_deadline"`  // My Activity Deadline
	Name string `json:"name"` // Name
	// NewPassword              string    `json:"new_password"`               // Set Password
	// NextAppraisalDate        time.Time `json:"next_appraisal_date"`        // Next Appraisal Date
	// NextRankID               any       `json:"next_rank_id"`               // Next Rank 📦 relation: many2one gamification.karma.rank
	NotificationType any `json:"notification_type"` // Notification ⭐ required
	// OcnToken                 string  `json:"ocn_token"`                  // OCN Token
	// OdoobotFailed            bool    `json:"odoobot_failed"`             // Odoobot Failed
	// OdoobotState any `json:"odoobot_state"` // OdooBot Status
	// OnTimeRate               float64 `json:"on_time_rate"`               // On-Time Delivery Rate
	// OnlinePartnerInformation string  `json:"online_partner_information"` // Online Partner Information
	// OpportunityCount         int     `json:"opportunity_count"`          // Opportunity
	// OpportunityIDs           any     `json:"opportunity_ids"`            // Opportunities 📦 relation: one2many crm.lead
	// OtherDependentPeople     bool    `json:"other_dependent_people"`     // Other Dependent People
	// OtherDisabledJuniorsDependent int       `json:"other_disabled_juniors_dependent"`  // # disabled people (<65)
	// OtherDisabledSeniorDependent  int       `json:"other_disabled_senior_dependent"`   // # disabled seniors (>=65)
	// OtherJuniorsDependent         int       `json:"other_juniors_dependent"`           // # people (<65)
	// OtherSeniorDependent          int       `json:"other_senior_dependent"`            // # seniors (>=65)
	// OutsideSalespersonID any    `json:"outside_salesperson_id"` // Outside Salesperson 📦 relation: many2one res.users
	// ParentID             any    `json:"parent_id"`              // Related Company 📦 relation: many2one res.partner
	// ParentName           string `json:"parent_name"`            // Parent name
	// PartnerGid           int    `json:"partner_gid"`            // Company database ID
	// PartnerID            any    `json:"partner_id"`             // Related Partner 📦 relation: many2one res.partner ⭐ required
	// PartnerLatitude               float64   `json:"partner_latitude"`                  // Geo Latitude
	// PartnerLongitude              float64   `json:"partner_longitude"`                 // Geo Longitude
	// PartnerShare                  bool      `json:"partner_share"`                     // Share Partner
	// PassportID                    string    `json:"passport_id"`                       // Passport No
	// Password                      string    `json:"password"`                          // Password
	// PaymentNextActionDate         time.Time `json:"payment_next_action_date"`          // Next Action Date
	// PaymentResponsibleID          any       `json:"payment_responsible_id"`            // Follow-up Responsible 📦 relation: many2one res.users
	// PaymentTokenCount             int       `json:"payment_token_count"`               // Payment Token Count
	// PaymentTokenIDs               any       `json:"payment_token_ids"`                 // Payment Tokens 📦 relation: one2many payment.token
	// PermitNo                      string    `json:"permit_no"`                         // Work Permit No
	Phone string `json:"phone"` // Phone
	// PhoneBlacklisted              bool      `json:"phone_blacklisted"`                 // Blacklisted Phone is Phone
	// PhoneMobileSearch             string    `json:"phone_mobile_search"`               // Phone/Mobile
	// PhoneSanitized                string    `json:"phone_sanitized"`                   // Sanitized Number
	// PhoneSanitizedBlacklisted     bool      `json:"phone_sanitized_blacklisted"`       // Phone Blacklisted
	// PickingWarn    any    `json:"picking_warn"`     // Stock Picking
	// PickingWarnMsg string `json:"picking_warn_msg"` // Message for Stock Picking
	Pin string `json:"pin"` // PIN
	// PlaceOfBirth string `json:"place_of_birth"` // Place of Birth
	// PlanToChangeBike              bool   `json:"plan_to_change_bike"`               // Plan To Change Bike
	// PlanToChangeCar               bool   `json:"plan_to_change_car"`                // Plan To Change Car
	// PrivateCity      string `json:"private_city"`       // Private City
	// PrivateCountryID any    `json:"private_country_id"` // Private Country 📦 relation: many2one res.country
	// PrivateEmail     string `json:"private_email"`      // Private Email
	// PrivateLang      any    `json:"private_lang"`       // Employee Lang
	// PrivateStateID   any    `json:"private_state_id"`   // Private State 📦 relation: many2one res.country.state
	// PrivateStreet    string `json:"private_street"`     // Private Street
	// PrivateStreet2   string `json:"private_street2"`    // Private Street2
	// PrivateZip       string `json:"private_zip"`        // Private Zip
	// PropertyAccountPayableID      any    `json:"property_account_payable_id"`       // Account Payable 📦 relation: many2one account.account ⭐ required
	// PropertyAccountPositionID     any    `json:"property_account_position_id"`      // Fiscal Position 📦 relation: many2one account.fiscal.position
	// PropertyAccountReceivableID   any    `json:"property_account_receivable_id"`    // Account Receivable 📦 relation: many2one account.account ⭐ required
	// PropertyDeliveryCarrierID     any    `json:"property_delivery_carrier_id"`      // Delivery Method 📦 relation: many2one delivery.carrier
	// PropertyPaymentMethodID       any    `json:"property_payment_method_id"`        // Payment Method 📦 relation: many2one account.payment.method
	// PropertyPaymentTermID         any    `json:"property_payment_term_id"`          // Customer Payment Terms 📦 relation: many2one account.payment.term
	// PropertyProductPricelist      any    `json:"property_product_pricelist"`        // Pricelist 📦 relation: many2one product.pricelist
	// PropertyPurchaseCurrencyID    any    `json:"property_purchase_currency_id"`     // Supplier Currency 📦 relation: many2one res.currency
	// PropertyStockCustomer         any    `json:"property_stock_customer"`           // Customer Location 📦 relation: many2one stock.location
	// PropertyStockSupplier         any    `json:"property_stock_supplier"`           // Vendor Location 📦 relation: many2one stock.location
	// PropertySupplierPaymentTermID any    `json:"property_supplier_payment_term_id"` // Vendor Payment Terms 📦 relation: many2one account.payment.term
	PropertyWarehouseID any `json:"property_warehouse_id"` // Default Warehouse 📦 relation: many2one stock.warehouse
	// PurchaseLineIDs               any    `json:"purchase_line_ids"`                 // Purchase Lines 📦 relation: one2many purchase.order.line
	// PurchaseOrderCount            int    `json:"purchase_order_count"`              // Purchase Order Count
	// PurchaseWarn                  any    `json:"purchase_warn"`                     // Purchase Order
	// PurchaseWarnMsg               string `json:"purchase_warn_msg"`                 // Message for Purchase Order
	// RankID                        any    `json:"rank_id"`                           // Rank 📦 relation: many2one gamification.karma.rank
	// ReceiptReminderEmail          bool   `json:"receipt_reminder_email"`            // Receipt Reminder
	Ref string `json:"ref"` // Reference
	// RefCompanyIDs                 any       `json:"ref_company_ids"`                   // Companies that refers to partner 📦 relation: one2many res.company
	// ReferralPointIDs              any       `json:"referral_point_ids"`                // Referral Point 📦 relation: one2many hr.referral.points
	// ReminderDateBeforeReceipt     int       `json:"reminder_date_before_receipt"`      // Days Before Receipt
	// RequestOvertime               bool      `json:"request_overtime"`                  // Request Overtime
	// ResUsersSettingsIDs           any       `json:"res_users_settings_ids"`            // Res Users Settings 📦 relation: one2many res.users.settings
	// ResidentBool                  bool      `json:"resident_bool"`                    // Nonresident
	// ResourceCalendarID            any       `json:"resource_calendar_id"`             // Default Working Hours 📦 relation: many2one resource.calendar
	// ResourceIDs                   any       `json:"resource_ids"`                     // Resources 📦 relation: one2many resource.resource
	// ResumeLineIDs                 any       `json:"resume_line_ids"`                  // Resumé lines 📦 relation: one2many hr.resume.line
	// RulesCount                    int       `json:"rules_count"`                      // # Record Rules
	// SaleOrderCount                int       `json:"sale_order_count"`                 // Sale Order Count
	// SaleOrderIDs                  any       `json:"sale_order_ids"`                   // Sales Order 📦 relation: one2many sale.order
	// SaleTeamID                    any       `json:"sale_team_id"`                     // User Sales Team 📦 relation: many2one crm.team
	// SaleWarn                      any       `json:"sale_warn"`                        // Sales Warnings
	// SaleWarnMsg                   string    `json:"sale_warn_msg"`                    // Message for Sales Order
	// SameVatPartnerID              any       `json:"same_vat_partner_id"`              // Partner with same Tax ID 📦 relation: many2one res.partner
	// Self                          any       `json:"self"`                             // Self 📦 relation: many2one res.partner
	// SeoName                       string    `json:"seo_name"`                         // Seo name
	// Share                         bool      `json:"share"`                            // Share User
	// ShowLeaves                    bool      `json:"show_leaves"`                      // Able to see Remaining Time Off
	// SignInitials                  any       `json:"sign_initials"`                    // Digitial Initials
	// SignRequestCount              int       `json:"sign_request_count"`               // Sign Request Count
	// SignSignature                 any       `json:"sign_signature"`                   // Digital Signature
	Signature string `json:"signature"` // Email Signature
	// SignatureCount                int       `json:"signature_count"`                  // # Signatures
	// SignupExpiration              time.Time `json:"signup_expiration"`                // Signup Expiration
	// SignupToken                   string    `json:"signup_token"`                     // Signup Token
	// SignupType                    string    `json:"signup_type"`                      // Signup Token Type
	// SignupUrl                     string    `json:"signup_url"`                       // Signup URL
	// SignupValid                   bool      `json:"signup_valid"`                     // Signup Token is Valid
	// SilverBadge                   int       `json:"silver_badge"`                     // Silver badges count
	// SlaIDs                        any       `json:"sla_ids"`                          // SLA Policies 📦 relation: many2many helpdesk.sla
	// SpouseBirthdate               time.Time `json:"spouse_birthdate"`                 // Spouse Birthdate
	// SpouseCompleteName            string    `json:"spouse_complete_name"`             // Spouse Complete Name
	// SpouseFiscalStatus            any       `json:"spouse_fiscal_status"`             // Tax status for spouse
	// SpouseFiscalStatusExplanation string    `json:"spouse_fiscal_status_explanation"` // Spouse Fiscal Status Explanation
	// State   any    `json:"state"`    // Status
	// StateID any    `json:"state_id"` // State 📦 relation: many2one res.country.state
	// Street  string `json:"street"`   // Street
	// Street2 string `json:"street2"`  // Street2
	// StudyField                    string    `json:"study_field"`                      // Field of Study
	// StudySchool                   string    `json:"study_school"`                     // School
	// SupplierInvoiceCount          int       `json:"supplier_invoice_count"`           // # Vendor Bills
	// SupplierRank                  int       `json:"supplier_rank"`                    // Supplier Rank
	// TargetSalesDone               int       `json:"target_sales_done"`                // Activities Done Target
	// TargetSalesInvoiced           int       `json:"target_sales_invoiced"`            // Invoiced in Sales Orders Target
	// TargetSalesWon                int       `json:"target_sales_won"`                 // Won in Opportunities Target
	// TaskCount                     int       `json:"task_count"`                       // # Tasks
	// TaskIDs                       any       `json:"task_ids"`                         // Tasks 📦 relation: one2many project.task
	// TeamID                        any       `json:"team_id"`                          // Sales Team 📦 relation: many2one crm.team
	// TicketCount                   int       `json:"ticket_count"`                     // Tickets
	// TimesheetManagerID            any       `json:"timesheet_manager_id"`             // Timesheet 📦 relation: many2one res.users
	Title any `json:"title"` // Title 📦 relation: many2one res.partner.title
	// TotalDue                      float64   `json:"total_due"`                        // Total Due
	// TotalInvoiced                 float64   `json:"total_invoiced"`                   // Total Invoiced
	// TotalOverdue                  float64   `json:"total_overdue"`                    // Total Overdue
	// TotalOvertime                 float64   `json:"total_overtime"`                   // Total Overtime
	// TotpEnabled                   bool      `json:"totp_enabled"`                     // Two-factor authentication
	// TotpTrustedDeviceIDs          any       `json:"totp_trusted_device_ids"`          // Trusted Devices 📦 relation: one2many auth_totp.device
	// Trust                         any       `json:"trust"`                            // Degree of trust you have in this debtor
	Type     any    `json:"type"`      // Address Type
	TZ       any    `json:"tz"`        // Timezone
	TZOffset string `json:"tz_offset"` // Timezone offset
	// UnpaidInvoices                any       `json:"unpaid_invoices"`                  // Unpaid Invoices 📦 relation: one2many account.move
	// UnreconciledAmlIDs            any       `json:"unreconciled_aml_ids"`             // Unreconciled Aml 📦 relation: one2many account.move.line
	// UserID                        any       `json:"user_id"`                          // Salesperson 📦 relation: many2one res.users
	// UserIDs                       any       `json:"user_ids"`                         // Users 📦 relation: one2many res.users
	// UserLivechatUsername          string    `json:"user_livechat_username"`           // User Livechat Username
	// UtmSourceID                   any       `json:"utm_source_id"`                    // Source 📦 relation: many2one utm.source
	VAT string `json:"vat"` // VAT/Tax ID
	// Vehicle    string    `json:"vehicle"`     // Company Vehicle
	// VisaExpire time.Time `json:"visa_expire"` // Visa Expiration Date
	// VisaNo     string    `json:"visa_no"`     // Visa No
	// VisitorIDs any       `json:"visitor_ids"` // Visitors 📦 relation: one2many website.visitor
	// Website    string    `json:"website"`     // Website Link
	// WebsiteDescription            string    `json:"website_description"`              // Website Partner Full Description
	// WebsiteID                     any       `json:"website_id"`                       // Website 📦 relation: many2one website
	// WebsiteMessageIDs             any       `json:"website_message_ids"`              // Website Messages 📦 relation: one2many mail.message
	// WebsiteMetaDescription        string    `json:"website_meta_description"`         // Website meta description
	// WebsiteMetaKeywords           string    `json:"website_meta_keywords"`            // Website meta keywords
	// WebsiteMetaOgImg              string    `json:"website_meta_og_img"`              // Website opengraph image
	// WebsiteMetaTitle              string    `json:"website_meta_title"`               // Website meta title
	// WebsitePublished              bool      `json:"website_published"`                // Visible on current website
	// WebsiteShortDescription       string    `json:"website_short_description"`        // Website Partner Short Description
	// WebsiteUrl                    string    `json:"website_url"`                      // Website URL
	// WorkEmail           string `json:"work_email"`            // Work Email
	WorkLocationID any    `json:"work_location_id"` // Work Location 📦 relation: many2one hr.work.location
	WorkPhone      string `json:"work_phone"`       // Work Phone
	// XOutsideSalesperson any    `json:"x_Outside_Salesperson"` // Not Use-Outside Salesperson 📦 relation: many2one res.users
	// XFax                string `json:"x_fax"`                 // Fax ⚠️ manual: True
	ZIP string `json:"zip"` // Zip
}

func (m *Model) ResUsers() {
	model := "res.users"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(ResUsers_150{})
	domain := []any{
		[]any{"name", "!=", "Administrator"},
	}

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, domain)
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var user ResUsers_150
		FillStruct(r, &user)
		// fmt.Println(prettyprint(user))

		ur := map[string]any{
			"name":                  user.Name,
			"login":                 user.Login,
			"email":                 user.Login,
			"lang":                  user.Lang,
			"tz":                    user.TZ,
			"notification_type":     user.NotificationType,
			"signature":             user.Signature,
			"work_phone":            user.WorkPhone,
			"work_location_id":      user.WorkLocationID,
			"vat":                   user.VAT,
			"zip":                   user.ZIP,
			"property_warehouse_id": user.PropertyWarehouseID,
			"avatar_128":            user.Avatar128,
			"avatar_256":            user.Avatar256,
			"avatar_512":            user.Avatar512,
			"avatar_1024":           user.Avatar1024,
			"avatar_1920":           user.Avatar1920,
			"image_128":             user.Image128,
			"image_256":             user.Image256,
			"image_512":             user.Image512,
			"image_1024":            user.Image1024,
			"image_1920":            user.Image1920,
			"image_medium":          user.ImageMedium,
		}
		if user.TZ == "Canada/Mountain" {
			ur["tz"] = "America/Edmonton"
		}
		// m.Log.Info(model, "ur", ur)

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", user.Name},
			[]any{"login", "=", user.Login},
		})

		// m.Log.Info(model, "rid", rid, "ur", ur)

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

func (m *Model) ResUsersCreateEmployee() {
	model := "res.users"
	fmt.Println(model, "func", trace())
	m.Log.Info(model, "func", trace())
	banner.Println(model, "create_employee")

	type tempResUsers struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	domain := []any{
		[]any{"name", "!=", "Administrator"},
	}

	records, err := m.Dest.SearchRead(model, 0, 0, []string{"id", "name"}, domain)
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	bar := progressbar.Default(-1, model)

	for _, r := range records {
		var user tempResUsers
		FillStruct(r, &user)

		rid, _ := m.Dest.GetID("hr.employee", []any{
			[]any{"name", "=", user.Name},
			[]any{"user_id", "=", user.ID},
		})

		if rid > -1 {
			continue
		}

		ur := map[string]any{
			"ids": []int{user.ID},
		}
		m.Dest.Execute(model, "action_create_employee", []any{ur})

		bar.Add(1)
	}
	bar.Finish()
}
