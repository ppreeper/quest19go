package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"quest19go/app"

	"github.com/ppreeper/str"
)

type Model app.App

type Many2One struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var banner = str.NewBanner()

var (
	// INSERT data
	INSERT = true
	// UPDATE data
	UPDATE = false
)

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fnameRaw := f.Name()
	fnameSplit := strings.Split(fnameRaw, ".")
	fname := fnameSplit[len(fnameSplit)-1]
	return fname
}

func mapKeyId(values []map[string]any, key string) map[string]int {
	m := map[string]int{}
	for _, v := range values {
		switch v[key].(type) {
		case string:
			m[v[key].(string)] = int(v["id"].(float64))
		}
	}
	return m
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// FillStruct fills a struct with values from a map based on json tags
func FillStruct(data map[string]interface{}, out interface{}) error {
	val := reflect.ValueOf(out)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("out must be a pointer to a struct")
	}
	val = val.Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}
		tag = strings.Split(tag, ",")[0] // remove omitempty etc.

		if v, ok := data[tag]; ok && v != nil {
			f := val.Field(i)
			if f.CanSet() {
				// convert types safely
				vv := reflect.ValueOf(v)
				if vv.Type().ConvertibleTo(f.Type()) {
					f.Set(vv.Convert(f.Type()))
				}
			}
		}
	}
	return nil
}

// ExtractJSONTags returns a slice of json tag names from a struct
func ExtractJSONTags(i any) []string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var tags []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			// remove options like omitempty, etc.
			tag = strings.Split(tag, ",")[0]
			tags = append(tags, tag)
		}
	}
	return tags
}

func (m *Model) writeRecord(model string, data map[string]any, recordID int) {
	if recordID == -1 {
		// fmt.Println("Create", model)
		row, err := m.Dest.Create(model, data)
		if err != nil {
			m.Log.Error(model, "mode", "create", "data", data, "row", row, "err", err)
		}
	} else {
		// fmt.Println("Update", model)
		res, err := m.Dest.Write(model, recordID, data)
		if err != nil {
			m.Log.Error(model, "mode", "write", "data", data, "row", recordID, "res", res, "err", err)
		}
	}
}

func prettyprint(i any) string {
	jsonData, err := json.MarshalIndent(i, "", "  ") // "" for prefix, "  " for indent
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return ""
	}
	return string(jsonData)
}

func ParseMany2One(v any) (id int, name string) {
	arr, ok := v.([]any)
	if !ok || len(arr) < 2 {
		return 0, ""
	}
	id = int(arr[0].(float64))
	name = arr[1].(string)
	return
}

func ExtractIDs(values []map[string]any) []int {
	ids := []int{}
	for _, v := range values {
		switch v["id"].(type) {
		case float64:
			ids = append(ids, int(v["id"].(float64)))
		}
	}
	return ids
}

func odoo_path(root string, paths ...string) string {
	fullPath := root
	for _, p := range paths {
		fullPath = strings.TrimRight(fullPath, "/") + " / " + strings.TrimLeft(p, "/")
	}
	return fullPath
}

func odoo_pathjoin(root string, paths ...string) string {
	fullPath := root
	for _, p := range paths {
		fullPath = strings.TrimRight(fullPath, "/") + "/" + strings.TrimLeft(p, "/")
	}
	return fullPath
}

func odoo_barcode(root string, paths ...string) string {
	fullPath := root
	for _, p := range paths {
		fullPath = strings.TrimRight(fullPath, "/") + "." + strings.TrimLeft(p, "/")
	}
	return strings.ReplaceAll(strings.ToUpper(fullPath), " ", "")
}

func (m *Model) getCountryFromSource(countryRecord any) int {
	_, country_name := ParseMany2One(countryRecord)
	country, err := m.Dest.SearchRead("res.country", 0, 0, []string{"name"}, []any{
		[]any{"name", "=", country_name},
	})
	if err != nil {
		m.Log.Error("res.country", "func", trace(), "err", err)
		return -1
	}
	countryIDf64, ok := country[0]["id"].(float64)
	if !ok {
		m.Log.Error("res.country", "func", trace(), "err", "invalid country id")
		return -1
	}
	countryID := int(countryIDf64)
	return countryID
}

func (m *Model) getCountryStateFromSource(countryRecord, stateRecord any) int {
	state_id, _ := ParseMany2One(stateRecord)
	state_src, err := m.Source.SearchRead("res.country.state", 0, 0, []string{}, []any{
		[]any{"id", "=", state_id},
	})
	if err != nil {
		m.Log.Error("res.country.state", "func", trace(), "err", err)
		return -1
	}
	state_src_name, ok := state_src[0]["name"].(string)
	if !ok {
		m.Log.Error("res.country.state", "func", trace(), "err", "invalid source state name")
		return -1
	}
	countryID := m.getCountryFromSource(countryRecord)
	state, err := m.Dest.SearchRead("res.country.state", 0, 0, []string{}, []any{
		[]any{"country_id", "=", countryID},
		[]any{"name", "=", state_src_name},
	})
	if err != nil {
		m.Log.Error("res.country.state", "func", trace(), "err", err)
		return -1
	}
	stateIDf64, ok := state[0]["id"].(float64)
	if !ok {
		m.Log.Error("res.country", "func", trace(), "err", "invalid country id")
		return -1
	}
	stateID := int(stateIDf64)
	return stateID
}

func (m *Model) getCurrencyFromSource(currencyRecord any) (currencyID int) {
	_, currency_name := ParseMany2One(currencyRecord)
	currencyID, err := m.Dest.GetID("res.currency", []any{
		[]any{"name", "=", currency_name},
	})
	if err != nil {
		m.Log.Error("res.currency", "func", trace(), "err", err)
		return -1
	}
	return currencyID
}

func (m *Model) getPaperFormatFromSource(paperFormatRecord any) (paperFormatID int) {
	_, paper_format_name := ParseMany2One(paperFormatRecord)
	paperFormatID, err := m.Dest.GetID("report.paperformat", []any{
		[]any{"name", "=", paper_format_name},
	})
	if err != nil {
		m.Log.Error("report.paperformat", "func", trace(), "err", err)
		return -1
	}
	return paperFormatID
}
