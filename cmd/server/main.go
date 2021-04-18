package main

import (
	"fmt"
	"net/http"

	"github.com/ourshoptools/gain-api/internal/database"
	transportHandler "github.com/ourshoptools/gain-api/internal/transport/http"
)

type App struct {}

func (app *App) Run() error {
	fmt.Println("-------Re:Gain Setting up Server-------")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		fmt.Println("-------Re:Gain Error Creating Database-------")
		return err
	}
	handler := transportHandler.NewHandler()
	handler.SetupRoutes()

	if err:= http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("-------Re:Gain Error Serving Router-------")
		return err
	}
	return nil
}

func main() {
	fmt.Println("-------Re:Gain API Server-------")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("-------Re:Gain Error Starting API Server-------")
		fmt.Println(err)
	}

}