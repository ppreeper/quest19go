package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type AccountAccount_150 struct {
	MessageIsFollower                    bool   `json:"message_is_follower"`                      // Is Follower
	MessageFollowerIDs                   any    `json:"message_follower_ids"`                     // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs                    any    `json:"message_partner_ids"`                      // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs                           any    `json:"message_ids"`                              // Messages 📦 Model Relation: mail.message
	HasMessage                           bool   `json:"has_message"`                              // Has Message
	MessageUnread                        bool   `json:"message_unread"`                           // Unread Messages
	MessageUnreadCounter                 int    `json:"message_unread_counter"`                   // Unread Messages Counter
	MessageNeedaction                    bool   `json:"message_needaction"`                       // Action Needed
	MessageNeedactionCounter             int    `json:"message_needaction_counter"`               // Number of Actions
	MessageHasError                      bool   `json:"message_has_error"`                        // Message Delivery error
	MessageHasErrorCounter               int    `json:"message_has_error_counter"`                // Number of errors
	MessageAttachmentCount               int    `json:"message_attachment_count"`                 // Attachment Count
	MessageMainAttachmentID              any    `json:"message_main_attachment_id"`               // Main Attachment 📦 Model Relation: ir.attachment
	WebsiteMessageIDs                    any    `json:"website_message_ids"`                      // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError                   bool   `json:"message_has_sms_error"`                    // SMS Delivery error
	Name                                 string `json:"name"`                                     // Account Name ⭐ Required
	CurrencyID                           any    `json:"currency_id"`                              // Account Currency 📦 Model Relation: res.currency
	Code                                 string `json:"code"`                                     // Code ⭐ Required
	Deprecated                           bool   `json:"deprecated"`                               // Deprecated
	Used                                 bool   `json:"used"`                                     // Used
	UserTypeID                           any    `json:"user_type_id"`                             // Type 📦 Model Relation: account.account.type ⭐ Required
	InternalType                         any    `json:"internal_type"`                            // Internal Type
	InternalGroup                        any    `json:"internal_group"`                           // Internal Group
	Reconcile                            bool   `json:"reconcile"`                                // Allow Reconciliation
	TaxIDs                               any    `json:"tax_ids"`                                  // Default Taxes 📦 Model Relation: account.tax
	Note                                 string `json:"note"`                                     // Internal Notes
	CompanyID                            any    `json:"company_id"`                               // Company 📦 Model Relation: res.company ⭐ Required
	TagIDs                               any    `json:"tag_ids"`                                  // Tags 📦 Model Relation: account.account.tag
	GroupID                              any    `json:"group_id"`                                 // Group 📦 Model Relation: account.group
	RootID                               any    `json:"root_id"`                                  // Root 📦 Model Relation: account.root
	AllowedJournalIDs                    any    `json:"allowed_journal_ids"`                      // Allowed Journals 📦 Model Relation: account.journal
	OpeningDebit                         any    `json:"opening_debit"`                            // Opening Debit
	OpeningCredit                        any    `json:"opening_credit"`                           // Opening Credit
	OpeningBalance                       any    `json:"opening_balance"`                          // Opening Balance
	IsOffBalance                         bool   `json:"is_off_balance"`                           // Is Off Balance
	CurrentBalance                       any    `json:"current_balance"`                          // Current Balance
	RelatedTaxesAmount                   int    `json:"related_taxes_amount"`                     // Related Taxes Amount
	ID                                   int    `json:"id"`                                       // ID
	LastUpdate                           any    `json:"__last_update"`                            // Last Modified on
	DisplayName                          string `json:"display_name"`                             // Display Name
	CreateUid                            any    `json:"create_uid"`                               // Created by 📦 Model Relation: res.users
	CreateDate                           any    `json:"create_date"`                              // Created on
	WriteUid                             any    `json:"write_uid"`                                // Last Updated by 📦 Model Relation: res.users
	WriteDate                            any    `json:"write_date"`                               // Last Updated on
	ExcludeProvisionCurrencyIDs          any    `json:"exclude_provision_currency_ids"`           // Exclude Provision Currency 📦 Model Relation: res.currency
	ExcludeFromAgedReports               bool   `json:"exclude_from_aged_reports"`                // Exclude From Aged Reports
	AssetModel                           any    `json:"asset_model"`                              // Asset Model 📦 Model Relation: account.asset
	CreateAsset                          any    `json:"create_asset"`                             // Create Asset ⭐ Required
	CanCreateAsset                       bool   `json:"can_create_asset"`                         // Can Create Asset
	FormViewRef                          string `json:"form_view_ref"`                            // Form View Ref
	AssetType                            any    `json:"asset_type"`                               // Asset Type
	MultipleAssetsPerLine                bool   `json:"multiple_assets_per_line"`                 // Multiple Assets per Line
	ConsolidationAccountIDs              any    `json:"consolidation_account_ids"`                // Consolidation Account 📦 Model Relation: consolidation.account
	ConsolidationAccountChartFilteredIDs any    `json:"consolidation_account_chart_filtered_ids"` // Consolidation Account Chart Filtered 📦 Model Relation: consolidation.account
	ConsolidationColor                   int    `json:"consolidation_color"`                      // Color
	XStudioMany2OneFieldP9Cwx            any    `json:"x_studio_many2one_field_p9cWX"`            // Account Group 📦 Model Relation: account.group ⚠️ Manual: True
}

