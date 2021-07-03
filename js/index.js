var machineIdSync = require('node-machine-id').machineIdSync

// get machine id
var machineId = machineIdSync(true)
console.log(machineId)
// get current time unix
var time = Date.now || function() {
    return +new Date;
};
var unix = Math.round(+new Date()/1000);
console.log(unix)
// encrypt message
var message = `${machineId}+${unix}`

console.log('----')
console.log(message)

