package models

import "fmt"

type ResCompany_150 struct {
	Name                                        string `json:"name"`                                             // Company Name ⭐ Required
	Sequence                                    int    `json:"sequence"`                                         // Sequence
	ParentID                                    any    `json:"parent_id"`                                        // Parent Company 📦 Model Relation: res.company
	ChildIDs                                    any    `json:"child_ids"`                                        // Child Companies 📦 Model Relation: res.company
	PartnerID                                   any    `json:"partner_id"`                                       // Partner 📦 Model Relation: res.partner ⭐ Required
	ReportHeader                                string `json:"report_header"`                                    // Company Tagline
	ReportFooter                                string `json:"report_footer"`                                    // Report Footer
	CompanyDetails                              string `json:"company_details"`                                  // Company Details
	Logo                                        any    `json:"logo"`                                             // Company Logo
	LogoWeb                                     any    `json:"logo_web"`                                         // Logo Web
	CurrencyID                                  any    `json:"currency_id"`                                      // Currency 📦 Model Relation: res.currency ⭐ Required
	UserIDs                                     any    `json:"user_ids"`                                         // Accepted Users 📦 Model Relation: res.users
	Street                                      string `json:"street"`                                           // Street
	Street2                                     string `json:"street2"`                                          // Street2
	Zip                                         string `json:"zip"`                                              // Zip
	City                                        string `json:"city"`                                             // City
	StateID                                     any    `json:"state_id"`                                         // Fed. State 📦 Model Relation: res.country.state
	BankIDs                                     any    `json:"bank_ids"`                                         // Banks 📦 Model Relation: res.partner.bank
	CountryID                                   any    `json:"country_id"`                                       // Country 📦 Model Relation: res.country
	Email                                       string `json:"email"`                                            // Email
	Phone                                       string `json:"phone"`                                            // Phone
	Mobile                                      string `json:"mobile"`                                           // Mobile
	Website                                     string `json:"website"`                                          // Website Link
	Vat                                         string `json:"vat"`                                              // Tax ID
	CompanyRegistry                             string `json:"company_registry"`                                 // Company Registry
	PaperformatID                               any    `json:"paperformat_id"`                                   // Paper format 📦 Model Relation: report.paperformat
	ExternalReportLayoutID                      any    `json:"external_report_layout_id"`                        // Document Template 📦 Model Relation: ir.ui.view
	BaseOnboardingCompanyState                  any    `json:"base_onboarding_company_state"`                    // State of the onboarding company step
	Favicon                                     any    `json:"favicon"`                                          // Company Favicon
	Font                                        any    `json:"font"`                                             // Font
	PrimaryColor                                string `json:"primary_color"`                                    // Primary Color
	SecondaryColor                              string `json:"secondary_color"`                                  // Secondary Color
	LayoutBackground                            any    `json:"layout_background"`                                // Layout Background ⭐ Required
	LayoutBackgroundImage                       any    `json:"layout_background_image"`                          // Background Image
	ID                                          int    `json:"id"`                                               // ID
	LastUpdate                                  any    `json:"__last_update"`                                    // Last Modified on
	DisplayName                                 string `json:"display_name"`                                     // Display Name
	CreateUid                                   any    `json:"create_uid"`                                       // Created by 📦 Model Relation: res.users
	CreateDate                                  any    `json:"create_date"`                                      // Created on
	WriteUid                                    any    `json:"write_uid"`                                        // Last Updated by 📦 Model Relation: res.users
	WriteDate                                   any    `json:"write_date"`                                       // Last Updated on
	SocialTwitter                               string `json:"social_twitter"`                                   // Twitter Account
	SocialFacebook                              string `json:"social_facebook"`                                  // Facebook Account
	SocialGithub                                string `json:"social_github"`                                    // GitHub Account
	SocialLinkedin                              string `json:"social_linkedin"`                                  // LinkedIn Account
	SocialYoutube                               string `json:"social_youtube"`                                   // Youtube Account
	SocialInstagram                             string `json:"social_instagram"`                                 // Instagram Account
	NomenclatureID                              any    `json:"nomenclature_id"`                                  // Nomenclature 📦 Model Relation: barcode.nomenclature
	ResourceCalendarIDs                         any    `json:"resource_calendar_ids"`                            // Working Hours 📦 Model Relation: resource.calendar
	ResourceCalendarID                          any    `json:"resource_calendar_id"`                             // Default Working Hours 📦 Model Relation: resource.calendar
	CatchallEmail                               string `json:"catchall_email"`                                   // Catchall Email
	CatchallFormatted                           string `json:"catchall_formatted"`                               // Catchall
	EmailFormatted                              string `json:"email_formatted"`                                  // Formatted Email
	HrPresenceControlEmailAmount                int    `json:"hr_presence_control_email_amount"`                 // # emails to send
	HrPresenceControlIpList                     string `json:"hr_presence_control_ip_list"`                      // Valid IP addresses
	LunchMinimumThreshold                       any    `json:"lunch_minimum_threshold"`                          // Lunch Minimum Threshold
	AppraisalPlan                               bool   `json:"appraisal_plan"`                                   // Automatically Generate Appraisals
	AssessmentNoteIDs                           any    `json:"assessment_note_ids"`                              // Assessment Note 📦 Model Relation: hr.appraisal.note
	AppraisalEmployeeFeedbackTemplate           string `json:"appraisal_employee_feedback_template"`             // Appraisal Employee Feedback Template
	AppraisalManagerFeedbackTemplate            string `json:"appraisal_manager_feedback_template"`              // Appraisal Manager Feedback Template
	AppraisalConfirmMailTemplate                any    `json:"appraisal_confirm_mail_template"`                  // Appraisal Confirm Mail Template 📦 Model Relation: mail.template
	DurationAfterRecruitment                    int    `json:"duration_after_recruitment"`                       // Create an Appraisal after recruitment
	DurationFirstAppraisal                      int    `json:"duration_first_appraisal"`                         // Create a first Appraisal after
	DurationNextAppraisal                       int    `json:"duration_next_appraisal"`                          // Create a second Appraisal after
	HrAttendanceOvertime                        bool   `json:"hr_attendance_overtime"`                           // Count Extra Hours
	OvertimeStartDate                           any    `json:"overtime_start_date"`                              // Extra Hours Starting Date
	OvertimeCompanyThreshold                    int    `json:"overtime_company_threshold"`                       // Tolerance Time In Favor Of Company
	OvertimeEmployeeThreshold                   int    `json:"overtime_employee_threshold"`                      // Tolerance Time In Favor Of Employee
	PartnerGid                                  int    `json:"partner_gid"`                                      // Company database ID
	IapEnrichAutoDone                           bool   `json:"iap_enrich_auto_done"`                             // Enrich Done
	SnailmailColor                              bool   `json:"snailmail_color"`                                  // Color
	SnailmailCover                              bool   `json:"snailmail_cover"`                                  // Add a Cover Page
	SnailmailDuplex                             bool   `json:"snailmail_duplex"`                                 // Both sides
	SignTerms                                   string `json:"sign_terms"`                                       // Sign Default Terms and Conditions
	SignTermsType                               any    `json:"sign_terms_type"`                                  // Sign Terms & Conditions format
	SignTermsHtml                               string `json:"sign_terms_html"`                                  // Sign Default Terms and Conditions as a Web page
	BackgroundImage                             any    `json:"background_image"`                                 // Home Menu Background Image
	FiscalyearLastDay                           int    `json:"fiscalyear_last_day"`                              // Fiscalyear Last Day ⭐ Required
	FiscalyearLastMonth                         any    `json:"fiscalyear_last_month"`                            // Fiscalyear Last Month ⭐ Required
	PeriodLockDate                              any    `json:"period_lock_date"`                                 // Lock Date for Non-Advisers
	FiscalyearLockDate                          any    `json:"fiscalyear_lock_date"`                             // Lock Date
	TaxLockDate                                 any    `json:"tax_lock_date"`                                    // Tax Lock Date
	TransferAccountID                           any    `json:"transfer_account_id"`                              // Inter-Banks Transfer Account 📦 Model Relation: account.account
	ExpectsChartOfAccounts                      bool   `json:"expects_chart_of_accounts"`                        // Expects a Chart of Accounts
	ChartTemplateID                             any    `json:"chart_template_id"`                                // Chart Template 📦 Model Relation: account.chart.template
	BankAccountCodePrefix                       string `json:"bank_account_code_prefix"`                         // Prefix of the bank accounts
	CashAccountCodePrefix                       string `json:"cash_account_code_prefix"`                         // Prefix of the cash accounts
	DefaultCashDifferenceIncomeAccountID        any    `json:"default_cash_difference_income_account_id"`        // Cash Difference Income Account 📦 Model Relation: account.account
	DefaultCashDifferenceExpenseAccountID       any    `json:"default_cash_difference_expense_account_id"`       // Cash Difference Expense Account 📦 Model Relation: account.account
	AccountJournalSuspenseAccountID             any    `json:"account_journal_suspense_account_id"`              // Journal Suspense Account 📦 Model Relation: account.account
	AccountJournalPaymentDebitAccountID         any    `json:"account_journal_payment_debit_account_id"`         // Journal Outstanding Receipts Account 📦 Model Relation: account.account
	AccountJournalPaymentCreditAccountID        any    `json:"account_journal_payment_credit_account_id"`        // Journal Outstanding Payments Account 📦 Model Relation: account.account
	TransferAccountCodePrefix                   string `json:"transfer_account_code_prefix"`                     // Prefix of the transfer accounts
	AccountSaleTaxID                            any    `json:"account_sale_tax_id"`                              // Default Sale Tax 📦 Model Relation: account.tax
	AccountPurchaseTaxID                        any    `json:"account_purchase_tax_id"`                          // Default Purchase Tax 📦 Model Relation: account.tax
	TaxCalculationRoundingMethod                any    `json:"tax_calculation_rounding_method"`                  // Tax Calculation Rounding Method
	CurrencyExchangeJournalID                   any    `json:"currency_exchange_journal_id"`                     // Exchange Gain or Loss Journal 📦 Model Relation: account.journal
	IncomeCurrencyExchangeAccountID             any    `json:"income_currency_exchange_account_id"`              // Gain Exchange Rate Account 📦 Model Relation: account.account
	ExpenseCurrencyExchangeAccountID            any    `json:"expense_currency_exchange_account_id"`             // Loss Exchange Rate Account 📦 Model Relation: account.account
	AngloSaxonAccounting                        bool   `json:"anglo_saxon_accounting"`                           // Use anglo-saxon accounting
	PropertyStockAccountInputCategID            any    `json:"property_stock_account_input_categ_id"`            // Input Account for Stock Valuation 📦 Model Relation: account.account
	PropertyStockAccountOutputCategID           any    `json:"property_stock_account_output_categ_id"`           // Output Account for Stock Valuation 📦 Model Relation: account.account
	PropertyStockValuationAccountID             any    `json:"property_stock_valuation_account_id"`              // Account Template for Stock Valuation 📦 Model Relation: account.account
	BankJournalIDs                              any    `json:"bank_journal_ids"`                                 // Bank Journals 📦 Model Relation: account.journal
	IncotermID                                  any    `json:"incoterm_id"`                                      // Default incoterm 📦 Model Relation: account.incoterms
	QrCode                                      bool   `json:"qr_code"`                                          // Display QR-code on invoices
	InvoiceIsEmail                              bool   `json:"invoice_is_email"`                                 // Email by default
	InvoiceIsPrint                              bool   `json:"invoice_is_print"`                                 // Print by default
	AccountOpeningMoveID                        any    `json:"account_opening_move_id"`                          // Opening Journal Entry 📦 Model Relation: account.move
	AccountOpeningJournalID                     any    `json:"account_opening_journal_id"`                       // Opening Journal 📦 Model Relation: account.journal
	AccountOpeningDate                          any    `json:"account_opening_date"`                             // Opening Entry ⭐ Required
	AccountSetupBankDataState                   any    `json:"account_setup_bank_data_state"`                    // State of the onboarding bank data step
	AccountSetupFyDataState                     any    `json:"account_setup_fy_data_state"`                      // State of the onboarding fiscal year step
	AccountSetupCoaState                        any    `json:"account_setup_coa_state"`                          // State of the onboarding charts of account step
	AccountSetupTaxesState                      any    `json:"account_setup_taxes_state"`                        // State of the onboarding Taxes step
	AccountOnboardingInvoiceLayoutState         any    `json:"account_onboarding_invoice_layout_state"`          // State of the onboarding invoice layout step
	AccountOnboardingCreateInvoiceState         any    `json:"account_onboarding_create_invoice_state"`          // State of the onboarding create invoice step
	AccountOnboardingSaleTaxState               any    `json:"account_onboarding_sale_tax_state"`                // State of the onboarding sale tax step
	AccountInvoiceOnboardingState               any    `json:"account_invoice_onboarding_state"`                 // State of the account invoice onboarding panel
	AccountDashboardOnboardingState             any    `json:"account_dashboard_onboarding_state"`               // State of the account dashboard onboarding panel
	InvoiceTerms                                string `json:"invoice_terms"`                                    // Default Terms and Conditions
	TermsType                                   any    `json:"terms_type"`                                       // Terms & Conditions format
	InvoiceTermsHtml                            string `json:"invoice_terms_html"`                               // Default Terms and Conditions as a Web page
	AccountSetupBillState                       any    `json:"account_setup_bill_state"`                         // State of the onboarding bill step
	AccountDefaultPosReceivableAccountID        any    `json:"account_default_pos_receivable_account_id"`        // Default PoS Receivable Account 📦 Model Relation: account.account
	ExpenseAccrualAccountID                     any    `json:"expense_accrual_account_id"`                       // Expense Accrual Account 📦 Model Relation: account.account
	RevenueAccrualAccountID                     any    `json:"revenue_accrual_account_id"`                       // Revenue Accrual Account 📦 Model Relation: account.account
	AutomaticEntryDefaultJournalID              any    `json:"automatic_entry_default_journal_id"`               // Automatic Entry Default Journal 📦 Model Relation: account.journal
	CountryCode                                 string `json:"country_code"`                                     // Country Code
	AccountFiscalCountryID                      any    `json:"account_fiscal_country_id"`                        // Fiscal Country 📦 Model Relation: res.country
	AccountEnabledTaxCountryIDs                 any    `json:"account_enabled_tax_country_ids"`                  // l10n-used countries 📦 Model Relation: res.country
	TaxExigibility                              bool   `json:"tax_exigibility"`                                  // Use Cash Basis
	TaxCashBasisJournalID                       any    `json:"tax_cash_basis_journal_id"`                        // Cash Basis Journal 📦 Model Relation: account.journal
	AccountCashBasisBaseAccountID               any    `json:"account_cash_basis_base_account_id"`               // Base Tax Received Account 📦 Model Relation: account.account
	FiscalPositionIDs                           any    `json:"fiscal_position_ids"`                              // Fiscal Position 📦 Model Relation: account.fiscal.position
	MultiVatForeignCountryIDs                   any    `json:"multi_vat_foreign_country_ids"`                    // Foreign VAT countries 📦 Model Relation: res.country
	PlanningGenerationInterval                  int    `json:"planning_generation_interval"`                     // Rate Of Shift Generation ⭐ Required
	PlanningAllowSelfUnassign                   bool   `json:"planning_allow_self_unassign"`                     // Can Employee Un-Assign Themselves?
	PlanningSelfUnassignDaysBefore              int    `json:"planning_self_unassign_days_before"`               // Days before shift for unassignment
	InternalTransitLocationID                   any    `json:"internal_transit_location_id"`                     // Internal Transit Location 📦 Model Relation: stock.location
	StockMoveEmailValidation                    bool   `json:"stock_move_email_validation"`                      // Email Confirmation picking
	StockMailConfirmationTemplateID             any    `json:"stock_mail_confirmation_template_id"`              // Email Template confirmation picking 📦 Model Relation: mail.template
	AnnualInventoryMonth                        any    `json:"annual_inventory_month"`                           // Annual Inventory Month
	AnnualInventoryDay                          int    `json:"annual_inventory_day"`                             // Day of the month
	WebsiteID                                   any    `json:"website_id"`                                       // Website 📦 Model Relation: website
	InvoicingSwitchThreshold                    any    `json:"invoicing_switch_threshold"`                       // Invoicing Switch Threshold
	AccountCheckPrintingLayout                  any    `json:"account_check_printing_layout"`                    // Check Layout
	AccountCheckPrintingDateLabel               bool   `json:"account_check_printing_date_label"`                // Print Date Label
	AccountCheckPrintingMultiStub               bool   `json:"account_check_printing_multi_stub"`                // Multi-Pages Check Stub
	AccountCheckPrintingMarginTop               any    `json:"account_check_printing_margin_top"`                // Check Top Margin
	AccountCheckPrintingMarginLeft              any    `json:"account_check_printing_margin_left"`               // Check Left Margin
	AccountCheckPrintingMarginRight             any    `json:"account_check_printing_margin_right"`              // Right Margin
	ExtractShowOcrOptionSelection               any    `json:"extract_show_ocr_option_selection"`                // Send mode on invoices attachments
	ExtractSingleLinePerTax                     bool   `json:"extract_single_line_per_tax"`                      // OCR Single Invoice Line Per Tax
	VatCheckVies                                bool   `json:"vat_check_vies"`                                   // Verify VAT Numbers
	CurrencyIntervalUnit                        any    `json:"currency_interval_unit"`                           // Interval Unit
	CurrencyNextExecutionDate                   any    `json:"currency_next_execution_date"`                     // Next Execution Date
	CurrencyProvider                            any    `json:"currency_provider"`                                // Service Provider
	DocumentsHrSettings                         bool   `json:"documents_hr_settings"`                            // Documents Hr Settings
	DocumentsHrFolder                           any    `json:"documents_hr_folder"`                              // hr Workspace 📦 Model Relation: documents.folder
	DocumentsProductSettings                    bool   `json:"documents_product_settings"`                       // Documents Product Settings
	ProductFolder                               any    `json:"product_folder"`                                   // Product Workspace 📦 Model Relation: documents.folder
	ProductTags                                 any    `json:"product_tags"`                                     // Product Tags 📦 Model Relation: documents.tag
	DocumentsProjectSettings                    bool   `json:"documents_project_settings"`                       // Documents Project Settings
	ProjectFolder                               any    `json:"project_folder"`                                   // Project Workspace 📦 Model Relation: documents.folder
	ProjectTags                                 any    `json:"project_tags"`                                     // Project Tags 📦 Model Relation: documents.tag
	DocumentsSpreadsheetFolderID                any    `json:"documents_spreadsheet_folder_id"`                  // Documents Spreadsheet Folder 📦 Model Relation: documents.folder
	ProjectTimeModeID                           any    `json:"project_time_mode_id"`                             // Project Time Unit 📦 Model Relation: uom.uom
	TimesheetEncodeUomID                        any    `json:"timesheet_encode_uom_id"`                          // Timesheet Encoding Unit 📦 Model Relation: uom.uom
	InternalProjectID                           any    `json:"internal_project_id"`                              // Internal Project 📦 Model Relation: project.project
	ManufacturingLead                           any    `json:"manufacturing_lead"`                               // Manufacturing Lead Time ⭐ Required
	PaymentAcquirerOnboardingState              any    `json:"payment_acquirer_onboarding_state"`                // State of the onboarding payment acquirer step
	PaymentOnboardingPaymentMethod              any    `json:"payment_onboarding_payment_method"`                // Selected onboarding payment method
	PoLead                                      any    `json:"po_lead"`                                          // Purchase Lead Time ⭐ Required
	PoLock                                      any    `json:"po_lock"`                                          // Purchase Order Modification
	PoDoubleValidation                          any    `json:"po_double_validation"`                             // Levels of Approvals
	PoDoubleValidationAmount                    any    `json:"po_double_validation_amount"`                      // Double validation amount
	InvoiceIsSnailmail                          bool   `json:"invoice_is_snailmail"`                             // Send by Post
	StockMoveSmsValidation                      bool   `json:"stock_move_sms_validation"`                        // SMS Confirmation
	StockSmsConfirmationTemplateID              any    `json:"stock_sms_confirmation_template_id"`               // SMS Template 📦 Model Relation: sms.template
	HasReceivedWarningStockSms                  bool   `json:"has_received_warning_stock_sms"`                   // Has Received Warning Stock Sms
	TotalsBelowSections                         bool   `json:"totals_below_sections"`                            // Add totals below sections
	AccountTaxPeriodicity                       any    `json:"account_tax_periodicity"`                          // Delay units ⭐ Required
	AccountTaxPeriodicityReminderDay            int    `json:"account_tax_periodicity_reminder_day"`             // Start from ⭐ Required
	AccountTaxPeriodicityJournalID              any    `json:"account_tax_periodicity_journal_id"`               // Journal 📦 Model Relation: account.journal
	AccountRevaluationJournalID                 any    `json:"account_revaluation_journal_id"`                   // Account Revaluation Journal 📦 Model Relation: account.journal
	AccountRevaluationExpenseProvisionAccountID any    `json:"account_revaluation_expense_provision_account_id"` // Expense Provision Account 📦 Model Relation: account.account
	AccountRevaluationIncomeProvisionAccountID  any    `json:"account_revaluation_income_provision_account_id"`  // Income Provision Account 📦 Model Relation: account.account
	AccountTaxUnitIDs                           any    `json:"account_tax_unit_ids"`                             // Tax Units 📦 Model Relation: account.tax.unit
	AccountRepresentativeID                     any    `json:"account_representative_id"`                        // Accounting Firm 📦 Model Relation: res.partner
	AccountDisplayRepresentativeField           bool   `json:"account_display_representative_field"`             // Account Display Representative Field
	DocumentsHrContractsTags                    any    `json:"documents_hr_contracts_tags"`                      // Documents Hr Contracts Tags 📦 Model Relation: documents.tag
	DocumentsHrPayslipsTags                     any    `json:"documents_hr_payslips_tags"`                       // Documents Hr Payslips Tags 📦 Model Relation: documents.tag
	DocumentsPayrollFolderID                    any    `json:"documents_payroll_folder_id"`                      // Documents Payroll Folder 📦 Model Relation: documents.folder
	DocumentsRecruitmentSettings                bool   `json:"documents_recruitment_settings"`                   // Documents Recruitment Settings
	RecruitmentFolderID                         any    `json:"recruitment_folder_id"`                            // Recruitment Workspace 📦 Model Relation: documents.folder
	RecruitmentTagIDs                           any    `json:"recruitment_tag_ids"`                              // Recruitment Tag 📦 Model Relation: documents.tag
	DeferredTimeOffManager                      any    `json:"deferred_time_off_manager"`                        // Deferred Time Off Manager 📦 Model Relation: res.users
	L10NCaPst                                   string `json:"l10n_ca_pst"`                                      // PST
	LeaveTimesheetTaskID                        any    `json:"leave_timesheet_task_id"`                          // Time Off Task 📦 Model Relation: project.task
	DaysToPurchase                              any    `json:"days_to_purchase"`                                 // Days to Purchase
	PortalConfirmationSign                      bool   `json:"portal_confirmation_sign"`                         // Online Signature
	PortalConfirmationPay                       bool   `json:"portal_confirmation_pay"`                          // Online Payment
	QuotationValidityDays                       int    `json:"quotation_validity_days"`                          // Default Quotation Validity (Days)
	SaleQuotationOnboardingState                any    `json:"sale_quotation_onboarding_state"`                  // State of the sale onboarding panel
	SaleOnboardingOrderConfirmationState        any    `json:"sale_onboarding_order_confirmation_state"`         // State of the onboarding confirmation order step
	SaleOnboardingSampleQuotationState          any    `json:"sale_onboarding_sample_quotation_state"`           // State of the onboarding sample quotation step
	SaleOnboardingPaymentMethod                 any    `json:"sale_onboarding_payment_method"`                   // Sale onboarding selected payment method
	TimesheetMailEmployeeAllow                  bool   `json:"timesheet_mail_employee_allow"`                    // Employee Reminder
	TimesheetMailEmployeeDelay                  int    `json:"timesheet_mail_employee_delay"`                    // Employee Reminder Days
	TimesheetMailEmployeeInterval               any    `json:"timesheet_mail_employee_interval"`                 // Employee Frequency ⭐ Required
	TimesheetMailEmployeeNextdate               any    `json:"timesheet_mail_employee_nextdate"`                 // Next scheduled date for employee reminder
	TimesheetMailManagerAllow                   bool   `json:"timesheet_mail_manager_allow"`                     // Manager Reminder
	TimesheetMailManagerDelay                   int    `json:"timesheet_mail_manager_delay"`                     // Manager Reminder Days
	TimesheetMailManagerInterval                any    `json:"timesheet_mail_manager_interval"`                  // Manager Reminder Frequency ⭐ Required
	TimesheetMailManagerNextdate                any    `json:"timesheet_mail_manager_nextdate"`                  // Next scheduled date for manager reminder
	GainAccountID                               any    `json:"gain_account_id"`                                  // Gain Account 📦 Model Relation: account.account
	LossAccountID                               any    `json:"loss_account_id"`                                  // Loss Account 📦 Model Relation: account.account
	ConsolidationColor                          int    `json:"consolidation_color"`                              // Accounts color
	AccountConsolidationCurrencyIsDifferent     bool   `json:"account_consolidation_currency_is_different"`      // Account Consolidation Currency Is Different
	ConsolidationDashboardOnboardingState       any    `json:"consolidation_dashboard_onboarding_state"`         // State Of The Account Consolidation Dashboard Onboarding Panel
	ConsolidationSetupConsolidationState        any    `json:"consolidation_setup_consolidation_state"`          // State Of The Onboarding Consolidation Step
	ConsolidationSetupCcoaState                 any    `json:"consolidation_setup_ccoa_state"`                   // State Of The Onboarding Consolidated Chart Of Account Step
	ConsolidationCreatePeriodState              any    `json:"consolidation_create_period_state"`                // State Of The Onboarding Create Period Step
	DocumentsAccountSettings                    bool   `json:"documents_account_settings"`                       // Documents Account Settings
	AccountFolder                               any    `json:"account_folder"`                                   // Accounting Workspace 📦 Model Relation: documents.folder
	ExpenseExtractShowOcrOptionSelection        any    `json:"expense_extract_show_ocr_option_selection"`        // Send mode on expense attachments
	OnssCompanyID                               string `json:"onss_company_id"`                                  // ONSS Company ID
	OnssRegistrationNumber                      string `json:"onss_registration_number"`                         // ONSS Registration Number
	DmfaEmployerClass                           string `json:"dmfa_employer_class"`                              // DMFA Employer Class
	DmfaLocationUnitIDs                         any    `json:"dmfa_location_unit_ids"`                           // Work address DMFA codes 📦 Model Relation: l10n_be.dmfa.location.unit
	L10NBeCompanyNumber                         string `json:"l10n_be_company_number"`                           // Company Number
	L10NBeRevenueCode                           string `json:"l10n_be_revenue_code"`                             // Revenue Code
	L10NBeFfeEmployerType                       any    `json:"l10n_be_ffe_employer_type"`                        // L10N Be Ffe Employer Type
	ManufacturingPeriod                         any    `json:"manufacturing_period"`                             // Manufacturing Period ⭐ Required
	ManufacturingPeriodToDisplay                int    `json:"manufacturing_period_to_display"`                  // Number of columns for the        given period to display in Master Production Schedule
	MrpMpsShowStartingInventory                 bool   `json:"mrp_mps_show_starting_inventory"`                  // Display Starting Inventory
	MrpMpsShowDemandForecast                    bool   `json:"mrp_mps_show_demand_forecast"`                     // Display Demand Forecast
	MrpMpsShowActualDemand                      bool   `json:"mrp_mps_show_actual_demand"`                       // Display Actual Demand
	MrpMpsShowIndirectDemand                    bool   `json:"mrp_mps_show_indirect_demand"`                     // Display Indirect Demand
	MrpMpsShowToReplenish                       bool   `json:"mrp_mps_show_to_replenish"`                        // Display To Replenish
	MrpMpsShowActualReplenishment               bool   `json:"mrp_mps_show_actual_replenishment"`                // Display Actual Replenishment
	MrpMpsShowSafetyStock                       bool   `json:"mrp_mps_show_safety_stock"`                        // Display Safety Stock
	MrpMpsShowAvailableToPromise                bool   `json:"mrp_mps_show_available_to_promise"`                // Display Available to Promise
	MrpMpsShowActualDemandYearMinus1            bool   `json:"mrp_mps_show_actual_demand_year_minus_1"`          // Display Actual Demand Last Year
	MrpMpsShowActualDemandYearMinus2            bool   `json:"mrp_mps_show_actual_demand_year_minus_2"`          // Display Actual Demand Before Year
	SaleOrderTemplateID                         any    `json:"sale_order_template_id"`                           // Default Sale Template 📦 Model Relation: sale.order.template
	SecurityLead                                any    `json:"security_lead"`                                    // Sales Safety Days ⭐ Required
	ShDisableFollowerConfirmSale                bool   `json:"sh_disable_follower_confirm_sale"`                 // Disable to add followers by Confirm Quotation
	ShDisableFollowerValidateInvoice            bool   `json:"sh_disable_follower_validate_invoice"`             // Disable to add Followers by Validate Invoice/Bill
	ShDisableFollowerEmail                      bool   `json:"sh_disable_follower_email"`                        // Disable to add Followers by Send by Email
	ShDisableFollowerCreatePicking              bool   `json:"sh_disable_follower_create_picking"`               // Disable to add Followers by create/update picking
	ShDisableFollowerSalesperson                bool   `json:"sh_disable_follower_salesperson"`                  // Disable to add Salesperson as followers
	ShDisableFollowerPrPurchase                 bool   `json:"sh_disable_follower_pr_purchase"`                  // Disable to add Purchase Representative as followers
	ShDisableFollowerResponsiblePicking         bool   `json:"sh_disable_follower_responsible_picking"`          // Disable to add Responsible as followers
	ShDisableFollowerSalespersonAccount         bool   `json:"sh_disable_follower_salesperson_account"`          // Disable to add Salesperson as followers
	RmaTemplate                                 bool   `json:"rma_template"`                                     // RMA Templates
	RmaTemplateID                               any    `json:"rma_template_id"`                                  // Rma Template 📦 Model Relation: mail.template
}

