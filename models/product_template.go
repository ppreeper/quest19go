package models

import (
	"slices"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// quest15data product.template model
type ProductTemplate_150 struct {
	// LastUpdate                             time.Time `json:"__last_update"`                              // Last Modified on
	// AccountTagIDs               any       `json:"account_tag_ids"`               // Account Tags 📦 relation: many2many account.account.tag
	// Active                      bool      `json:"active"`                        // Active
	// ActivityCalendarEventID     any       `json:"activity_calendar_event_id"`    // Next Activity Calendar Event 📦 relation: many2one calendar.event
	// ActivityDateDeadline        time.Time `json:"activity_date_deadline"`        // Next Activity Deadline
	// ActivityExceptionDecoration any       `json:"activity_exception_decoration"` // Activity Exception Decoration
	// ActivityExceptionIcon       string    `json:"activity_exception_icon"`       // Icon
	// ActivityIDs                 any       `json:"activity_ids"`                  // Activities 📦 relation: one2many mail.activity
	// ActivityState               any       `json:"activity_state"`                // Activity State
	// ActivitySummary             string    `json:"activity_summary"`              // Next Activity Summary
	// ActivityTypeIcon            string    `json:"activity_type_icon"`            // Activity Type Icon
	// ActivityTypeID              any       `json:"activity_type_id"`              // Next Activity Type 📦 relation: many2one mail.activity.type
	// ActivityUserID              any       `json:"activity_user_id"`              // Responsible User 📦 relation: many2one res.users
	// Astm string `json:"astm"` // ASTM
	// AttributeLineIDs     any     `json:"attribute_line_ids"`       // Product Attributes 📦 relation: one2many product.template.attribute.line
	Barcode string `json:"barcode"` // Barcode
	// BomCount             int     `json:"bom_count"`                // # Bill of Material
	// BomIDs               any     `json:"bom_ids"`                  // Bill of Materials 📦 relation: one2many mrp.bom
	// BomLineIDs           any     `json:"bom_line_ids"`             // BoM Components 📦 relation: one2many mrp.bom.line
	// BoxQty        float64 `json:"box_qty"`         // Box Quantity
	CanBeExpensed bool `json:"can_be_expensed"` // Can be Expensed
	// CanImage1024BeZoomed bool    `json:"can_image_1024_be_zoomed"` // Can Image 1024 be zoomed
	CategID any `json:"categ_id"` // Product Category 📦 relation: many2one product.category ⭐ required
	// Color     int `json:"color"`      // Color Index
	// CompanyID any `json:"company_id"` // Company 📦 relation: many2one res.company
	// CostCurrencyID any `json:"cost_currency_id"` // Cost Currency 📦 relation: many2one res.currency
	// CostMethod     any `json:"cost_method"`      // Costing Method
	// CreateDate                             time.Time `json:"create_date"`                                // Created on
	// CreateUid                              any       `json:"create_uid"`                                 // Created by 📦 relation: many2one res.users
	// CurrencyID            any    `json:"currency_id"`            // Currency 📦 relation: many2one res.currency
	DefaultCode string `json:"default_code"` // Internal Reference
	Description string `json:"description"`  // Description
	// DescriptionPicking    string `json:"description_picking"`    // Description on Picking
	// DescriptionPickingin  string `json:"description_pickingin"`  // Description on Receptions
	// DescriptionPickingout string `json:"description_pickingout"` // Description on Delivery Orders
	// DescriptionPurchase   string `json:"description_purchase"`   // Purchase Description
	// DescriptionSale       string `json:"description_sale"`       // Sales Description
	// DetailedType          any    `json:"detailed_type"`          // Product Type ⭐ required
	// DisplayName           string `json:"display_name"`           // Display Name
	// EcoCount                               int       `json:"eco_count"`                                  // # ECOs
	// EcoIDs                                 any       `json:"eco_ids"`                                    // ECOs 📦 relation: one2many mrp.eco
	// ExpensePolicy                          any       `json:"expense_policy"`                             // Re-Invoice Expenses
	// HasAvailableRouteIDs                   bool      `json:"has_available_route_ids"`                    // Routes can be selected on this product
	// HasConfigurableAttributes              bool      `json:"has_configurable_attributes"`                // Is a configurable product
	// HasMessage                             bool      `json:"has_message"`                                // Has Message
	HsCode string `json:"hs_code"` // HS Code
	ID     int    `json:"id"`      // ID
	// Image1024 any    `json:"image_1024"` // Image 1024
	// Image128  any    `json:"image_128"`  // Image 128
	// Image1920 any    `json:"image_1920"` // Image
	// Image256  any    `json:"image_256"`  // Image 256
	// Image512  any    `json:"image_512"`  // Image 512
	// IncomingQty                            float64   `json:"incoming_qty"`                               // Incoming
	// InvoicePolicy                          any       `json:"invoice_policy"`                             // Invoicing Policy
	// IsKits bool `json:"is_kits"` // Is Kits
	// IsProductVariant                       bool      `json:"is_product_variant"`                         // Is a product variant
	ListPrice float64 `json:"list_price"` // Sales Price
	// LocationID                             any       `json:"location_id"`                                // Location 📦 relation: many2one stock.location
	// MessageAttachmentCount                 int       `json:"message_attachment_count"`                   // Attachment Count
	// MessageFollowerIDs                     any       `json:"message_follower_ids"`                       // Followers 📦 relation: one2many mail.followers
	// MessageHasError                        bool      `json:"message_has_error"`                          // Message Delivery error
	// MessageHasErrorCounter                 int       `json:"message_has_error_counter"`                  // Number of errors
	// MessageHasSmsError                     bool      `json:"message_has_sms_error"`                      // SMS Delivery error
	// MessageIDs                             any       `json:"message_ids"`                                // Messages 📦 relation: one2many mail.message
	// MessageIsFollower                      bool      `json:"message_is_follower"`                        // Is Follower
	// MessageMainAttachmentID                any       `json:"message_main_attachment_id"`                 // Main Attachment 📦 relation: many2one ir.attachment
	// MessageNeedaction                      bool      `json:"message_needaction"`                         // Action Needed
	// MessageNeedactionCounter               int       `json:"message_needaction_counter"`                 // Number of Actions
	// MessagePartnerIDs                      any       `json:"message_partner_ids"`                        // Followers (Partners) 📦 relation: many2many res.partner
	// MessageUnread                          bool      `json:"message_unread"`                             // Unread Messages
	// MessageUnreadCounter                   int       `json:"message_unread_counter"`                     // Unread Messages Counter
	// MrpProductQty                          float64   `json:"mrp_product_qty"`                            // Manufactured
	// MyActivityDateDeadline                 time.Time `json:"my_activity_date_deadline"`                  // My Activity Deadline
	Name string `json:"name"` // Name ⭐ required
	// NbrMovesIn                             int     `json:"nbr_moves_in"`                               // Nbr Moves In
	// NbrMovesOut                            int     `json:"nbr_moves_out"`                              // Nbr Moves Out
	// NbrReorderingRules                     int     `json:"nbr_reordering_rules"`                       // Reordering Rules
	// OutgoingQty                            float64 `json:"outgoing_qty"`                               // Outgoing
	// PackagingIDs                           any     `json:"packaging_ids"`                              // Product Packages 📦 relation: one2many product.packaging
	// PlanningEnabled                        bool    `json:"planning_enabled"`                           // Plan Services
	// PlanningRoleID                         any     `json:"planning_role_id"`                           // Planning Role 📦 relation: many2one planning.role
	// Price                                  float64 `json:"price"`                                      // Price
	// PricelistID                            any     `json:"pricelist_id"`                               // Pricelist 📦 relation: many2one product.pricelist
	// PricelistItemCount                     int     `json:"pricelist_item_count"`                       // Number of price rules
	// Priority                               any     `json:"priority"`                                   // Favorite
	// ProduceDelay                           float64 `json:"produce_delay"`                              // Manufacturing Lead Time
	// ProductTooltip                         string  `json:"product_tooltip"`                            // Product Tooltip
	// ProductVariantCount                    int     `json:"product_variant_count"`                      // # Product Variants
	// ProductVariantID                       any     `json:"product_variant_id"`                         // Product 📦 relation: many2one product.product
	// ProductVariantIDs                      any     `json:"product_variant_ids"`                        // Products 📦 relation: one2many product.product ⭐ required
	// ProjectID                              any     `json:"project_id"`                                 // Project 📦 relation: many2one project.project
	// ProjectTemplateID                      any     `json:"project_template_id"`                        // Project Template 📦 relation: many2one project.project
	// PropertyAccountCreditorPriceDifference any     `json:"property_account_creditor_price_difference"` // Price Difference Account 📦 relation: many2one account.account
	// PropertyAccountExpenseID               any     `json:"property_account_expense_id"`                // Expense Account 📦 relation: many2one account.account
	// PropertyAccountIncomeID                any     `json:"property_account_income_id"`                 // Income Account 📦 relation: many2one account.account
	// PropertyStockInventory                 any     `json:"property_stock_inventory"`                   // Inventory Location 📦 relation: many2one stock.location
	// PropertyStockProduction                any     `json:"property_stock_production"`                  // Production Location 📦 relation: many2one stock.location
	// PurchaseLineWarn                       any     `json:"purchase_line_warn"`                         // Purchase Order Line Warning ⭐ required
	PurchaseLineWarnMsg string `json:"purchase_line_warn_msg"` // Message for Purchase Order Line
	// PurchaseMethod                         any     `json:"purchase_method"`                            // Control Policy
	// PurchaseOk                             bool    `json:"purchase_ok"`                                // Can be Purchased
	// PurchaseRequisition                    any     `json:"purchase_requisition"`                       // Procurement
	// PurchasedProductQty                    float64 `json:"purchased_product_qty"`                      // Purchased
	// QtyAvailable                         float64 `json:"qty_available"`                             // Quantity On Hand
	// QualityControlPointQty               int     `json:"quality_control_point_qty"`                 // Quality Control Point Qty
	// QualityFailQty                       int     `json:"quality_fail_qty"`                          // Quality Fail Qty
	// QualityPassQty                       int     `json:"quality_pass_qty"`                          // Quality Pass Qty
	// ReorderingMaxQty                     float64 `json:"reordering_max_qty"`                        // Reordering Max Qty
	// ReorderingMinQty                     float64 `json:"reordering_min_qty"`                        // Reordering Min Qty
	// ResponsibleID                        any     `json:"responsible_id"`                            // Responsible 📦 relation: many2one res.users
	// RouteFromCategIDs                    any     `json:"route_from_categ_ids"`                      // Category Routes 📦 relation: many2many stock.location.route
	RouteIDs  any     `json:"route_ids"`  // Routes 📦 relation: many2many stock.location.route
	SaleDelay float64 `json:"sale_delay"` // Customer Lead Time
	// SaleLineWarn                         any     `json:"sale_line_warn"`                            // Sales Order Line ⭐ required
	SaleLineWarnMsg string `json:"sale_line_warn_msg"` // Message for Sales Order Line
	// SaleOk                               bool    `json:"sale_ok"`                                   // Can be Sold
	// SalesCount                           float64 `json:"sales_count"`                               // Sold
	// ScheduleCount                        int     `json:"schedule_count"`                            // Schedules
	// SellerIDs                            any     `json:"seller_ids"`                                // Vendors 📦 relation: one2many product.supplierinfo
	// Sequence                             int     `json:"sequence"`                                  // Sequence
	// ServicePolicy                        any     `json:"service_policy"`                            // Service Invoicing Policy
	// ServiceToPurchase                    bool    `json:"service_to_purchase"`                       // Subcontract Service
	// ServiceTracking                      any     `json:"service_tracking"`                          // Create on Order
	// ServiceType                          any     `json:"service_type"`                              // Track Service
	// ServiceUpsellThreshold               float64 `json:"service_upsell_threshold"`                  // Threshold
	// ServiceUpsellThresholdRatio          string  `json:"service_upsell_threshold_ratio"`            // Service Upsell Threshold Ratio
	// ShowForecastedQtyStatusButton        bool    `json:"show_forecasted_qty_status_button"`         // Show Forecasted Qty Status Button
	// ShowOnHandQtyStatusButton            bool    `json:"show_on_hand_qty_status_button"`            // Show On Hand Qty Status Button
	// SizeDimension                        string  `json:"size_dimension"`                            // Size / Dimension
	StandardPrice float64 `json:"standard_price"` // Cost
	// SupplierTaxesID                      any     `json:"supplier_taxes_id"`                         // Vendor Taxes 📦 relation: many2many account.tax
	// TaxString                            string  `json:"tax_string"`                                // Tax String
	// TaxesID                              any     `json:"taxes_id"`                                  // Customer Taxes 📦 relation: many2many account.tax
	// Tracking                             any     `json:"tracking"`                                  // Tracking ⭐ required
	// Type                                 any     `json:"type"`                                      // Type
	UomID any `json:"uom_id"` // Unit of Measure 📦 relation: many2one uom.uom ⭐ required
	// UomName                              string  `json:"uom_name"`                                  // Unit of Measure Name
	// UomPoID                              any     `json:"uom_po_id"`                                 // Purchase UoM 📦 relation: many2one uom.uom ⭐ required
	// UsedInBomCount                       int     `json:"used_in_bom_count"`                         // # of BoM Where is Used
	// ValidProductTemplateAttributeLineIDs any     `json:"valid_product_template_attribute_line_ids"` // Valid Product Attribute Lines 📦 relation: many2many product.template.attribute.line
	// Valuation                            any     `json:"valuation"`                                 // Inventory Valuation
	// VariantSellerIDs                     any     `json:"variant_seller_ids"`                        // Variant Seller 📦 relation: one2many product.supplierinfo
	// Version                              int     `json:"version"`                                   // Version
	// VirtualAvailable                     float64 `json:"virtual_available"`                         // Forecasted Quantity
	// VisibleExpensePolicy                 bool    `json:"visible_expense_policy"`                    // Re-Invoice Policy visible
	// VisibleQtyConfigurator               bool    `json:"visible_qty_configurator"`                  // Quantity visible in configurator
	Volume float64 `json:"volume"` // Volume
	// VolumeUomName                        string  `json:"volume_uom_name"`                           // Volume unit of measure label
	// WarehouseID                          any     `json:"warehouse_id"`                              // Warehouse 📦 relation: many2one stock.warehouse
	// WebsiteMessageIDs                    any     `json:"website_message_ids"`                       // Website Messages 📦 relation: one2many mail.message
	Weight float64 `json:"weight"` // Weight
	// WeightUomName                        string  `json:"weight_uom_name"`                           // Weight unit of measure label
	// WriteDate                              time.Time `json:"write_date"`                                 // Last Updated on
	// WriteUid                               any       `json:"write_uid"`                                  // Last Updated by 📦 relation: many2one res.users
}

// make sure that the MTO is set if it already set on the routes

func (m *Model) ProductTemplateStockable() {
	model := "product.template"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// 28751

	pcMap := m.productCategoryMap()

	products, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductTemplate_150{}), []any{
		[]any{"detailed_type", "=", "product"},
		// []any{"name", "like", "M999%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(products) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	source_mto_id, _ := m.Source.GetID("stock.location.route", []any{[]any{"name", "=", "Manufacture To Order"}})
	dest_mto_id, _ := m.Dest.GetID("stock.route", []any{[]any{"name", "=", "Replenish on Order (MTO)"}})

	bar := progressbar.Default(int64(len(products)))

	sem := make(chan int, 8)
	var wg sync.WaitGroup
	for _, r := range products {
		wg.Go(func() {
			sem <- 1
			defer func() { bar.Add(1); <-sem }()
			var p ProductTemplate_150
			FillStruct(r, &p)

			ur := map[string]any{
				"name":                   p.Name,
				"default_code":           p.DefaultCode,
				"barcode":                p.Barcode,
				"list_price":             p.ListPrice,
				"standard_price":         p.StandardPrice,
				"type":                   "consu",
				"is_storable":            true,
				"can_be_expensed":        p.CanBeExpensed,
				"sale_ok":                true,
				"purchase_ok":            true,
				"description":            p.Description,
				"sale_line_warn_msg":     p.SaleLineWarnMsg,
				"purchase_line_warn_msg": p.PurchaseLineWarnMsg,
				"volume":                 p.Volume,
				"weight":                 p.Weight,
				"sale_delay":             p.SaleDelay,
				"hs_code":                p.HsCode,
			}
			uom_id, _ := m.Dest.GetID("uom.uom", []any{[]any{"name", "=", p.UomID}})
			if uom_id != -1 {
				ur["uom_id"] = uom_id
			}

			_, categ_name := ParseMany2One(p.CategID)
			if categ_id, ok := pcMap[categ_name]; ok {
				ur["categ_id"] = categ_id
			}

			// RouteIDs processing to ensure MTO is included
			// m.Log.Info(model, "Processing RouteIDs for product", p.RouteIDs)
			routeIDs := anytointslice(p.RouteIDs)
			if slices.Contains(routeIDs, source_mto_id) {
				ur["route_ids"] = []int{dest_mto_id}
			}

			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", p.Name},
				[]any{"default_code", "=", p.DefaultCode},
				[]any{"barcode", "=", p.Barcode},
			})

			m.writeRecord(model, ur, rid)
			// bar.Add(1)
		})
	}
	wg.Wait()
	bar.Finish()
}

