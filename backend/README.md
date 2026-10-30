# Warehouse Express 3D - Backend Services

Modul backend untuk game **Warehouse Express 3D: Connected City**. Menyediakan engine validasi pergerakan anti-cheat, dispatcher acak anti-duplicate berbasis mutex, katalog kendaraan dealer, serta sinkronisasi state pemain.

---

## 🛠️ Tech Stack Backend

Backend ini dirancang dengan dua opsi runtime:
1. **Golang (Primary/Planned Game Server)**:
   - High concurrency, zero-allocation memory performance.
   - Built-in mutex synchronization untuk `Dispatcher.NextTargetHouse`.
   - Menggunakan `net/http` standard library (Zero external vendor dependencies).
2. **Node.js (Quickstart / Development Runtime)**:
   - Runtime yang sudah tersedia langsung di lingkungan developer.
   - Native HTTP server tanpa perlu `npm install` library pihak ketiga.

---

## 🚀 Cara Menjalankan

### Opsi A: Menggunakan Node.js (Langsung jalan)
Pastikan berada di folder `backend/`:
```bash
node server.js
```
Server akan aktif di `http://localhost:8080`.

### Opsi B: Menggunakan Golang (Direkomendasikan untuk Production)
Pastikan compiler Go terpasang:
```bash
go run cmd/server/main.go
```
Atau build binary:
```bash
go build -o bin/server.exe cmd/server/main.go
./bin/server.exe
```

---

## 📡 Daftar Endpoint API

| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/health` | Status health check server |
| `GET` | `/api/v1/vehicles` | Katalog kendaraan dealer mobil |
| `POST` | `/api/v1/dispatch/next` | Menentukan rumah target berikutnya (Anti-Duplicate) |
| `POST` | `/api/v1/telemetry/verify` | Validasi posisi pemain & anti-cheat speed multiplier |
| `POST` | `/api/v1/player/save` | Menyimpan progres profil, saldo, level, mobil |
| `GET` | `/api/v1/player/load?playerId=...` | Memuat data progres pemain |
| `GET` | `/api/v1/leaderboard` | Papan peringkat kurir terbaik |
