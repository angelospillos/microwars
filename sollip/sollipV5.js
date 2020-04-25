const { Worker, isMainThread } = require('worker_threads');
var cluster = require('cluster');
var numCPUs = require('os').cpus().length;
var processes = 3;
var threads = 3;

if (cluster.isMaster) {

  for (var i = 0; i < numCPUs; i++) {
    for (var i = 0; i < processes; i++) {
      cluster.fork();
    }
  }

  cluster.on('exit', function (worker, code, signal) {
    console.log('worker ' + worker.process.pid + ' died');
  });

} else {
  if (isMainThread) {
    for (var i = 0; i < threads; i++) {
      new Worker("./sollipV4.js");
    }
  } else {
    require("./sollipV4.js");
  }
}
