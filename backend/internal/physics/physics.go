package physics

import (
	"math"
)

// SurfaceType mendefinisikan jenis permukaan jalan di Warehouse Express 3D
type SurfaceType string

const (
	SurfaceAsphalt  SurfaceType = "asphalt"  // 100% Speed (0.56 max base speed)
	SurfaceConcrete SurfaceType = "concrete" // 80% Speed  (0.45)
	SurfaceDirt     SurfaceType = "dirt"     // 60% Speed  (0.34)
	SurfaceGrass    SurfaceType = "grass"    // 35% Speed  (0.20)
)

// GetSurfaceMaxSpeed mengembalikan kecepatan batas wajar berdasarkan tipe permukaan
func GetSurfaceMaxSpeed(s SurfaceType) float64 {
	switch s {
	case SurfaceAsphalt:
		return 0.56
	case SurfaceConcrete:
		return 0.45
	case SurfaceDirt:
		return 0.34
	case SurfaceGrass:
		return 0.20
	default:
		return 0.20
	}
}

// Courier merepresentasikan state posisi kendaraan kurir
type Courier struct {
	ID   string  `json:"id"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Z    float64 `json:"z"`
	RotY float64 `json:"rot_y"`
}

// ValidatePlayerStep memvalidasi perpindahan pemain untuk mencegah teleportasi / speed-hack
func ValidatePlayerStep(p *Courier, nextX, nextZ, dt float64, surf SurfaceType) (bool, float64, float64) {
	dist := math.Hypot(nextX-p.X, nextZ-p.Z)
	// Toleransi 1.25x untuk lag jitter / framing variatif (dt dalam ms, standar 16.67ms = 60fps)
	maxAllowed := GetSurfaceMaxSpeed(surf) * (dt / 16.67) * 1.25

	isValid := dist <= maxAllowed
	return isValid, dist, maxAllowed
}