func (m *Model) ResCompany() {
	model := "res.company"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(ResCompany_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	fmt.Println(model, "records found:", len(records))

	for _, r := range records {
		var record ResCompany_150
		FillStruct(r, &record)

		ur := map[string]any{
			"name":     record.Name,
			"street":   record.Street,
			"street2":  record.Street2,
			"zip":      record.Zip,
			"city":     record.City,
			"phone":    record.Phone,
			"website":  record.Website,
			"vat":      record.Vat,
			"logo":     record.Logo,
			"logo_web": record.LogoWeb,
			"email":    record.Email,
		}

		countryID := m.getCountryFromSource(record.CountryID)
		if countryID != -1 {
			ur["country_id"] = countryID
		}
		stateID := m.getCountryStateFromSource(record.CountryID, record.StateID)
		if stateID != -1 {
			ur["state_id"] = stateID
		}
		currencyID := m.getCurrencyFromSource(record.CurrencyID)
		if currencyID != -1 {
			ur["currency_id"] = currencyID
		}
		paperFormatID := m.getPaperFormatFromSource(record.PaperformatID)
		if paperFormatID != -1 {
			ur["paperformat_id"] = paperFormatID
		}

		if record.ID == 1 {
			// m.Log.Info(model, "rid", 1, "ur", ur)
			// fmt.Println("base record", prettyprint(ur))
			m.writeRecord(model, ur, 1)
		} else {
			_, err := m.Dest.GetID(model, []any{
				[]any{"name", "=", record.Name},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				return
			}
			// fmt.Println(prettyprint(ur))
			// m.Log.Info(model, "rid", rid, "ur", ur)
			// m.writeRecord(model, ur, rid)
		}

	}
}
