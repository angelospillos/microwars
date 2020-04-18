const uuid = require('uuid-random');
const fs = require('fs');

const serverHostname = '127.0.0.1';
const serverPort = 3000;
const serverLog = fs.createWriteStream('log.txt', { flags: 'a' });

const statusPath = '/status';
const testPath = '/test';
const combatPath = '/combat';
const jabPath = '/jab';
const crossPath = '/cross';
const hookPath = '/hook';
const uppercutPath = '/uppercut';

const opponentUrl = 'http://127.0.0.1:3001'
const opponentStatusUrl = opponentUrl + statusPath;
const opponentJabUrl = opponentUrl + jabPath;
const opponentCrossUrl = opponentUrl + crossPath;
const opponentHookUrl = opponentUrl + hookPath;
const opponentUppercutUrl = opponentUrl + uppercutPath;

require('uWebSockets.js').App()
  .get(statusPath, (res, req) => {
    ok(req, res);
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
  res.setHeader('Content-Type', 'application/json');
  res.write('{ "message": "Hello World" }');
  res.end();
  executeOpponentJab();
  executeOpponentJab();
}

async function cross(req, res) {
  let uuid = uuid();
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  res.write('{ "uuid": "' + uuid + '" }');
  res.end();
  executeOpponentJab();
  executeOpponentJab();
  executeOpponentCross();
  log(uuid);
}

async function hook(req, res) {
  let uuid = uuid();
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  res.write('{ "uuid": "' + uuid + '" }');
  res.end();
  executeOpponentHook();
  executeOpponentHook();
  executeOpponentUppercut();
  log(uuid);
}

async function uppercut(req, res) {
  fib(20).then((fibonacci) => {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json');
    res.write('{ "fibonacci": "' + fibonacci + '" }');
    res.end();
    log(fibonacci);
  });
  executeOpponentCross();
  executeOpponentHook();
  executeOpponentUppercut();
}

async function executeOpponentStatus() {
  http.get(opponentStatusUrl, (response) => {
    let data = '';
    response.on('data', (chunk) => {
      data += chunk;
    });
    response.on('end', () => {
      log(data);
    });
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentJab() {
  http.get(opponentJabUrl, (response) => {
    let data = '';
    response.on('data', (chunk) => {
      data += chunk;
    });
    response.on('end', () => {
      log(data);
    });
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentCross() {
  http.get(opponentCrossUrl, (response) => {
    let data = '';
    response.on('data', (chunk) => {
      data += chunk;
    });
    response.on('end', () => {
      log(data);
    });
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentHook() {
  http.get(opponentHookUrl, (response) => {
    let data = '';
    response.on('data', (chunk) => {
      data += chunk;
    });
    response.on('end', () => {
      log(data);
    });
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function executeOpponentUppercut() {
  http.get(opponentUppercutUrl, (response) => {
    let data = '';
    response.on('data', (chunk) => {
      data += chunk;
    });
    response.on('end', () => {
      log(data);
    });
  }).on("error", (error) => {
    log("Error: " + error.message);
  });
}

async function log(message) {
  serverLog.write(message);
}

async function fib(n, sum = 0, prev = 1) {
  if (n <= 1) return sum;
  return fib(n - 1, prev + sum, sum);
}