func (m *Model) ProductTemplateConsumable() {
	model := "product.template"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// 2977

	pcMap := m.productCategoryMap()

	products, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductTemplate_150{}), []any{
		[]any{"detailed_type", "=", "consu"},
		// []any{"name", "like", "M999%"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(products) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	source_mto_id, _ := m.Source.GetID("stock.location.route", []any{[]any{"name", "=", "Manufacture To Order"}})
	dest_mto_id, _ := m.Dest.GetID("stock.route", []any{[]any{"name", "=", "Replenish on Order (MTO)"}})

	bar := progressbar.Default(int64(len(products)))

	sem := make(chan int, 8)
	var wg sync.WaitGroup
	for _, r := range products {
		wg.Go(func() {
			sem <- 1
			defer func() { bar.Add(1); <-sem }()
			var p ProductTemplate_150
			FillStruct(r, &p)

			ur := map[string]any{
				"name":                   p.Name,
				"default_code":           p.DefaultCode,
				"barcode":                p.Barcode,
				"list_price":             p.ListPrice,
				"standard_price":         p.StandardPrice,
				"type":                   "consu",
				"is_storable":            false,
				"can_be_expensed":        p.CanBeExpensed,
				"sale_ok":                true,
				"purchase_ok":            true,
				"description":            p.Description,
				"sale_line_warn_msg":     p.SaleLineWarnMsg,
				"purchase_line_warn_msg": p.PurchaseLineWarnMsg,
				"volume":                 p.Volume,
				"weight":                 p.Weight,
				"sale_delay":             p.SaleDelay,
				"hs_code":                p.HsCode,
			}
			uom_id, _ := m.Dest.GetID("uom.uom", []any{[]any{"name", "=", p.UomID}})
			if uom_id != -1 {
				ur["uom_id"] = uom_id
			}

			_, categ_name := ParseMany2One(p.CategID)
			if categ_id, ok := pcMap[categ_name]; ok {
				ur["categ_id"] = categ_id
			}

			// RouteIDs processing to ensure MTO is included
			// m.Log.Info(model, "Processing RouteIDs for product", p.RouteIDs)
			routeIDs := anytointslice(p.RouteIDs)
			if slices.Contains(routeIDs, source_mto_id) {
				ur["route_ids"] = []int{dest_mto_id}
			}

			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", p.Name},
				[]any{"default_code", "=", p.DefaultCode},
				[]any{"barcode", "=", p.Barcode},
			})

			// m.Log.Info(model, "rid", rid)
			// fmt.Println(model, "rid", rid, "ur", ur)

			m.writeRecord(model, ur, rid)

			// bar.Add(1)
		})
	}
	wg.Wait()
	bar.Finish()
}

