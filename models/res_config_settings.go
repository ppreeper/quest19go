package models

import "fmt"

type ResConfigSettings_150 struct {
	ID                                    int    `json:"id"`                                        // ID
	LastUpdate                            any    `json:"__last_update"`                             // Last Modified on
	DisplayName                           string `json:"display_name"`                              // Display Name
	CreateUid                             any    `json:"create_uid"`                                // Created by 📦 Model Relation: res.users
	CreateDate                            any    `json:"create_date"`                               // Created on
	WriteUid                              any    `json:"write_uid"`                                 // Last Updated by 📦 Model Relation: res.users
	WriteDate                             any    `json:"write_date"`                                // Last Updated on
	CompanyID                             any    `json:"company_id"`                                // Company 📦 Model Relation: res.company ⭐ Required
	UserDefaultRights                     bool   `json:"user_default_rights"`                       // Default Access Rights
	ExternalEmailServerDefault            bool   `json:"external_email_server_default"`             // Custom Email Servers
	ModuleBaseImport                      bool   `json:"module_base_import"`                        // Allow users to import data from CSV/XLS/XLSX/ODS files
	ModuleGoogleCalendar                  bool   `json:"module_google_calendar"`                    // Allow the users to synchronize their calendar  with Google Calendar
	ModuleMicrosoftCalendar               bool   `json:"module_microsoft_calendar"`                 // Allow the users to synchronize their calendar with Outlook Calendar
	ModuleMailPlugin                      bool   `json:"module_mail_plugin"`                        // Allow integration with the mail plugins
	ModuleGoogleDrive                     bool   `json:"module_google_drive"`                       // Attach Google documents to any record
	ModuleGoogleSpreadsheet               bool   `json:"module_google_spreadsheet"`                 // Google Spreadsheet
	ModuleAuthOauth                       bool   `json:"module_auth_oauth"`                         // Use external authentication providers (OAuth)
	ModuleAuthLdap                        bool   `json:"module_auth_ldap"`                          // LDAP Authentication
	ModuleBaseGengo                       bool   `json:"module_base_gengo"`                         // Translate Your Website with Gengo
	ModuleAccountInterCompanyRules        bool   `json:"module_account_inter_company_rules"`        // Manage Inter Company
	ModulePad                             bool   `json:"module_pad"`                                // Collaborative Pads
	ModuleVoip                            bool   `json:"module_voip"`                               // Asterisk (VoIP)
	ModuleWebUnsplash                     bool   `json:"module_web_unsplash"`                       // Unsplash Image Library
	ModulePartnerAutocomplete             bool   `json:"module_partner_autocomplete"`               // Partner Autocomplete
	ModuleBaseGeolocalize                 bool   `json:"module_base_geolocalize"`                   // GeoLocalize
	ModuleGoogleRecaptcha                 bool   `json:"module_google_recaptcha"`                   // reCAPTCHA
	ReportFooter                          string `json:"report_footer"`                             // Custom Report Footer
	GroupMultiCurrency                    bool   `json:"group_multi_currency"`                      // Multi-Currencies
	ExternalReportLayoutID                any    `json:"external_report_layout_id"`                 // Document Template 📦 Model Relation: ir.ui.view
	ShowEffect                            bool   `json:"show_effect"`                               // Show Effect
	CompanyCount                          int    `json:"company_count"`                             // Number of Companies
	ActiveUserCount                       int    `json:"active_user_count"`                         // Number of Active Users
	LanguageCount                         int    `json:"language_count"`                            // Number of Languages
	CompanyName                           string `json:"company_name"`                              // Company Name
	CompanyInformations                   string `json:"company_informations"`                      // Company Informations
	ProfilingEnabledUntil                 any    `json:"profiling_enabled_until"`                   // Profiling enabled until
	ModuleProductImages                   bool   `json:"module_product_images"`                     // Get product pictures using barcode
	Minlength                             int    `json:"minlength"`                                 // Minimum Password Length
	GeolocProviderID                      any    `json:"geoloc_provider_id"`                        // API 📦 Model Relation: base.geo_provider
	GeolocProviderTechname                string `json:"geoloc_provider_techname"`                  // Tech Name
	GeolocProviderGooglemapKey            string `json:"geoloc_provider_googlemap_key"`             // Google Map API Key
	RecaptchaPublicKey                    string `json:"recaptcha_public_key"`                      // Site Key
	RecaptchaPrivateKey                   string `json:"recaptcha_private_key"`                     // Secret Key
	RecaptchaMinScore                     any    `json:"recaptcha_min_score"`                       // Minimum score
	KsToggleColorFieldChange              string `json:"ks_toggle_color_field_change"`              // Toggle Color
	KsHeaderColorFieldChange              string `json:"ks_header_color_field_change"`              // Header Color
	KsHeaderTextColorFieldChange          string `json:"ks_header_text_color_field_change"`         // Header Text Color
	KsSerialNumber                        bool   `json:"ks_serial_number"`                          // Serial Number
	KsO2MSerialNumber                     bool   `json:"ks_o2m_serial_number"`                      // One2Many List Serial Number
	FailCounter                           int    `json:"fail_counter"`                              // Fail Mail
	AliasDomain                           string `json:"alias_domain"`                              // Alias Domain
	RestrictTemplateRendering             bool   `json:"restrict_template_rendering"`               // Restrict Template Rendering
	UseTwilioRtcServers                   bool   `json:"use_twilio_rtc_servers"`                    // Use Twilio ICE servers
	TwilioAccountSid                      string `json:"twilio_account_sid"`                        // Twilio Account SID
	TwilioAccountToken                    string `json:"twilio_account_token"`                      // Twilio Account Auth Token
	MapBoxToken                           string `json:"map_box_token"`                             // Token Map Box
	UnsplashAccessKey                     string `json:"unsplash_access_key"`                       // Access Key
	GroupAnalyticAccounting               bool   `json:"group_analytic_accounting"`                 // Analytic Accounting
	AuthSignupResetPassword               bool   `json:"auth_signup_reset_password"`                // Enable password reset from Login page
	AuthSignupUninvited                   any    `json:"auth_signup_uninvited"`                     // Customer Account
	AuthSignupTemplateUserID              any    `json:"auth_signup_template_user_id"`              // Template user for new users created through signup 📦 Model Relation: res.users
	DelayAlertContract                    int    `json:"delay_alert_contract"`                      // Delay alert contract outdated
	GoogleDriveAuthorizationCode          string `json:"google_drive_authorization_code"`           // Authorization Code
	GoogleDriveUri                        string `json:"google_drive_uri"`                          // URI
	IsGoogleDriveTokenGenerated           bool   `json:"is_google_drive_token_generated"`           // Refresh Token Generated
	ResourceCalendarID                    any    `json:"resource_calendar_id"`                      // Company Working Hours 📦 Model Relation: resource.calendar
	ModuleHrPresence                      bool   `json:"module_hr_presence"`                        // Advanced Presence Control
	ModuleHrSkills                        bool   `json:"module_hr_skills"`                          // Skills Management
	HrPresenceControlLogin                bool   `json:"hr_presence_control_login"`                 // Based on user status in system
	HrPresenceControlEmail                bool   `json:"hr_presence_control_email"`                 // Based on number of emails sent
	HrPresenceControlIp                   bool   `json:"hr_presence_control_ip"`                    // Based on IP Address
	ModuleHrAttendance                    bool   `json:"module_hr_attendance"`                      // Based on attendances
	HrPresenceControlEmailAmount          int    `json:"hr_presence_control_email_amount"`          // # emails to send
	HrPresenceControlIpList               string `json:"hr_presence_control_ip_list"`               // Valid IP addresses
	HrEmployeeSelfEdit                    bool   `json:"hr_employee_self_edit"`                     // Employee Editing
	CurrencyID                            any    `json:"currency_id"`                               // Currency 📦 Model Relation: res.currency ⭐ Required
	CompanyLunchMinimumThreshold          any    `json:"company_lunch_minimum_threshold"`           // Maximum Allowed Overdraft
	GroupDiscountPerSoLine                bool   `json:"group_discount_per_so_line"`                // Discounts
	GroupUom                              bool   `json:"group_uom"`                                 // Units of Measure
	GroupProductVariant                   bool   `json:"group_product_variant"`                     // Variants
	ModuleSaleProductConfigurator         bool   `json:"module_sale_product_configurator"`          // Product Configurator
	ModuleSaleProductMatrix               bool   `json:"module_sale_product_matrix"`                // Sales Grid Entry
	GroupStockPackaging                   bool   `json:"group_stock_packaging"`                     // Product Packagings
	GroupProductPricelist                 bool   `json:"group_product_pricelist"`                   // Pricelists
	GroupSalePricelist                    bool   `json:"group_sale_pricelist"`                      // Advanced Pricelists
	ProductPricelistSetting               any    `json:"product_pricelist_setting"`                 // Pricelists Method
	ProductWeightInLbs                    any    `json:"product_weight_in_lbs"`                     // Weight unit of measure
	ProductVolumeVolumeInCubicFeet        any    `json:"product_volume_volume_in_cubic_feet"`       // Volume unit of measure
	GoogleDriveUriCopy                    string `json:"google_drive_uri_copy"`                     // URI Copy
	AppraisalPlan                         bool   `json:"appraisal_plan"`                            // Automatically Generate Appraisals
	AssessmentNoteIDs                     any    `json:"assessment_note_ids"`                       // Evaluation Scale 📦 Model Relation: hr.appraisal.note
	AppraisalEmployeeFeedbackTemplate     string `json:"appraisal_employee_feedback_template"`      // Appraisal Employee Feedback Template
	AppraisalManagerFeedbackTemplate      string `json:"appraisal_manager_feedback_template"`       // Appraisal Manager Feedback Template
	DurationAfterRecruitment              int    `json:"duration_after_recruitment"`                // Create an Appraisal after recruitment
	DurationFirstAppraisal                int    `json:"duration_first_appraisal"`                  // Create a first Appraisal after
	DurationNextAppraisal                 int    `json:"duration_next_appraisal"`                   // Create a second Appraisal after
	ModuleHrAppraisalSurvey               bool   `json:"module_hr_appraisal_survey"`                // 360 Feedback
	GroupAttendanceUsePin                 bool   `json:"group_attendance_use_pin"`                  // Employee PIN
	HrAttendanceOvertime                  bool   `json:"hr_attendance_overtime"`                    // Count Extra Hours
	OvertimeStartDate                     any    `json:"overtime_start_date"`                       // Extra Hours Starting Date
	OvertimeCompanyThreshold              int    `json:"overtime_company_threshold"`                // Tolerance Time In Favor Of Company
	OvertimeEmployeeThreshold             int    `json:"overtime_employee_threshold"`               // Tolerance Time In Favor Of Employee
	DisableRedirectFirebaseDynamicLink    bool   `json:"disable_redirect_firebase_dynamic_link"`    // Disable link redirection to mobile app
	EnableOcn                             bool   `json:"enable_ocn"`                                // Push Notifications
	PartnerAutocompleteInsufficientCredit bool   `json:"partner_autocomplete_insufficient_credit"`  // Insufficient credit
	PortalAllowApiKeys                    bool   `json:"portal_allow_api_keys"`                     // Customer API Keys
	SnailmailColor                        bool   `json:"snailmail_color"`                           // Print In Color
	SnailmailCover                        bool   `json:"snailmail_cover"`                           // Add a Cover Page
	SnailmailDuplex                       bool   `json:"snailmail_duplex"`                          // Print Both sides
	ModuleSocialDemo                      bool   `json:"module_social_demo"`                        // Enable Demo Mode
	DigestEmails                          bool   `json:"digest_emails"`                             // Digest Emails
	DigestID                              any    `json:"digest_id"`                                 // Digest Email 📦 Model Relation: digest.digest
	SignTerms                             string `json:"sign_terms"`                                // Sign Terms & Conditions
	SignTermsHtml                         string `json:"sign_terms_html"`                           // Sign Terms & Conditions as a Web page
	SignTermsType                         any    `json:"sign_terms_type"`                           // Sign Terms & Conditions format
	SignPreviewReady                      bool   `json:"sign_preview_ready"`                        // Display sign preview button
	UseSignTerms                          bool   `json:"use_sign_terms"`                            // Sign Default Terms & Conditions
	FacebookUseOwnAccount                 bool   `json:"facebook_use_own_account"`                  // Use your own Facebook Account
	FacebookAppID                         string `json:"facebook_app_id"`                           // Facebook App ID
	FacebookClientSecret                  string `json:"facebook_client_secret"`                    // Facebook App Secret
	InstagramUseOwnAccount                bool   `json:"instagram_use_own_account"`                 // Use your own Instagram Account
	InstagramAppID                        string `json:"instagram_app_id"`                          // Instagram App ID
	InstagramClientSecret                 string `json:"instagram_client_secret"`                   // Instagram App Secret
	LinkedinUseOwnAccount                 bool   `json:"linkedin_use_own_account"`                  // Use your own LinkedIn Account
	LinkedinAppID                         string `json:"linkedin_app_id"`                           // App ID
	LinkedinClientSecret                  string `json:"linkedin_client_secret"`                    // App Secret
	TwitterUseOwnAccount                  bool   `json:"twitter_use_own_account"`                   // Use your own Twitter Account
	TwitterConsumerKey                    string `json:"twitter_consumer_key"`                      // Twitter Consumer Key
	TwitterConsumerSecretKey              string `json:"twitter_consumer_secret_key"`               // Twitter Consumer Secret Key
	YoutubeUseOwnAccount                  bool   `json:"youtube_use_own_account"`                   // Use your own YouTube Account
	YoutubeOauthClientID                  string `json:"youtube_oauth_client_id"`                   // YouTube OAuth Client ID
	YoutubeOauthClientSecret              string `json:"youtube_oauth_client_secret"`               // YouTube OAuth Client Secret
	HasAccountingEntries                  bool   `json:"has_accounting_entries"`                    // Has Accounting Entries
	CurrencyExchangeJournalID             any    `json:"currency_exchange_journal_id"`              // Currency Exchange Journal 📦 Model Relation: account.journal
	IncomeCurrencyExchangeAccountID       any    `json:"income_currency_exchange_account_id"`       // Gain Account 📦 Model Relation: account.account
	ExpenseCurrencyExchangeAccountID      any    `json:"expense_currency_exchange_account_id"`      // Loss Account 📦 Model Relation: account.account
	HasChartOfAccounts                    bool   `json:"has_chart_of_accounts"`                     // Company has a chart of accounts
	ChartTemplateID                       any    `json:"chart_template_id"`                         // Template 📦 Model Relation: account.chart.template
	SaleTaxID                             any    `json:"sale_tax_id"`                               // Default Sale Tax 📦 Model Relation: account.tax
	PurchaseTaxID                         any    `json:"purchase_tax_id"`                           // Default Purchase Tax 📦 Model Relation: account.tax
	TaxCalculationRoundingMethod          any    `json:"tax_calculation_rounding_method"`           // Tax calculation rounding method
	AccountJournalSuspenseAccountID       any    `json:"account_journal_suspense_account_id"`       // Bank Suspense Account 📦 Model Relation: account.account
	AccountJournalPaymentDebitAccountID   any    `json:"account_journal_payment_debit_account_id"`  // Outstanding Receipts Account 📦 Model Relation: account.account
	AccountJournalPaymentCreditAccountID  any    `json:"account_journal_payment_credit_account_id"` // Outstanding Payments Account 📦 Model Relation: account.account
	TransferAccountID                     any    `json:"transfer_account_id"`                       // Internal Transfer Account 📦 Model Relation: account.account
	ModuleAccountAccountant               bool   `json:"module_account_accountant"`                 // Accounting
	GroupAnalyticTags                     bool   `json:"group_analytic_tags"`                       // Analytic Tags
	GroupWarningAccount                   bool   `json:"group_warning_account"`                     // Warnings in Invoices
	GroupCashRounding                     bool   `json:"group_cash_rounding"`                       // Cash Rounding
	GroupShowLineSubtotalsTaxExcluded     bool   `json:"group_show_line_subtotals_tax_excluded"`    // Show line subtotals without taxes (B2B)
	GroupShowLineSubtotalsTaxIncluded     bool   `json:"group_show_line_subtotals_tax_included"`    // Show line subtotals with taxes (B2C)
	GroupShowSaleReceipts                 bool   `json:"group_show_sale_receipts"`                  // Sale Receipt
	GroupShowPurchaseReceipts             bool   `json:"group_show_purchase_receipts"`              // Purchase Receipt
	ShowLineSubtotalsTaxSelection         any    `json:"show_line_subtotals_tax_selection"`         // Line Subtotals Tax Display ⭐ Required
	ModuleAccountBudget                   bool   `json:"module_account_budget"`                     // Budget Management
	ModuleAccountPayment                  bool   `json:"module_account_payment"`                    // Invoice Online Payment
	ModuleAccountReports                  bool   `json:"module_account_reports"`                    // Dynamic Reports
	ModuleAccountCheckPrinting            bool   `json:"module_account_check_printing"`             // Allow check printing and deposits
	ModuleAccountBatchPayment             bool   `json:"module_account_batch_payment"`              // Use batch payments
	ModuleAccountSepa                     bool   `json:"module_account_sepa"`                       // SEPA Credit Transfer (SCT)
	ModuleAccountSepaDirectDebit          bool   `json:"module_account_sepa_direct_debit"`          // Use SEPA Direct Debit
	ModuleL10NFrFecImport                 bool   `json:"module_l10n_fr_fec_import"`                 // Import FEC files
	ModuleAccountBankStatementImportQif   bool   `json:"module_account_bank_statement_import_qif"`  // Import .qif files
	ModuleAccountBankStatementImportOfx   bool   `json:"module_account_bank_statement_import_ofx"`  // Import in .ofx format
	ModuleAccountBankStatementImportCsv   bool   `json:"module_account_bank_statement_import_csv"`  // Import in .csv format
	ModuleAccountBankStatementImportCamt  bool   `json:"module_account_bank_statement_import_camt"` // Import in CAMT.053 format
	ModuleCurrencyRateLive                bool   `json:"module_currency_rate_live"`                 // Automatic Currency Rates
	ModuleAccountIntrastat                bool   `json:"module_account_intrastat"`                  // Intrastat
	ModuleProductMargin                   bool   `json:"module_product_margin"`                     // Allow Product Margin
	ModuleL10NEuOss                       bool   `json:"module_l10n_eu_oss"`                        // EU Intra-community Distance Selling
	ModuleAccountTaxcloud                 bool   `json:"module_account_taxcloud"`                   // Account TaxCloud
	ModuleAccountInvoiceExtract           bool   `json:"module_account_invoice_extract"`            // Bill Digitalization
	ModuleSnailmailAccount                bool   `json:"module_snailmail_account"`                  // Snailmail
	TaxExigibility                        bool   `json:"tax_exigibility"`                           // Cash Basis
	TaxCashBasisJournalID                 any    `json:"tax_cash_basis_journal_id"`                 // Tax Cash Basis Journal 📦 Model Relation: account.journal
	AccountCashBasisBaseAccountID         any    `json:"account_cash_basis_base_account_id"`        // Base Tax Received Account 📦 Model Relation: account.account
	AccountFiscalCountryID                any    `json:"account_fiscal_country_id"`                 // Fiscal Country 📦 Model Relation: res.country
	QrCode                                bool   `json:"qr_code"`                                   // Display SEPA QR-code
	InvoiceIsPrint                        bool   `json:"invoice_is_print"`                          // Print
	InvoiceIsEmail                        bool   `json:"invoice_is_email"`                          // Send Email
	IncotermID                            any    `json:"incoterm_id"`                               // Default incoterm 📦 Model Relation: account.incoterms
	InvoiceTerms                          string `json:"invoice_terms"`                             // Terms & Conditions
	InvoiceTermsHtml                      string `json:"invoice_terms_html"`                        // Terms & Conditions as a Web page
	TermsType                             any    `json:"terms_type"`                                // Terms & Conditions format
	PreviewReady                          bool   `json:"preview_ready"`                             // Display preview button
	UseInvoiceTerms                       bool   `json:"use_invoice_terms"`                         // Default Terms & Conditions
	CountryCode                           string `json:"country_code"`                              // Country Code
	GroupUseLead                          bool   `json:"group_use_lead"`                            // Leads
	GroupUseRecurringRevenues             bool   `json:"group_use_recurring_revenues"`              // Recurring Revenues
	IsMembershipMulti                     bool   `json:"is_membership_multi"`                       // Multi Teams
	CrmUseAutoAssignment                  bool   `json:"crm_use_auto_assignment"`                   // Rule-Based Assignment
	CrmAutoAssignmentAction               any    `json:"crm_auto_assignment_action"`                // Auto Assignment Action
	CrmAutoAssignmentIntervalType         any    `json:"crm_auto_assignment_interval_type"`         // Auto Assignment Interval Unit
	CrmAutoAssignmentIntervalNumber       int    `json:"crm_auto_assignment_interval_number"`       // Repeat every
	CrmAutoAssignmentRunDatetime          any    `json:"crm_auto_assignment_run_datetime"`          // Auto Assignment Next Execution Date
	ModuleCrmIapMine                      bool   `json:"module_crm_iap_mine"`                       // Generate new leads based on their country, industries, size, etc.
	ModuleCrmIapEnrich                    bool   `json:"module_crm_iap_enrich"`                     // Enrich your leads automatically with company data based on their email address.
	ModuleWebsiteCrmIapReveal             bool   `json:"module_website_crm_iap_reveal"`             // Create Leads/Opportunities from your website's traffic
	LeadEnrichAuto                        any    `json:"lead_enrich_auto"`                          // Enrich lead automatically
	LeadMiningInPipeline                  bool   `json:"lead_mining_in_pipeline"`                   // Create a lead mining request directly from the opportunity pipeline.
	PredictiveLeadScoringStartDate        any    `json:"predictive_lead_scoring_start_date"`        // Lead Scoring Starting Date
	PredictiveLeadScoringStartDateStr     string `json:"predictive_lead_scoring_start_date_str"`    // Lead Scoring Starting Date in String
	PredictiveLeadScoringFields           any    `json:"predictive_lead_scoring_fields"`            // Lead Scoring Frequency Fields 📦 Model Relation: crm.lead.scoring.frequency.field
	PredictiveLeadScoringFieldsStr        string `json:"predictive_lead_scoring_fields_str"`        // Lead Scoring Frequency Fields in String
	PredictiveLeadScoringFieldLabels      string `json:"predictive_lead_scoring_field_labels"`      // Predictive Lead Scoring Field Labels
	ModuleWebsiteHrRecruitment            bool   `json:"module_website_hr_recruitment"`             // Online Posting
	ModuleHrRecruitmentSurvey             bool   `json:"module_hr_recruitment_survey"`              // Interview Forms
	GroupMassMailingCampaign              bool   `json:"group_mass_mailing_campaign"`               // Mailing Campaigns
	MassMailingOutgoingMailServer         bool   `json:"mass_mailing_outgoing_mail_server"`         // Dedicated Server
	MassMailingMailServerID               any    `json:"mass_mailing_mail_server_id"`               // Mail Server 📦 Model Relation: ir.mail_server
	ShowBlacklistButtons                  bool   `json:"show_blacklist_buttons"`                    // Blacklist Option when Unsubscribing
	PlanningGenerationInterval            int    `json:"planning_generation_interval"`              // Rate Of Shift Generation ⭐ Required
	PlanningAllowSelfUnassign             bool   `json:"planning_allow_self_unassign"`              // Allow Unassignment
	PlanningSelfUnassignDaysBefore        int    `json:"planning_self_unassign_days_before"`        // Days before shift for unassignment
	ModuleProjectForecast                 bool   `json:"module_project_forecast"`                   // Planning
	ModuleHrTimesheet                     bool   `json:"module_hr_timesheet"`                       // Task Logs
	GroupSubtaskProject                   bool   `json:"group_subtask_project"`                     // Sub-tasks
	GroupProjectRating                    bool   `json:"group_project_rating"`                      // Customer Ratings
	GroupProjectStages                    bool   `json:"group_project_stages"`                      // Project Stages
	GroupProjectRecurringTasks            bool   `json:"group_project_recurring_tasks"`             // Recurring Tasks
	GroupProjectTaskDependencies          bool   `json:"group_project_task_dependencies"`           // Task Dependencies
	ModuleProductExpiry                   bool   `json:"module_product_expiry"`                     // Expiration Dates
	GroupStockProductionLot               bool   `json:"group_stock_production_lot"`                // Lots & Serial Numbers
	GroupLotOnDeliverySlip                bool   `json:"group_lot_on_delivery_slip"`                // Display Lots & Serial Numbers on Delivery Slips
	GroupStockTrackingLot                 bool   `json:"group_stock_tracking_lot"`                  // Packages
	GroupStockTrackingOwner               bool   `json:"group_stock_tracking_owner"`                // Consignment
	GroupStockAdvLocation                 bool   `json:"group_stock_adv_location"`                  // Multi-Step Routes
	GroupWarningStock                     bool   `json:"group_warning_stock"`                       // Warnings for Stock
	GroupStockSignDelivery                bool   `json:"group_stock_sign_delivery"`                 // Signature
	ModuleStockPickingBatch               bool   `json:"module_stock_picking_batch"`                // Batch Transfers
	GroupStockPickingWave                 bool   `json:"group_stock_picking_wave"`                  // Wave Transfers
	ModuleStockBarcode                    bool   `json:"module_stock_barcode"`                      // Barcode Scanner
	StockMoveEmailValidation              bool   `json:"stock_move_email_validation"`               // Email Confirmation picking
	StockMailConfirmationTemplateID       any    `json:"stock_mail_confirmation_template_id"`       // Email Template confirmation picking 📦 Model Relation: mail.template
	ModuleStockSms                        bool   `json:"module_stock_sms"`                          // SMS Confirmation
	ModuleDelivery                        bool   `json:"module_delivery"`                           // Delivery Methods
	ModuleDeliveryDhl                     bool   `json:"module_delivery_dhl"`                       // DHL Express Connector
	ModuleDeliveryFedex                   bool   `json:"module_delivery_fedex"`                     // FedEx Connector
	ModuleDeliveryUps                     bool   `json:"module_delivery_ups"`                       // UPS Connector
	ModuleDeliveryUsps                    bool   `json:"module_delivery_usps"`                      // USPS Connector
	ModuleDeliveryBpost                   bool   `json:"module_delivery_bpost"`                     // bpost Connector
	ModuleDeliveryEasypost                bool   `json:"module_delivery_easypost"`                  // Easypost Connector
	ModuleQualityControl                  bool   `json:"module_quality_control"`                    // Quality
	ModuleQualityControlWorksheet         bool   `json:"module_quality_control_worksheet"`          // Quality Worksheet
	GroupStockMultiLocations              bool   `json:"group_stock_multi_locations"`               // Storage Locations
	GroupStockStorageCategories           bool   `json:"group_stock_storage_categories"`            // Storage Categories
	AnnualInventoryMonth                  any    `json:"annual_inventory_month"`                    // Annual Inventory Month
	AnnualInventoryDay                    int    `json:"annual_inventory_day"`                      // Day of the month
	GroupStockReceptionReport             bool   `json:"group_stock_reception_report"`              // Reception Report
	GroupStockAutoReceptionReport         bool   `json:"group_stock_auto_reception_report"`         // Show Reception Report at Validation
	WebsiteID                             any    `json:"website_id"`                                // website 📦 Model Relation: website
	WebsiteName                           string `json:"website_name"`                              // Website Name
	WebsiteDomain                         string `json:"website_domain"`                            // Website Domain
	WebsiteCountryGroupIDs                any    `json:"website_country_group_ids"`                 // Country Groups 📦 Model Relation: res.country.group
	WebsiteCompanyID                      any    `json:"website_company_id"`                        // Website Company 📦 Model Relation: res.company
	WebsiteLogo                           any    `json:"website_logo"`                              // Website Logo
	LanguageIDs                           any    `json:"language_ids"`                              // Languages 📦 Model Relation: res.lang
	WebsiteLanguageCount                  int    `json:"website_language_count"`                    // Number of languages
	WebsiteDefaultLangID                  any    `json:"website_default_lang_id"`                   // Default language 📦 Model Relation: res.lang
	WebsiteDefaultLangCode                string `json:"website_default_lang_code"`                 // Default language code
	SpecificUserAccount                   bool   `json:"specific_user_account"`                     // Specific User Account
	WebsiteCookiesBar                     bool   `json:"website_cookies_bar"`                       // Cookies Bar
	GoogleAnalyticsKey                    string `json:"google_analytics_key"`                      // Google Analytics Key
	GoogleManagementClientID              string `json:"google_management_client_id"`               // Google Client ID
	GoogleManagementClientSecret          string `json:"google_management_client_secret"`           // Google Client Secret
	GoogleSearchConsole                   string `json:"google_search_console"`                     // Google Search Console
	CdnActivated                          bool   `json:"cdn_activated"`                             // Content Delivery Network (CDN)
	CdnUrl                                string `json:"cdn_url"`                                   // CDN Base URL
	CdnFilters                            string `json:"cdn_filters"`                               // CDN Filters
	SocialTwitter                         string `json:"social_twitter"`                            // Twitter Account
	SocialFacebook                        string `json:"social_facebook"`                           // Facebook Account
	SocialGithub                          string `json:"social_github"`                             // GitHub Account
	SocialLinkedin                        string `json:"social_linkedin"`                           // LinkedIn Account
	SocialYoutube                         string `json:"social_youtube"`                            // Youtube Account
	SocialInstagram                       string `json:"social_instagram"`                          // Instagram Account
	HasSocialNetwork                      bool   `json:"has_social_network"`                        // Configure Social Network
	Favicon                               any    `json:"favicon"`                                   // Favicon
	SocialDefaultImage                    any    `json:"social_default_image"`                      // Default Social Share Image
	GoogleMapsApiKey                      string `json:"google_maps_api_key"`                       // Google Maps API Key
	GroupMultiWebsite                     bool   `json:"group_multi_website"`                       // Multi-website
	HasGoogleAnalytics                    bool   `json:"has_google_analytics"`                      // Google Analytics
	HasGoogleAnalyticsDashboard           bool   `json:"has_google_analytics_dashboard"`            // Google Analytics Dashboard
	HasGoogleMaps                         bool   `json:"has_google_maps"`                           // Google Maps
	HasDefaultShareImage                  bool   `json:"has_default_share_image"`                   // Use a image by default for sharing
	HasGoogleSearchConsole                bool   `json:"has_google_search_console"`                 // Console Google Search
	FiscalyearLastDay                     int    `json:"fiscalyear_last_day"`                       // Fiscalyear Last Day ⭐ Required
	FiscalyearLastMonth                   any    `json:"fiscalyear_last_month"`                     // Fiscalyear Last Month ⭐ Required
	PeriodLockDate                        any    `json:"period_lock_date"`                          // Lock Date for Non-Advisers
	FiscalyearLockDate                    any    `json:"fiscalyear_lock_date"`                      // Lock Date for All Users
	TaxLockDate                           any    `json:"tax_lock_date"`                             // Tax Lock Date
	UseAngloSaxon                         bool   `json:"use_anglo_saxon"`                           // Anglo-Saxon Accounting
	ModuleAccountPredictiveBills          bool   `json:"module_account_predictive_bills"`           // Account Predictive Bills
	InvoicingSwitchThreshold              any    `json:"invoicing_switch_threshold"`                // Invoicing Switch Threshold
	GroupFiscalYear                       bool   `json:"group_fiscal_year"`                         // Fiscal Years
	AccountCheckPrintingLayout            any    `json:"account_check_printing_layout"`             // Check Layout
	AccountCheckPrintingDateLabel         bool   `json:"account_check_printing_date_label"`         // Print Date Label
	AccountCheckPrintingMultiStub         bool   `json:"account_check_printing_multi_stub"`         // Multi-Pages Check Stub
	AccountCheckPrintingMarginTop         any    `json:"account_check_printing_margin_top"`         // Check Top Margin
	AccountCheckPrintingMarginLeft        any    `json:"account_check_printing_margin_left"`        // Check Left Margin
	AccountCheckPrintingMarginRight       any    `json:"account_check_printing_margin_right"`       // Check Right Margin
	ExtractShowOcrOptionSelection         any    `json:"extract_show_ocr_option_selection"`         // Processing Option
	ExtractSingleLinePerTax               bool   `json:"extract_single_line_per_tax"`               // OCR Single Invoice Line Per Tax
	VatCheckVies                          bool   `json:"vat_check_vies"`                            // Verify VAT Numbers
	CurrencyIntervalUnit                  any    `json:"currency_interval_unit"`                    // Interval Unit
	CurrencyProvider                      any    `json:"currency_provider"`                         // Service Provider
	CurrencyNextExecutionDate             any    `json:"currency_next_execution_date"`              // Next Execution Date
	DocumentsHrSettings                   bool   `json:"documents_hr_settings"`                     // Human Resources
	DocumentsHrFolder                     any    `json:"documents_hr_folder"`                       // hr default workspace 📦 Model Relation: documents.folder
	DocumentsProductSettings              bool   `json:"documents_product_settings"`                // Product
	ProductFolder                         any    `json:"product_folder"`                            // product default workspace 📦 Model Relation: documents.folder
	ProductTags                           any    `json:"product_tags"`                              // Product Tags 📦 Model Relation: documents.tag
	DocumentsProjectSettings              bool   `json:"documents_project_settings"`                // Project
	ProjectFolder                         any    `json:"project_folder"`                            // project default workspace 📦 Model Relation: documents.folder
	ProjectTags                           any    `json:"project_tags"`                              // Project Tags 📦 Model Relation: documents.tag
	DocumentsSpreadsheetFolderID          any    `json:"documents_spreadsheet_folder_id"`           // Documents Spreadsheet Folder 📦 Model Relation: documents.folder
	ExpenseAliasPrefix                    string `json:"expense_alias_prefix"`                      // Default Alias Name for Expenses
	UseMailgateway                        bool   `json:"use_mailgateway"`                           // Let your employees record expenses by email
	ModuleHrPayrollExpense                bool   `json:"module_hr_payroll_expense"`                 // Reimburse Expenses in Payslip
	ModuleHrExpenseExtract                bool   `json:"module_hr_expense_extract"`                 // Send bills to OCR to generate expenses
	ModuleL10NFrHrPayroll                 bool   `json:"module_l10n_fr_hr_payroll"`                 // French Payroll
	ModuleL10NBeHrPayroll                 bool   `json:"module_l10n_be_hr_payroll"`                 // Belgium Payroll
	ModuleL10NInHrPayroll                 bool   `json:"module_l10n_in_hr_payroll"`                 // Indian Payroll
	ModuleHrPayrollAccount                bool   `json:"module_hr_payroll_account"`                 // Payroll with Accounting
	ModuleHrPayrollAccountSepa            bool   `json:"module_hr_payroll_account_sepa"`            // Payroll with SEPA payment
	ModuleProjectTimesheetSynchro         bool   `json:"module_project_timesheet_synchro"`          // Awesome Timesheet
	ModuleProjectTimesheetHolidays        bool   `json:"module_project_timesheet_holidays"`         // Time Off
	ReminderUserAllow                     bool   `json:"reminder_user_allow"`                       // Employee Reminder
	ReminderManagerAllow                  bool   `json:"reminder_manager_allow"`                    // Manager Reminder
	ProjectTimeModeID                     any    `json:"project_time_mode_id"`                      // Project Time Unit 📦 Model Relation: uom.uom
	TimesheetEncodeUomID                  any    `json:"timesheet_encode_uom_id"`                   // Encoding Unit 📦 Model Relation: uom.uom
	IsEncodeUomDays                       bool   `json:"is_encode_uom_days"`                        // Is Encode Uom Days
	ManufacturingLead                     any    `json:"manufacturing_lead"`                        // Manufacturing Lead Time
	UseManufacturingLead                  bool   `json:"use_manufacturing_lead"`                    // Default Manufacturing Lead Time
	GroupMrpByproducts                    bool   `json:"group_mrp_byproducts"`                      // By-Products
	ModuleMrpMps                          bool   `json:"module_mrp_mps"`                            // Master Production Schedule
	ModuleMrpPlm                          bool   `json:"module_mrp_plm"`                            // Product Lifecycle Management (PLM)
	ModuleMrpWorkorder                    bool   `json:"module_mrp_workorder"`                      // Work Orders
	ModuleMrpSubcontracting               bool   `json:"module_mrp_subcontracting"`                 // Subcontracting
	GroupMrpRoutings                      bool   `json:"group_mrp_routings"`                        // MRP Work Orders
	GroupUnlockedByDefault                bool   `json:"group_unlocked_by_default"`                 // Unlock Manufacturing Orders
	LockConfirmedPo                       bool   `json:"lock_confirmed_po"`                         // Lock Confirmed Orders
	PoLock                                any    `json:"po_lock"`                                   // Purchase Order Modification *
	PoOrderApproval                       bool   `json:"po_order_approval"`                         // Purchase Order Approval
	PoDoubleValidation                    any    `json:"po_double_validation"`                      // Levels of Approvals *
	PoDoubleValidationAmount              any    `json:"po_double_validation_amount"`               // Minimum Amount
	CompanyCurrencyID                     any    `json:"company_currency_id"`                       // Company Currency 📦 Model Relation: res.currency
	DefaultPurchaseMethod                 any    `json:"default_purchase_method"`                   // Bill Control
	GroupWarningPurchase                  bool   `json:"group_warning_purchase"`                    // Purchase Warnings
	ModuleAccount3WayMatch                bool   `json:"module_account_3way_match"`                 // 3-way matching: purchases, receptions and bills
	ModulePurchaseRequisition             bool   `json:"module_purchase_requisition"`               // Purchase Agreements
	ModulePurchaseProductMatrix           bool   `json:"module_purchase_product_matrix"`            // Purchase Grid Entry
	PoLead                                any    `json:"po_lead"`                                   // Purchase Lead Time
	UsePoLead                             bool   `json:"use_po_lead"`                               // Security Lead Time for Purchase
	GroupSendReminder                     bool   `json:"group_send_reminder"`                       // Receipt Reminder
	InvoiceIsSnailmail                    bool   `json:"invoice_is_snailmail"`                      // Send by Post
	FirebaseEnablePushNotifications       bool   `json:"firebase_enable_push_notifications"`        // Enable Web Push Notifications
	FirebaseUseOwnAccount                 bool   `json:"firebase_use_own_account"`                  // Use your own Firebase account
	FirebaseProjectID                     string `json:"firebase_project_id"`                       // Firebase Project ID
	FirebaseWebApiKey                     string `json:"firebase_web_api_key"`                      // Firebase Web API Key
	FirebasePushCertificateKey            string `json:"firebase_push_certificate_key"`             // Firebase Push Certificate Key
	FirebaseSenderID                      string `json:"firebase_sender_id"`                        // Firebase Sender ID
	FirebaseAdminKeyFile                  any    `json:"firebase_admin_key_file"`                   // Firebase Admin Key File
	NotificationRequestTitle              string `json:"notification_request_title"`                // Notification Request Title
	NotificationRequestBody               string `json:"notification_request_body"`                 // Notification Request Text
	NotificationRequestDelay              int    `json:"notification_request_delay"`                // Notification Request Delay (seconds)
	NotificationRequestIcon               any    `json:"notification_request_icon"`                 // Notification Request Icon
	ModuleStockLandedCosts                bool   `json:"module_stock_landed_costs"`                 // Landed Costs
	GroupLotOnInvoice                     bool   `json:"group_lot_on_invoice"`                      // Display Lots & Serial Numbers on Invoices
	BarcodeNomenclatureID                 any    `json:"barcode_nomenclature_id"`                   // Nomenclature 📦 Model Relation: barcode.nomenclature
	StockBarcodeDemoActive                bool   `json:"stock_barcode_demo_active"`                 // Demo Data Active
	ShowBarcodeNomenclature               bool   `json:"show_barcode_nomenclature"`                 // Show Barcode Nomenclature
	StockMoveSmsValidation                bool   `json:"stock_move_sms_validation"`                 // SMS Validation with stock move
	StockSmsConfirmationTemplateID        any    `json:"stock_sms_confirmation_template_id"`        // SMS Template 📦 Model Relation: sms.template
	CrmDefaultTeamID                      any    `json:"crm_default_team_id"`                       // Default Sales Team 📦 Model Relation: crm.team
	CrmDefaultUserID                      any    `json:"crm_default_user_id"`                       // Default Salesperson 📦 Model Relation: res.users
	ChannelID                             any    `json:"channel_id"`                                // Website Live Channel 📦 Model Relation: im_livechat.channel
	TotalsBelowSections                   bool   `json:"totals_below_sections"`                     // Add totals below sections
	AccountTaxPeriodicity                 any    `json:"account_tax_periodicity"`                   // Periodicity ⭐ Required
	AccountTaxPeriodicityReminderDay      int    `json:"account_tax_periodicity_reminder_day"`      // Reminder ⭐ Required
	AccountTaxPeriodicityJournalID        any    `json:"account_tax_periodicity_journal_id"`        // Journal 📦 Model Relation: account.journal
	DocumentsHrContractsTags              any    `json:"documents_hr_contracts_tags"`               // Contracts 📦 Model Relation: documents.tag
	DocumentsPayrollFolderID              any    `json:"documents_payroll_folder_id"`               // Documents Payroll Folder 📦 Model Relation: documents.folder
	DocumentsHrPayslipsTags               any    `json:"documents_hr_payslips_tags"`                // Payslip 📦 Model Relation: documents.tag
	DocumentsRecruitmentSettings          bool   `json:"documents_recruitment_settings"`            // Recruitment
	RecruitmentFolderID                   any    `json:"recruitment_folder_id"`                     // Recruitment default workspace 📦 Model Relation: documents.folder
	RecruitmentTagIDs                     any    `json:"recruitment_tag_ids"`                       // Recruitment Tags 📦 Model Relation: documents.tag
	DeferredTimeOffManager                any    `json:"deferred_time_off_manager"`                 // Deferred Time Off Manager 📦 Model Relation: res.users
	GroupMrpWoTabletTimer                 bool   `json:"group_mrp_wo_tablet_timer"`                 // Timer
	InternalProjectID                     any    `json:"internal_project_id"`                       // Internal Project 📦 Model Relation: project.project ⭐ Required
	LeaveTimesheetTaskID                  any    `json:"leave_timesheet_task_id"`                   // Time Off Task 📦 Model Relation: project.task
	ModuleStockDropshipping               bool   `json:"module_stock_dropshipping"`                 // Dropshipping
	DaysToPurchase                        any    `json:"days_to_purchase"`                          // Days to Purchase
	IsInstalledSale                       bool   `json:"is_installed_sale"`                         // Is the Sale Module Installed
	GroupAutoDoneSetting                  bool   `json:"group_auto_done_setting"`                   // Lock Confirmed Sales
	ModuleSaleMargin                      bool   `json:"module_sale_margin"`                        // Margins
	QuotationValidityDays                 int    `json:"quotation_validity_days"`                   // Default Quotation Validity (Days)
	UseQuotationValidityDays              bool   `json:"use_quotation_validity_days"`               // Default Quotation Validity
	GroupWarningSale                      bool   `json:"group_warning_sale"`                        // Sale Order Warnings
	PortalConfirmationSign                bool   `json:"portal_confirmation_sign"`                  // Online Signature
	PortalConfirmationPay                 bool   `json:"portal_confirmation_pay"`                   // Online Payment
	GroupSaleDeliveryAddress              bool   `json:"group_sale_delivery_address"`               // Customer Addresses
	GroupProformaSales                    bool   `json:"group_proforma_sales"`                      // Pro-Forma Invoice
	DefaultInvoicePolicy                  any    `json:"default_invoice_policy"`                    // Invoicing Policy
	DepositDefaultProductID               any    `json:"deposit_default_product_id"`                // Deposit Product 📦 Model Relation: product.product
	ModuleProductEmailTemplate            bool   `json:"module_product_email_template"`             // Specific Email
	ModuleSaleCoupon                      bool   `json:"module_sale_coupon"`                        // Coupons & Promotions
	ModuleSaleAmazon                      bool   `json:"module_sale_amazon"`                        // Amazon Sync
	AutomaticInvoice                      bool   `json:"automatic_invoice"`                         // Automatic Invoice
	InvoiceMailTemplateID                 any    `json:"invoice_mail_template_id"`                  // Invoice Email Template 📦 Model Relation: mail.template
	ConfirmationMailTemplateID            any    `json:"confirmation_mail_template_id"`             // Confirmation Email Template 📦 Model Relation: mail.template
	ReminderUserDelay                     int    `json:"reminder_user_delay"`                       // Days to Remind User
	ReminderUserInterval                  any    `json:"reminder_user_interval"`                    // User Reminder Frequency ⭐ Required
	ReminderManagerDelay                  int    `json:"reminder_manager_delay"`                    // Days to Remind Manager
	ReminderManagerInterval               any    `json:"reminder_manager_interval"`                 // Manager Reminder Frequency ⭐ Required
	TimesheetMinDuration                  int    `json:"timesheet_min_duration"`                    // Minimal Duration
	TimesheetRounding                     int    `json:"timesheet_rounding"`                        // Round up
	DocumentsAccountSettings              bool   `json:"documents_account_settings"`                // Accounting
	AccountFolder                         any    `json:"account_folder"`                            // account default folder 📦 Model Relation: documents.folder
	ExpenseExtractShowOcrOptionSelection  any    `json:"expense_extract_show_ocr_option_selection"` // Expense processing option
	ModuleIndustryFsmReport               bool   `json:"module_industry_fsm_report"`                // Worksheets
	ModuleIndustryFsmSale                 bool   `json:"module_industry_fsm_sale"`                  // Time and Material
	GroupIndustryFsmQuotations            bool   `json:"group_industry_fsm_quotations"`             // Extra Quotations
	DefaultCommissionOnTarget             any    `json:"default_commission_on_target"`              // Commission on Target
	DefaultFuelCard                       any    `json:"default_fuel_card"`                         // Fuel Card
	DefaultRepresentationFees             any    `json:"default_representation_fees"`               // Representation Fees
	DefaultInternet                       any    `json:"default_internet"`                          // Internet
	DefaultMobile                         any    `json:"default_mobile"`                            // Mobile
	DefaultMealVoucherAmount              any    `json:"default_meal_voucher_amount"`               // Meal Vouchers
	DefaultEcoChecks                      any    `json:"default_eco_checks"`                        // Eco Vouchers
	OnssCompanyID                         string `json:"onss_company_id"`                           // ONSS Company ID
	OnssRegistrationNumber                string `json:"onss_registration_number"`                  // ONSS Registration Number
	DmfaEmployerClass                     string `json:"dmfa_employer_class"`                       // DMFA Employer Class
	L10NBeCompanyNumber                   string `json:"l10n_be_company_number"`                    // Company Number
	L10NBeRevenueCode                     string `json:"l10n_be_revenue_code"`                      // Revenue Code
	HospitalInsuranceAmountChild          any    `json:"hospital_insurance_amount_child"`           // Hospital Insurance Amount per Child
	HospitalInsuranceAmountAdult          any    `json:"hospital_insurance_amount_adult"`           // Hospital Insurance Amount per Adult
	L10NBeFfeEmployerType                 any    `json:"l10n_be_ffe_employer_type"`                 // L10N Be Ffe Employer Type
	ManufacturingPeriod                   any    `json:"manufacturing_period"`                      // Manufacturing Period
	ManufacturingPeriodToDisplay          int    `json:"manufacturing_period_to_display"`           // Number of Manufacturing Period Columns
	GroupSaleOrderTemplate                bool   `json:"group_sale_order_template"`                 // Quotation Templates
	CompanySoTemplateID                   any    `json:"company_so_template_id"`                    // Default Template 📦 Model Relation: sale.order.template
	ModuleSaleQuotationBuilder            bool   `json:"module_sale_quotation_builder"`             // Quotation Builder
	SecurityLead                          any    `json:"security_lead"`                             // Security Lead Time
	GroupDisplayIncoterm                  bool   `json:"group_display_incoterm"`                    // Incoterms
	UseSecurityLead                       bool   `json:"use_security_lead"`                         // Security Lead Time for Sales
	DefaultPickingPolicy                  any    `json:"default_picking_policy"`                    // Picking Policy ⭐ Required
	MaxUnusedCars                         int    `json:"max_unused_cars"`                           // Maximum unused cars
	ShDisableFollowerConfirmSale          bool   `json:"sh_disable_follower_confirm_sale"`          // Disable to add followers by Confirm Quotation
	ShDisableFollowerValidateInvoice      bool   `json:"sh_disable_follower_validate_invoice"`      // Disable to add Followers by Validate Invoice/Bill
	ShDisableFollowerEmail                bool   `json:"sh_disable_follower_email"`                 // Disable to add Followers by Send by Email
	ShDisableFollowerCreatePicking        bool   `json:"sh_disable_follower_create_picking"`        // Disable to add Followers by create/update picking
	ShDisableFollowerSalesperson          bool   `json:"sh_disable_follower_salesperson"`           // Disable to add Salesperson as followers
	ShDisableFollowerPrPurchase           bool   `json:"sh_disable_follower_pr_purchase"`           // Disable to add Purchase Representative as followers
	ShDisableFollowerResponsiblePicking   bool   `json:"sh_disable_follower_responsible_picking"`   // Disable to add Responsible as followers
	ShDisableFollowerSalespersonAccount   bool   `json:"sh_disable_follower_salesperson_account"`   // Disable to add Salesperson as followers
	ModuleRmaEnterpriseEpt                bool   `json:"module_rma_enterprise_ept"`                 // Module Rma Enterprise Ept
	RmaTemplate                           bool   `json:"rma_template"`                              // RMA Templates
	RmaTemplateID                         any    `json:"rma_template_id"`                           // Rma Template 📦 Model Relation: mail.template
	InvoicePolicy                         bool   `json:"invoice_policy"`                            // Invoice Policy
	InvoicedTimesheet                     any    `json:"invoiced_timesheet"`                        // Timesheets Invoicing
	MtrWorkspaceID                        any    `json:"mtr_workspace_id"`                          // MTR Folder 📦 Model Relation: documents.folder
}

func (m *Model) ResConfigSettings() {
	model := "res.config.settings"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// No data to migrate
	sourceFields := ExtractJSONTags(ResConfigSettings_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println(model, "records found:", len(records))

	for _, r := range records {
		var record ResConfigSettings_150
		FillStruct(r, &record)

		fmt.Println(prettyprint(record))
	}
}
