package models

import (
	"strings"

	"github.com/schollz/progressbar/v3"
)

// quest15data res.partner model
type ResPartner_150 struct {
	// Avatar1024                    any     `json:"avatar_1024"`                       // Avatar 1024
	// Avatar128                     any     `json:"avatar_128"`                        // Avatar 128
	// Avatar1920                    any     `json:"avatar_1920"`                       // Avatar
	// Avatar256                     any     `json:"avatar_256"`                        // Avatar 256
	// Avatar512                     any     `json:"avatar_512"`                        // Avatar 512
	// BankIDs     any     `json:"bank_ids"`     // Banks 📦 relation: one2many res.partner.bank
	Barcode string `json:"barcode"` // Barcode
	// CategoryID  any     `json:"category_id"`  // Tags 📦 relation: many2many res.partner.category
	City        string  `json:"city"`         // City
	Color       int     `json:"color"`        // Color Index
	Comment     string  `json:"comment"`      // Notes
	CompanyType any     `json:"company_type"` // Company Type
	CountryCode string  `json:"country_code"` // Country Code
	CountryID   any     `json:"country_id"`   // Country 📦 relation: many2one res.country
	Credit      float64 `json:"credit"`       // Total Receivable
	CreditLimit float64 `json:"credit_limit"` // Credit Limit
	CurrencyID  any     `json:"currency_id"`  // Currency 📦 relation: many2one res.currency
	Debit       float64 `json:"debit"`        // Total Payable
	DebitLimit  float64 `json:"debit_limit"`  // Payable Limit
	DisplayName string  `json:"display_name"` // Display Name
	Email       string  `json:"email"`        // Email
	Employee    bool    `json:"employee"`     // Employee
	ID          int     `json:"id"`           // ID
	// Image1024                     any     `json:"image_1024"`                        // Image 1024
	// Image128                      any     `json:"image_128"`                         // Image 128
	Image1920 any `json:"image_1920"` // Image
	// Image256                      any     `json:"image_256"`                         // Image 256
	// Image512                      any     `json:"image_512"`                         // Image 512
	// ImageMedium                   any     `json:"image_medium"`                      // Medium-sized image
	IndustryID           any    `json:"industry_id"`            // Industry 📦 relation: many2one res.partner.industry
	InvoiceWarnMsg       string `json:"invoice_warn_msg"`       // Message for Invoice
	IsCompany            bool   `json:"is_company"`             // Is a Company
	IsPublished          bool   `json:"is_published"`           // Is Published
	L10NCaPst            string `json:"l10n_ca_pst"`            // PST
	Lang                 any    `json:"lang"`                   // Language
	Mobile               string `json:"mobile"`                 // Mobile
	Name                 string `json:"name"`                   // Name
	OutsideSalespersonID any    `json:"outside_salesperson_id"` // Outside Salesperson 📦 relation: many2one res.users
	ParentID             any    `json:"parent_id"`              // Related Company 📦 relation: many2one res.partner
	// ParentName                    string `json:"parent_name"`                       // Parent name
	Phone                         string `json:"phone"`                             // Phone
	PickingWarnMsg                string `json:"picking_warn_msg"`                  // Message for Stock Picking
	PropertyAccountPayableID      any    `json:"property_account_payable_id"`       // Account Payable 📦 relation: many2one account.account ⭐ required
	PropertyAccountPositionID     any    `json:"property_account_position_id"`      // Fiscal Position 📦 relation: many2one account.fiscal.position
	PropertyAccountReceivableID   any    `json:"property_account_receivable_id"`    // Account Receivable 📦 relation: many2one account.account ⭐ required
	PropertyDeliveryCarrierID     any    `json:"property_delivery_carrier_id"`      // Delivery Method 📦 relation: many2one delivery.carrier
	PropertyPaymentMethodID       any    `json:"property_payment_method_id"`        // Payment Method 📦 relation: many2one account.payment.method
	PropertyPaymentTermID         any    `json:"property_payment_term_id"`          // Customer Payment Terms 📦 relation: many2one account.payment.term
	PropertyProductPricelist      any    `json:"property_product_pricelist"`        // Pricelist 📦 relation: many2one product.pricelist
	PropertyPurchaseCurrencyID    any    `json:"property_purchase_currency_id"`     // Supplier Currency 📦 relation: many2one res.currency
	PropertyStockCustomer         any    `json:"property_stock_customer"`           // Customer Location 📦 relation: many2one stock.location
	PropertyStockSupplier         any    `json:"property_stock_supplier"`           // Vendor Location 📦 relation: many2one stock.location
	PropertySupplierPaymentTermID any    `json:"property_supplier_payment_term_id"` // Vendor Payment Terms 📦 relation: many2one account.payment.term
	PurchaseWarnMsg               string `json:"purchase_warn_msg"`                 // Message for Purchase Order
	Ref                           string `json:"ref"`                               // Reference
	SaleWarnMsg                   string `json:"sale_warn_msg"`                     // Message for Sales Order
	StateID                       any    `json:"state_id"`                          // State 📦 relation: many2one res.country.state
	Street                        string `json:"street"`                            // Street
	Street2                       string `json:"street2"`                           // Street2
	// SupplierRank                  int    `json:"supplier_rank"`                     // Supplier Rank
	TeamID any `json:"team_id"` // Sales Team 📦 relation: many2one crm.team
	Title  any `json:"title"`   // Title 📦 relation: many2one res.partner.title
	// Trust                         any    `json:"trust"`                             // Degree of trust you have in this debtor
	Type     any    `json:"type"`      // Address Type
	TZ       any    `json:"tz"`        // Timezone
	TZOffset string `json:"tz_offset"` // Timezone offset
	// UnreconciledAmlIDs            any    `json:"unreconciled_aml_ids"`              // Unreconciled Aml 📦 relation: one2many account.move.line
	UserID              any    `json:"user_id"`               // Salesperson 📦 relation: many2one res.users
	UserIDs             any    `json:"user_ids"`              // Users 📦 relation: one2many res.users
	VAT                 string `json:"vat"`                   // VAT/Tax ID
	Website             string `json:"website"`               // Website Link
	XOutsideSalesperson any    `json:"x_Outside_Salesperson"` // Not Use-Outside Salesperson 📦 relation: many2one res.users ⚠️ manual: True
	XFax                string `json:"x_fax"`                 // Fax ⚠️ manual: True
	ZIP                 string `json:"zip"`                   // Zip
}

