const http = require('http');
const uuid = require('uuid');
const fs = require('fs');
const fibonator = require('fibonator');

const serverHostname = '127.0.0.1';
const serverPort = 8000;
const serverLog = fs.createWriteStream('log.txt', { flags: 'a' });

const statusPath = '/status';
const testPath = '/test';
const combatPath = '/combat';
const jabPath = '/jab';
const crossPath = '/cross';
const hookPath = '/hook';
const uppercutPath = '/uppercut';

const DEFAULT_RESP = '{ "status": "ok" }';
const TIMEOUT = parseInt(process.env.TIMEOUT, 10) || 4000;
const opponentAddr = process.env.OPPONENT_ADDR || '127.0.0.1';

const refereeStatusRequest = {
  host: 'enqfc8y2t9fo.x.pipedream.net',
  port: 80,
  method: 'GET',
  path: '/alive',
  timeout: TIMEOUT,
  headers: {
    'user-agent': 'sollipV5',
  }
}

const refereeCheckRequest = {
  host: 'enqfc8y2t9fo.x.pipedream.net',
  port: 80,
  method: 'GET',
  path: '/introduce',
  timeout: TIMEOUT,
  headers: {
    'user-agent': 'sollipV5',
  }
}

const refereeKnockoutRequest = {
  host: 'enqfc8y2t9fo.x.pipedream.net',
  port: 80,
  method: 'GET',
  path: '/knockout',
  timeout: TIMEOUT,
  headers: {
    'user-agent': 'sollipV5',
  }
}

const opponentStatusRequest = {
  host: opponentAddr,
  port: 8000,
  method: 'GET',
  path: statusPath,
  timeout: TIMEOUT,
}

const opponentJabRequest = {
  host: opponentAddr,
  port: 8000,
  method: 'GET',
  path: jabPath,
  timeout: TIMEOUT,
}

const opponentCrossRequest = {
  host: opponentAddr,
  port: 8000,
  method: 'GET',
  path: crossPath,
  timeout: TIMEOUT,
}
const opponentHookRequest = {
  host: opponentAddr,
  port: 8000,
  method: 'GET',
  path: hookPath,
  timeout: TIMEOUT,
}

const opponentUppercutRequest = {
  host: opponentAddr,
  port: 8000,
  method: 'GET',
  path: uppercutPath,
  timeout: TIMEOUT,
}

// warmup the engine 
fibonator.fib(8);
fibonator.fibrec(8);

require('uWebSockets.js').App()
  .get(statusPath, (res, req) => {
    ok(req, res);
  })
  .get("/work", (res, req) => {
    res.statusCode = 200;
    res.end(`{ "uuid": "${uuid.v4()}", "fib": "${fibonator.fibrec(16)}" }`);
  })
  .get(testPath, (res, req) => {
    test(req, res);
  })
  .get(combatPath, (res, req) => {
    combat(req, res);
  })
  .get(jabPath, (res, req) => {
    jab(req, res);
  })
  .get(crossPath, (res, req) => {
    cross(req, res);
  })
  .get(hookPath, (res, req) => {
    hook(req, res);
  })
  .get(uppercutPath, (res, req) => {
    uppercut(req, res);
  })
  .any('/*', (res, req) => {
    res.statusCode = 404;
    res.writeHeader('Content-Type', 'application/json');
    res.end('{ "status": "not_found" }');
  })
  .listen(serverPort, (token) => {
    if (token) {
      console.log(`Server running at http://${serverHostname}:${serverPort}/`);
    } else {
      console.log(`Server NOT running at http://${serverHostname}:${serverPort}/`);
    }
  });

async function ok(req, res) {
  res.statusCode = 200;
  res.end(DEFAULT_RESP);
  http.get(refereeStatusRequest);
}

async function test(req, res) {
  executeOpponentStatus();
  res.statusCode = 200;
  res.end(DEFAULT_RESP);
}

async function combat(req, res) {
  executeOpponentJab();
  executeOpponentHook();
  res.statusCode = 200;
  res.end(DEFAULT_RESP);
}

async function jab(req, res) {
  res.statusCode = 200;
  res.end(`{ "uuid": "${uuid.v4()}", "fib": ${fibonator.fibrec(2)} }`);
  executeOpponentJab();
  executeOpponentJab();
}

async function cross(req, res) {
  res.statusCode = 200;
  res.end(`{ "uuid": "${uuid.v4()}", "fib": ${fibonator.fibrec(4)} }`);
  executeOpponentJab();
  executeOpponentJab();
  executeOpponentCross();
}

async function hook(req, res) {
  res.statusCode = 200;
  res.end(`{ "uuid": "${uuid.v4()}", "fib": ${fibonator.fibrec(8)} }`);
  executeOpponentHook();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function uppercut(req, res) {
  res.statusCode = 200;
  res.end(`{ "uuid": "${uuid.v4()}", "fib": ${fibonator.fibrec(16)} }`);
  executeOpponentCross();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function executeOpponentStatus() {
  http.get(opponentStatusRequest, (response) => {
    const { statusCode } = response;
    if (statusCode !== 200) {
      log("Error: " + error.message);
    } else{
      http.get(refereeCheckRequest);
    }
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentJab() {
  http.get(opponentJabRequest, (response) => {
    const { statusCode } = response;
    if (statusCode !== 200) {
      log("Error: " + error.message);
    }
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentCross() {
  http.get(opponentCrossRequest, (response) => {
    const { statusCode } = response;
    if (statusCode !== 200) {
      log("Error: " + error.message);
    }
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentHook() {
  http.get(opponentHookRequest, (response) => {
    const { statusCode } = response;
    if (statusCode !== 200) {
      log("Error: " + error.message);
    }
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentUppercut() {
  http.get(opponentUppercutRequest, (response) => {
    const { statusCode } = response;
    if (statusCode !== 200) {
      log("Error: " + error.message);
    }
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function log(message) {
  http.get(refereeKnockoutRequest);
  serverLog.write(`${message}\n`, () => { });
}