type AccountAccount_190 struct {
	ID                              any `json:"id"`                                  // ID
	DisplayName                     any `json:"display_name"`                        // Display Name
	ActivityIDs                     any `json:"activity_ids"`                        // Activities 📦 Model Relation: mail.activity
	ActivityState                   any `json:"activity_state"`                      // Activity State
	ActivityUserID                  any `json:"activity_user_id"`                    // Responsible User 📦 Model Relation: res.users
	ActivityTypeID                  any `json:"activity_type_id"`                    // Next Activity Type 📦 Model Relation: mail.activity.type
	ActivityTypeIcon                any `json:"activity_type_icon"`                  // Activity Type Icon
	ActivityDateDeadline            any `json:"activity_date_deadline"`              // Next Activity Deadline
	MyActivityDateDeadline          any `json:"my_activity_date_deadline"`           // My Activity Deadline
	ActivitySummary                 any `json:"activity_summary"`                    // Next Activity Summary
	ActivityExceptionDecoration     any `json:"activity_exception_decoration"`       // Activity Exception Decoration
	ActivityExceptionIcon           any `json:"activity_exception_icon"`             // Icon
	ActivityCalendarEventID         any `json:"activity_calendar_event_id"`          // Next Activity Calendar Event 📦 Model Relation: calendar.event
	MessageIsFollower               any `json:"message_is_follower"`                 // Is Follower
	MessageFollowerIDs              any `json:"message_follower_ids"`                // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs               any `json:"message_partner_ids"`                 // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs                      any `json:"message_ids"`                         // Messages 📦 Model Relation: mail.message
	HasMessage                      any `json:"has_message"`                         // Has Message
	MessageNeedaction               any `json:"message_needaction"`                  // Action Needed
	MessageNeedactionCounter        any `json:"message_needaction_counter"`          // Number of Actions
	MessageHasError                 any `json:"message_has_error"`                   // Message Delivery error
	MessageHasErrorCounter          any `json:"message_has_error_counter"`           // Number of errors
	MessageAttachmentCount          any `json:"message_attachment_count"`            // Attachment Count
	RatingIDs                       any `json:"rating_ids"`                          // Ratings 📦 Model Relation: rating.rating
	WebsiteMessageIDs               any `json:"website_message_ids"`                 // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError              any `json:"message_has_sms_error"`               // SMS Delivery error
	Name                            any `json:"name"`                                // Account Name ⭐ Required
	Description                     any `json:"description"`                         // Description
	CurrencyID                      any `json:"currency_id"`                         // Account Currency 📦 Model Relation: res.currency
	CompanyCurrencyID               any `json:"company_currency_id"`                 // Company Currency 📦 Model Relation: res.currency
	CompanyFiscalCountryCode        any `json:"company_fiscal_country_code"`         // Company Fiscal Country Code
	Code                            any `json:"code"`                                // Code
	CodeStore                       any `json:"code_store"`                          // Code Store
	PlaceholderCode                 any `json:"placeholder_code"`                    // Display code
	Active                          any `json:"active"`                              // Active
	Used                            any `json:"used"`                                // Used
	AccountType                     any `json:"account_type"`                        // Type ⭐ Required
	IncludeInitialBalance           any `json:"include_initial_balance"`             // Bring Accounts Balance Forward
	InternalGroup                   any `json:"internal_group"`                      // Internal Group
	Reconcile                       any `json:"reconcile"`                           // Allow Reconciliation
	TaxIDs                          any `json:"tax_ids"`                             // Default Taxes 📦 Model Relation: account.tax
	Note                            any `json:"note"`                                // Internal Notes
	CompanyIDs                      any `json:"company_ids"`                         // Companies 📦 Model Relation: res.company ⭐ Required
	CodeMappingIDs                  any `json:"code_mapping_ids"`                    // Code Mapping 📦 Model Relation: account.code.mapping
	TagIDs                          any `json:"tag_ids"`                             // Tags 📦 Model Relation: account.account.tag
	GroupID                         any `json:"group_id"`                            // Group 📦 Model Relation: account.group
	RootID                          any `json:"root_id"`                             // Root 📦 Model Relation: account.root
	OpeningDebit                    any `json:"opening_debit"`                       // Opening Debit
	OpeningCredit                   any `json:"opening_credit"`                      // Opening Credit
	OpeningBalance                  any `json:"opening_balance"`                     // Opening Balance
	CurrentBalance                  any `json:"current_balance"`                     // Current Balance
	RelatedTaxesAmount              any `json:"related_taxes_amount"`                // Related Taxes Amount
	NonTrade                        any `json:"non_trade"`                           // Non Trade
	DisplayMappingTab               any `json:"display_mapping_tab"`                 // Display Mapping Tab
	CreateUid                       any `json:"create_uid"`                          // Created by 📦 Model Relation: res.users
	CreateDate                      any `json:"create_date"`                         // Created on
	WriteUid                        any `json:"write_uid"`                           // Last Updated by 📦 Model Relation: res.users
	WriteDate                       any `json:"write_date"`                          // Last Updated on
	AccountStockVariationID         any `json:"account_stock_variation_id"`          // Variation Account 📦 Model Relation: account.account
	AccountStockExpenseID           any `json:"account_stock_expense_id"`            // Expense Account 📦 Model Relation: account.account
	ExcludeProvisionCurrencyIDs     any `json:"exclude_provision_currency_ids"`      // Exclude Provision Currency 📦 Model Relation: res.currency
	BudgetItemIDs                   any `json:"budget_item_ids"`                     // Budget Item 📦 Model Relation: account.report.budget.item
	AuditDebit                      any `json:"audit_debit"`                         // Debit
	AuditCredit                     any `json:"audit_credit"`                        // Credit
	AuditBalance                    any `json:"audit_balance"`                       // Balance
	AuditBalanceShowWarning         any `json:"audit_balance_show_warning"`          // Audit Balance Show Warning
	AuditPreviousBalance            any `json:"audit_previous_balance"`              // Balance N-1
	AuditPreviousBalanceShowWarning any `json:"audit_previous_balance_show_warning"` // Audit Previous Balance Show Warning
	AuditVarN1                      any `json:"audit_var_n_1"`                       // Var N-1
	AuditVarPercentage              any `json:"audit_var_percentage"`                // Var %
	AuditStatus                     any `json:"audit_status"`                        // Status
	AccountStatus                   any `json:"account_status"`                      // Account Status 📦 Model Relation: account.audit.account.status
	LastMessage                     any `json:"last_message"`                        // Last Message
	FiscalCategoryID                any `json:"fiscal_category_id"`                  // Fiscal Category 📦 Model Relation: account.fiscal.category
	RateIDs                         any `json:"rate_ids"`                            // Rate 📦 Model Relation: account.account.fiscal.rate
	AssetModelIDs                   any `json:"asset_model_ids"`                     // Asset Model 📦 Model Relation: account.asset
	CreateAsset                     any `json:"create_asset"`                        // Create Asset ⭐ Required
	CanCreateAsset                  any `json:"can_create_asset"`                    // Can Create Asset
	FormViewRef                     any `json:"form_view_ref"`                       // Form View Ref
	MultipleAssetsPerLine           any `json:"multiple_assets_per_line"`            // Multiple Assets per Line
}