func (m *Model) ResPartner() {
	model := "res.partner"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	currency_map := m.currencyMap()
	pricelist_map := m.pricelistMap()
	old_pricelist_map := m.oldPricelistMap()

	res_users, _ := m.Source.SearchRead("res.users", 0, 0, []string{"id", "partner_id"}, []any{
		[]any{"name", "!=", "Administrator"},
	})
	source_res_partner_ids := []int{}
	for _, ru := range res_users {
		partner_id := int(ru["partner_id"].([]any)[0].(float64))
		source_res_partner_ids = append(source_res_partner_ids, partner_id)
	}
	source_res_partner_ids = append(source_res_partner_ids, 1) // add "My Company (San Francisco)" to exclude

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ResPartner_150{}), []any{
		[]any{"id", "not in", source_res_partner_ids},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	// Companies
	bar := progressbar.Default(int64(len(records)), "company_type: company")
	for _, r := range records {
		var p ResPartner_150
		FillStruct(r, &p)

		if p.CompanyType != "company" {
			bar.Add(1)
			continue
		}
		// fmt.Println(prettyprint(p))
		if strings.TrimSpace(p.Name) == "" {
			bar.Add(1)
			continue
		}

		ur := map[string]any{
			"name":              p.Name,
			"company_type":      p.CompanyType,
			"street":            p.Street,
			"street2":           p.Street2,
			"city":              p.City,
			"zip":               p.ZIP,
			"vat":               p.VAT,
			"l10n_ca_pst":       p.L10NCaPst,
			"phone":             p.Phone,
			"email":             p.Email,
			"website":           p.Website,
			"ref":               p.Ref,
			"comment":           p.Comment,
			"credit_limit":      p.CreditLimit,
			"sale_warn_msg":     p.SaleWarnMsg,
			"purchase_warn_msg": p.PurchaseWarnMsg,
			// "user_id":           p.UserID,
			// "invoice_warn_msg":  p.InvoiceWarnMsg, does not exist in v19
			"picking_warn_msg": p.PickingWarnMsg,
		}
		if p.Type == "private" {
			ur["type"] = "other"
		} else {
			ur["type"] = p.Type
		}
		if p.Image1920 != nil {
			ur["image_1920"] = p.Image1920
		}

		if p.XFax != "" {
			ur["fax"] = p.XFax
		}
		if p.Mobile != "" {
			ur["mobile"] = p.Mobile
		}

		// country
		_, country := ParseMany2One(p.CountryID)
		cid := m.getCountry(country)
		if cid != -1 {
			ur["country_id"] = cid
		}

		// state
		_, state := ParseMany2One(p.StateID)
		stid, err := m.Dest.GetID("res.country.state", []any{
			[]any{"display_name", "=", state},
			[]any{"country_id", "=", cid},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if stid != -1 {
			ur["state_id"] = stid
		}

		// fiscal position
		fiscID := m.getFiscalPositions(cid, stid)
		if fiscID != -1 {
			ur["property_account_position_id"] = fiscID
		}

		// purchase currency
		_, purchase_currency := ParseMany2One(p.PropertyPurchaseCurrencyID)
		purchase_curid := currency_map[purchase_currency]
		if purchase_curid != -1 {
			ur["property_purchase_currency_id"] = purchase_curid
		}

		// pricelist
		_, pricelist := ParseMany2One(p.PropertyProductPricelist)
		old_pricelist_name := old_pricelist_map[pricelist]
		pricelist_id := pricelist_map[old_pricelist_name]
		if pricelist_id != -1 {
			ur["property_product_pricelist"] = pricelist_id
		}

		// inside salesperson
		_, user_name := ParseMany2One(p.UserID)
		inside_salesperson_id, err := m.Dest.GetID("res.users", []any{
			[]any{"name", "=", user_name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if inside_salesperson_id != -1 {
			ur["user_id"] = inside_salesperson_id
		}

		// outside salesperson
		_, outside_salesperson := ParseMany2One(p.OutsideSalespersonID)
		outside_salesperson_id, err := m.Dest.GetID("res.users", []any{
			[]any{"name", "=", outside_salesperson},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if outside_salesperson_id != -1 {
			ur["outside_salesperson_id"] = outside_salesperson_id
		}

		// sales team
		_, s_team := ParseMany2One(p.TeamID)
		s_team_id, err := m.Dest.GetID("crm.team", []any{
			[]any{"name", "=", s_team},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if s_team_id != -1 {
			ur["team_id"] = s_team_id
		}

		// payment terms
		_, payment_terms := ParseMany2One(p.PropertyPaymentTermID)
		payment_term_id, err := m.Dest.GetID("account.payment.term", []any{
			[]any{"name", "=", payment_terms},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if payment_term_id != -1 {
			ur["property_payment_term_id"] = payment_term_id
		}

		// purchase payment terms
		_, purchase_payment_terms := ParseMany2One(p.PropertySupplierPaymentTermID)
		purchase_payment_term_id, err := m.Dest.GetID("account.payment.term", []any{
			[]any{"name", "=", purchase_payment_terms},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if purchase_payment_term_id != -1 {
			ur["property_supplier_payment_term_id"] = purchase_payment_term_id
		}

		// if p.PropertyPaymentMethodID != false {
		// 	fmt.Println("property_payment_method_id", p.PropertyPaymentMethodID)
		// 	p_id, p_name := ParseMany2One(p.PropertyPaymentMethodID)
		// 	switch p_name {
		// 	case "Manual":
		// 		ur["property_outbound_payment_method_line_id"] = p_id
		// 	case "Check":
		// 		ur["property_outbound_payment_method_line_id"] = p_id
		// 	}
		// } else {
		// 	ur["property_outbound_payment_method_line_id"] = 2
		// }

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", p.Name},
			[]any{"type", "=", p.Type},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	// Contacts
	bar = progressbar.Default(int64(len(records)), "company_type: person")
	for _, r := range records {
		var p ResPartner_150
		FillStruct(r, &p)

		if p.CompanyType != "person" {
			bar.Add(1)
			continue
		}
		// fmt.Println(prettyprint(p))
		if strings.TrimSpace(p.Name) == "" {
			bar.Add(1)
			continue
		}

		ur := map[string]any{
			"name":              p.Name,
			"company_type":      p.CompanyType,
			"street":            p.Street,
			"street2":           p.Street2,
			"city":              p.City,
			"zip":               p.ZIP,
			"vat":               p.VAT,
			"l10n_ca_pst":       p.L10NCaPst,
			"phone":             p.Phone,
			"email":             p.Email,
			"website":           p.Website,
			"ref":               p.Ref,
			"comment":           p.Comment,
			"credit_limit":      p.CreditLimit,
			"sale_warn_msg":     p.SaleWarnMsg,
			"purchase_warn_msg": p.PurchaseWarnMsg,
			// "user_id":           p.UserID,
			// "invoice_warn_msg":  p.InvoiceWarnMsg, does not exist in v19
			"picking_warn_msg": p.PickingWarnMsg,
		}
		if p.Type == "private" {
			ur["type"] = "other"
		} else {
			ur["type"] = p.Type
		}
		if p.Image1920 != nil {
			ur["image_1920"] = p.Image1920
		}

		_, parent := ParseMany2One(p.ParentID)
		parent_id, _ := m.Dest.GetID("res.partner", []any{
			[]any{"name", "=", parent},
		})
		if parent_id != -1 {
			ur["parent_id"] = parent_id
		}

		// country
		_, country := ParseMany2One(p.CountryID)
		cid := m.getCountry(country)
		if cid != -1 {
			ur["country_id"] = cid
		}

		// state
		_, state := ParseMany2One(p.StateID)
		_, stid := m.getCountryState(country, state)
		if stid != -1 {
			ur["state_id"] = stid
		}

		// fiscal position
		fiscID := m.getFiscalPositions(cid, stid)
		if fiscID != -1 {
			ur["property_account_position_id"] = fiscID
		}

		// purchase currency
		_, purchase_currency := ParseMany2One(p.PropertyPurchaseCurrencyID)
		purchase_curid := currency_map[purchase_currency]
		if purchase_curid != -1 {
			ur["property_purchase_currency_id"] = purchase_curid
		}

		// outside salesperson
		_, outside_salesperson := ParseMany2One(p.OutsideSalespersonID)
		outside_salesperson_id, err := m.Dest.GetID("res.users", []any{
			[]any{"name", "=", outside_salesperson},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if outside_salesperson_id != -1 {
			ur["outside_salesperson_id"] = outside_salesperson_id
		}

		// sales team
		_, s_team := ParseMany2One(p.TeamID)
		s_team_id, err := m.Dest.GetID("crm.team", []any{
			[]any{"name", "=", s_team},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if s_team_id != -1 {
			ur["team_id"] = s_team_id
		}

		// payment terms
		_, payment_terms := ParseMany2One(p.PropertyPaymentTermID)
		payment_term_id, err := m.Dest.GetID("account.payment.term", []any{
			[]any{"name", "=", payment_terms},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if payment_term_id != -1 {
			ur["property_payment_term_id"] = payment_term_id
		}

		// purchase payment terms
		_, purchase_payment_terms := ParseMany2One(p.PropertySupplierPaymentTermID)
		purchase_payment_term_id, err := m.Dest.GetID("account.payment.term", []any{
			[]any{"name", "=", purchase_payment_terms},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}
		if purchase_payment_term_id != -1 {
			ur["property_supplier_payment_term_id"] = purchase_payment_term_id
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", p.Name},
			[]any{"type", "=", p.Type},
			[]any{"parent_id", "=", parent_id},
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
