package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diffdiff/foodji/app/config"
	"github.com/diffdiff/foodji/app/handler"
	"github.com/diffdiff/foodji/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/products", a.GetAllProducts)
	a.Post("/products", a.CreateProduct)
	a.Get("/products/{title}", a.GetProduct)
	a.Put("/products/{title}", a.UpdateProduct)
	a.Delete("/products/{title}", a.DeleteProduct)

	a.Get("/manufacturers", a.GetAllManufacturers)
	a.Post("/manufacturers", a.CreateManufacturer)
	a.Get("/manufacturers/{title}", a.GetManufacturer)
	a.Put("/manufacturers/{title}", a.UpdateManufacturer)
	a.Delete("/manufacturers/{title}", a.DeleteManufacturer)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Product Data
func (a *App) GetAllManufacturers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllManufactures(a.DB, w, r)
}

func (a *App) CreateManufacturer(w http.ResponseWriter, r *http.Request) {
	handler.CreateManufacturer(a.DB, w, r)
}

func (a *App) GetManufacturer(w http.ResponseWriter, r *http.Request) {
	handler.GetManufacturer(a.DB, w, r)
}

func (a *App) UpdateManufacturer(w http.ResponseWriter, r *http.Request) {
	handler.UpdateManufacturer(a.DB, w, r)
}

func (a *App) DeleteManufacturer(w http.ResponseWriter, r *http.Request) {
	handler.DeleteManufacturer(a.DB, w, r)
}

// Handlers to manage Products Data
func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
