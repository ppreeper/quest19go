package models

import (
	"github.com/schollz/progressbar/v3"
)

type (
	AccountJournal_150 struct {
		Active bool   `json:"active"` // Active
		Name   string `json:"name"`   // Journal Name ⭐ required
		Type   any    `json:"type"`   // Journal Type ⭐ required 📦 selection: account.journal.type
		// Journal Entries Tab
		DefaultAccountID     any    `json:"default_account_id"`     // Default Debit/Credit Account 📦 relation: many2one account.account
		SuspenseAccountID    any    `json:"suspense_account_id"`    // Suspense Account 📦 relation: many2one account.account
		ProfitAccountID      any    `json:"profit_account_id"`      // Profit Account 📦 relation: many2one account.account
		LossAccountID        any    `json:"loss_account_id"`        // Loss Account 📦 relation: many2one account.account
		Code                 string `json:"code"`                   // Journal Code ⭐ required
		CurrencyID           any    `json:"currency_id"`            // Currency 📦 relation: many2one res.currency
		BankAccountID        any    `json:"bank_account_id"`        // Bank Account 📦 relation: many2one account.bank.statement
		BankStatementsSource any    `json:"bank_statements_source"` // Bank Statements Source 📦 selection: account.journal.bank.statements.source
		// Incoming Payments Tab
		// Outgoing Payments Tab
		// Advanced Settings Tab
	}
	AccountJournal_190 struct{}
)

func (m *Model) AccountJournalRename() {
	model := "account.journal"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	rename_map := map[string]string{
		"Miscellaneous Operations": "GENERAL JOURNAL",
		"Inventory Valuation":      "Stock Journal",
	}

	for old_name, new_name := range rename_map {
		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", old_name},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if rid == -1 {
			continue
		}
		ur := make(map[string]any)
		ur["name"] = new_name
		m.writeRecord(model, ur, rid)
	}
}

func (m *Model) AccountJournal() {
	model := "account.journal"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	m.AccountJournalRename()

	currency_map := m.currencyMap()

	sourceFields := ExtractJSONTags(AccountJournal_150{})

	journals, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"name", "not like", "x%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	banner.Println(model, "journals")
	bar := progressbar.Default(-1, "journals")
	for _, r := range journals {
		var journal AccountJournal_150
		FillStruct(r, &journal)
		// fmt.Println(prettyprint(journal))

		ur := make(map[string]any)
		ur["name"] = journal.Name
		ur["type"] = journal.Type
		ur["code"] = journal.Code
		ur["active"] = journal.Active

		if journal.Type == "bank" {
			ur["bank_statements_source"] = journal.BankStatementsSource
			if journal.BankStatementsSource == false {
				ur["bank_statements_source"] = "undefined"
			}
		}

		// currency
		if journal.CurrencyID != false {
			_, currency_name := ParseMany2One(journal.CurrencyID)
			currency_id := currency_map[currency_name]
			if currency_id != -1 {
				ur["currency_id"] = currency_id
			}
		}

		// default_account_id
		if journal.DefaultAccountID != false {
			daid := m.AccountAccountLookup(ParseMany2One(journal.DefaultAccountID))
			if daid != -1 {
				ur["default_account_id"] = daid
			}
		}

		// fmt.Println(prettyprint(ur))

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", journal.Name},
			[]any{"type", "=", journal.Type},
			[]any{"code", "=", journal.Code},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		// fmt.Println("rid", rid)
		m.writeRecord("account.journal", ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}
