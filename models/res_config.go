package models

type ResConfig_150 struct {
	ID          int    `json:"id"`            // ID
	LastUpdate  any    `json:"__last_update"` // Last Modified on
	DisplayName string `json:"display_name"`  // Display Name
	CreateUid   any    `json:"create_uid"`    // Created by 📦 Model Relation: res.users
	CreateDate  any    `json:"create_date"`   // Created on
	WriteUid    any    `json:"write_uid"`     // Last Updated by 📦 Model Relation: res.users
	WriteDate   any    `json:"write_date"`    // Last Updated on
}

func (m *Model) ResConfig() {
	model := "res.config"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())
}
