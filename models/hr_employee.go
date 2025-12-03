package models

import (
	"fmt"
)

type HrEmployee_150 struct {
	Image1920                         any    `json:"image_1920"`                            // Image
	Image1024                         any    `json:"image_1024"`                            // Image 1024
	Image512                          any    `json:"image_512"`                             // Image 512
	Image256                          any    `json:"image_256"`                             // Image 256
	Image128                          any    `json:"image_128"`                             // Image 128
	Avatar1920                        any    `json:"avatar_1920"`                           // Avatar
	Avatar1024                        any    `json:"avatar_1024"`                           // Avatar 1024
	Avatar512                         any    `json:"avatar_512"`                            // Avatar 512
	Avatar256                         any    `json:"avatar_256"`                            // Avatar 256
	Avatar128                         any    `json:"avatar_128"`                            // Avatar 128
	ResourceID                        any    `json:"resource_id"`                           // Resource 📦 Model Relation: resource.resource ⭐ Required
	CompanyID                         any    `json:"company_id"`                            // Company 📦 Model Relation: res.company ⭐ Required
	ResourceCalendarID                any    `json:"resource_calendar_id"`                  // Working Hours 📦 Model Relation: resource.calendar
	Tz                                any    `json:"tz"`                                    // Timezone
	ActivityIDs                       any    `json:"activity_ids"`                          // Activities 📦 Model Relation: mail.activity
	ActivityState                     any    `json:"activity_state"`                        // Activity State
	ActivityUserID                    any    `json:"activity_user_id"`                      // Responsible User 📦 Model Relation: res.users
	ActivityTypeID                    any    `json:"activity_type_id"`                      // Next Activity Type 📦 Model Relation: mail.activity.type
	ActivityTypeIcon                  string `json:"activity_type_icon"`                    // Activity Type Icon
	ActivityDateDeadline              any    `json:"activity_date_deadline"`                // Next Activity Deadline
	MyActivityDateDeadline            any    `json:"my_activity_date_deadline"`             // My Activity Deadline
	ActivitySummary                   string `json:"activity_summary"`                      // Next Activity Summary
	ActivityExceptionDecoration       any    `json:"activity_exception_decoration"`         // Activity Exception Decoration
	ActivityExceptionIcon             string `json:"activity_exception_icon"`               // Icon
	ActivityCalendarEventID           any    `json:"activity_calendar_event_id"`            // Next Activity Calendar Event 📦 Model Relation: calendar.event
	MessageIsFollower                 bool   `json:"message_is_follower"`                   // Is Follower
	MessageFollowerIDs                any    `json:"message_follower_ids"`                  // Followers 📦 Model Relation: mail.followers
	MessagePartnerIDs                 any    `json:"message_partner_ids"`                   // Followers (Partners) 📦 Model Relation: res.partner
	MessageIDs                        any    `json:"message_ids"`                           // Messages 📦 Model Relation: mail.message
	HasMessage                        bool   `json:"has_message"`                           // Has Message
	MessageUnread                     bool   `json:"message_unread"`                        // Unread Messages
	MessageUnreadCounter              int    `json:"message_unread_counter"`                // Unread Messages Counter
	MessageNeedaction                 bool   `json:"message_needaction"`                    // Action Needed
	MessageNeedactionCounter          int    `json:"message_needaction_counter"`            // Number of Actions
	MessageHasError                   bool   `json:"message_has_error"`                     // Message Delivery error
	MessageHasErrorCounter            int    `json:"message_has_error_counter"`             // Number of errors
	MessageAttachmentCount            int    `json:"message_attachment_count"`              // Attachment Count
	MessageMainAttachmentID           any    `json:"message_main_attachment_id"`            // Main Attachment 📦 Model Relation: ir.attachment
	WebsiteMessageIDs                 any    `json:"website_message_ids"`                   // Website Messages 📦 Model Relation: mail.message
	MessageHasSmsError                bool   `json:"message_has_sms_error"`                 // SMS Delivery error
	Name                              string `json:"name"`                                  // Employee Name
	Active                            bool   `json:"active"`                                // Active
	Color                             int    `json:"color"`                                 // Color Index
	DepartmentID                      any    `json:"department_id"`                         // Department 📦 Model Relation: hr.department
	JobID                             any    `json:"job_id"`                                // Job Position 📦 Model Relation: hr.job
	JobTitle                          string `json:"job_title"`                             // Job Title
	AddressID                         any    `json:"address_id"`                            // Work Address 📦 Model Relation: res.partner
	WorkPhone                         string `json:"work_phone"`                            // Work Phone
	MobilePhone                       string `json:"mobile_phone"`                          // Work Mobile
	WorkEmail                         string `json:"work_email"`                            // Work Email
	WorkLocationID                    any    `json:"work_location_id"`                      // Work Location 📦 Model Relation: hr.work.location
	UserID                            any    `json:"user_id"`                               // User 📦 Model Relation: res.users
	ParentID                          any    `json:"parent_id"`                             // Manager 📦 Model Relation: hr.employee
	CoachID                           any    `json:"coach_id"`                              // Coach 📦 Model Relation: hr.employee
	HrPresenceState                   any    `json:"hr_presence_state"`                     // Hr Presence State
	LastActivity                      any    `json:"last_activity"`                         // Last Activity
	LastActivityTime                  string `json:"last_activity_time"`                    // Last Activity Time
	HrIconDisplay                     any    `json:"hr_icon_display"`                       // Hr Icon Display
	EmployeeType                      any    `json:"employee_type"`                         // Employee Type ⭐ Required
	ParentUserID                      any    `json:"parent_user_id"`                        // Parent User 📦 Model Relation: res.users
	LastAppraisalID                   any    `json:"last_appraisal_id"`                     // Last Appraisal 📦 Model Relation: hr.appraisal
	GoalIDs                           any    `json:"goal_ids"`                              // Employee HR Goals 📦 Model Relation: gamification.goal
	BadgeIDs                          any    `json:"badge_ids"`                             // Employee Badges 📦 Model Relation: gamification.badge.user
	HasBadges                         bool   `json:"has_badges"`                            // Has Badges
	DirectBadgeIDs                    any    `json:"direct_badge_ids"`                      // Direct Badge 📦 Model Relation: gamification.badge.user
	LeaveManagerID                    any    `json:"leave_manager_id"`                      // Time Off 📦 Model Relation: res.users
	RemainingLeaves                   any    `json:"remaining_leaves"`                      // Remaining Paid Time Off
	CurrentLeaveState                 any    `json:"current_leave_state"`                   // Current Time Off Status
	LeaveDateFrom                     any    `json:"leave_date_from"`                       // From Date
	LeaveDateTo                       any    `json:"leave_date_to"`                         // To Date
	LeavesCount                       any    `json:"leaves_count"`                          // Number of Time Off
	AllocationCount                   any    `json:"allocation_count"`                      // Total number of days allocated.
	AllocationsCount                  int    `json:"allocations_count"`                     // Total number of allocations
	AllocationUsedCount               any    `json:"allocation_used_count"`                 // Total number of days off used
	ShowLeaves                        bool   `json:"show_leaves"`                           // Able to see Remaining Time Off
	IsAbsent                          bool   `json:"is_absent"`                             // Absent Today
	AllocationDisplay                 string `json:"allocation_display"`                    // Allocation Display
	AllocationUsedDisplay             string `json:"allocation_used_display"`               // Allocation Used Display
	ChildAllCount                     int    `json:"child_all_count"`                       // Indirect Subordinates Count
	UserPartnerID                     any    `json:"user_partner_id"`                       // User's partner 📦 Model Relation: res.partner
	CompanyCountryID                  any    `json:"company_country_id"`                    // Company Country 📦 Model Relation: res.country
	CompanyCountryCode                string `json:"company_country_code"`                  // Country Code
	AddressHomeID                     any    `json:"address_home_id"`                       // Address 📦 Model Relation: res.partner
	IsAddressHomeACompany             bool   `json:"is_address_home_a_company"`             // The employee address has a company linked
	PrivateEmail                      string `json:"private_email"`                         // Private Email
	Lang                              any    `json:"lang"`                                  // Lang
	CountryID                         any    `json:"country_id"`                            // Nationality (Country) 📦 Model Relation: res.country
	Gender                            any    `json:"gender"`                                // Gender
	Marital                           any    `json:"marital"`                               // Marital Status
	SpouseCompleteName                string `json:"spouse_complete_name"`                  // Spouse Complete Name
	SpouseBirthdate                   any    `json:"spouse_birthdate"`                      // Spouse Birthdate
	Children                          int    `json:"children"`                              // Number of Children
	PlaceOfBirth                      string `json:"place_of_birth"`                        // Place of Birth
	CountryOfBirth                    any    `json:"country_of_birth"`                      // Country of Birth 📦 Model Relation: res.country
	Birthday                          any    `json:"birthday"`                              // Date of Birth
	Ssnid                             string `json:"ssnid"`                                 // SSN No
	Sinid                             string `json:"sinid"`                                 // SIN No
	IdentificationID                  string `json:"identification_id"`                     // Identification No
	PassportID                        string `json:"passport_id"`                           // Passport No
	BankAccountID                     any    `json:"bank_account_id"`                       // Bank Account Number 📦 Model Relation: res.partner.bank
	PermitNo                          string `json:"permit_no"`                             // Work Permit No
	VisaNo                            string `json:"visa_no"`                               // Visa No
	VisaExpire                        any    `json:"visa_expire"`                           // Visa Expiration Date
	WorkPermitExpirationDate          any    `json:"work_permit_expiration_date"`           // Work Permit Expiration Date
	HasWorkPermit                     any    `json:"has_work_permit"`                       // Work Permit
	WorkPermitScheduledActivity       bool   `json:"work_permit_scheduled_activity"`        // Work Permit Scheduled Activity
	AdditionalNote                    string `json:"additional_note"`                       // Additional Note
	Certificate                       any    `json:"certificate"`                           // Certificate Level
	StudyField                        string `json:"study_field"`                           // Field of Study
	StudySchool                       string `json:"study_school"`                          // School
	EmergencyContact                  string `json:"emergency_contact"`                     // Emergency Contact
	EmergencyPhone                    string `json:"emergency_phone"`                       // Emergency Phone
	KmHomeWork                        int    `json:"km_home_work"`                          // Home-Work Distance
	Phone                             string `json:"phone"`                                 // Private Phone
	ChildIDs                          any    `json:"child_ids"`                             // Direct subordinates 📦 Model Relation: hr.employee
	CategoryIDs                       any    `json:"category_ids"`                          // Tags 📦 Model Relation: hr.employee.category
	Notes                             string `json:"notes"`                                 // Notes
	Barcode                           string `json:"barcode"`                               // Badge ID
	Pin                               string `json:"pin"`                                   // PIN
	DepartureReasonID                 any    `json:"departure_reason_id"`                   // Departure Reason 📦 Model Relation: hr.departure.reason
	DepartureDescription              string `json:"departure_description"`                 // Additional Information
	DepartureDate                     any    `json:"departure_date"`                        // Departure Date
	IdCard                            any    `json:"id_card"`                               // ID Card Copy
	DrivingLicense                    any    `json:"driving_license"`                       // Driving License
	ID                                int    `json:"id"`                                    // ID
	LastUpdate                        any    `json:"__last_update"`                         // Last Modified on
	DisplayName                       string `json:"display_name"`                          // Display Name
	CreateUid                         any    `json:"create_uid"`                            // Created by 📦 Model Relation: res.users
	CreateDate                        any    `json:"create_date"`                           // Created on
	WriteUid                          any    `json:"write_uid"`                             // Last Updated by 📦 Model Relation: res.users
	WriteDate                         any    `json:"write_date"`                            // Last Updated on
	NextAppraisalDate                 any    `json:"next_appraisal_date"`                   // Next Appraisal Date
	LastAppraisalDate                 any    `json:"last_appraisal_date"`                   // Last Appraisal Date
	RelatedPartnerID                  any    `json:"related_partner_id"`                    // Related Partner 📦 Model Relation: res.partner
	AppraisalCount                    int    `json:"appraisal_count"`                       // Appraisal Count
	UncompleteGoalsCount              int    `json:"uncomplete_goals_count"`                // Uncomplete Goals Count
	AppraisalChildIDs                 any    `json:"appraisal_child_ids"`                   // Appraisal Child 📦 Model Relation: hr.employee
	AppraisalIDs                      any    `json:"appraisal_ids"`                         // Appraisal 📦 Model Relation: hr.appraisal
	AttendanceIDs                     any    `json:"attendance_ids"`                        // Attendance 📦 Model Relation: hr.attendance
	LastAttendanceID                  any    `json:"last_attendance_id"`                    // Last Attendance 📦 Model Relation: hr.attendance
	LastCheckIn                       any    `json:"last_check_in"`                         // Check In
	LastCheckOut                      any    `json:"last_check_out"`                        // Check Out
	AttendanceState                   any    `json:"attendance_state"`                      // Attendance Status
	HoursLastMonth                    any    `json:"hours_last_month"`                      // Hours Last Month
	HoursToday                        any    `json:"hours_today"`                           // Hours Today
	HoursLastMonthDisplay             string `json:"hours_last_month_display"`              // Hours Last Month Display
	OvertimeIDs                       any    `json:"overtime_ids"`                          // Overtime 📦 Model Relation: hr.attendance.overtime
	TotalOvertime                     any    `json:"total_overtime"`                        // Total Overtime
	Vehicle                           string `json:"vehicle"`                               // Company Vehicle
	ContractIDs                       any    `json:"contract_ids"`                          // Employee Contracts 📦 Model Relation: hr.contract
	ContractID                        any    `json:"contract_id"`                           // Current Contract 📦 Model Relation: hr.contract
	CalendarMismatch                  bool   `json:"calendar_mismatch"`                     // Calendar Mismatch
	ContractsCount                    int    `json:"contracts_count"`                       // Contract Count
	ContractWarning                   bool   `json:"contract_warning"`                      // Contract Warning
	FirstContractDate                 any    `json:"first_contract_date"`                   // First Contract Date
	EmployeeCarsCount                 int    `json:"employee_cars_count"`                   // Cars
	CarIDs                            any    `json:"car_ids"`                               // Vehicles (private) 📦 Model Relation: fleet.vehicle
	MobilityCard                      string `json:"mobility_card"`                         // Mobility Card
	CurrentLeaveID                    any    `json:"current_leave_id"`                      // Current Time Off Type 📦 Model Relation: hr.leave.type
	EquipmentIDs                      any    `json:"equipment_ids"`                         // Equipment 📦 Model Relation: maintenance.equipment
	EquipmentCount                    int    `json:"equipment_count"`                       // Equipments
	SubordinateIDs                    any    `json:"subordinate_ids"`                       // Subordinates 📦 Model Relation: hr.employee
	ResumeLineIDs                     any    `json:"resume_line_ids"`                       // Resumé lines 📦 Model Relation: hr.resume.line
	EmployeeSkillIDs                  any    `json:"employee_skill_ids"`                    // Skills 📦 Model Relation: hr.employee.skill
	SignRequestCount                  int    `json:"sign_request_count"`                    // Sign Request Count
	NewlyHiredEmployee                bool   `json:"newly_hired_employee"`                  // Newly hired employee
	ApplicantID                       any    `json:"applicant_id"`                          // Applicant 📦 Model Relation: hr.applicant
	DefaultPlanningRoleID             any    `json:"default_planning_role_id"`              // Default Planning Role 📦 Model Relation: planning.role
	PlanningRoleIDs                   any    `json:"planning_role_ids"`                     // Planning Roles 📦 Model Relation: planning.role
	EmployeeToken                     string `json:"employee_token"`                        // Security Token
	DocumentCount                     int    `json:"document_count"`                        // Document Count
	ExpenseManagerID                  any    `json:"expense_manager_id"`                    // Expense 📦 Model Relation: res.users
	SlipIDs                           any    `json:"slip_ids"`                              // Payslips 📦 Model Relation: hr.payslip
	PayslipCount                      int    `json:"payslip_count"`                         // Payslip Count
	RegistrationNumber                string `json:"registration_number"`                   // Registration Number of the Employee
	SalaryAttachmentIDs               any    `json:"salary_attachment_ids"`                 // Salary Attachments 📦 Model Relation: hr.salary.attachment
	SalaryAttachmentCount             int    `json:"salary_attachment_count"`               // Salary Attachment Count
	MobileInvoice                     any    `json:"mobile_invoice"`                        // Mobile Subscription Invoice
	SimCard                           any    `json:"sim_card"`                              // SIM Card Copy
	InternetInvoice                   any    `json:"internet_invoice"`                      // Internet Subscription Invoice
	TimesheetCost                     any    `json:"timesheet_cost"`                        // Timesheet Cost
	CurrencyID                        any    `json:"currency_id"`                           // Currency 📦 Model Relation: res.currency
	TimesheetManagerID                any    `json:"timesheet_manager_id"`                  // Timesheet 📦 Model Relation: res.users
	Niss                              string `json:"niss"`                                  // NISS Number
	SpouseFiscalStatus                any    `json:"spouse_fiscal_status"`                  // Tax status for spouse
	SpouseFiscalStatusExplanation     string `json:"spouse_fiscal_status_explanation"`      // Spouse Fiscal Status Explanation
	Disabled                          bool   `json:"disabled"`                              // Disabled
	DisabledSpouseBool                bool   `json:"disabled_spouse_bool"`                  // Disabled Spouse
	DisabledChildrenBool              bool   `json:"disabled_children_bool"`                // Disabled Children
	ResidentBool                      bool   `json:"resident_bool"`                         // Nonresident
	DisabledChildrenNumber            int    `json:"disabled_children_number"`              // Number of disabled children
	DependentChildren                 int    `json:"dependent_children"`                    // Considered number of dependent children
	L10NBeDependentChildrenAttachment int    `json:"l10n_be_dependent_children_attachment"` // # dependent children for salary attachement
	OtherDependentPeople              bool   `json:"other_dependent_people"`                // Other Dependent People
	OtherSeniorDependent              int    `json:"other_senior_dependent"`                // # seniors (>=65)
	OtherDisabledSeniorDependent      int    `json:"other_disabled_senior_dependent"`       // # disabled seniors (>=65)
	OtherJuniorsDependent             int    `json:"other_juniors_dependent"`               // # people (<65)
	OtherDisabledJuniorsDependent     int    `json:"other_disabled_juniors_dependent"`      // # disabled people (<65)
	DependentSeniors                  int    `json:"dependent_seniors"`                     // Considered number of dependent seniors
	DependentJuniors                  int    `json:"dependent_juniors"`                     // Considered number of dependent juniors
	StartNoticePeriod                 any    `json:"start_notice_period"`                   // Start notice period
	EndNoticePeriod                   any    `json:"end_notice_period"`                     // End notice period
	FirstContractInCompany            any    `json:"first_contract_in_company"`             // First contract in company
	HasBicycle                        bool   `json:"has_bicycle"`                           // Bicycle to work
	L10NBeScaleSeniority              int    `json:"l10n_be_scale_seniority"`               // Seniority at Hiring
	DoublePayLineIDs                  any    `json:"double_pay_line_ids"`                   // Previous Occupations 📦 Model Relation: l10n.be.double.pay.recovery.line
	L10NBeHolidayPayToRecoverN        any    `json:"l10n_be_holiday_pay_to_recover_n"`      // Simple Holiday Pay to Recover (N)
	L10NBeHolidayPayNumberOfDaysN     any    `json:"l10n_be_holiday_pay_number_of_days_n"`  // Number of days to recover (N)
	L10NBeHolidayPayRecoveredN        any    `json:"l10n_be_holiday_pay_recovered_n"`       // Recovered Simple Holiday Pay (N)
	L10NBeHolidayPayToRecoverN1       any    `json:"l10n_be_holiday_pay_to_recover_n1"`     // Simple Holiday Pay to Recover (N-1)
	L10NBeHolidayPayNumberOfDaysN1    any    `json:"l10n_be_holiday_pay_number_of_days_n1"` // Number of days to recover (N-1)
	L10NBeHolidayPayRecoveredN1       any    `json:"l10n_be_holiday_pay_recovered_n1"`      // Recovered Simple Holiday Pay (N-1)
}

