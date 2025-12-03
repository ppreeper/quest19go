package models

import (
	"github.com/schollz/progressbar/v3"
)

type DecimalPrecision_150 struct {
	Name        string `json:"name"`          // Usage ⭐ Required
	Digits      int    `json:"digits"`        // Digits ⭐ Required
	ID          int    `json:"id"`            // ID
	LastUpdate  any    `json:"__last_update"` // Last Modified on
	DisplayName string `json:"display_name"`  // Display Name
	CreateUid   any    `json:"create_uid"`    // Created by 📦 Model Relation: res.users
	CreateDate  any    `json:"create_date"`   // Created on
	WriteUid    any    `json:"write_uid"`     // Last Updated by 📦 Model Relation: res.users
	WriteDate   any    `json:"write_date"`    // Last Updated on
}

// Decimal Precisions v15
// Account	2
// Discount	5
// Payment Terms	6
// Payroll	2
// Payroll Rate	4
// Product Price	2
// Product Unit of Measure	3
// Product UoS	3
// Quality Tests	2
// Stock Weight	2
// Volume 2

// Decimal Precisions v19
// Discount	2
// Payment Terms	6
// Percentage Analytic	2
// Product Price	2
// Product Unit	2
// Quality Tests	2
// Stock Weight	2
// Volume	2

var precision_map = map[string]string{
	"Account":                 "Account",
	"Discount":                "Discount",
	"Payment Terms":           "Payment Terms",
	"Payroll":                 "Payroll",
	"Payroll Rate":            "Payroll Rate",
	"Product Price":           "Product Price",
	"Product Unit of Measure": "Product Unit",
	"Product UoS":             "Product UoS",
	"Quality Tests":           "Quality Tests",
	"Stock Weight":            "Stock Weight",
	"Volume":                  "Volume",
}

func (m *Model) DecimalPrecision() {
	model := "decimal.precision"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(DecimalPrecision_150{})

	precisions, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	banner.Println(model, "precisions")
	bar := progressbar.Default(-1, "precisions")
	for _, precision := range precisions {
		var root DecimalPrecision_150
		FillStruct(precision, &root)

		ur := map[string]any{
			"digits": root.Digits,
		}
		name := precision_map[root.Name]
		if name == "" {
			name = root.Name
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", name},
		})

		if rid == -1 {
			continue
		}
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}
