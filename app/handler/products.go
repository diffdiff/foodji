package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diffdiff/foodji/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Products := []model.Product{}
	db.Find(&Products)
	respondJSON(w, http.StatusOK, Products)
}

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, Product)
}

func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Product := getProductOr404(db, id, w, r)
	if Product == nil {
		return
	}
	respondJSON(w, http.StatusOK, Product)
}

func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Product := getProductOr404(db, id, w, r)
	if Product == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, Product)
}

func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Product := getProductOr404(db, id, w, r)
	if Product == nil {
		return
	}
	if err := db.Delete(&Product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getProductOr404 gets a Product instance if exists, or respond the 404 error otherwise
func getProductOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Product {
	Product := model.Product{}
	productID, err := strconv.Atoi(id)

	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid product id")
		return nil
	}

	if err := db.First(&Product, productID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &Product
}