func anytointslice(a any) []int {
	ints := []int{}
	switch v := a.(type) {
	case []any:
		for _, item := range v {
			switch id := item.(type) {
			case float64:
				ints = append(ints, int(id))
			}
		}
	case []int:
		ints = v
	}
	return ints
}

func (m *Model) ProductTemplateService() {
	model := "product.template"
	banner.Println(model, trace())
	m.Log.Info(model, "func", trace())

	// 120

	pcMap := m.productCategoryMap()

	products, err := m.Source.SearchRead(model, 0, 0, ExtractJSONTags(ProductTemplate_150{}), []any{
		[]any{"detailed_type", "=", "service"},
	})
	if err != nil {
		m.Log.Error(model, "func", trace(), "err", err)
		return
	}
	if len(products) == 0 {
		m.Log.Info(model, "func", trace(), "err", "no record found")
		return
	}

	source_mto_id, _ := m.Source.GetID("stock.location.route", []any{[]any{"name", "=", "Manufacture To Order"}})
	dest_mto_id, _ := m.Dest.GetID("stock.route", []any{[]any{"name", "=", "Replenish on Order (MTO)"}})

	bar := progressbar.Default(int64(len(products)))

	sem := make(chan int, 8)
	var wg sync.WaitGroup
	for _, r := range products {
		wg.Go(func() {
			sem <- 1
			defer func() { bar.Add(1); <-sem }()
			var p ProductTemplate_150
			FillStruct(r, &p)

			ur := map[string]any{
				"name":                   p.Name,
				"default_code":           p.DefaultCode,
				"barcode":                p.Barcode,
				"list_price":             p.ListPrice,
				"standard_price":         p.StandardPrice,
				"type":                   "service",
				"can_be_expensed":        p.CanBeExpensed,
				"sale_ok":                true,
				"purchase_ok":            true,
				"description":            p.Description,
				"sale_line_warn_msg":     p.SaleLineWarnMsg,
				"purchase_line_warn_msg": p.PurchaseLineWarnMsg,
				"volume":                 p.Volume,
				"weight":                 p.Weight,
				"sale_delay":             p.SaleDelay,
				"hs_code":                p.HsCode,
			}
			uom_id, _ := m.Dest.GetID("uom.uom", []any{[]any{"name", "=", p.UomID}})
			if uom_id != -1 {
				ur["uom_id"] = uom_id
			}

			_, categ_name := ParseMany2One(p.CategID)
			if categ_id, ok := pcMap[categ_name]; ok {
				ur["categ_id"] = categ_id
			}

			// RouteIDs processing to ensure MTO is included
			// m.Log.Info(model, "Processing RouteIDs for product", p.RouteIDs)
			routeIDs := anytointslice(p.RouteIDs)
			if slices.Contains(routeIDs, source_mto_id) {
				ur["route_ids"] = []int{dest_mto_id}
			}

			rid, _ := m.Dest.GetID(model, []any{
				[]any{"name", "=", p.Name},
				[]any{"default_code", "=", p.DefaultCode},
				[]any{"barcode", "=", p.Barcode},
			})

			// m.Log.Info(model, "rid", rid)

			m.writeRecord(model, ur, rid)

			// bar.Add(1)
		})
	}
	wg.Wait()
	bar.Finish()
}
