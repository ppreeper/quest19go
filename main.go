package main

import (
	"fmt"
	"time"

	"quest19go/app"
	"quest19go/models"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// app.Ctx = context.Background()
	// app.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("failed to connect database: %v", err)
	// }
	// Migrate the schema
	// app.DB.AutoMigrate(&models.StockLocation_150{})

	start := time.Now()
	startTime := time.Since(start)
	app.Log.Info("odooimport", "start", startTime)
	// Make sure all sequences are set
	model := models.Model(*app)

	// model.AccountPaymentTermLine,   // Payment Term Lines

	// Add NCR Manager and MTR Manager permissions before loading data

	fns := []func(){
		model.DecimalPrecision, // Decimal Precision
		model.ResCurrency,      // Currencies
		model.ResCompany,       // Company
		model.Uomuom,           // Units of Measure
		model.ResBank,          // Banks
		model.AccountAccountCleanup,
		model.AccountAccountNewAccounts, // Account New Accounts (clearing accounts)
		model.AccountAccountRename,      // Account Rename
		model.AccountAccount,            // Chart Of Accounts
		model.AccountJournal,            // Journals
		model.AccountPaymentTermRename,  // Payment Terms Rename
		model.AccountPaymentTerm,        // Payment Terms
		model.ResUsers,                  // Employees
		model.ResUsersCreateEmployee,
		model.HRDepartment,
		model.HRJob,
		model.HREmployee,
		model.HRDepartment,
		model.CrmTeam,           // Sales Teams
		model.CrmTeamMember,     // Sales Team Members
		model.ProductCategories, // Product Categories
		model.StockWarehouse,    // Warehouses
		model.StockLocation,     // Stock Locations
		model.ProductTemplateStockable,
		model.ProductTemplateConsumable,
		model.ProductTemplateService,
		model.ProductPricelist,     // Customer Pricelists
		model.ProductPricelistItem, // Customer Pricelist items
		model.ResPartner,           // Partners - Companies and Contacts
		model.ProductSupplierinfo,  // Vendor Pricelists
		model.MRPWorkcenterTag,
		model.MRPWorkcenter,
		model.MRPBom,               // BOM
		model.MRPRoutingWorkcenter, // BOM Operations
		model.MRPBomLine,           // BOM Components
		model.NCRCategory,          // NCR Categories
		model.NCRClaim,             // NCR Claims
		model.NCRProductLine,       // NCR Claim Product Lines
		model.MtrTemplate,          // MTR Templates
	}
	for _, fn := range fns {
		step := time.Now()
		fn()
		fmt.Printf("%s step: total: %v\n", time.Since(step), time.Since(start))
	}
	endTime := time.Since(start)
	app.Log.Info("odooimport", "end", endTime)
	fmt.Printf("odooimport end: %v\n", endTime)
}
