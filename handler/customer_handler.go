package handler

import (
	"encoding/json"
	"f2_gc1/model"
	"f2_gc1/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerUseCase usecase.ICustomerUseCase
}

func NewCustomerHandler(customerUseCase usecase.ICustomerUseCase) customerHandler {
	return customerHandler{
		customerUseCase: customerUseCase,
	}
}

func (c *customerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload model.Customer

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}

	err = c.customerUseCase.Create(payload)
	if err != nil {
		_ = fmt.Errorf("error:%v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Success"))
}

func (c *customerHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payload model.Customer
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	payload.ID = id // Inject ID dari URL ke payload

	if err := c.customerUseCase.Update(payload); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update customer: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated successfully"})
}

func (c *customerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.customerUseCase.Delete(id); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete customer: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
}

func (c *customerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := c.customerUseCase.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch customers: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (c *customerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	customer, err := c.customerUseCase.GetByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Customer not found: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
