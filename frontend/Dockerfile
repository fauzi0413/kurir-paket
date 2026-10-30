# =========================================================================
# Production Dockerfile untuk Frontend (Three.js Web Game)
# Menggunakan Nginx Alpine ringan (<25 MB) dengan Gzip & Reverse Proxy
# =========================================================================

FROM nginx:alpine

# Hapus default website Nginx
RUN rm -rf /usr/share/nginx/html/*

# Salin konfigurasi Nginx kustom
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Salin file frontend statis ke webroot
COPY . /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
