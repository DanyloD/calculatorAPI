package main

import (
	"backendCalculator/internal/operations"
	"encoding/json"
	"fmt"
	"net/http"
)

type CalculationHandler struct {
	calc operations.CalculationService
}

func NewCalculationHandler(calc operations.CalculationService) *CalculationHandler {
	return &CalculationHandler{
		calc: calc,
	}
}

type requestData struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type responseData struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func (h *CalculationHandler) decodeRequest(w http.ResponseWriter, r *http.Request) (requestData, bool) {
	var req requestData

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseData{Error: "Incorrect JSON: " + err.Error()})
		return req, false
	}
	return req, true
}

func (h *CalculationHandler) HandleAdd(w http.ResponseWriter, r *http.Request) {
	req, ok := h.decodeRequest(w, r)
	if !ok {
		return
	}

	res := h.calc.Add(req.A, req.B)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData{Result: res})
}

func (h *CalculationHandler) HandleSub(w http.ResponseWriter, r *http.Request) {
	req, ok := h.decodeRequest(w, r)
	if !ok {
		return
	}

	res := h.calc.Sub(req.A, req.B)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData{Result: res})
}

func (h *CalculationHandler) HandleMul(w http.ResponseWriter, r *http.Request) {
	req, ok := h.decodeRequest(w, r)
	if !ok {
		return
	}

	res := h.calc.Mul(req.A, req.B)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData{Result: res})
}

func (h *CalculationHandler) HandleDiv(w http.ResponseWriter, r *http.Request) {
	req, ok := h.decodeRequest(w, r)
	if !ok {
		return
	}

	res, err := h.calc.Div(req.A, req.B)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(responseData{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData{Result: res})
}

func main() {
	calcService := operations.ToTie{}
	calcHandler := NewCalculationHandler(calcService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add", calcHandler.HandleAdd)
	mux.HandleFunc("POST /sub", calcHandler.HandleSub)
	mux.HandleFunc("POST /mul", calcHandler.HandleMul)
	mux.HandleFunc("POST /div", calcHandler.HandleDiv)

	fmt.Println("Сервер запущено на :8080")
	http.ListenAndServe(":8080", mux)
}
