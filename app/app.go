package app

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/ppreeper/odoorpc"
	"github.com/ppreeper/odoorpc/odoojrpc"
	"github.com/ppreeper/odoorpc/odoojson"
)

type App struct {
	Source odoorpc.Odoo
	Dest   odoorpc.Odoo
	Log    *slog.Logger
	// DB     *gorm.DB
	// Ctx    context.Context
}

// 5ef6cf19b1c270ad4489095e9839b80aec61d32e

// q15 = JRPC(
//
//	hostname="quest15data.odoopro.ca",
//	port=443,
//	schema="https",
//	database="quest15_data",
//	username="aventador",
//	password="57c850aeec0310c8c67bc167701c1145ecfff111",
//
// )
// q15.Login()
// q19 = JRPC(
//
//	hostname="quest19data.odoopro.ca",
//	port=443,
//	schema="https",
//	database="quest19_data",
//	username="admin",
//	password="5142bde98870da39d83f2f33fb60f07fbf28c6f7",
func NewApp() (*App, error) {
	conn := &App{}
	conn.Source = odoojrpc.NewOdoo().
		WithHostname("quest15data.odoopro.ca").
		WithPort(443).WithSchema("https").
		WithDatabase("quest15_data").
		WithUsername("aventador").
		WithPassword("z4ok8Zj2SX9hfPI5CCeM")

	conn.Dest = odoojson.NewOdoo().
		WithHostname("quest19data.odoopro.ca").
		WithPort(443).WithSchema("https").
		WithDatabase("quest19_data").
		WithAPIKey("5142bde98870da39d83f2f33fb60f07fbf28c6f7")

	conn.setupLogging("import.log")

	// odoo login to get uid
	if err := conn.Source.Login(); err != nil {
		return &App{}, err
	}
	if err := conn.Dest.Login(); err != nil {
		return &App{}, err
	}
	return conn, nil
}

func (a *App) CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		a.Log.Error(err.Error())
	}
}

func (a *App) FatalErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		a.Log.Error(err.Error())
		os.Exit(1)
	}
}

func (a *App) setupLogging(logName string) {
	// check for file existence
	_, err := os.Stat(logName)
	if !errors.Is(err, os.ErrNotExist) {
		// if exists truncate
		err = os.Truncate(logName, 0)
		if err != nil {
			fmt.Println(err.Error())
			slog.Error(err.Error())
		}
	}

	f, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err.Error())
		slog.Error(err.Error())
		os.Exit(1)
	}
	logwriter := io.Writer(f)
	a.Log = slog.New(slog.NewTextHandler(logwriter, nil))
	slog.SetDefault(a.Log)
}
