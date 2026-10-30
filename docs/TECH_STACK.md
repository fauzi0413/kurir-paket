# Spesifikasi & Rencana Tech Stack (Warehouse Express 3D)

Dokumen ini memuat detail rencana *tech stack* dan evolusi teknologi untuk proyek **Warehouse Express 3D: Connected City**.

---

## 1. Ringkasan Tech Stack

| Lapisan Sistem | Teknologi / Framework | Alasan & Fungsi |
| :--- | :--- | :--- |
| **Frontend 3D Core** | **Three.js (r128+)** | Rendering WebGL performa tinggi, scene graph, pencahayaan, bayangan dinamis, dan navigasi 3D marker. |
| **Frontend Styling** | **Tailwind CSS + Vanilla CSS** | Utility-first untuk HUD adaptif & responsive, dipadukan Vanilla CSS untuk efek Glassmorphism & animasi keyframes. |
| **Frontend Audio** | **Web Audio API (Procedural)** | Sintesis suara mesin, sirine, koin, dan klakson secara dinamis tanpa perlu aset audio audio berat. |
| **Backend Core (Planned)** | **Golang (Go 1.20+)** | Server game performa tinggi, konkurensi goroutine + mutex aman untuk dispatcher dan anti-cheat. |
| **Backend Runtime (Dev/Mock)**| **Node.js (Built-in HTTP)** | Implementasi referensi instan tanpa instalasi dependensi pihak ketiga, siap dijalankan langsung di mesin lokal. |
| **Penyimpanan Data** | **LocalStorage (Client) & SQLite / PostgreSQL (Server)** | Dual-mode: Offline play via LocalStorage dan cloud syncing via REST API. |
| **Protokol Jaringan** | **HTTP/1.1 REST API & WebSocket (Rencana)** | REST API untuk sinkronisasi akun & penentuan target paket; WebSocket untuk telemetry multiplayer mendatang. |
| **DevOps & Kontainerisasi** | **Docker, Docker Compose, Nginx Alpine, GitHub Actions** | Multi-stage build Go & Nginx, reverse proxy, health checks, dan CI lint/build testing. |
| **Cloud Hosting (PaaS)** | **Vercel, Netlify, Railway, Render, VPS** | Konfigurasi deployment fleksibel (`vercel.json`, `netlify.toml`, `docker-compose.yml`). |

---

## 2. Rincian Lapisan Frontend

### Three.js 3D Engine
- **Scene & Camera**: `PerspectiveCamera` (FOV 45°) dengan algoritma kamera pengekor dinamis (*Third-Person Follow Camera*) dengan peredaman halus (*lerp*).
- **Lighting & Shadow**:
  - `DirectionalLight` (Matahari) dengan `PCFSoftShadowMap` untuk bayangan realistis mobil dan gedung.
  - `AmbientLight` untuk menerangi area bayangan agar tetap estetik.
- **Procedural World Generation**:
  - Geometri kota disusun secara algoritmik per tingkat level.
  - Distribusi perumahan, jalur aspal utama, lorong tanah, dan pepohonan diposisikan tanpa memerlukan model 3D berat (.gltf/.obj).

### Audio Synthesizer
- `OscillatorNode`: Menghasilkan frekuensi pitch mesin yang bergeser mengikuti rasio kecepatan kendaraan (`speedRatio`).
- `BiquadFilterNode`: Low-pass filter untuk menyaring frekuensi kasar sehingga suara mesin terasa tebal dan realistis.

---

## 3. Rincian Lapisan Backend

### Golang Game Server (`backend/cmd/server`)
- Dipilih sebagai bahasa server utama karena karakteristik:
  1. **Ultra Low Latency**: Eksekusi kode biner asli tanpa overhead Garbage Collector runtime berat.
  2. **Thread Safety**: Manajemen mutasi state bersama (`Dispatcher.lastDeliveredHouseID`) terlindungi penuh dengan `sync.Mutex`.
  3. **Authoritative Anti-Cheat**: Validasi kalkulasi fisika di sisi server sebelum mencatat pencapaian ke papan skor.

### Node.js Development Server (`backend/server.js`)
- Disediakan sebagai pelengkap cepat untuk pengujian lokal:
  - Menggunakan API bawaan `http` dan `URL` (Zero-dependency).
  - Port default: `8080`.

---

## 4. Rencana Roadmap Pengembangan Mendatang

1. **Fase 1 (Selesai)**:
   - Pemisahan modular Frontend (`frontend/`) dan Backend (`backend/`).
   - Ekstraksi styling CSS terpisah.
   - Implementasi Go Game Server & Node.js Server referensi.
   - Dokumentasi lengkap arsitektur dan kontrak API.
2. **Fase 2 (Berikutnya)**:
   - Integrasi koneksi API aktif dari client ke backend (Fetch API untuk penentuan order paket dan verifikasi langkah).
   - Pengaktifan fitur sinkronisasi cloud login / simpan progres online.
3. **Fase 3 (Lanjutan)**:
   - Mode Multiplayer real-time kurir berbasis WebSocket.
   - Papan peringkat global lintas pemain.
