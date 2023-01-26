package main

import (
	"os"
	"testing"

	"github.com/danielsmith0630/go-boilerplate/app"
)

var testApp *app.App

func TestMain(m *testing.M) {
	testApp = app.NewApp()
	testApp.InitializeDB()
	defer testApp.DB.Close()

	testApp.InitializeRoutes()

	code := m.Run()

	os.Exit(code)
}
