# Spesifikasi API Kontrak (Warehouse Express 3D)

Dokumen ini memuat spesifikasi endpoint REST API untuk integrasi antara Frontend Game dan Backend Game Server.

- **Base URL Lokal**: `http://localhost:8080`
- **Header Standar**: `Content-Type: application/json`
- **CORS**: Didukung (`Access-Control-Allow-Origin: *`)

---

## 1. Health Check

Memeriksa ketersediaan dan status server.

- **URL**: `/health`
- **Method**: `GET`
- **Response `200 OK`**:
```json
{
  "status": "healthy",
  "service": "Warehouse Express 3D Game Server",
  "timestamp": 1793352594
}
```

---

## 2. Katalog Kendaraan Dealer

Mengambil daftar kendaraan yang tersedia untuk dibeli pemain.

- **URL**: `/api/v1/vehicles`
- **Method**: `GET`
- **Response `200 OK`**:
```json
[
  {
    "id": "courier_van",
    "name": "Eco Van 1.0",
    "tier": "Starter",
    "price": 0,
    "topSpeed": 0.56,
    "accel": 0.015,
    "capacity": 3,
    "colorHex": "#2563eb",
    "description": "Mobil box standar pengantaran kota."
  },
  {
    "id": "turbo_pickup",
    "name": "Turbo Pickup X",
    "tier": "Mid",
    "price": 1500000,
    "topSpeed": 0.65,
    "accel": 0.02,
    "capacity": 4,
    "colorHex": "#059669",
    "description": "Pick-up bertenaga tinggi untuk rute menanjak."
  }
]
```

---

## 3. Anti-Duplicate Order Dispatcher

Menentukan rumah tujuan pengantaran paket berikutnya secara acak tanpa menduplikasi rumah sebelumnya.

- **URL**: `/api/v1/dispatch/next`
- **Method**: `POST`
- **Request Body**:
```json
{
  "playerId": "player_123",
  "levelId": 3,
  "houseIds": [1, 2, 3, 4, 5, 6, 7, 8]
}
```
- **Response `200 OK`**:
```json
{
  "targetHouseId": 5,
  "message": "Target dispatch generated via Go Dispatcher"
}
```

---

## 4. Validasi Telemetri & Anti-Cheat

Memverifikasi jarak langkah kendaraan pemain berdasarkan batas wajar tipe permukaan dan selang waktu ($\Delta t$).

- **URL**: `/api/v1/telemetry/verify`
- **Method**: `POST`
- **Request Body**:
```json
{
  "playerId": "player_123",
  "prevX": 10.5,
  "prevZ": -20.2,
  "nextX": 10.9,
  "nextZ": -20.4,
  "deltaMs": 16.67,
  "surface": "asphalt"
}
```
- **Response `200 OK`**:
```json
{
  "valid": true,
  "distanceMoved": 0.4472,
  "maxAllowedDist": 0.7000
}
```

---

## 5. Simpan Progres Pemain (Cloud Save)

Menyimpan data saldo gaji, level yang dicapai, dan kendaraan pemain.

- **URL**: `/api/v1/player/save`
- **Method**: `POST`
- **Request Body**:
```json
{
  "playerId": "fauzi_driver",
  "playerName": "Fauzi",
  "totalSalary": 3450000,
  "currentLevel": 4,
  "equippedVehicle": "turbo_pickup",
  "purchasedVehicles": ["courier_van", "turbo_pickup"]
}
```
- **Response `200 OK`**:
```json
{
  "success": true,
  "message": "Player profile saved successfully",
  "data": { ... }
}
```

---

## 6. Muat Progres Pemain (Cloud Load)

Mengambil data simpanan pemain berdasarkan ID.

- **URL**: `/api/v1/player/load?playerId=fauzi_driver`
- **Method**: `GET`
- **Response `200 OK`**:
```json
{
  "exists": true,
  "data": {
    "playerId": "fauzi_driver",
    "playerName": "Fauzi",
    "totalSalary": 3450000,
    "currentLevel": 4,
    "equippedVehicle": "turbo_pickup",
    "purchasedVehicles": ["courier_van", "turbo_pickup"],
    "updatedTimestamp": 1793352600
  }
}
```

---

## 7. Global Leaderboard

Mengambil daftar skor dan pencapaian tertinggi.

- **URL**: `/api/v1/leaderboard`
- **Method**: `GET`
- **Response `200 OK`**:
```json
[
  {
    "rank": 1,
    "playerName": "SpeedyCourier",
    "highestLevel": 15,
    "totalSalary": 12500000,
    "vehicleUsed": "Hyper Delivery 2077"
  }
]
```
