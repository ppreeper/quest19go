package models

import (
	"github.com/schollz/progressbar/v3"
)

type MrpWorkcenterTag_150 struct {
	Name        string `json:"name"`          // Tag Name ⭐ Required
	Color       int    `json:"color"`         // Color Index
	ID          int    `json:"id"`            // ID
	LastUpdate  any    `json:"__last_update"` // Last Modified on
	DisplayName string `json:"display_name"`  // Display Name
	CreateUid   any    `json:"create_uid"`    // Created by 📦 Model Relation: res.users
	CreateDate  any    `json:"create_date"`   // Created on
	WriteUid    any    `json:"write_uid"`     // Last Updated by 📦 Model Relation: res.users
	WriteDate   any    `json:"write_date"`    // Last Updated on
}

func (m *Model) MRPWorkcenterTag() {
	model := "mrp.workcenter.tag"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MrpWorkcenterTag_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var record MrpWorkcenterTag_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record))

		ur := map[string]any{
			"name":  record.Name,
			"color": record.Color,
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", record.Name},
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
