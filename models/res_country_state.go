package models

// quest19data res.country.state model
type ResCountryState_190 struct {
	Code        string `json:"code"`         // State Code ⭐ required
	CountryID   any    `json:"country_id"`   // Country 📦 relation: many2one res.country ⭐ required
	DisplayName string `json:"display_name"` // Display Name
	ID          int    `json:"id"`           // ID
	Name        string `json:"name"`         // State Name ⭐ required
}
