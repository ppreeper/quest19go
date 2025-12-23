package models

import "github.com/schollz/progressbar/v3"

func (m *Model) FreightCharges() {
	model := "sale.freight"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	bar := progressbar.Default(int64(len(saleFrieghts)), "FreightCharges")
	for _, fc := range saleFrieghts {
		ur := map[string]any{
			"name": fc,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", fc},
		})
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

var saleFrieghts = []string{
	"***3RD PARTY BILLING***",
	"CALL FOR PICK UP",
	"CALL FOR PICK UP RUSH",
	"COLLECT",
	"COLLECT - DIRECT RUSH",
	"COLLECT - EXW CHANGZHOU",
	"COLLECT - EXW UK",
	"COLLECT - EXWORKS CHENNAI",
	"COLLECT - FOB MUMBAI",
	"COLLECT - FOB NANJING",
	"COLLECT - FOB NINGBO",
	"COLLECT - RUSH",
	"CUSTOMER PICK UP",
	"CUSTOMER PICK UP 10:30AM",
	"CUSTOMER PICK UP 10AM",
	"CUSTOMER PICK UP 11:30AM",
	"CUSTOMER PICK UP 12PM",
	"CUSTOMER PICK UP 1PM",
	"CUSTOMER PICK UP 2PM",
	"CUSTOMER PICK UP 3PM",
	"CUSTOMER PICK UP 8AM",
	"CUSTOMER PICK UP 9:30AM",
	"CUSTOMER PICK UP 9AM",
	"CUSTOMER PICK UP AM",
	"CUSTOMER PICK UP RUSH",
	"DELIVER",
	"DELIVERED",
	"HOLD FOR PICK UP",
	"MLT",
	"N/A",
	"PICKED UP BY CUSTOMER",
	"PLEASE SEE COMMENTS",
	"PPD",
	"PREPAID",
	"PREPAID - DIRECT RUSH",
	"PREPAID & CHARGE ***",
	"SHIP WOG",
	"TBA",
	"THIRD PARTY",
}
