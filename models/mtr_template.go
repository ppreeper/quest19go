package models

import (
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

// quest15data mtr.template model
type MtrTemplate_150 struct {
	Active                          bool      `json:"active"`                             // Active
	AluminiumRequirement            string    `json:"aluminium_requirement"`              // ALUMINIUM ( Al ) REQUIREMENT
	AluminiumValue                  string    `json:"aluminium_value"`                    // ALUMINIUM ( Al ) VALUE
	CarbonRequirement               string    `json:"carbon_requirement"`                 // CARBON ( C ) REQUIREMENT
	CarbonValue                     string    `json:"carbon_value"`                       // CARBON ( C ) VALUE
	Category                        any       `json:"category"`                           // CATEGORY ⭐ required
	CertificateNumber               string    `json:"certificate_number"`                 // CERTIFICATE NUMBER
	CharpyRequirement               string    `json:"charpy_requirement"`                 // CHARPY REQUIREMENT
	CharpyValue                     string    `json:"charpy_value"`                       // CHARPY VALUE
	CompletedNutHardnessRequirement string    `json:"completed_nut_hardness_requirement"` // COMPLETED NUT HARDNESS REQUIREMENT
	CompletedNutHardnessValue       string    `json:"completed_nut_hardness_value"`       // COMPLETED NUT HARDNESS VALUE
	CopperRequirement               string    `json:"copper_requirement"`                 // COPPER ( Cu ) REQUIREMENT
	CopperValue                     string    `json:"copper_value"`                       // COPPER ( Cu ) VALUE
	CountryOfOrigin                 string    `json:"country_of_origin"`                  // COUNTRY OF ORIGIN
	CrPercentRequirement            string    `json:"cr_percent_requirement"`             // Cr % REQUIREMENT
	CrPercentValue                  string    `json:"cr_percent_value"`                   // Cr % VALUE
	CreateDate                      time.Time `json:"create_date"`                        // Created on
	CreateUid                       any       `json:"create_uid"`                         // Created by 📦 relation: many2one res.users
	Description                     string    `json:"description"`                        // DESCRIPTION
	Diameter                        string    `json:"diameter"`                           // DIAMETER
	Dimension                       string    `json:"dimension"`                          // DIMENSION
	DimensionalInspection           string    `json:"dimensional_inspection"`             // DIMENSIONAL INSPECTION
	DisplayName                     string    `json:"display_name"`                       // Display Name
	ElongationRequirement           string    `json:"elongation_requirement"`             // ELONGATION REQUIREMENT
	ElongationValue                 string    `json:"elongation_value"`                   // ELONGATION VALUE
	HardnessRequirement             string    `json:"hardness_requirement"`               // HARDNESS REQUIREMENT
	HardnessValue                   string    `json:"hardness_value"`                     // HARDNESS VALUE
	HasMessage                      bool      `json:"has_message"`                        // Has Message
	HeatNumber                      string    `json:"heat_number"`                        // HEAT NUMBER ⭐ required
	HeatTreatmentRequirement        string    `json:"heat_treatment_requirement"`         // HEAT TREATMENT REQUIREMENT
	HeatTreatmentValue              string    `json:"heat_treatment_value"`               // HEAT TREATMENT VALUE
	ID                              int       `json:"id"`                                 // ID
	InternalCertificate             string    `json:"internal_certificate"`               // INTERNAL CERTIFICATE
	IronRequirement                 string    `json:"iron_requirement"`                   // IRON ( Fe ) REQUIREMENT
	IronValue                       string    `json:"iron_value"`                         // IRON ( Fe ) VALUE
	Length                          string    `json:"length"`                             // LENGTH
	LotBatchNumber                  string    `json:"lot_batch_number"`                   // LOT/BATCH NUMBER
	Macroetch                       string    `json:"macroetch"`                          // MACROETCH
	ManganeseRequirement            string    `json:"manganese_requirement"`              // MANGANESE ( Mn ) REQUIREMENT
	ManganeseValue                  string    `json:"manganese_value"`                    // MANGANESE ( Mn ) VALUE
	ManufacturersID                 string    `json:"manufacturers_id"`                   // MANUFACTURERS ID
	Marking                         string    `json:"marking"`                            // MARKING
	Material                        string    `json:"material"`                           // MATERIAL
	MessageAttachmentCount          int       `json:"message_attachment_count"`           // Attachment Count
	MessageFollowerIDs              any       `json:"message_follower_ids"`               // Followers 📦 relation: one2many mail.followers
	MessageHasError                 bool      `json:"message_has_error"`                  // Message Delivery error
	MessageHasErrorCounter          int       `json:"message_has_error_counter"`          // Number of errors
	MessageHasSmsError              bool      `json:"message_has_sms_error"`              // SMS Delivery error
	MessageIDs                      any       `json:"message_ids"`                        // Messages 📦 relation: one2many mail.message
	MessageIsFollower               bool      `json:"message_is_follower"`                // Is Follower
	MessageMainAttachmentID         any       `json:"message_main_attachment_id"`         // Main Attachment 📦 relation: many2one ir.attachment
	MessageNeedaction               bool      `json:"message_needaction"`                 // Action Needed
	MessageNeedactionCounter        int       `json:"message_needaction_counter"`         // Number of Actions
	MessagePartnerIDs               any       `json:"message_partner_ids"`                // Followers (Partners) 📦 relation: many2many res.partner
	MessageUnread                   bool      `json:"message_unread"`                     // Unread Messages
	MessageUnreadCounter            int       `json:"message_unread_counter"`             // Unread Messages Counter
	MoPercentRequirement            string    `json:"mo_percent_requirement"`             // Mo % REQUIREMENT
	MoPercentValue                  string    `json:"mo_percent_value"`                   // Mo % VALUE
	MtrOriginalDocumentID           any       `json:"mtr_original_document_id"`           // MTR Original Document 📦 relation: many2one documents.document
	MyActivityDateDeadline          any       `json:"my_activity_date_deadline"`          // My Activity Deadline
	NaceValue                       string    `json:"nace_value"`                         // NACE
	Name                            string    `json:"name"`                               // Name
	NickelRequirement               string    `json:"nickel_requirement"`                 // NICKEL ( Ni )% REQUIREMENT
	NickelValue                     string    `json:"nickel_value"`                       // NICKEL ( Ni )% VALUE
	NitrogenRequirement             string    `json:"nitrogen_requirement"`               // NITROGEN ( N )% REQUIREMENT
	NitrogenValue                   string    `json:"nitrogen_value"`                     // NITROGEN ( N )% VALUE
	OriginMtrID                     any       `json:"origin_mtr_id"`                      // Origin MTR 📦 relation: many2one mtr.template
	PhosphorusRequirement           string    `json:"phosphorus_requirement"`             // PHOSPHORUS ( P ) REQUIREMENT
	PhosphorusValue                 string    `json:"phosphorus_value"`                   // PHOSPHORUS ( P ) VALUE
	PoNumber                        string    `json:"po_number"`                          // PO NUMBERS
	ProductID                       any       `json:"product_id"`                         // PRODUCT 📦 relation: many2one product.product ⭐ required
	ProductType                     string    `json:"product_type"`                       // TYPE
	ProofLoadRequirement            string    `json:"proof_load_requirement"`             // PROOF LOAD REQUIREMENT
	ProofLoadValue                  string    `json:"proof_load_value"`                   // PROOF LOAD VALUE
	QuenchingRequirement            string    `json:"quenching_requirement"`              // QUENCHING REQUIREMENT
	QuenchingValue                  string    `json:"quenching_value"`                    // QUENCHING VALUE
	ReductionRequirement            string    `json:"reduction_requirement"`              // REDUCTION REQUIREMENT
	ReductionValue                  string    `json:"reduction_value"`                    // REDUCTION VALUE
	Reference                       string    `json:"reference"`                          // REFERENCE
	SampleNutHardnessRequirement    string    `json:"sample_nut_hardness_requirement"`    // SAMPLE NUT HARDNESS REQUIREMENT
	SampleNutHardnessValue          string    `json:"sample_nut_hardness_value"`          // SAMPLE NUT HARDNESS VALUE
	SiliconRequirement              string    `json:"silicon_requirement"`                // SILICON ( Si ) REQUIREMENT
	SiliconValue                    string    `json:"silicon_value"`                      // SILICON ( Si ) VALUE
	Standard                        string    `json:"standard"`                           // STANDARD
	SulfurRequirement               string    `json:"sulfur_requirement"`                 // SULFUR ( S ) REQUIREMENT
	SulfurValue                     string    `json:"sulfur_value"`                       // SULFUR ( S ) VALUE
	SupplierCertificate             string    `json:"supplier_certificate"`               // SUPPLIER CERTIFICATE
	TemperingRequirement            string    `json:"tempering_requirement"`              // TEMPERING REQUIREMENT
	TemperingValue                  string    `json:"tempering_value"`                    // TEMPERING VALUE
	TensileStrengthRequirement      string    `json:"tensile_strength_requirement"`       // TENSILE STRENGTH REQUIREMENT
	TensileStrengthValue            string    `json:"tensile_strength_value"`             // TENSILE STRENGTH VALUE
	TestDate                        string    `json:"test_date"`                          // TEST DATE
	ThreeDigitTraceabilityNo        string    `json:"three_digit_traceability_no"`        // 3 DIGIT TRACEABILITY NO.
	TitaniumRequirement             string    `json:"titanium_requirement"`               // TITANIUM ( Ti ) REQUIREMENT
	TitaniumValue                   string    `json:"titanium_value"`                     // TITANIUM ( Ti ) VALUE
	VanadiumRequirement             string    `json:"vanadium_requirement"`               // VANADIUM ( V ) REQUIREMENT
	VanadiumValue                   string    `json:"vanadium_value"`                     // VANADIUM ( V ) VALUE
	VisualInspection                string    `json:"visual_inspection"`                  // VISUAL INSPECTION
	WebsiteMessageIDs               any       `json:"website_message_ids"`                // Website Messages 📦 relation: one2many mail.message
	Year                            string    `json:"year"`                               // YEAR
	YieldStrengthRequirement        string    `json:"yield_strength_requirement"`         // YIELD STRENGTH REQUIREMENT
	YieldStrengthValue              string    `json:"yield_strength_value"`               // YIELD STRENGTH VALUE
}

// quest19data mtr.template model
type MtrTemplate_190 struct {
	Active              bool   `json:"active"`                // Active
	DisplayName         string `json:"display_name"`          // Display Name
	HasMessage          bool   `json:"has_message"`           // Has Message
	ID                  int    `json:"id"`                    // ID
	OriginalMtr         any    `json:"original_mtr"`          // Original MTR
	OriginalMtrFilename string `json:"original_mtr_filename"` // Original MTR Filename
	RatingIDs           any    `json:"rating_ids"`            // Ratings 📦 relation: one2many rating.rating
	WebsiteMessageIDs   any    `json:"website_message_ids"`   // Website Messages 📦 relation: one2many mail.message
}

func (m *Model) MtrTemplate() {
	model := "mtr.template"
	// banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	sourceFields := ExtractJSONTags(MtrTemplate_150{})

	// root
	records, err := m.Source.SearchRead(model, 0, 0, sourceFields, []any{
		[]any{"product_id.active", "=", true},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	sem := make(chan int, 8)
	var wg sync.WaitGroup
	bar := progressbar.Default(int64(len(records)), model)
	for _, r := range records {
		wg.Go(func() {
			sem <- 1
			defer func() { bar.Add(1); <-sem }()
			var record MtrTemplate_150
			FillStruct(r, &record)
			// fmt.Println(prettyprint(record))
			ur := map[string]any{
				"name":        record.Name,
				"category":    record.Category,
				"heat_number": record.HeatNumber,
			}
			pid_source, _ := ParseMany2One(record.ProductID)
			ps, _ := m.Source.SearchRead("product.product", 0, 0, []string{"id", "name"}, []any{[]any{"id", "=", pid_source}})
			if len(ps) > 0 {
				pid := -1
				if pp_name_source, ok := ps[0]["name"].(string); ok {
					pid, _ = m.Dest.GetID("product.product", []any{[]any{"name", "=", pp_name_source}})
				}
				if pid != -1 {
					ur["product_id"] = pid
				}
			}

			// bar/stud info
			//
			// column 1
			//
			// Description
			if record.Description != "" {
				ur["description"] = record.Description
			}
			// Dimension
			if record.Dimension != "" {
				ur["dimension"] = record.Dimension
			}
			// Diameter
			if record.Diameter != "" {
				ur["diameter"] = record.Diameter
			}
			// Length
			if record.Length != "" {
				ur["length"] = record.Length
			}
			// Type
			if record.ProductType != "" {
				ur["product_type"] = record.ProductType
			}
			// CountryOfOrigin
			if record.CountryOfOrigin != "" {
				ur["country_of_origin"] = record.CountryOfOrigin
			}
			// ManufacturerID
			if record.ManufacturersID != "" {
				ur["manufacturer_id"] = record.ManufacturersID
			}
			// Standard
			if record.Standard != "" {
				ur["standard"] = record.Standard
			}
			// Material
			if record.Material != "" {
				ur["material"] = record.Material
			}
			// Marking
			if record.Marking != "" {
				ur["marking"] = record.Marking
			}
			// LotBatchNumber
			if record.LotBatchNumber != "" {
				ur["lot_batch_number"] = record.LotBatchNumber
			}
			// Reference
			if record.Reference != "" {
				ur["reference"] = record.Reference
			}
			//
			// column 2
			//
			// Year
			if record.Year != "" {
				ur["year"] = record.Year
			}
			// TestDate
			if record.TestDate != "" {
				ur["test_date"] = record.TestDate
			}
			// CertificateNumber
			if record.CertificateNumber != "" {
				ur["certificate_number"] = record.CertificateNumber
			}
			// SupplierCertificate
			if record.SupplierCertificate != "" {
				ur["supplier_certificate"] = record.SupplierCertificate
			}
			// Macroetch
			if record.Macroetch != "" {
				ur["macroetch"] = record.Macroetch
			}
			// DimensionalInspection
			if record.DimensionalInspection != "" {
				ur["dimensional_inspection"] = record.DimensionalInspection
			}
			// VisualInspection
			if record.VisualInspection != "" {
				ur["visual_inspection"] = record.VisualInspection
			}
			// Nace
			if record.NaceValue != "" {
				ur["nace_value"] = record.NaceValue
			}
			// 3 Digit Traceability No.
			if record.ThreeDigitTraceabilityNo != "" {
				ur["three_digit_traceability"] = record.ThreeDigitTraceabilityNo
			}
			// PO Number
			if record.PoNumber != "" {
				ur["po_number"] = record.PoNumber
			}

			// BAR/STUD Physical Properties

			// charpy_requirement
			if record.CharpyRequirement != "" {
				ur["charpy_requirement"] = record.CharpyRequirement
			}
			// charpy_value
			if record.CharpyValue != "" {
				ur["charpy_value"] = record.CharpyValue
			}

			// completed_nut_hardness_requirement
			if record.CompletedNutHardnessRequirement != "" {
				ur["completed_nut_hardness_requirement"] = record.CompletedNutHardnessRequirement
			}
			// completed_nut_hardness_value
			if record.CompletedNutHardnessValue != "" {
				ur["completed_nut_hardness_value"] = record.CompletedNutHardnessValue
			}

			// elongation_requirement
			if record.ElongationRequirement != "" {
				ur["elongation_requirement"] = record.ElongationRequirement
			}
			// elongation_value
			if record.ElongationValue != "" {
				ur["elongation_value"] = record.ElongationValue
			}

			// hardness_requirement
			if record.HardnessRequirement != "" {
				ur["hardness_requirement"] = record.HardnessRequirement
			}
			// hardness_value
			if record.HardnessValue != "" {
				ur["hardness_value"] = record.HardnessValue
			}

			// heat_treatment_requirement
			if record.HeatTreatmentRequirement != "" {
				ur["heat_treatment_requirement"] = record.HeatTreatmentRequirement
			}
			// heat_treatment_value
			if record.HeatTreatmentValue != "" {
				ur["heat_treatment_value"] = record.HeatTreatmentValue
			}

			// proof_load_requirement
			if record.ProofLoadRequirement != "" {
				ur["proof_load_requirement"] = record.ProofLoadRequirement
			}
			// proof_load_value
			if record.ProofLoadValue != "" {
				ur["proof_load_value"] = record.ProofLoadValue
			}

			// quenching_requirement
			if record.QuenchingRequirement != "" {
				ur["quenching_requirement"] = record.QuenchingRequirement
			}
			// quenching_value
			if record.QuenchingValue != "" {
				ur["quenching_value"] = record.QuenchingValue
			}

			// reduction_requirement
			if record.ReductionRequirement != "" {
				ur["reduction_requirement"] = record.ReductionRequirement
			}
			// reduction_value
			if record.ReductionValue != "" {
				ur["reduction_value"] = record.ReductionValue
			}

			// sample_nut_hardness_requirement
			if record.SampleNutHardnessRequirement != "" {
				ur["sample_nut_hardness_requirement"] = record.SampleNutHardnessRequirement
			}
			// sample_nut_hardness_value
			if record.SampleNutHardnessValue != "" {
				ur["sample_nut_hardness_value"] = record.SampleNutHardnessValue
			}

			// tempering_requirement
			if record.TemperingRequirement != "" {
				ur["tempering_requirement"] = record.TemperingRequirement
			}
			// tempering_value
			if record.TemperingValue != "" {
				ur["tempering_value"] = record.TemperingValue
			}

			// tensile_strength_requirement
			if record.TensileStrengthRequirement != "" {
				ur["tensile_strength_requirement"] = record.TensileStrengthRequirement
			}
			// tensile_strength_value
			if record.TensileStrengthValue != "" {
				ur["tensile_strength_value"] = record.TensileStrengthValue
			}

			// yield_strength_requirement
			if record.YieldStrengthRequirement != "" {
				ur["yield_strength_requirement"] = record.YieldStrengthRequirement
			}
			// yield_strength_value
			if record.YieldStrengthValue != "" {
				ur["yield_strength_value"] = record.YieldStrengthValue
			}

			// BAR/STUD Chemical Properties
			// aluminium_requirement
			if record.AluminiumRequirement != "" {
				ur["aluminium_requirement"] = record.AluminiumRequirement
			}
			// aluminium_value
			if record.AluminiumValue != "" {
				ur["aluminium_value"] = record.AluminiumValue
			}

			// carbon_requirement
			if record.CarbonRequirement != "" {
				ur["carbon_requirement"] = record.CarbonRequirement
			}
			// carbon_value
			if record.CarbonValue != "" {
				ur["carbon_value"] = record.CarbonValue
			}

			// cr_percent_requirement
			if record.CrPercentRequirement != "" {
				ur["chromium_requirement"] = record.CrPercentRequirement
			}
			// cr_percent_value
			if record.CrPercentValue != "" {
				ur["chromium_value"] = record.CrPercentValue
			}

			// copper_requirement
			if record.CopperRequirement != "" {
				ur["copper_requirement"] = record.CopperRequirement
			}
			// copper_value
			if record.CopperValue != "" {
				ur["copper_value"] = record.CopperValue
			}

			// iron_requirement
			if record.IronRequirement != "" {
				ur["iron_requirement"] = record.IronRequirement
			}
			// iron_value
			if record.IronValue != "" {
				ur["iron_value"] = record.IronValue
			}

			// manganese_requirement
			if record.ManganeseRequirement != "" {
				ur["manganese_requirement"] = record.ManganeseRequirement
			}
			// manganese_value
			if record.ManganeseValue != "" {
				ur["manganese_value"] = record.ManganeseValue
			}

			// mo_percent_requirement
			if record.MoPercentRequirement != "" {
				ur["molybdenum_requirement"] = record.MoPercentRequirement
			}
			// mo_percent_value
			if record.MoPercentValue != "" {
				ur["molybdenum_value"] = record.MoPercentValue
			}

			// nickel_requirement
			if record.NickelRequirement != "" {
				ur["nickel_requirement"] = record.NickelRequirement
			}
			// nickel_value
			if record.NickelValue != "" {
				ur["nickel_value"] = record.NickelValue
			}

			// nitrogen_requirement
			if record.NitrogenRequirement != "" {
				ur["nitrogen_requirement"] = record.NitrogenRequirement
			}
			// nitrogen_value
			if record.NitrogenValue != "" {
				ur["nitrogen_value"] = record.NitrogenValue
			}

			// phosphorus_requirement
			if record.PhosphorusRequirement != "" {
				ur["phosphorus_requirement"] = record.PhosphorusRequirement
			}
			// phosphorus_value
			if record.PhosphorusValue != "" {
				ur["phosphorus_value"] = record.PhosphorusValue
			}

			// silicon_requirement
			if record.SiliconRequirement != "" {
				ur["silicon_requirement"] = record.SiliconRequirement
			}
			// silicon_value
			if record.SiliconValue != "" {
				ur["silicon_value"] = record.SiliconValue
			}

			// sulfur_requirement
			if record.SulfurRequirement != "" {
				ur["sulfur_requirement"] = record.SulfurRequirement
			}
			// sulfur_value
			if record.SulfurValue != "" {
				ur["sulfur_value"] = record.SulfurValue
			}

			// titanium_requirement
			if record.TitaniumRequirement != "" {
				ur["titanium_requirement"] = record.TitaniumRequirement
			}
			// titanium_value
			if record.TitaniumValue != "" {
				ur["titanium_value"] = record.TitaniumValue
			}

			// vanadium_requirement
			if record.VanadiumRequirement != "" {
				ur["vanadium_requirement"] = record.VanadiumRequirement
			}
			// vanadium_value
			if record.VanadiumValue != "" {
				ur["vanadium_value"] = record.VanadiumValue
			}

			rid, err := m.Dest.GetID(model, []any{
				[]any{"name", "=", record.Name},
				[]any{"category", "=", record.Category},
				[]any{"heat_number", "=", record.HeatNumber},
				[]any{"product_id", "=", ur["product_id"]},
			})
			if err != nil {
				m.Log.Error(model, "func", trace(), "err", err)
			}

			m.writeRecord(model, ur, rid)
			// bar.Add(1)
		})
	}
	wg.Wait()
	bar.Finish()
}
