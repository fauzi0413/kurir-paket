package dispatcher

import (
	"math/rand"
	"sync"
	"time"
)

// Dispatcher mengelola penugasan rumah target secara thread-safe tanpa duplikasi beruntun
type Dispatcher struct {
	mu                   sync.Mutex
	lastDeliveredHouseID int
	rng                  *rand.Rand
}

// NewDispatcher menginisialisasi dispatcher baru
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		lastDeliveredHouseID: -1,
		rng:                  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NextTargetHouse memilih rumah berikutnya secara acak tanpa mengulang rumah terakhir
func (d *Dispatcher) NextTargetHouse(houseIDs []int) int {
	d.mu.Lock()
	defer d.mu.Unlock()

	if len(houseIDs) == 0 {
		return -1
	}
	if len(houseIDs) == 1 {
		d.lastDeliveredHouseID = houseIDs[0]
		return houseIDs[0]
	}

	candidates := make([]int, 0, len(houseIDs))
	for _, id := range houseIDs {
		if id != d.lastDeliveredHouseID {
			candidates = append(candidates, id)
		}
	}

	if len(candidates) == 0 {
		candidates = houseIDs
	}

	picked := candidates[d.rng.Intn(len(candidates))]
	d.lastDeliveredHouseID = picked
	return picked
}

// Reset mereset riwayat pengantaran terakhir (misal saat restart level)
func (d *Dispatcher) Reset() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.lastDeliveredHouseID = -1
}