func (m *Model) HREmployee() {
	model := "hr.employee"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(HrEmployee_150{})

	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"name", "!=", "Administrator"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(records) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}
	fmt.Println(model, "records found:", len(records))

	// bar := progressbar.Default(int64(len(records)))

	for _, r := range records {
		// fmt.Println(prettyprint(r))
		var employee HrEmployee_150
		FillStruct(r, &employee)
		// fmt.Println(prettyprint(employee))
		// fmt.Println()
		// fmt.Println(employee.Name)

		ur := map[string]any{
			"name":       employee.Name,
			"email":      employee.WorkEmail,
			"work_phone": employee.WorkPhone,
			"job_title":  employee.JobTitle,
			// "work_location_id":      employee.WorkLocationID,
			// "private_phone":         employee.Phone,
			// "emergency_contact":     employee.EmergencyContact,
			// "emergency_phone":       employee.EmergencyPhone,
			// "country_id":            employee.CountryID,
			// "identification_id":     employee.IdentificationID,
			// "ssnid":                 employee.IdentificationID,
			// "passport_id":           employee.PassportID,
			// "marital":               employee.Marital,
			// "children":              employee.Children,
			// "legal_name":            employee.Name,
			// "birthday":              employee.Birthday,
			// "place_of_birth":        employee.PlaceOfBirth,
			// "country_of_birth":      employee.CountryOfBirth,
			// "gender":                employee.Gender,
			// "visa_no":               employee.VisaNo,
			// "permit_no":             employee.PermitNo,
			// "private_street":        "",
			// "private_street2":       "",
			// "private_city":          "",
			// "private_zip":           "",
			// "distance_home_work":    employee.KmHomeWork,
			// "certificate":           employee.Certificate,
			// "study_field":           employee.StudyField,
			"pin": employee.Pin,
			// "barcode": employee.Barcode,
			"avatar_1024": employee.Avatar1024,
			"avatar_128":  employee.Avatar128,
			"avatar_1920": employee.Avatar1920,
			"avatar_256":  employee.Avatar256,
			"avatar_512":  employee.Avatar512,
		}
		_, dep_name := ParseMany2One(employee.DepartmentID)
		departmentID, _ := m.Dest.GetID("hr.department", []any{
			[]any{"name", "=", dep_name},
		})
		if departmentID != -1 {
			ur["department_id"] = departmentID
		}

		_, parent_name := ParseMany2One(employee.ParentID)
		parentID, _ := m.Dest.GetID("hr.employee", []any{
			[]any{"name", "=", parent_name},
		})
		if parentID != -1 {
			ur["parent_id"] = parentID
		}

		// _, expense_manager_name := ParseMany2One(employee.ExpenseManagerID)
		// expense_managerID, _ := m.Dest.GetID("hr.employee", []any{
		// 	[]any{"name", "=", expense_manager_name},
		// })
		// if expense_managerID != -1 {
		// 	ur["expense_manager_id"] = expense_managerID
		// }

		// _, leave_manager_name := ParseMany2One(employee.LeaveManagerID)
		// leave_managerID, _ := m.Dest.GetID("hr.employee", []any{
		// 	[]any{"name", "=", leave_manager_name},
		// })
		// if leave_managerID != -1 {
		// 	ur["leave_manager_id"] = leave_managerID
		// }

		// _, attendance_manager_name := ParseMany2One(employee.TimesheetManagerID)
		// attendance_managerID, _ := m.Dest.GetID("hr.employee", []any{
		// 	[]any{"name", "=", attendance_manager_name},
		// })
		// if attendance_managerID != -1 {
		// 	ur["attendance_manager_id"] = attendance_managerID
		// }

		_, job_name := ParseMany2One(employee.JobID)
		jobID, _ := m.Dest.GetID("hr.job", []any{
			[]any{"name", "=", job_name},
		})
		if jobID != -1 {
			ur["job_id"] = jobID
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", employee.Name},
		})

		// fmt.Println("rid", rid, "ur", prettyprint(ur))
		m.writeRecord(model, ur, rid)

		// 	bar.Add(1)
	}
	// bar.Finish()
}
