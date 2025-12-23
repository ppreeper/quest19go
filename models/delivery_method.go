package models

import "github.com/schollz/progressbar/v3"

type DeliveryCarrier struct {
	Name   string  `json:"name"`    // Carrier Name ⭐ Required
	CutOff float64 `json:"cut_off"` // Cut Off ⭐ Required
}

func (m *Model) DeliveryMethods() {
	model := "delivery.carrier"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	m.loadDeliveryProducts()

	bar := progressbar.Default(int64(len(deliveryCarriers)), "DeliveryMethods")
	for _, dp := range deliveryCarriers {

		productID, _ := m.Dest.GetID("product.template", []any{
			[]any{"name", "=", dp.Name},
		})

		ur := map[string]any{
			"name":       dp.Name,
			"cut_off":    dp.CutOff,
			"product_id": productID,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", dp.Name},
			[]any{"cut_off", "=", dp.CutOff},
			[]any{"product_id", "=", productID},
		})
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

func (m *Model) loadDeliveryProducts() {
	model := "product.template"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	deliveryPCID, _ := m.Dest.GetID("product.category", []any{
		[]any{"name", "=", "Deliveries"},
	})

	bar := progressbar.Default(int64(len(deliveryCarriers)), "deliveryProducts")
	for _, dp := range deliveryCarriers {
		ur := map[string]any{
			"name":           dp.Name,
			"type":           "service",
			"sale_ok":        false,
			"purchase_ok":    false,
			"categ_id":       deliveryPCID,
			"list_price":     0.0,
			"standard_price": 0.0,
		}

		rid, _ := m.Dest.GetID(model, []any{
			[]any{"name", "=", dp.Name},
			[]any{"categ_id", "=", deliveryPCID},
		})
		m.writeRecord(model, ur, rid)
		bar.Add(1)
	}
	bar.Finish()
}

var deliveryCarriers = []DeliveryCarrier{
	{"360", 15},
	{"5 STAR COURIER", 15},
	{"ABF FREIGHT", 15},
	{"ACE COURIER", 13},
	{"ACTION", 14},
	{"AIR CANADA", 15},
	{"AIR CARGO", 15},
	{"AIR NORTH", 15},
	{"AIRDRIE COURIER", 15},
	{"APPLE EXPRESS", 13},
	{"ATRIPCO DELIVERY SERVICES", 12},
	{"B&R ECKEL'S", 15},
	{"BADGER", 15},
	{"BAILEY'S HOT SHOT", 15},
	{"BANDSTRA TRANSPORT", 15},
	{"BARREL TAXI", 15},
	{"BCM", 15},
	{"BERT BAXTER", 15},
	{"BESTWAY", 15},
	{"BLACKGOLD HOTSHOT", 15},
	{"BUFFALO AIR", 15},
	{"CANADA POST", 15},
	{"CANADIAN FREIGHTWAYS", 14},
	{"CAPITAL CITY COURIER", 15},
	{"CENTRAL CARRIERS", 17},
	{"CHARIOT", 15},
	{"COLD SHOT", 14},
	{"COLE INTERNATIONAL INC.", 15},
	{"COMCO DISPATCH", 15},
	{"CRADLE TRANSPORT", 13},
	{"CUSTOMER PICK UP", 15},
	{"DAVIS & DUNN TRUCKING", 15},
	{"DAY & ROSS", 14},
	{"DEFENDER TRUCKING", 14},
	{"DHL", 14.5},
	{"DMX EXPRESS", 15},
	{"DRIVER DIRECT", 10},
	{"DUNRITE", 13},
	{"FAST LANE TRANSPORT", 15},
	{"FEDEX", 13},
	{"FLEX", 13},
	{"G & D LIGHT HAUL", 15},
	{"GERTZ DELIVERIES", 15},
	{"GHOST TRANSPORT", 14},
	{"GOING PLACES COURIER", 9},
	{"GOOD-2-GO", 15},
	{"GRIMSHAW", 15},
	{"HCL", 15},
	{"HERCULES FREIGHT", 15},
	{"HIGH COUNTRY", 15},
	{"HIWAY 9", 18},
	{"IMPACT FREIGHT", 15.5},
	{"JJ TRANSPORT", 0},
	{"LA CRETE TRANSPORT", 17},
	{"LAC LA BICHE TRANSPORT", 15},
	{"LCS", 15},
	{"LG COURIER", 15},
	{"LLB TRANSPORT", 15},
	{"LOOMIS", 14.5},
	{"MANITOULIN", 17},
	{"MEL MARTINS", 15},
	{"MLT", 18},
	{"MONARCH", 15},
	{"MSCG CARGO", 15},
	{"MSCJ VENTURES", 10},
	{"MURRAYS", 13},
	{"MUSTANG", 17},
	{"NISKU DISPATCH", 15},
	{"NOMAD", 15},
	{"OCEAN FREIGHT", 15},
	{"ONE TRANSPORT", 15},
	{"ONTIME EXPRESS", 15},
	{"OVERLAND WEST", 13.5},
	{"OVERNIGHT EXPEDITE", 11},
	{"POWER EXPRESS", 14},
	{"PRIVATEERS", 13},
	{"PUNISHER TRANSPORT", 10},
	{"PUROLATOR", 15.5},
	{"QUEST TRUCK", 15},
	{"RCT DELIVERY", 14},
	{"REDCO", 15},
	{"ROCKET", 17},
	{"ROO RUNNER", 15},
	{"ROSENAU", 17},
	{"ROYAL DELIVERY", 15},
	{"RUNNERS EXPRESS", 15},
	{"RXO CONNECT", 15},
	{"SEE NOTES", 15},
	{"SELECT COURIER", 14},
	{"SLICK TRANSPORT", 17},
	{"SOUTHVIEW TRUCKING", 15},
	{"STAMPEDE MESSENGER", 14},
	{"STREAMLINE", 8.5},
	{"STS", 17},
	{"SUREFIRE EXPRESS", 15},
	{"TARGET TRANSPORTATION", 14},
	{"TF TRUCKLOAD", 15},
	{"TRUCK ALL", 14},
	{"TST-CF EXPRESS", 15},
	{"TUBOSCOPE", 15},
	{"UPS", 14},
	{"VANKAM", 8.5},
	{"VITRAN TRANSPORT", 13},
	{"WEST DIRECT", 13},
	{"WESTERN DISPATCH", 17},
	{"WHITECOURT TRANSPORT", 15},
	{"WILLY'S TRUCKING", 17.5},
	{"WOLF TRANSPORT", 10},
	{"WTX", 13},
}
