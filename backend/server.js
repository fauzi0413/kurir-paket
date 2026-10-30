/**
 * Warehouse Express 3D - Game Server (Node.js Reference Implementation)
 * Zero external dependencies: Uses native Node.js http, url, and crypto modules.
 */

const http = require('http');
const url = require('url');

const PORT = process.env.PORT || 8080;

// State in-memory
const state = {
  lastDeliveredHouseId: -1,
  profiles: new Map(),
  vehicles: [
    { id: "courier_van", name: "Eco Van 1.0", tier: "Starter", price: 0, topSpeed: 0.56, accel: 0.015, capacity: 3, colorHex: "#2563eb", description: "Mobil box standar pengantaran kota." },
    { id: "turbo_pickup", name: "Turbo Pickup X", tier: "Mid", price: 1500000, topSpeed: 0.65, accel: 0.020, capacity: 4, colorHex: "#059669", description: "Pick-up bertenaga tinggi untuk rute menanjak." },
    { id: "electric_truck", name: "Volt Cargo Heavy", tier: "Pro", price: 3500000, topSpeed: 0.72, accel: 0.024, capacity: 5, colorHex: "#d97706", description: "Truk kargo listrik dengan kapasitas muatan besar." },
    { id: "cyber_express", name: "Hyper Delivery 2077", tier: "Legendary", price: 8000000, topSpeed: 0.85, accel: 0.030, capacity: 6, colorHex: "#7c3aed", description: "Kendaraan prototipe futuristik kecepatan maksimal." }
  ],
  leaderboard: [
    { rank: 1, playerName: "SpeedyCourier", highestLevel: 15, totalSalary: 12500000, vehicleUsed: "Hyper Delivery 2077" },
    { rank: 2, playerName: "CargoKing", highestLevel: 11, totalSalary: 8400000, vehicleUsed: "Volt Cargo Heavy" },
    { rank: 3, playerName: "DriftRunner", highestLevel: 8, totalSalary: 5100000, vehicleUsed: "Turbo Pickup X" },
    { rank: 4, playerName: "UrbanExpress", highestLevel: 5, totalSalary: 2300000, vehicleUsed: "Eco Van 1.0" }
  ]
};

// Physics surface speed limits
const SURFACE_MAX_SPEED = {
  asphalt: 0.56,
  concrete: 0.45,
  dirt: 0.34,
  grass: 0.20
};

// Mutex-Safe Anti-Duplicate Dispatcher
function nextTargetHouse(houseIds) {
  if (!houseIds || houseIds.length === 0) return -1;
  if (houseIds.length === 1) {
    state.lastDeliveredHouseId = houseIds[0];
    return houseIds[0];
  }

  const candidates = houseIds.filter(id => id !== state.lastDeliveredHouseId);
  const pool = candidates.length > 0 ? candidates : houseIds;
  const picked = pool[Math.floor(Math.random() * pool.length)];
  state.lastDeliveredHouseId = picked;
  return picked;
}

// Anti-cheat verification
function validatePlayerStep(prevX, prevZ, nextX, nextZ, deltaMs, surface) {
  const dist = Math.hypot(nextX - prevX, nextZ - prevZ);
  const maxSpeed = SURFACE_MAX_SPEED[surface] || 0.20;
  const maxAllowed = maxSpeed * (deltaMs / 16.67) * 1.25;
  return {
    valid: dist <= maxAllowed,
    distanceMoved: Number(dist.toFixed(4)),
    maxAllowedDist: Number(maxAllowed.toFixed(4))
  };
}

function getCORSOrigin() {
  return process.env.CORS_ALLOWED_ORIGINS || '*';
}

function sendJSON(res, statusCode, data) {
  res.writeHead(statusCode, {
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': getCORSOrigin(),
    'Access-Control-Allow-Methods': 'GET, POST, OPTIONS',
    'Access-Control-Allow-Headers': 'Content-Type, Authorization'
  });
  res.end(JSON.stringify(data));
}

function parseBody(req) {
  return new Promise((resolve, reject) => {
    let body = '';
    req.on('data', chunk => { body += chunk; });
    req.on('end', () => {
      try {
        resolve(body ? JSON.parse(body) : {});
      } catch (err) {
        reject(err);
      }
    });
    req.on('error', reject);
  });
}

const server = http.createServer(async (req, res) => {
  const parsedUrl = new URL(req.url, `http://${req.headers.host || 'localhost'}`);
  const path = parsedUrl.pathname;
  const method = req.method;

  // Handle CORS Preflight
  if (method === 'OPTIONS') {
    res.writeHead(204, {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Methods': 'GET, POST, OPTIONS',
      'Access-Control-Allow-Headers': 'Content-Type, Authorization'
    });
    res.end();
    return;
  }

  try {
    // 1. Health check
    if (path === '/health' && method === 'GET') {
      return sendJSON(res, 200, {
        status: 'healthy',
        service: 'Warehouse Express 3D Game Server (Node.js Runtime)',
        timestamp: Math.floor(Date.now() / 1000)
      });
    }

    // 2. Vehicle Catalogue
    if (path === '/api/v1/vehicles' && method === 'GET') {
      return sendJSON(res, 200, state.vehicles);
    }

    // 3. Dispatch Target House (Anti-duplicate)
    if (path === '/api/v1/dispatch/next' && method === 'POST') {
      const body = await parseBody(req);
      const targetHouseId = nextTargetHouse(body.houseIds);
      return sendJSON(res, 200, {
        targetHouseId,
        message: 'Target dispatch generated via Anti-Duplicate Dispatcher'
      });
    }

    // 4. Telemetry Anti-Cheat Verification
    if (path === '/api/v1/telemetry/verify' && method === 'POST') {
      const body = await parseBody(req);
      const result = validatePlayerStep(
        body.prevX,
        body.prevZ,
        body.nextX,
        body.nextZ,
        body.deltaMs || 16.67,
        body.surface || 'asphalt'
      );
      return sendJSON(res, 200, result);
    }

    // 5. Save Player Profile
    if (path === '/api/v1/player/save' && method === 'POST') {
      const body = await parseBody(req);
      const playerId = body.playerId || 'player_default';
      const profile = {
        ...body,
        playerId,
        updatedTimestamp: Math.floor(Date.now() / 1000)
      };
      state.profiles.set(playerId, profile);
      return sendJSON(res, 200, {
        success: true,
        message: 'Player profile saved successfully',
        data: profile
      });
    }

    // 6. Load Player Profile
    if (path === '/api/v1/player/load' && method === 'GET') {
      const playerId = parsedUrl.searchParams.get('playerId') || 'player_default';
      if (!state.profiles.has(playerId)) {
        return sendJSON(res, 200, {
          exists: false,
          message: 'Profile not found'
        });
      }
      return sendJSON(res, 200, {
        exists: true,
        data: state.profiles.get(playerId)
      });
    }

    // 7. Leaderboard
    if (path === '/api/v1/leaderboard' && method === 'GET') {
      return sendJSON(res, 200, state.leaderboard);
    }

    // 404 Route Not Found
    return sendJSON(res, 404, { error: 'Endpoint not found' });
  } catch (err) {
    return sendJSON(res, 500, { error: 'Internal Server Error', message: err.message });
  }
});

server.listen(PORT, () => {
  console.log('=====================================================');
  console.log(` Warehouse Express 3D - Game Server (Node.js)`);
  console.log(` Running on http://localhost:${PORT}`);
  console.log(` Health check: http://localhost:${PORT}/health`);
  console.log('=====================================================');
});