type AccountType struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	DisplayName           string `json:"display_name"`
	IncludeInitialBalance bool   `json:"include_initial_balance"`
	InternalGroup         string `json:"internal_group"`
	Note                  string `json:"note"`
	Type                  string `json:"type"`
}

func (m *Model) AccountAccount() {
	model := "account.account"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(AccountAccount_150{})

	accounts, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	banner.Println(model, "accounts")
	bar := progressbar.Default(-1, "accounts")
	for _, account := range accounts {
		var root AccountAccount_150
		FillStruct(account, &root)
		// fmt.Println(prettyprint(root))

		_, user_type_name := ParseMany2One(root.UserTypeID)
		// fmt.Println("user_type_id", user_type_id, "user_type_name", user_type_name, "mapped", quest_account_type[user_type_name])
		account_type := quest_account_type[user_type_name]
		// fmt.Println("account_type", account_type)

		lookup_code := questmap[root.Code]
		if lookup_code == "" {
			lookup_code = root.Code
		}

		rid := -1
		rida, _ := m.Dest.GetID(model, []any{
			[]any{"code", "=", lookup_code},
		})
		ridb, _ := m.Dest.GetID(model, []any{
			[]any{"code", "=", root.Code},
		})
		if rida != -1 {
			rid = rida
		}
		if rida == -1 && ridb != -1 {
			rid = ridb
		}
		// fmt.Println("rida", rida, "ridb", ridb, "rid", rid)

		ur := map[string]any{
			"name":      root.Name,
			"code":      root.Code,
			"reconcile": root.Reconcile,
		}
		if account_type != "" {
			ur["account_type"] = account_type
		}

		// m.Log.Info(model, "rid", rid, "ur", ur)
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

// TODO: skip tax accounts for ar and ap

var questmap = map[string]string{
	// 111101	Liquidity Transfer
	// 111211	Cash
	// 111311	Bank
	// 111312	Bank Suspense Account
	"1": "111312",
	// 112110	Trade Accounts Receivable
	"112100": "112110",
	// 113110	Equity Investment in Canadian Corporations, valued at cost of shares
	// 118100	GST receivable
	"113100": "118100",
	// 118200	PST/QST receivable
	"113110": "118200",
	// 118300	HST receivable - 13%
	// 118400	HST receivable - 14%
	// 118500	HST receivable - 15%
	// 212100	Non-Mortgage Loans, Bank Loans - Overdrafts and Loans from Chartered Banks and Branches of Foreign Banks Operating in Canada (excluding Lien Notes Payable)
	// 221110	Accounts Payable - Retail and Wholesale Trade Accounts
	"211100": "221110",
	// 231000	GST to pay
	"212100": "231000",
	// 232000	PST/QST to pay
	"212110": "232000",
	// 233000	HST to pay - 13%
	// 234000	HST to pay - 14%
	// 235000	HST to pay - 15%
	// 411100	Sales of Goods
	"43": "411100",
	// 423100	Cash Discount Gain
	// 423200	Discounts Taken
	// 511210	Purchases
	"53010": "511210",
	// 522100	Cash Discount Loss
	// 522200	Discounts Given
	// 999001	Cash Difference Gain
	// 999002	Cash Difference Loss

}

var quest_account_type = map[string]string{
	"Bank and Cash":           "asset_cash",
	"Credit Card":             "liability_credit_card",
	"Receivable":              "asset_receivable",
	"Payable":                 "liability_payable",
	"Current Assets":          "asset_current",
	"Non-current Assets":      "asset_non_current",
	"Prepayments":             "asset_prepayments",
	"Fixed Assets":            "asset_fixed",
	"Current Liabilities":     "liability_current",
	"Non-current Liabilities": "liability_non_current",
	"Equity":                  "equity",
	"Current Year Earnings":   "equity_unaffected",
	"Other Income":            "income_other",
	"Income":                  "income",
	"Depreciation":            "expense_depreciation",
	"Expenses":                "expense",
	"Cost of Revenue":         "expense_direct_cost",
	"Off-Balance Sheet":       "off_balance",
}

func (m *Model) AccountAccountCleanup() {
	model := "account.account"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(AccountAccount_190{})

	accounts, err := m.Dest.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"code", "not in", odoo19excludeAccounts},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	accounts_ids := ExtractIDs(accounts)

	bar := progressbar.Default(-1, "accounts cleanup")
	for _, account := range accounts_ids {
		// fmt.Println("unlink account", account)
		_, err := m.Dest.Unlink(model, []int{account})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
		}
		// fmt.Println("unlink", account, result)
		bar.Add(1)
	}
	bar.Finish()
}

