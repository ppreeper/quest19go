package models

type HrDepartment_150 struct {
	MessageIsFollower           bool   `json:"message_is_follower"`             // Is Follower
	MessageFollowerIDs          any    `json:"message_follower_ids"`            // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs           any    `json:"message_partner_ids"`             // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs                  any    `json:"message_ids"`                     // Messages 📦 Model Relation: mail.message
	HasMessage                  bool   `json:"has_message"`                     // Has Message
	MessageUnread               bool   `json:"message_unread"`                  // Unread Messages
	MessageUnreadCounter        int    `json:"message_unread_counter"`          // Unread Messages Counter
	MessageNeedaction           bool   `json:"message_needaction"`              // Action Needed
	MessageNeedactionCounter    int    `json:"message_needaction_counter"`      // Number of Actions
	MessageHasError             bool   `json:"message_has_error"`               // Message Delivery error
	MessageHasErrorCounter      int    `json:"message_has_error_counter"`       // Number of errors
	MessageAttachmentCount      int    `json:"message_attachment_count"`        // Attachment Count
	MessageMainAttachmentID     any    `json:"message_main_attachment_id"`      // Main Attachment 📦 Model Relation: ir.attachment
	WebsiteMessageIDs           any    `json:"website_message_ids"`             // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError          bool   `json:"message_has_sms_error"`           // SMS Delivery error
	Name                        string `json:"name"`                            // Department Name ⭐ Required
	CompleteName                string `json:"complete_name"`                   // Complete Name
	Active                      bool   `json:"active"`                          // Active
	CompanyID                   any    `json:"company_id"`                      // Company 📦 Model Relation: res.company
	ParentID                    any    `json:"parent_id"`                       // Parent Department 📦 Model Relation: hr.department
	ChildIDs                    any    `json:"child_ids"`                       // Child Departments 📦 Model Relation: hr.department
	ManagerID                   any    `json:"manager_id"`                      // Manager 📦 Model Relation: hr.employee
	MemberIDs                   any    `json:"member_ids"`                      // Members 📦 Model Relation: hr.employee
	TotalEmployee               int    `json:"total_employee"`                  // Total Employee
	JobsIDs                     any    `json:"jobs_ids"`                        // Jobs 📦 Model Relation: hr.job
	Note                        string `json:"note"`                            // Note
	Color                       int    `json:"color"`                           // Color Index
	ID                          int    `json:"id"`                              // ID
	LastUpdate                  any    `json:"__last_update"`                   // Last Modified on
	DisplayName                 string `json:"display_name"`                    // Display Name
	CreateUid                   any    `json:"create_uid"`                      // Created by 📦 Model Relation: res.users
	CreateDate                  any    `json:"create_date"`                     // Created on
	WriteUid                    any    `json:"write_uid"`                       // Last Updated by 📦 Model Relation: res.users
	WriteDate                   any    `json:"write_date"`                      // Last Updated on
	AppraisalsToProcessCount    int    `json:"appraisals_to_process_count"`     // Appraisals to Process
	EmployeeFeedbackTemplate    string `json:"employee_feedback_template"`      // Employee Feedback Template
	ManagerFeedbackTemplate     string `json:"manager_feedback_template"`       // Manager Feedback Template
	CustomAppraisalTemplates    bool   `json:"custom_appraisal_templates"`      // Custom Appraisal Templates
	AbsenceOfToday              int    `json:"absence_of_today"`                // Absence by Today
	LeaveToApproveCount         int    `json:"leave_to_approve_count"`          // Time Off to Approve
	AllocationToApproveCount    int    `json:"allocation_to_approve_count"`     // Allocation to Approve
	NewApplicantCount           int    `json:"new_applicant_count"`             // New Applicant
	NewHiredEmployee            int    `json:"new_hired_employee"`              // New Hired Employee
	ExpectedEmployee            int    `json:"expected_employee"`               // Expected Employee
	ExpenseSheetsToApproveCount int    `json:"expense_sheets_to_approve_count"` // Expenses Reports to Approve
}

func (m *Model) HRDepartment() {
	model := "hr.department"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourcefields := ExtractJSONTags(HrDepartment_150{})

	departments, err := m.Source.SearchRead(model, 0, 0, sourcefields, []any{})
	if err != nil {
		m.Log.Error(model, "SearchRead", err)
		return
	}

	for _, u := range departments {
		var department HrDepartment_150
		FillStruct(u, &department)
		// fmt.Println(prettyprint(department.Name))

		ur := map[string]any{
			"name":  department.Name,
			"color": department.Color,
			"note":  department.Note,
		}

		parent, _ := m.Source.SearchRead(model, 0, 0, []string{"id", "name"}, []any{
			[]any{"id", "=", department.ParentID},
		})
		if len(parent) != 0 {
			parent_id, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", parent[0]["name"]},
			})
			ur["parent_id"] = parent_id
		}

		manager, _ := m.Source.SearchRead("hr.employee", 0, 0, []string{"id", "name", "user_id"}, []any{
			[]any{"id", "=", department.ManagerID.([]any)[0]},
		})
		if len(manager) != 0 {
			manager_id, _ := m.Dest.GetID("hr.employee", []any{
				[]any{"name", "=", manager[0]["user_id"].([]any)[1]},
			})
			ur["manager_id"] = manager_id
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", department.Name},
		})
		// fmt.Println("rid", rid)

		m.writeRecord(model, ur, rid)

	}
}
