package models

import "github.com/schollz/progressbar/v3"

// quest15data uom.category model
type UomCategory_150 struct {
	DisplayName    string `json:"display_name"`     // Display Name
	ID             int    `json:"id"`               // ID
	Name           string `json:"name"`             // Unit of Measure Category ⭐ required
	ReferenceUomID any    `json:"reference_uom_id"` // Reference UoM 📦 relation: many2one uom.uom
	UomIDs         any    `json:"uom_ids"`          // Uom 📦 relation: one2many uom.uom
}

// quest15data uom.uom model
type UomUom_150 struct {
	Active          bool    `json:"active"`           // Active
	CategoryID      any     `json:"category_id"`      // Category 📦 relation: many2one uom.category ⭐ required
	Color           int     `json:"color"`            // Color
	DisplayName     string  `json:"display_name"`     // Display Name
	Factor          float64 `json:"factor"`           // Ratio ⭐ required
	FactorInv       float64 `json:"factor_inv"`       // Bigger Ratio ⭐ required
	ID              int     `json:"id"`               // ID
	Name            string  `json:"name"`             // Unit of Measure ⭐ required
	Ratio           float64 `json:"ratio"`            // Combined Ratio
	Rounding        float64 `json:"rounding"`         // Rounding Precision ⭐ required
	TimesheetWidget string  `json:"timesheet_widget"` // Widget
	UomType         any     `json:"uom_type"`         // Type ⭐ required
}

// quest19data uom.uom model
type UomUom_190 struct {
	Active             bool    `json:"active"`               // Active
	DisplayName        string  `json:"display_name"`         // Display Name
	Factor             float64 `json:"factor"`               // Absolute Quantity
	FiscalCountryCodes string  `json:"fiscal_country_codes"` // Fiscal Country Codes
	ID                 int     `json:"id"`                   // ID
	Name               string  `json:"name"`                 // Unit Name ⭐ required
	PackageTypeID      any     `json:"package_type_id"`      // Package Type 📦 relation: many2one stock.package.type
	ParentPath         string  `json:"parent_path"`          // Parent Path
	ProductUomIDs      any     `json:"product_uom_ids"`      // Barcodes 📦 relation: one2many product.uom
	RelatedUomIDs      any     `json:"related_uom_ids"`      // Related UoMs 📦 relation: one2many uom.uom
	RelativeFactor     float64 `json:"relative_factor"`      // Contains ⭐ required
	RelativeUomID      any     `json:"relative_uom_id"`      // Reference Unit 📦 relation: many2one uom.uom
	Rounding           float64 `json:"rounding"`             // Rounding Precision
	RouteIDs           any     `json:"route_ids"`            // Routes 📦 relation: many2many stock.route
	Sequence           int     `json:"sequence"`             // Sequence
	TimesheetWidget    string  `json:"timesheet_widget"`     // Widget
}

// "fiscal_country_codes": "CA",
// "package_type_id": false,
// "parent_path": "14/15/16/",

