package handler
 
import (
	"encoding/json"
	"net/http"
 
	"github.com/diffdiff/foodji/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
 
func GetAllManufactures(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Manufactures := []model.Manufacturer{}
	db.Find(&Manufactures)
	respondJSON(w, http.StatusOK, Manufactures)
}
 
func CreateManufacturer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Manufacturer := model.Manufacturer{}
 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Manufacturer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
 
	if err := db.Save(&Manufacturer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, Manufacturer)
}
 
func GetManufacturer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
 
	name := vars["name"]
	Manufacturer := getManufacturerOr404(db, name, w, r)
	if Manufacturer == nil {
		return
	}
	respondJSON(w, http.StatusOK, Manufacturer)
}
 
func UpdateManufacturer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
 
	name := vars["name"]
	Manufacturer := getManufacturerOr404(db, name, w, r)
	if Manufacturer == nil {
		return
	}
 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Manufacturer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
 
	if err := db.Save(&Manufacturer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, Manufacturer)
}
 
func DeleteManufacturer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
 
	name := vars["name"]
	Manufacturer := getManufacturerOr404(db, name, w, r)
	if Manufacturer == nil {
		return
	}
	if err := db.Delete(&Manufacturer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
 
 
// getManufacturerOr404 gets a Manufacturer instance if exists, or respond the 404 error otherwise
func getManufacturerOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Manufacturer {
	Manufacturer := model.Manufacturer{}
	if err := db.First(&Manufacturer, model.Manufacturer{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &Manufacturer
}