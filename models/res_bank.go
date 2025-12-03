package models

import (
	"strings"

	"github.com/schollz/progressbar/v3"
)

type ResBank_150 struct {
	Name        string `json:"name"`          // Name ⭐ Required
	Street      string `json:"street"`        // Street
	Street2     string `json:"street2"`       // Street2
	Zip         string `json:"zip"`           // Zip
	City        string `json:"city"`          // City
	State       any    `json:"state"`         // Fed. State 📦 Model Relation: res.country.state
	Country     any    `json:"country"`       // Country 📦 Model Relation: res.country
	Email       string `json:"email"`         // Email
	Phone       string `json:"phone"`         // Phone
	Active      bool   `json:"active"`        // Active
	Bic         string `json:"bic"`           // Bank Identifier Code
	ID          int    `json:"id"`            // ID
	LastUpdate  any    `json:"__last_update"` // Last Modified on
	DisplayName string `json:"display_name"`  // Display Name
	CreateUid   any    `json:"create_uid"`    // Created by 📦 Model Relation: res.users
	CreateDate  any    `json:"create_date"`   // Created on
	WriteUid    any    `json:"write_uid"`     // Last Updated by 📦 Model Relation: res.users
	WriteDate   any    `json:"write_date"`    // Last Updated on
}

func (m *Model) ResBank() {
	model := "res.bank"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	records, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ResBank_150{}), []any{
		// []any{"name", "=", "CANADIAN WESTERN BANK"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}
	bar := progressbar.Default(-1, "banks")
	for _, record := range records {
		var p ResBank_150
		FillStruct(record, &p)
		// fmt.Println(prettyprint(p))

		ur := map[string]any{
			"name": p.Name,
			"bic":  strings.ToUpper(p.Bic),
		}

		if p.Phone != "" {
			ur["phone"] = p.Phone
		}
		if p.Email != "" {
			ur["email"] = p.Email
		}
		if p.Street != "" {
			ur["street"] = p.Street
		}
		if p.Street2 != "" {
			ur["street2"] = p.Street2
		}
		if p.Zip != "" {
			ur["zip"] = p.Zip
		}
		if p.City != "" {
			ur["city"] = p.City
		}

		_, country := ParseMany2One(p.Country)
		// fmt.Println("Country Name:", country)
		cid, err := m.Dest.GetID("res.country", []any{
			[]any{"name", "=", country},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if cid > 0 {
			ur["country"] = cid
		}

		_, state := ParseMany2One(p.State)
		// fmt.Println("State Name:", state)
		stid, err := m.Dest.GetID("res.country.state", []any{
			[]any{"country_id", "=", cid},
			[]any{"name", "=", state},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if stid > 0 {
			ur["state_id"] = stid
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", p.Name},
			[]any{"bic", "=", strings.ToUpper(p.Bic)},
		})

		// fmt.Println(model, "rid", rid, "ur", ur)

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}
