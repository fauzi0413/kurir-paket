package models

// PlayerProfile merepresentasikan data profil dan progres pemain
type PlayerProfile struct {
	PlayerID         string   `json:"playerId"`
	PlayerName       string   `json:"playerName"`
	TotalSalary      int      `json:"totalSalary"`
	CurrentLevel     int      `json:"currentLevel"`
	EquippedVehicle  string   `json:"equippedVehicle"`
	PurchasedVehicles []string `json:"purchasedVehicles"`
	UpdatedTimestamp int64    `json:"updatedTimestamp"`
}

// VehicleItem merepresentasikan spesifikasi kendaraan di dealer
type VehicleItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Tier        string  `json:"tier"`
	Price       int     `json:"price"`
	TopSpeed    float64 `json:"topSpeed"`
	Accel       float64 `json:"accel"`
	Capacity    int     `json:"capacity"`
	ColorHex    string  `json:"colorHex"`
	Description string  `json:"description"`
}

// DispatchRequest merepresentasikan payload request penentuan tujuan berikutnya
type DispatchRequest struct {
	PlayerID string `json:"playerId"`
	LevelID  int    `json:"levelId"`
	HouseIDs []int  `json:"houseIds"`
}

// DispatchResponse merepresentasikan respon penentuan target rumah
type DispatchResponse struct {
	TargetHouseID int    `json:"targetHouseId"`
	Message       string `json:"message"`
}

// TelemetryVerifyRequest memverifikasi langkah kendaraan untuk anti-cheat
type TelemetryVerifyRequest struct {
	PlayerID string  `json:"playerId"`
	PrevX    float64 `json:"prevX"`
	PrevZ    float64 `json:"prevZ"`
	NextX    float64 `json:"nextX"`
	NextZ    float64 `json:"nextZ"`
	DeltaMs  float64 `json:"deltaMs"`
	Surface  string  `json:"surface"`
}

// TelemetryVerifyResponse respon verifikasi langkah
type TelemetryVerifyResponse struct {
	Valid          bool    `json:"valid"`
	DistanceMoved  float64 `json:"distanceMoved"`
	MaxAllowedDist float64 `json:"maxAllowedDist"`
}

// LeaderboardEntry merepresentasikan catatan skor tertinggi
type LeaderboardEntry struct {
	Rank         int    `json:"rank"`
	PlayerName   string `json:"playerName"`
	HighestLevel int    `json:"highestLevel"`
	TotalSalary  int    `json:"totalSalary"`
	VehicleUsed  string `json:"vehicleUsed"`
}
