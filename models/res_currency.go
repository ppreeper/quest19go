package models

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data res.currency model
type ResCurrency_150 struct {
	Active                   bool      `json:"active"`                      // Active
	CurrencySubunitLabel     string    `json:"currency_subunit_label"`      // Currency Subunit
	CurrencyUnitLabel        string    `json:"currency_unit_label"`         // Currency Unit
	Date                     time.Time `json:"date"`                        // Date
	DecimalPlaces            int       `json:"decimal_places"`              // Decimal Places
	DisplayName              string    `json:"display_name"`                // Display Name
	DisplayRoundingWarning   bool      `json:"display_rounding_warning"`    // Display Rounding Warning
	FullName                 string    `json:"full_name"`                   // Name
	ID                       int       `json:"id"`                          // ID
	InverseRate              float64   `json:"inverse_rate"`                // Inverse Rate
	IsCurrentCompanyCurrency bool      `json:"is_current_company_currency"` // Is Current Company Currency
	Name                     string    `json:"name"`                        // Currency ⭐ required
	Position                 any       `json:"position"`                    // Symbol Position
	Rate                     float64   `json:"rate"`                        // Current Rate
	RateIDs                  any       `json:"rate_ids"`                    // Rates 📦 relation: one2many res.currency.rate
	RateString               string    `json:"rate_string"`                 // Rate String
	Rounding                 float64   `json:"rounding"`                    // Rounding Factor
	Symbol                   string    `json:"symbol"`                      // Symbol ⭐ required
}

// quest19data res.currency model
type ResCurrency_190 struct {
	Active                   bool      `json:"active"`                      // Active
	CurrencySubunitLabel     string    `json:"currency_subunit_label"`      // Currency Subunit
	CurrencyUnitLabel        string    `json:"currency_unit_label"`         // Currency Unit
	Date                     time.Time `json:"date"`                        // Date
	DecimalPlaces            int       `json:"decimal_places"`              // Decimal Places
	DisplayName              string    `json:"display_name"`                // Display Name
	DisplayRoundingWarning   bool      `json:"display_rounding_warning"`    // Display Rounding Warning
	FiscalCountryCodes       string    `json:"fiscal_country_codes"`        // Fiscal Country Codes
	FullName                 string    `json:"full_name"`                   // Name
	ID                       int       `json:"id"`                          // ID
	InverseRate              float64   `json:"inverse_rate"`                // Inverse Rate
	IsCurrentCompanyCurrency bool      `json:"is_current_company_currency"` // Is Current Company Currency
	IsoNumeric               int       `json:"iso_numeric"`                 // Currency numeric code.
	Name                     string    `json:"name"`                        // Currency ⭐ required
	Position                 any       `json:"position"`                    // Symbol Position
	Rate                     float64   `json:"rate"`                        // Current Rate
	RateIDs                  any       `json:"rate_ids"`                    // Rates 📦 relation: one2many res.currency.rate
	RateString               string    `json:"rate_string"`                 // Rate String
	Rounding                 float64   `json:"rounding"`                    // Rounding Factor
	Symbol                   string    `json:"symbol"`                      // Symbol ⭐ required
}

func (m *Model) ResCurrency() {
	model := "res.currency"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(ResCurrency_150{})

	currencies, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"active", "=", true},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	banner.Println(model, "currencies")
	bar := progressbar.Default(-1, "currencies")
	for _, currency := range currencies {
		var root ResCurrency_150
		FillStruct(currency, &root)

		ur := map[string]any{
			"active":         true,
			"symbol":         root.Symbol,
			"rounding":       root.Rounding,
			"decimal_places": root.DecimalPlaces,
			"position":       root.Position,
		}

		rida, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", root.Name},
		})
		ridi, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", root.Name},
			[]any{"active", "=", false},
		})
		var rid int
		if rida != -1 {
			rid = rida
		}
		if rida == -1 && ridi != -1 {
			rid = ridi
		}
		if rida == -1 && ridi == -1 {
			rid = -1
		}

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}
