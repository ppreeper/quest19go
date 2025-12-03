package models

func (m *Model) currencyMap() map[string]int {
	currencies_list, err := m.Dest.SearchRead("res.currency", 0, 0, ExtractJSONTags(ResCurrency_190{}), []any{})
	if err != nil {
		m.Log.Error("res.currency", "func", trace(), "err", err)
		return nil
	}
	currencyMap := make(map[string]int)
	for _, r := range currencies_list {
		var c ResCurrency_190
		FillStruct(r, &c)
		currencyMap[c.Name] = int(r["id"].(float64))
	}
	return currencyMap
}

func (m *Model) pricelistMap() map[string]int {
	pricelist_list, err := m.Dest.SearchRead("product.pricelist", 0, 0, ExtractJSONTags(ProductPricelist_190{}), []any{})
	if err != nil {
		m.Log.Error("product.pricelist", "func", trace(), "err", err)
		return nil
	}
	pricelistMap := make(map[string]int)
	for _, r := range pricelist_list {
		var p ProductPricelist_190
		FillStruct(r, &p)
		pricelistMap[p.Name] = int(r["id"].(float64))
	}
	return pricelistMap
}

func (m *Model) oldPricelistMap() map[string]string {
	pricelist_list, err := m.Source.SearchRead("product.pricelist", 0, 0, ExtractJSONTags(ProductPricelist_150{}), []any{})
	if err != nil {
		m.Log.Error("product.pricelist", "func", trace(), "err", err)
		return nil
	}
	pricelistMap := make(map[string]string)
	for _, r := range pricelist_list {
		var p ProductPricelist_150
		FillStruct(r, &p)
		pricelistMap[p.DisplayName] = r["name"].(string)
	}
	return pricelistMap
}

func (m *Model) countryMap() map[string]int {
	countries_list, err := m.Dest.SearchRead("res.country", 0, 0, ExtractJSONTags(ResCountry_190{}), []any{})
	if err != nil {
		m.Log.Error("res.country", "func", trace(), "err", err)
		return nil
	}
	countryMap := make(map[string]int)
	for _, r := range countries_list {
		var c ResCountry_190
		FillStruct(r, &c)
		countryMap[c.Name] = int(r["id"].(float64))
	}
	return countryMap
}

func (m *Model) getCountry(countryName string) int {
	countryID, err := m.Dest.GetID("res.country", []any{
		[]any{"name", "=", countryName},
	})
	if err != nil {
		m.Log.Error("res.country", "func", trace(), "err", err)
		return -1
	}
	return countryID
}

func (m *Model) getCountryState(countryName, stateName string) (int, int) {
	countryID := m.getCountry(countryName)
	if countryID == -1 {
		return -1, -1
	}
	stateID, err := m.Dest.GetID("res.country.state", []any{
		[]any{"country_id", "=", countryID},
		[]any{"name", "=", stateName},
	})
	if err != nil {
		m.Log.Error("res.country.state", "func", trace(), "err", err)
		return -1, -1
	}
	return countryID, stateID
}

func (m *Model) getFiscalPositions(countryID, stateID int) int {
	intlID, err := m.Dest.GetID("account.fiscal.position", []any{
		[]any{"country_id", "!=", countryID},
	})
	if err != nil {
		m.Log.Error("account.fiscal.position", "func", trace(), "err", err)
		return -1
	}
	fiscalID, err := m.Dest.GetID("account.fiscal.position", []any{
		[]any{"country_id", "=", countryID},
		[]any{"state_ids", "in", stateID},
	})
	if err != nil {
		m.Log.Error("account.fiscal.position", "func", trace(), "err", err)
		return -1
	}
	if fiscalID == -1 {
		return intlID
	}
	return fiscalID
}

func (m *Model) uomMAP() map[string]int {
	uom_list, err := m.Dest.SearchRead("uom.uom", 0, 0, ExtractJSONTags(UomUom_190{}), []any{})
	if err != nil {
		m.Log.Error("uom.uom", "func", trace(), "err", err)
		return nil
	}
	uomMap := make(map[string]int)
	for _, r := range uom_list {
		var u UomUom_190
		FillStruct(r, &u)
		uomMap[u.Name] = int(r["id"].(float64))
	}
	return uomMap
}

func (m *Model) productCategoryMap() map[string]int {
	type ProductCategoryTemp struct {
		ID           int    `json:"id"`
		CompleteName string `json:"complete_name"`
	}
	pc_list, err := m.Dest.SearchRead("product.category", 0, 0, ExtractJSONTags(ProductCategoryTemp{}), []any{})
	if err != nil {
		m.Log.Error("product.category", "func", trace(), "err", err)
		return nil
	}
	pcMap := make(map[string]int)
	for _, r := range pc_list {
		var u ProductCategoryTemp
		FillStruct(r, &u)
		pcMap[u.CompleteName] = u.ID
	}
	return pcMap
}