var odoo19excludeAccounts = []string{
	"111101", //	Liquidity Transfer
	"111211", //	Cash
	"111311", //	Bank
	"111312", //	Bank Suspense Account
	"112110", //	Trade Accounts Receivable
	"118100", //	GST receivable
	"118200", //	PST/QST receivable
	"118300", //	HST receivable - 13%
	"118400", //	HST receivable - 14%
	"118500", //	HST receivable - 15%
	"121120", //	Goods for Sale Inventory
	"221110", //	Accounts Payable - Retail and Wholesale Trade Accounts
	"231000", //	GST to pay
	"232000", //	PST/QST to pay
	"233000", //	HST to pay - 13%
	"234000", //	HST to pay - 14%
	"235000", //	HST to pay - 15%
	"298000", //	All Other Current Liabilities
	"411100", //	Sales of Goods
	"423100", //	Cash Discount Gain
	"423200", //	Discounts Taken
	"511120", //	Goods/Work in Process
	"511320", //	Goods/Work in Process
	"511210", //	Purchases
	"522100", //	Cash Discount Loss
	"522200", //	Discounts Given
	"999001", //	Cash Difference Gain
	"999002", //	Cash Difference Loss
}

// map[account_type:asset_current code:113110 name:HST Receivable Tax (Paid) reconcile:false]" row=97 res=false err="write failed: request failed with status 422: 422 Unprocessable Entity Account codes must be unique. You can't create accounts with these duplicate codes: 113110"
// map[account_type:liability_current code:212100 name:GST Payable Tax (Collected) reconcile:false]" row=253 res=false err="write failed: request failed with status 422: 422 Unprocessable Entity Account codes must be unique. You can't create accounts with these duplicate codes: 212100"

