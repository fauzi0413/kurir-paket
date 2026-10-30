# 🚚 Warehouse Express 3D: Connected City

Sebuah game simulasi pengantaran paket 3D berbasis WebGL (*Three.js*) dengan fisika permukaan dinamis, kontur bukit bertingkat (*plateau 3D*), sistem audio sintetis prosedural, dan arsitektur backend yang dirancang untuk validasi anti-cheat serta order dispatcher thread-safe.

Repositori ini telah dikonfigurasi **Production-Ready** untuk deployment ke berbagai platform (Docker Compose, Vercel, Netlify, Railway, Render, dan VPS).

---

## 📁 Struktur Repositori Terpisah (Frontend & Backend)

```text
kurir-paket/
├── frontend/                     # Modul Frontend (Client-side Game)
│   ├── index.html                # Entry point utama aplikasi game 3D
│   ├── css/
│   │   └── style.css             # Desain Glassmorphism, animations, & custom scrollbar
│   ├── js/
│   │   ├── config.js             # Konfigurasi deteksi environment & API Base URL dinamis
│   │   └── ...                   # Modul logika permainan
│   ├── assets/                   # Folder aset statis (tekstur, ikon, audio)
│   ├── nginx.conf                # Konfigurasi Nginx produksi (Gzip, caching, reverse proxy /api/)
│   └── Dockerfile                # Image Nginx Alpine (<25 MB)
│
├── backend/                      # Modul Backend (Game Server & API)
│   ├── cmd/server/
│   │   └── main.go               # Entry point Go Game Server (Graceful shutdown & static serve)
│   ├── internal/
│   │   ├── api/                  # Handler REST API, CORS dinamis, & route mappings
│   │   ├── dispatcher/           # Algoritma penentuan target tanpa duplikasi (Mutex-safe)
│   │   ├── models/               # Model data struct untuk profil, kargo, & telemetri
│   │   └── physics/              # Authoritative speed multiplier & anti-cheat validator
│   ├── server.js                 # Implementasi referensi Node.js (Zero-dependency quickstart)
│   ├── package.json              # Konfigurasi runtime Node.js
│   ├── go.mod                    # Modul Go backend
│   ├── .env.example              # Template variabel lingkungan backend
│   ├── Dockerfile                # Multi-stage build Go minimalis (~15 MB) & non-root user
│   ├── Dockerfile.node           # Dockerfile alternatif untuk runtime Node.js
│   └── README.md                 # Panduan teknis backend
│
├── docs/                         # Dokumentasi Arsitektur, Spesifikasi, & Deployment
│   ├── ARCHITECTURE.md           # Diagram dan detail pemisahan FE-BE & sistem fisika
│   ├── TECH_STACK.md             # Rencana tech stack lengkap dan roadmap
│   ├── API_SPEC.md               # Kontrak spesifikasi REST API
│   └── DEPLOYMENT.md             # Panduan deployment lengkap (Docker, Vercel, Netlify, VPS)
│
├── .github/workflows/
│   └── ci.yml                    # Pipeline CI GitHub Actions untuk uji frontend & backend
├── docker-compose.yml            # Orkestrasi Docker multi-container (Frontend + Backend)
├── vercel.json                   # Konfigurasi deployment frontend di Vercel
├── netlify.toml                  # Konfigurasi deployment frontend di Netlify
├── index.html                    # Entry point fallback root
├── .gitignore                    # Konfigurasi file yang diabaikan git
└── README.md                     # Dokumentasi utama proyek
```

---

## ⚡ Quick Start (Pengembangan Lokal)

### 1. Menjalankan Frontend
Buka file `frontend/index.html` langsung di browser, atau jalankan menggunakan local server:
```bash
npx serve frontend -p 3000
```
Buka peramban di `http://localhost:3000`.

### 2. Menjalankan Backend
Pilih salah satu runtime:

**Opsi A: Node.js (Sudah terpasang di sistem)**
```bash
cd backend
node server.js
```
Server aktif di `http://localhost:8080`.

**Opsi B: Golang (Kompilasi performa tinggi)**
```bash
cd backend
go run cmd/server/main.go
```

---

## 🚀 Deployment ke Produksi

Projek ini mendukung 3 metode deployment:

### 1. Docker Compose (1 Perintah di Server VPS)
```bash
docker-compose up -d --build
```
- Frontend otomatis berjalan di port `80` (dan `3000`).
- Backend otomatis berjalan di port `8080`.
- Nginx pada frontend otomatis meneruskan request `/api/` ke backend sehingga bebas masalah CORS.

### 2. Vercel / Netlify (Frontend) + Railway / Render (Backend)
- **Frontend**: Hubungkan repo ke Vercel atau Netlify. Konfigurasi [vercel.json](file:///c:/Fauzi/kurir-paket/vercel.json) dan [netlify.toml](file:///c:/Fauzi/kurir-paket/netlify.toml) sudah tersedia otomatis.
- **Backend**: Hubungkan folder `backend/` ke Railway atau Render.
- Baca panduan lengkap langkah-demi-langkah di [docs/DEPLOYMENT.md](file:///c:/Fauzi/kurir-paket/docs/DEPLOYMENT.md).

---

## 🎮 Kontrol Permainan

| Tombol / Aksi | Fungsi |
| :--- | :--- |
| `W` / `Panah Atas` | Gas maju (Akselerasi) |
| `S` / `Panah Bawah`| Rem / Mundur |
| `A` / `Panah Kiri` | Belok kiri |
| `D` / `Panah Kanan`| Belok kanan |
| `Spasi` | Rem tangan / Tahan untuk memuat paket di gudang atau drop di rumah |
| `R` | Reset posisi mobil jika terbalik |
| `ESC` | Jeda (Pause) permainan / Buka menu opsi |
| `Touch Joystick` | Kemudi layar sentuh (Mobile friendly) |

---

## 📖 Dokumentasi Lengkap

- [Panduan Deployment Produksi](file:///c:/Fauzi/kurir-paket/docs/DEPLOYMENT.md)
- [Detail Arsitektur Sistem](file:///c:/Fauzi/kurir-paket/docs/ARCHITECTURE.md)
- [Rencana & Spesifikasi Tech Stack](file:///c:/Fauzi/kurir-paket/docs/TECH_STACK.md)
- [Spesifikasi Kontrak REST API](file:///c:/Fauzi/kurir-paket/docs/API_SPEC.md)
- [Petunjuk Modul Backend](file:///c:/Fauzi/kurir-paket/backend/README.md)
