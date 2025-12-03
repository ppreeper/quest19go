package models

// quest15data account.fiscal.position model
type AccountFiscalPosition_150 struct {
	AccountIDs           any    `json:"account_ids"`             // Account Mapping 📦 relation: one2many account.fiscal.position.account
	Active               bool   `json:"active"`                  // Active
	AutoApply            bool   `json:"auto_apply"`              // Detect Automatically
	CompanyCountryID     any    `json:"company_country_id"`      // Company Country 📦 relation: many2one res.country
	CompanyID            any    `json:"company_id"`              // Company 📦 relation: many2one res.company ⭐ required
	CountryGroupID       any    `json:"country_group_id"`        // Country Group 📦 relation: many2one res.country.group
	CountryID            any    `json:"country_id"`              // Country 📦 relation: many2one res.country
	DisplayName          string `json:"display_name"`            // Display Name
	ForeignVat           string `json:"foreign_vat"`             // Foreign Tax ID
	ForeignVatHeaderMode any    `json:"foreign_vat_header_mode"` // Foreign Vat Header Mode
	ID                   int    `json:"id"`                      // ID
	Name                 string `json:"name"`                    // Fiscal Position ⭐ required
	Note                 string `json:"note"`                    // Notes
	Sequence             int    `json:"sequence"`                // Sequence
	StateIDs             any    `json:"state_ids"`               // Federal States 📦 relation: many2many res.country.state
	StatesCount          int    `json:"states_count"`            // States Count
	TaxIDs               any    `json:"tax_ids"`                 // Tax Mapping 📦 relation: one2many account.fiscal.position.tax
	VATRequired          bool   `json:"vat_required"`            // VAT required
	ZIPFrom              string `json:"zip_from"`                // Zip Range From
	ZIPTo                string `json:"zip_to"`                  // Zip Range To
}