func (m *Model) AccountAccountRename() {
	model := "account.account"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())
	type acct struct {
		OldCode string
		Name    string
		Code    string
	}
	questList := []acct{
		{"111312", "Bank Suspense Account", "1"},
		{"112110", "Trade Accounts Receivable", "112100"},
		{"221110", "Accounts Payable CAD", "211100"},
		{"411100", "Other Sales Income", "43"},
		{"511210", "Purchases COGS", "53010"},
	}

	for _, acct := range questList {
		rid, err := m.Dest.GetID(model, []any{
			[]any{"code", "=", acct.OldCode},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if rid == -1 {
			continue
		}
		ur := make(map[string]any)
		ur["name"] = acct.Name
		ur["code"] = acct.Code
		m.writeRecord(model, ur, rid)
	}
}

func (m *Model) AccountAccountNewAccounts() {
	model := "account.account"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())
	type acct struct {
		Name string
		Code string
		Type string
	}

	questList := []acct{
		{"Open AR Clearing", "800001", "income_other"},
		{"Open AP Clearing", "800002", "expense_other"},
	}

	for _, new_acct := range questList {
		rid, err := m.Dest.GetID(model, []any{
			[]any{"code", "=", new_acct.Code},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		ur := make(map[string]any)
		ur["name"] = new_acct.Name
		ur["code"] = new_acct.Code
		ur["account_type"] = new_acct.Type
		m.writeRecord(model, ur, rid)
	}
}

func (m *Model) AccountAccountLookup(old_id int, old_name string) int {
	// fmt.Println("old_id", old_id, "old_name", old_name)

	old_accounts, _ := m.Source.SearchRead("account.account", 0, 1, []string{"name", "code"}, []any{
		[]any{"id", "=", old_id},
	})
	if len(old_accounts) == 0 {
		m.Log.Error("account.account", "func", trace(), "err", fmt.Sprintf("account.account with id %d not found", old_id))
		return -1
	}
	old_account := old_accounts[0]
	old_account_name := old_account["name"].(string)
	old_account_code := old_account["code"].(string)
	// fmt.Println("old_account_name", old_account_name, "old_account_code", old_account_code)

	new_account_id, err := m.Dest.GetID("account.account", []any{
		[]any{"name", "=", old_account_name},
		[]any{"code", "=", old_account_code},
	})
	if err != nil {
		m.Log.Error("account.account", "func", trace(), "err", err)
		return -1
	}
	// fmt.Println("new_account_id", new_account_id)
	return new_account_id
}
