package api

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"kurir-paket/backend/internal/dispatcher"
	"kurir-paket/backend/internal/models"
	"kurir-paket/backend/internal/physics"
)

type ServerHandler struct {
	dispatcher *dispatcher.Dispatcher
	mu         sync.RWMutex
	profiles   map[string]*models.PlayerProfile
	vehicles   []models.VehicleItem
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{
		dispatcher: dispatcher.NewDispatcher(),
		profiles:   make(map[string]*models.PlayerProfile),
		vehicles: []models.VehicleItem{
			{ID: "courier_van", Name: "Eco Van 1.0", Tier: "Starter", Price: 0, TopSpeed: 0.56, Accel: 0.015, Capacity: 3, ColorHex: "#2563eb", Description: "Mobil box standar pengantaran kota."},
			{ID: "turbo_pickup", Name: "Turbo Pickup X", Tier: "Mid", Price: 1500000, TopSpeed: 0.65, Accel: 0.020, Capacity: 4, ColorHex: "#059669", Description: "Pick-up bertenaga tinggi untuk rute menanjak."},
			{ID: "electric_truck", Name: "Volt Cargo Heavy", Tier: "Pro", Price: 3500000, TopSpeed: 0.72, Accel: 0.024, Capacity: 5, ColorHex: "#d97706", Description: "Truk kargo listrik dengan kapasitas muatan besar."},
			{ID: "cyber_express", Name: "Hyper Delivery 2077", Tier: "Legendary", Price: 8000000, TopSpeed: 0.85, Accel: 0.030, Capacity: 6, ColorHex: "#7c3aed", Description: "Kendaraan prototipe futuristik kecepatan maksimal."},
		},
	}
}

func EnableCORS(w *http.ResponseWriter) {
	allowed := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowed == "" {
		allowed = "*"
	}
	(*w).Header().Set("Access-Control-Allow-Origin", allowed)
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func (h *ServerHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"service":   "Warehouse Express 3D Game Server",
		"timestamp": time.Now().Unix(),
	})
}

func (h *ServerHandler) GetVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.vehicles)
}

func (h *ServerHandler) DispatchNextHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.DispatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	pickedID := h.dispatcher.NextTargetHouse(req.HouseIDs)

	res := models.DispatchResponse{
		TargetHouseID: pickedID,
		Message:       "Target dispatch generated via Go Dispatcher",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ServerHandler) TelemetryVerifyHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.TelemetryVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	courier := &physics.Courier{
		ID: req.PlayerID,
		X:  req.PrevX,
		Z:  req.PrevZ,
	}

	valid, distMoved, maxAllowed := physics.ValidatePlayerStep(
		courier,
		req.NextX,
		req.NextZ,
		req.DeltaMs,
		physics.SurfaceType(req.Surface),
	)

	res := models.TelemetryVerifyResponse{
		Valid:          valid,
		DistanceMoved:  distMoved,
		MaxAllowedDist: maxAllowed,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ServerHandler) PlayerSaveHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var profile models.PlayerProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if profile.PlayerID == "" {
		profile.PlayerID = "player_default"
	}
	profile.UpdatedTimestamp = time.Now().Unix()

	h.mu.Lock()
	h.profiles[profile.PlayerID] = &profile
	h.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Player profile saved successfully",
		"data":    profile,
	})
}

func (h *ServerHandler) PlayerLoadHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}

	playerID := r.URL.Query().Get("playerId")
	if playerID == "" {
		playerID = "player_default"
	}

	h.mu.RLock()
	profile, exists := h.profiles[playerID]
	h.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if !exists {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"exists":  false,
			"message": "Profile not found, using default starting profile",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"exists": true,
		"data":   profile,
	})
}

func (h *ServerHandler) LeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	EnableCORS(&w)
	if r.Method == http.MethodOptions {
		return
	}

	leaderboard := []models.LeaderboardEntry{
		{Rank: 1, PlayerName: "SpeedyCourier", HighestLevel: 15, TotalSalary: 12500000, VehicleUsed: "Hyper Delivery 2077"},
		{Rank: 2, PlayerName: "CargoKing", HighestLevel: 11, TotalSalary: 8400000, VehicleUsed: "Volt Cargo Heavy"},
		{Rank: 3, PlayerName: "DriftRunner", HighestLevel: 8, TotalSalary: 5100000, VehicleUsed: "Turbo Pickup X"},
		{Rank: 4, PlayerName: "UrbanExpress", HighestLevel: 5, TotalSalary: 2300000, VehicleUsed: "Eco Van 1.0"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}
