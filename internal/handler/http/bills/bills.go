package bills

import (
	"encoding/json"
	"net/http"

	"users/internal/domain/bills/dto"
	"users/internal/service"
	"users/pkgs/logger"
)

type BillH struct {
	BillService service.BillService
	logger      logger.Logger
}

func NewBillH(bs service.BillService, log logger.Logger) *BillH {
	return &BillH{
		BillService: bs,
		logger:      log,
	}
}

func (h *BillH) GetBills(w http.ResponseWriter, r *http.Request) {
	var req dto.BillrequestCheck
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bills, err := h.BillService.GetBills(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: "bills fetched successfully",
		Data:    bills,
	})
	h.logger.Infof("bills fetched successfully")
}

func (h *BillH) PayBills(w http.ResponseWriter, r *http.Request) {
	var req dto.BillPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.BillService.PayBills(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: "Bill paid successfully",
		Data:    resp,
	})
}
