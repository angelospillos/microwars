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

const TIMEOUT = parseInt(process.env.TIMEOUT, 10) || 4000;
const opponentAddr = process.env.OPPONENT_ADDR || '127.0.0.1';

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
    res.write(`{ "uuid": "${uuid.v4()}", "fib": "${fibonator.fibrec(8)}" }`);
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
  res.writeHeader('Content-Type', 'application/json');
  res.write('{ "status": "ok" }');
  res.end();
}

async function test(req, res) {
  executeOpponentStatus();
}

async function combat(req, res) {
  executeOpponentJab();
  executeOpponentHook();
}

async function jab(req, res) {
  res.statusCode = 200;
  // no gains from native function on less than fib(8)
  res.write(`{ "uuid": "${uuid.v4()}", "fib": "${slowfib(2)}" }`);
  res.end();
  executeOpponentJab();
  executeOpponentJab();
}

async function cross(req, res) {
  res.statusCode = 200;
  // no gains from native function on less than fib(8)
  res.write(`{ "uuid": "${uuid.v4()}", "fib": "${slowfib(4)}" }`);
  res.end();
  executeOpponentJab();
  executeOpponentJab();
  executeOpponentCross();
}

async function hook(req, res) {
  let uuid = uuid();
  res.statusCode = 200;
  // no gains from native function on less than fib(8)
  // res.write(`{ "uuid": "${uuid.v4()}", "fib": "${slowfib(8)}" }`);
  res.write(`{ "uuid": "${uuid.v4()}", "fib": "${fibonator.fibrec(8)}" }`);
  res.end();
  executeOpponentHook();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function uppercut(req, res) {
  fib(20).then((fibonacci) => {
    res.statusCode = 200;
    res.write(`{ "uuid": "${uuid.v4()}", "fib": "${fibonator.fibrec(16)}" }`);
    res.end();
    log(fibonacci);
  });
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
  serverLog.write(message);
}

function slowfib(n) {
  if (n <= 1) return n;
  return slowfib(n - 1) + slowfib(n - 2);
}