func (m *Model) Uomuom() {
	model := "uom.uom"
	uomcat_model := "uom.category"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	//
	// "Unit"
	//
	uom_cat_unit, _ := m.Source.GetID(uomcat_model, []any{
		[]any{"name", "=", "Unit"},
	})

	uoms15_units, _ := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(UomUom_150{}), []any{
		[]any{"category_id", "=", uom_cat_unit},
	})

	// RENAME Units to EA (id 1)
	m.writeRecord(model, map[string]any{"name": "EA"}, 1)
	uom19_unit_ref, _ := m.Dest.GetID(model, []any{
		[]any{"name", "=", "EA"},
	})
	bar := progressbar.Default(-1, "unit")
	for _, r := range uoms15_units {
		var uom15 UomUom_150
		FillStruct(r, &uom15)
		// fmt.Println(prettyprint(uom15))

		ur := map[string]any{
			"name":            uom15.Name,
			"active":          uom15.Active,
			"relative_factor": uom15.Ratio,
			"rounding":        uom15.Rounding,
		}
		if uom15.UomType != "reference" {
			ur["relative_uom_id"] = uom19_unit_ref
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", uom15.Name},
		})
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	//
	// "Weight"
	//

	weightConversions := []UomUom_190{
		{Name: "oz(s)", Factor: 28.3495, Rounding: 0.001},
		{Name: "LB", Factor: 453.592, Rounding: 0.001},
		{Name: "5LB", Factor: 2267.96, Rounding: 0.001},
		{Name: "10LB", Factor: 4535.92, Rounding: 0.001},
	}

	uom19_weight_ref, _ := m.Dest.GetID(model, []any{
		[]any{"name", "=", "g"},
	})

	bar = progressbar.Default(-1, "weight")
	for _, uom := range weightConversions {
		ur := map[string]any{
			"name":            uom.Name,
			"relative_factor": uom.Factor,
			"rounding":        uom.Rounding,
			"relative_uom_id": uom19_weight_ref,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", uom.Name},
		})
		// fmt.Println("rid", rid, "ur", ur)
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	//
	// "Working Time"
	//

	//
	// "Length / Distance"
	//

	lengthConversions := []UomUom_190{
		{Name: "cm", Factor: 10.0, Rounding: 0.01},
		{Name: "in", Factor: 25.4, Rounding: 0.01},
		{Name: "ft", Factor: 304.8, Rounding: 0.01},
		{Name: "yd", Factor: 914.4, Rounding: 0.01},
	}

	uom19_length_ref, _ := m.Dest.GetID(model, []any{
		[]any{"name", "=", "mm"},
	})

	bar = progressbar.Default(-1, "length")
	for _, uom := range lengthConversions {

		ur := map[string]any{
			"name":            uom.Name,
			"relative_factor": uom.Factor,
			"rounding":        uom.Rounding,
			"relative_uom_id": uom19_length_ref,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", uom.Name},
		})
		// fmt.Println("rid", rid, "ur", ur)
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	//
	// "Volume"
	//

	volumeConversions := []UomUom_190{
		{Name: "fl oz", Factor: 29.5735, Rounding: 0.01},
		{Name: "pt", Factor: 473.176, Rounding: 0.01},
		{Name: "qt", Factor: 946.352, Rounding: 0.01},
		{Name: "gal(s)", Factor: 3785.41, Rounding: 0.01},
		{Name: "in³", Factor: 16.3871, Rounding: 0.01},
		{Name: "ft³", Factor: 28316.8, Rounding: 0.01},
		{Name: "yd³", Factor: 764555, Rounding: 0.01},
		{Name: "m³", Factor: 1000000.0, Rounding: 0.01},
	}

	uom19_volume_ref, _ := m.Dest.GetID(model, []any{
		[]any{"name", "=", "ml"},
	})
	bar = progressbar.Default(-1, "volume")
	for _, uom := range volumeConversions {

		ur := map[string]any{
			"name":            uom.Name,
			"relative_factor": uom.Factor,
			"rounding":        uom.Rounding,
			"relative_uom_id": uom19_volume_ref,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", uom.Name},
		})
		// fmt.Println("rid", rid, "ur", ur)
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()

	//
	// "Area"
	//

	areaConversions := []UomUom_190{
		{Name: "m²", Factor: 10.7639, Rounding: 0.01},
		{Name: "SqIN", Factor: 144.00, Rounding: 0.01},
		{Name: "SqYD", Factor: 9, Rounding: 0.01},
		{Name: "SH 3x3", Factor: 9, Rounding: 0.01},
		{Name: "SH 3x4", Factor: 12.00000, Rounding: 0.01000},
		{Name: "SH 3x5", Factor: 15.00000, Rounding: 0.01000},
		{Name: "SHTa", Factor: 16.00000, Rounding: 0.01000},
		{Name: "SH 5x5", Factor: 25.00000, Rounding: 0.01000},
		{Name: "SH 4x8", Factor: 32.00000, Rounding: 0.01000},
		{Name: "SH 5x7", Factor: 35.00000, Rounding: 0.01000},
		{Name: "SH 4x9", Factor: 36.00000, Rounding: 0.01000},
		{Name: "SH 4x10", Factor: 40.00000, Rounding: 0.01000},
		{Name: "SH 5x10", Factor: 50.00000, Rounding: 0.01000},
	}

	// change m² to SFT
	m.writeRecord(model, map[string]any{"name": "SFT"}, 10)
	uom19_area_ref := 10
	// uom19_area_ref, _ := m.Dest.GetID(model, []any{
	// 	[]any{"name", "=", "SFT"},
	// })

	bar = progressbar.Default(-1, "area")
	for _, uom := range areaConversions {

		ur := map[string]any{
			"name":            uom.Name,
			"relative_factor": uom.Factor,
			"rounding":        uom.Rounding,
			"relative_uom_id": uom19_area_ref,
		}
		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", uom.Name},
		})
		// fmt.Println("rid", rid, "ur", ur)
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

func quickconvertuom(input string) string {
	switch input {
	case "IN":
		return "in"
	case "FT":
		return "ft"
	case "LFT":
		return "ft"
	case "YD":
		return "yd"
	case "LYD":
		return "yd"
	case "Gram":
		return "g"
	default:
		return input
	}
}
