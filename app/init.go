package app

import (
	"os"

	"github.com/nicolasgaraza/petConnect/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //MySQL driver for gorm
	"github.com/revel/revel"
)

//Gdb it is the main access to the database
var Gdb *gorm.DB

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	//revel.InterceptMethod((*gromcontroller.GormController).Begin, revel.BEFORE)
	//revel.InterceptMethod((*gromcontroller.GormController).Commit, revel.AFTER)
	//revel.InterceptMethod((*gromcontroller.GormController).Rollback, revel.FINALLY)
	//revel.OnAppStart(FillCache)
}

func InitDB() {

	var connectionString string

	if os.Getenv("TRAVIS_GO_VERSION") != "" {
		connectionString = "travis:@/petconnect?charset=utf8&parseTime=True&loc=Local"
	} else {
		connectionString = "root:1q2w3e4r5t6y@/petconnect?charset=utf8&parseTime=True&loc=Local"
	}

	var err error
	Gdb, err = gorm.Open("mysql", connectionString)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	revel.INFO.Println("About to automigrate model")

	if revel.DevMode {
		initDevData(Gdb)
	} else {
		Gdb.AutoMigrate(&models.Service{})
	}

}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
