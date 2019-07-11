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
	"github.com/rs/cors"
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
	api := a.Router.PathPrefix("/api/v1/").Subrouter()

	// Routing for handling the projects
	api.HandleFunc("/products", a.GetAllProducts).Methods("GET")
	api.HandleFunc("/products", a.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", a.GetProduct).Methods("GET")
	api.HandleFunc("/products/{id}", a.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", a.DeleteProduct).Methods("DELETE")

	api.HandleFunc("/manufacturers", a.GetAllManufacturers).Methods("GET")
	api.HandleFunc("/manufacturers", a.CreateManufacturer).Methods("POST")
	api.HandleFunc("/manufacturers/{id}", a.GetManufacturer).Methods("GET")
	api.HandleFunc("/manufacturers/{id}", a.UpdateManufacturer).Methods("PUT")
	api.HandleFunc("/manufacturers/{id}", a.DeleteManufacturer).Methods("DELETE")

	// Serve static assets directly.
	a.Router.PathPrefix("/*").Handler(http.FileServer(http.Dir("./dist/ui")))
	a.Router.PathPrefix("/").HandlerFunc(IndexHandler("./ui/dist/ui/index.html"))
}

// IndexHandler entry point for web app
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

// Get App Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

//Post App Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

//Put App  Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete App Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// GetAllManufacturers Handlers to manage Product Data
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
	c := cors.AllowAll()

	handler := c.Handler(a.Router)

	// start server listen
	log.Fatal(http.ListenAndServe(host, handler))
}
