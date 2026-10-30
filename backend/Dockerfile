# =========================================================================
# Production Dockerfile untuk Go Game Server
# Multi-stage build menghasilkan image ultra-ringan (~15 MB) & aman
# =========================================================================

# Stage 1: Build binary Go
FROM golang:1.20-alpine AS builder

WORKDIR /build

# Copy file dependency
COPY go.mod ./
# Download dependencies bila nanti ada package eksternal
RUN go mod download

# Copy source code backend
COPY . .

# Kompilasi binary berekstensi statis tanpa CGO
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /build/bin/server cmd/server/main.go

# Stage 2: Minimalist runtime environment
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

# Buat non-root user untuk alasan keamanan deployment
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Salin binary dari stage builder
COPY --from=builder /build/bin/server /app/server

USER appuser

EXPOSE 8080

# Health check otomatis untuk orchestrator (Docker / Kubernetes)
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/server"]
