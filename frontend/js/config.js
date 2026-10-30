/**
 * Warehouse Express 3D - Client Environment & API Configuration
 * Mendeteksi environment secara dinamis untuk deployment produksi, staging, maupun development.
 */

(function () {
  const isLocalhost = Boolean(
    window.location.hostname === 'localhost' ||
    window.location.hostname === '127.0.0.1' ||
    window.location.hostname === '[::1]' ||
    window.location.protocol === 'file:'
  );

  // Jika ada custom runtime injection (misal window.__ENV__.API_BASE_URL)
  const injectedApiUrl = window.__ENV__ && window.__ENV__.API_BASE_URL;

  // Aturan resolusi endpoint API:
  // 1. Injected URL (jika ada)
  // 2. Localhost -> default ke backend port 8080
  // 3. Produksi -> menggunakan relative path '/api' (direkomendasikan via reverse proxy / domain sama)
  let resolvedApiBaseUrl = 'http://localhost:8080';

  if (injectedApiUrl) {
    resolvedApiBaseUrl = injectedApiUrl;
  } else if (!isLocalhost) {
    // Di lingkungan produksi (Vercel, Netlify, Docker Nginx, VPS)
    // gunakan origin yang sama atau sub-path '/api'
    resolvedApiBaseUrl = window.location.origin;
  }

  window.GAME_CONFIG = {
    ENV: isLocalhost ? 'development' : 'production',
    API_BASE_URL: resolvedApiBaseUrl,
    ENDPOINTS: {
      HEALTH: `${resolvedApiBaseUrl}/health`,
      VEHICLES: `${resolvedApiBaseUrl}/api/v1/vehicles`,
      DISPATCH_NEXT: `${resolvedApiBaseUrl}/api/v1/dispatch/next`,
      TELEMETRY_VERIFY: `${resolvedApiBaseUrl}/api/v1/telemetry/verify`,
      PLAYER_SAVE: `${resolvedApiBaseUrl}/api/v1/player/save`,
      PLAYER_LOAD: `${resolvedApiBaseUrl}/api/v1/player/load`,
      LEADERBOARD: `${resolvedApiBaseUrl}/api/v1/leaderboard`
    },
    VERSION: '1.0.0'
  };

  console.log(`[GameConfig] Initialized in ${window.GAME_CONFIG.ENV} mode.`);
  console.log(`[GameConfig] API Base URL: ${window.GAME_CONFIG.API_BASE_URL}`);
})();
