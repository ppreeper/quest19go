package models

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

type MrpWorkcenter_150 struct {
	ResourceID               any    `json:"resource_id"`                // Resource 📦 Model Relation: resource.resource ⭐ Required
	CompanyID                any    `json:"company_id"`                 // Company 📦 Model Relation: res.company
	ResourceCalendarID       any    `json:"resource_calendar_id"`       // Working Hours 📦 Model Relation: resource.calendar
	Tz                       any    `json:"tz"`                         // Timezone
	Name                     string `json:"name"`                       // Work Center
	TimeEfficiency           any    `json:"time_efficiency"`            // Time Efficiency
	Active                   bool   `json:"active"`                     // Active
	Code                     string `json:"code"`                       // Code
	Note                     string `json:"note"`                       // Description
	Capacity                 any    `json:"capacity"`                   // Capacity
	Sequence                 int    `json:"sequence"`                   // Sequence ⭐ Required
	Color                    int    `json:"color"`                      // Color
	CostsHour                any    `json:"costs_hour"`                 // Cost per hour
	TimeStart                any    `json:"time_start"`                 // Setup Time
	TimeStop                 any    `json:"time_stop"`                  // Cleanup Time
	RoutingLineIDs           any    `json:"routing_line_ids"`           // Routing Lines 📦 Model Relation: mrp.routing.workcenter
	OrderIDs                 any    `json:"order_ids"`                  // Orders 📦 Model Relation: mrp.workorder
	WorkorderCount           int    `json:"workorder_count"`            // # Work Orders
	WorkorderReadyCount      int    `json:"workorder_ready_count"`      // # Read Work Orders
	WorkorderProgressCount   int    `json:"workorder_progress_count"`   // Total Running Orders
	WorkorderPendingCount    int    `json:"workorder_pending_count"`    // Total Pending Orders
	WorkorderLateCount       int    `json:"workorder_late_count"`       // Total Late Orders
	TimeIDs                  any    `json:"time_ids"`                   // Time Logs 📦 Model Relation: mrp.workcenter.productivity
	WorkingState             any    `json:"working_state"`              // Workcenter Status
	BlockedTime              any    `json:"blocked_time"`               // Blocked Time
	ProductiveTime           any    `json:"productive_time"`            // Productive Time
	Oee                      any    `json:"oee"`                        // Oee
	OeeTarget                any    `json:"oee_target"`                 // OEE Target
	Performance              int    `json:"performance"`                // Performance
	WorkcenterLoad           any    `json:"workcenter_load"`            // Work Center Load
	AlternativeWorkcenterIDs any    `json:"alternative_workcenter_ids"` // Alternative Workcenters 📦 Model Relation: mrp.workcenter
	TagIDs                   any    `json:"tag_ids"`                    // Tag 📦 Model Relation: mrp.workcenter.tag
	ID                       int    `json:"id"`                         // ID
	LastUpdate               any    `json:"__last_update"`              // Last Modified on
	DisplayName              string `json:"display_name"`               // Display Name
	CreateUid                any    `json:"create_uid"`                 // Created by 📦 Model Relation: res.users
	CreateDate               any    `json:"create_date"`                // Created on
	WriteUid                 any    `json:"write_uid"`                  // Last Updated by 📦 Model Relation: res.users
	WriteDate                any    `json:"write_date"`                 // Last Updated on
	CostsHourAccountID       any    `json:"costs_hour_account_id"`      // Analytic Account 📦 Model Relation: account.analytic.account
	EquipmentIDs             any    `json:"equipment_ids"`              // Maintenance Equipment 📦 Model Relation: maintenance.equipment
}

func (m *Model) MRPWorkcenter() {
	model := "mrp.workcenter"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MrpWorkcenter_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		// []any{"deprecated", "=", false},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	bar := progressbar.Default(-1, model)
	for _, r := range records {
		var record MrpWorkcenter_150
		FillStruct(r, &record)
		// fmt.Println(prettyprint(record.TagIDs))

		ur := map[string]any{
			"name":            record.Name,
			"code":            record.Code,
			"costs_hour":      record.CostsHour,
			"time_efficiency": record.TimeEfficiency,
			"oee_target":      record.OeeTarget,
			"time_start":      record.TimeStart,
			"time_stop":       record.TimeStop,
			"note":            record.Note,
			"sequence":        record.Sequence,
		}
		// Tags
		var tags []string
		for _, v := range record.TagIDs.([]any) {
			tagSource, _ := m.Source.SearchRead("mrp.workcenter.tag", 0, 1, []string{"name"}, []any{
				[]any{"id", "=", v},
			})
			for _, tag := range tagSource {
				tags = append(tags, tag["name"].(string))
			}
		}
		var tagIDs []any
		for _, tag := range tags {
			tid, err := m.Dest.GetID("mrp.workcenter.tag", []any{
				[]any{"name", "=", tag},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
				continue
			}
			if tid == -1 {
				m.Log.Error(model, "func", trace(), "err", fmt.Sprintf("Tag %s not found", tag))
				continue
			}
			tagIDs = append(tagIDs, tid)
		}
		ur["tag_ids"] = tagIDs

		rid, err := m.Dest.GetID(model, []any{
			[]any{"name", "=", record.Name},
			[]any{"code", "=", record.Code},
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
