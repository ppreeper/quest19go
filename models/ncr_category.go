package models

import (
	"strings"

	"github.com/schollz/progressbar/v3"
)

// quest15data ncr.category model
type NcrCategory_150 struct {
	DisplayName      string `json:"display_name"`       // Display Name
	ID               int    `json:"id"`                 // ID
	Name             string `json:"name"`               // Category ⭐ required
	ParentCategoryID any    `json:"parent_category_id"` // Parent Category 📦 relation: many2one ncr.category
}

// quest19data ncr.category model
type NcrCategory_190 struct {
	DisplayName      string `json:"display_name"`       // Display Name
	ID               int    `json:"id"`                 // ID
	Name             string `json:"name"`               // Category ⭐ required
	ParentCategoryID any    `json:"parent_category_id"` // Parent Category 📦 relation: many2one ncr.category
}

func (m *Model) NCRCategory() {
	model := "ncr.category"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(NcrCategory_150{})

	// root
	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"parent_category_id", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}

	banner.Println(model, "layer0")
	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var record NcrCategory_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record.Name))

		ur := map[string]any{
			"name": strings.ToUpper(record.Name),
		}

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", strings.ToUpper(record.Name)},
			[]any{"parent_category_id", "=", false},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}

		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	// layer 1
	banner.Println(model, "layer1")
	for _, r0 := range records {
		var record0 NcrCategory_150
		FillStruct(r0, &record0)

		records1, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
			[]any{"parent_category_id", "=", record0.ID},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		if len(records1) == 0 {
			continue
		}
		bar1 := progressbar.Default(-1, model)
		for _, r1 := range records1 {
			var record1 NcrCategory_150
			FillStruct(r1, &record1)
			// fmt.Println(prettyprint(record1))

			parentID, err := m.Dest.GetID(model, []any{
				[]any{"name", "=", strings.ToUpper(record0.Name)},
				[]any{"parent_category_id", "=", false},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}

			ur := map[string]any{
				"name":               strings.ToUpper(record1.Name),
				"parent_category_id": parentID,
			}

			rid, err := m.Dest.GetID(model, []any{
				[]any{"name", "=", strings.ToUpper(record1.Name)},
				[]any{"parent_category_id", "=", parentID},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}

			m.writeRecord(model, ur, rid)
			bar1.Add(1)
		}
		bar1.Finish()
	}

	// layer 2
	banner.Println(model, "layer2")
	for _, r0 := range records {
		var record0 NcrCategory_150
		FillStruct(r0, &record0)

		records1, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
			[]any{"parent_category_id", "=", record0.ID},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		for _, r1 := range records1 {
			var record1 NcrCategory_150
			FillStruct(r1, &record1)

			records2, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
				[]any{"parent_category_id", "=", record1.ID},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if len(records2) == 0 {
				continue
			}
			bar2 := progressbar.Default(-1, model)
			for _, r2 := range records2 {
				var record2 NcrCategory_150
				FillStruct(r2, &record2)
				// fmt.Println(prettyprint(record2))

				parent1ID, err := m.Dest.GetID(model, []any{
					[]any{"name", "=", strings.ToUpper(record0.Name)},
					[]any{"parent_category_id", "=", false},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}

				parent2ID, err := m.Dest.GetID(model, []any{
					[]any{"name", "=", strings.ToUpper(record1.Name)},
					[]any{"parent_category_id", "=", parent1ID},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}

				ur := map[string]any{
					"name":               strings.ToUpper(record2.Name),
					"parent_category_id": parent2ID,
				}

				rid, err := m.Dest.GetID(model, []any{
					[]any{"name", "=", strings.ToUpper(record2.Name)},
					[]any{"parent_category_id", "=", parent2ID},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}

				m.writeRecord(model, ur, rid)
				bar2.Add(1)
			}
			bar2.Finish()
		}
	}

	// layer 3
	banner.Println(model, "layer3")
	for _, r0 := range records {
		var record0 NcrCategory_150
		FillStruct(r0, &record0)

		records1, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
			[]any{"parent_category_id", "=", record0.ID},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		for _, r1 := range records1 {
			var record1 NcrCategory_150
			FillStruct(r1, &record1)

			records2, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
				[]any{"parent_category_id", "=", record1.ID},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			for _, r2 := range records2 {
				var record2 NcrCategory_150
				FillStruct(r2, &record2)

				records3, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
					[]any{"parent_category_id", "=", record2.ID},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}
				if len(records3) == 0 {
					continue
				}
				bar3 := progressbar.Default(-1, model)
				for _, r3 := range records3 {
					var record3 NcrCategory_150
					FillStruct(r3, &record3)
					// fmt.Println(prettyprint(record3))

					parent1ID, err := m.Dest.GetID(model, []any{
						[]any{"name", "=", strings.ToUpper(record0.Name)},
						[]any{"parent_category_id", "=", false},
					})
					if err != nil {
						m.Log.Error(model, "func", trace(), "err", err)
						continue
					}

					parent2ID, err := m.Dest.GetID(model, []any{
						[]any{"name", "=", strings.ToUpper(record1.Name)},
						[]any{"parent_category_id", "=", parent1ID},
					})
					if err != nil {
						m.Log.Error(model, "func", trace(), "err", err)
						continue
					}

					parent3ID, err := m.Dest.GetID(model, []any{
						[]any{"name", "=", strings.ToUpper(record2.Name)},
						[]any{"parent_category_id", "=", parent2ID},
					})
					if err != nil {
						m.Log.Error(model, "func", trace(), "err", err)
						continue
					}

					ur := map[string]any{
						"name":               strings.ToUpper(record3.Name),
						"parent_category_id": parent3ID,
					}

					rid, err := m.Dest.GetID(model, []any{
						[]any{"name", "=", strings.ToUpper(record3.Name)},
						[]any{"parent_category_id", "=", parent3ID},
					})
					if err != nil {
						m.Log.Error(model, "func", trace(), "err", err)
						continue
					}

					m.writeRecord(model, ur, rid)
					bar3.Add(1)
				}
				bar3.Finish()
			}
		}
	}

	// layer 4
	banner.Println(model, "layer4")
	for _, r0 := range records {
		var record0 NcrCategory_150
		FillStruct(r0, &record0)

		records1, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
			[]any{"parent_category_id", "=", record0.ID},
		})
		if err != nil {
			m.Log.Error(model, "func", trace(), "err", err)
			continue
		}
		for _, r1 := range records1 {
			var record1 NcrCategory_150
			FillStruct(r1, &record1)

			records2, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
				[]any{"parent_category_id", "=", record1.ID},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			for _, r2 := range records2 {
				var record2 NcrCategory_150
				FillStruct(r2, &record2)

				records3, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
					[]any{"parent_category_id", "=", record2.ID},
				})
				if err != nil {
					m.Log.Error(model, "func", trace(), "err", err)
					continue
				}
				for _, r3 := range records3 {
					var record3 NcrCategory_150
					FillStruct(r3, &record3)

					records4, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
						[]any{"parent_category_id", "=", record3.ID},
					})
					if err != nil {
						m.Log.Error(model, "func", trace(), "err", err)
						continue
					}
					if len(records4) == 0 {
						continue
					}
					bar4 := progressbar.Default(-1, model)
					for _, r4 := range records4 {
						var record4 NcrCategory_150
						FillStruct(r4, &record4)
						// fmt.Println(prettyprint(record4))

						parent1ID, err := m.Dest.GetID(model, []any{
							[]any{"name", "=", strings.ToUpper(record0.Name)},
							[]any{"parent_category_id", "=", false},
						})
						if err != nil {
							m.Log.Error(model, "func", trace(), "err", err)
							continue
						}

						parent2ID, err := m.Dest.GetID(model, []any{
							[]any{"name", "=", strings.ToUpper(record1.Name)},
							[]any{"parent_category_id", "=", parent1ID},
						})
						if err != nil {
							m.Log.Error(model, "func", trace(), "err", err)
							continue
						}

						parent3ID, err := m.Dest.GetID(model, []any{
							[]any{"name", "=", strings.ToUpper(record2.Name)},
							[]any{"parent_category_id", "=", parent2ID},
						})
						if err != nil {
							m.Log.Error(model, "func", trace(), "err", err)
							continue
						}

						parent4ID, err := m.Dest.GetID(model, []any{
							[]any{"name", "=", strings.ToUpper(record3.Name)},
							[]any{"parent_category_id", "=", parent3ID},
						})
						if err != nil {
							m.Log.Error(model, "func", trace(), "err", err)
							continue
						}

						ur := map[string]any{
							"name":               strings.ToUpper(record4.Name),
							"parent_category_id": parent4ID,
						}

						rid, err := m.Dest.GetID(model, []any{
							[]any{"name", "=", strings.ToUpper(record4.Name)},
							[]any{"parent_category_id", "=", parent4ID},
						})
						if err != nil {
							m.Log.Error(model, "func", trace(), "err", err)
							continue
						}

						m.writeRecord(model, ur, rid)
						bar4.Add(1)
					}
					bar4.Finish()
				}
			}
		}
	}
}
