package models

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data account.payment.term model
type AccountPaymentTerm_150 struct {
	CompanyID   any    `json:"company_id"`   // Company 📦 relation: many2one res.company
	DisplayName string `json:"display_name"` // Display Name
	ID          int    `json:"id"`           // ID
	LineIDs     any    `json:"line_ids"`     // Terms 📦 relation: one2many account.payment.term.line
	Name        string `json:"name"`         // Payment Terms ⭐ required
	Note        string `json:"note"`         // Description on the Invoice
	Sequence    int    `json:"sequence"`     // Sequence ⭐ required
}

// quest19data account.payment.term model
type AccountPaymentTerm_190 struct {
	CompanyID                   any       `json:"company_id"`                     // Company 📦 relation: many2one res.company
	CurrencyID                  any       `json:"currency_id"`                    // Currency 📦 relation: many2one res.currency
	DiscountDays                int       `json:"discount_days"`                  // Discount Days
	DiscountPercentage          float64   `json:"discount_percentage"`            // Discount %
	DisplayName                 string    `json:"display_name"`                   // Display Name
	DisplayOnInvoice            bool      `json:"display_on_invoice"`             // Show installment dates
	EarlyDiscount               bool      `json:"early_discount"`                 // Early Discount
	EarlyPayDiscountComputation any       `json:"early_pay_discount_computation"` // Cash Discount Tax Reduction
	ExampleAmount               float64   `json:"example_amount"`                 // Example Amount
	ExampleDate                 time.Time `json:"example_date"`                   // Date example
	ExampleInvalid              bool      `json:"example_invalid"`                // Example Invalid
	ExamplePreview              string    `json:"example_preview"`                // Example Preview
	ExamplePreviewDiscount      string    `json:"example_preview_discount"`       // Example Preview Discount
	FiscalCountryCodes          string    `json:"fiscal_country_codes"`           // Fiscal Country Codes
	ID                          int       `json:"id"`                             // ID
	LineIDs                     any       `json:"line_ids"`                       // Terms 📦 relation: one2many account.payment.term.line
	Name                        string    `json:"name"`                           // Payment Terms ⭐ required
	Note                        string    `json:"note"`                           // Description on the Invoice
	Sequence                    int       `json:"sequence"`                       // Sequence ⭐ required
}

// 19 existing mapping
// Immediate Payment | Immediate Payment
// 15 Days | Net 15 > RENAME
// 21 Days | 21 Days
// 30 Days | Net 30 > RENAME
// 45 Days | Net 45 > RENAME
// End of Following Month | End of Following Month
// 10 Days after End of Next Month | NA
// 30% Now, Balance 60 Days | 30% Now, Balance 60 Days
// 2/7 Net 30 | NA
// 90 days, on the 10th | NA

func (m *Model) AccountPaymentTermRename() {
	model := "account.payment.term"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// rename 15 Days to Net 15
	tid, err := m.Dest.GetID(model, []any{[]any{"name", "=", "15 Days"}})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if tid != -1 {
		ur := make(map[string]any)
		ur["name"] = "Net 15"
		m.writeRecord(model, ur, tid)
	}

	// rename 30 Days to Net 30
	tid, err = m.Dest.GetID(model, []any{[]any{"name", "=", "30 Days"}})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if tid != -1 {
		ur := make(map[string]any)
		ur["name"] = "Net 30"
		m.writeRecord(model, ur, tid)
	}

	// rename 45 Days to Net 45
	tid, err = m.Dest.GetID(model, []any{[]any{"name", "=", "45 Days"}})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if tid != -1 {
		ur := make(map[string]any)
		ur["name"] = "Net 45"
		m.writeRecord(model, ur, tid)
	}
}

func (m *Model) AccountPaymentTerm() {
	model := "account.payment.term"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(AccountPaymentTerm_150{})

	// root
	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"parent_category_id", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var record AccountPaymentTerm_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record))

		ur := make(map[string]any)
		ur["name"] = record.Name
		ur["note"] = record.Note
		ur["sequence"] = record.Sequence

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", record.Name},
			[]any{"sequence", "=", record.Sequence},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

// quest15data account.payment.term.line model
type AccountPaymentTermLine_150 struct {
	DayOfTheMonth int     `json:"day_of_the_month"` // Day of the month
	Days          int     `json:"days"`             // Number of Days ⭐ required
	DisplayName   string  `json:"display_name"`     // Display Name
	ID            int     `json:"id"`               // ID
	Option        any     `json:"option"`           // Options ⭐ required
	PaymentID     any     `json:"payment_id"`       // Payment Terms 📦 relation: many2one account.payment.term ⭐ required
	Sequence      int     `json:"sequence"`         // Sequence
	Value         any     `json:"value"`            // Type ⭐ required
	ValueAmount   float64 `json:"value_amount"`     // Value
}

// quest19data account.payment.term.line model
type AccountPaymentTermLine_190 struct {
	DaysNextMonth        string  `json:"days_next_month"`         // Days on the next month
	DelayType            any     `json:"delay_type"`              // Delay Type ⭐ required
	DisplayDaysNextMonth bool    `json:"display_days_next_month"` // Display Days Next Month
	DisplayName          string  `json:"display_name"`            // Display Name
	ID                   int     `json:"id"`                      // ID
	NbDays               int     `json:"nb_days"`                 // Days
	PaymentID            any     `json:"payment_id"`              // Payment Terms 📦 relation: many2one account.payment.term ⭐ required
	Value                any     `json:"value"`                   // Value ⭐ required
	ValueAmount          float64 `json:"value_amount"`            // Due
}

// Delay Type options
// [days_after] - Days after invoice date
// [days_after_end_of_month] - Days after end of month
// [days_after_end_of_next_month] - Days after end of next month
// [days_end_of_month_on_the] - Day of the month on the

func (m *Model) AccountPaymentTermLine() {
	model := "account.payment.term.line"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// root
	terms, err := m.Source.SearchRead("account.payment.term", 0, 0, ExtractJSONTags(AccountPaymentTerm_150{}), []any{
		// []any{"parent_category_id", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	for _, t := range terms {
		var term AccountPaymentTerm_150
		FillStruct(t, &term)
		// fmt.Println(prettyprint(term))

		lines, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(AccountPaymentTermLine_150{}), []any{
			[]any{"id", "in", term.LineIDs},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			return
		}

		for _, r := range lines {
			var record AccountPaymentTermLine_150
			FillStruct(r, &record)
			fmt.Println(prettyprint(record))

			// var record AccountPaymentTermLine_150
			// FillStruct(r, &record)
			// fmt.Println(prettyprint(record))

			searchTerm := []any{}

			ur := make(map[string]any)
			ur["value"] = record.Value
			ur["value_amount"] = record.ValueAmount
			ur["nb_days"] = record.Days
			ur["delay_type"] = record.Option
			ur["sequence"] = record.Sequence

			pid, err := m.Dest.GetID("account.payment.term", []any{
				[]any{"name", "=", term.Name},
				[]any{"sequence", "=", term.Sequence},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if pid != -1 {
				ur["payment_id"] = pid
				searchTerm = append(searchTerm, []any{"payment_id", "=", pid})
			}

			searchTerm = append(searchTerm, []any{"value", "=", record.Value})
			searchTerm = append(searchTerm, []any{"sequence", "=", record.Sequence})

			rid, err := m.Dest.GetID("account.payment.term.line", searchTerm)
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}

			m.writeRecord(model, ur, rid)

		}
	}
}
