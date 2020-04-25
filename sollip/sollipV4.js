const http = require('http');
const uuid = require('uuid');
const fs = require('fs');
const fibonator = require('fibonator');
const querystring = require('querystring');

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
  host: 'http://enqfc8y2t9fo.x.pipedream.net',
  port: 80,
  method: 'GET',
  path: '',
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
slowfib(8);
fibonator.fib(8);
fibonator.fibrec(8);

require('uWebSockets.js').App()
  .get(statusPath, (res, req) => {
    ok(req, res);
  })
  .get("/work", (res, req) => {
    res.statusCode = 200;
    res.write(`{ "uuid": "${uuid.v4()}", "fib": "${fibonator.fibrec(16)}" }`);
    res.end();
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
    res.write('{ "status": "not_found" }');
    res.end();
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
  res.write(DEFAULT_RESP);
  res.end();
  log('status');
}

async function test(req, res) {
  executeOpponentStatus();
  res.statusCode = 200;
  res.write(DEFAULT_RESP);
  res.end();
}

async function combat(req, res) {
  executeOpponentJab();
  executeOpponentHook();
  res.statusCode = 200;
  res.write(DEFAULT_RESP);
  res.end();
}

async function jab(req, res) {
  // no gains from native function on less than fib(8)
  res.statusCode = 200;
  res.write(`{ "uuid": "${uuid.v4()}", "fib": ${slowfib(2)} }`);
  res.end();
  executeOpponentJab();
  executeOpponentJab();
}

async function cross(req, res) {
  // no gains from native function on less than fib(8)
  res.statusCode = 200;
  res.write(`{ "uuid": "${uuid.v4()}", "fib": ${slowfib(4)} }`);
  res.end();
  executeOpponentJab();
  executeOpponentJab();
  executeOpponentCross();
}

async function hook(req, res) {
  // no gains from native function on less than fib(8)
  res.statusCode = 200;
  res.write(`{ "uuid": "${uuid.v4()}", "fib": ${slowfib(8)} }`);
  res.end();
  executeOpponentHook();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function uppercut(req, res) {
  res.statusCode = 200;
  res.write(`{ "uuid": "${uuid.v4()}", "fib": ${fibonator.fibrec(16)} }`);
  res.end();
  executeOpponentCross();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function executeOpponentStatus() {
  http.get(opponentStatusRequest, (response) => {
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentJab() {
  http.get(opponentJabRequest, (response) => {
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentCross() {
  http.get(opponentCrossRequest, (response) => {
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentHook() {
  http.get(opponentHookRequest, (response) => {
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentUppercut() {
  http.get(opponentUppercutRequest, (response) => {
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function log(message) {
  http.get(refereeStatusRequest, (response) => {});
  serverLog.write(`${message}\n`, () => {});
}

function slowfib(n) {
  if (n <= 1) return n;
  return slowfib(n - 1) + slowfib(n - 2);
}