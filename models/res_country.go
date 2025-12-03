package models

// quest19data res.country model
type ResCountry_190 struct {
	AddressFormat                 string `json:"address_format"`                    // Layout in Reports
	AddressViewID                 any    `json:"address_view_id"`                   // Input View 📦 relation: many2one ir.ui.view
	Code                          string `json:"code"`                              // Country Code ⭐ required
	CountryGroupCodes             any    `json:"country_group_codes"`               // Country Group Codes
	CountryGroupIDs               any    `json:"country_group_ids"`                 // Country Groups 📦 relation: many2many res.country.group
	CurrencyID                    any    `json:"currency_id"`                       // Currency 📦 relation: many2one res.currency
	DisplayName                   string `json:"display_name"`                      // Display Name
	ID                            int    `json:"id"`                                // ID
	ImageUrl                      string `json:"image_url"`                         // Flag
	IsMercadoPagoSupportedCountry bool   `json:"is_mercado_pago_supported_country"` // Is Mercado Pago Supported Country
	IsStripeSupportedCountry      bool   `json:"is_stripe_supported_country"`       // Is Stripe Supported Country
	Name                          string `json:"name"`                              // Country Name ⭐ required
	NamePosition                  any    `json:"name_position"`                     // Customer Name Position
	PhoneCode                     int    `json:"phone_code"`                        // Country Calling Code
	StateIDs                      any    `json:"state_ids"`                         // States 📦 relation: one2many res.country.state
	StateRequired                 bool   `json:"state_required"`                    // State Required
	VATLabel                      string `json:"vat_label"`                         // Vat Label
	ZIPRequired                   bool   `json:"zip_required"`                      // Zip Required
}
