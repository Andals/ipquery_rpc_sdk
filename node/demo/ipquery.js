var ipquery = require('../ipquery.js');

var client = new ipquery.Client('localhost:9090');

client.ipqueryVersion(function(err, data) {
    console.log('[ipqueryVersion()]');
    console.log(data);
});

client.ipquery("61.189.127.109", "prov_city", function(err, data) {
    console.log('[ipquery()]');
    console.log(data);
});

var ipList = [
    "60.25.71.203",
    "110.231.68.79",
    "113.12.144.23",    
];

client.ipqueryEx(ipList, "prov_city", function(err, data) {
    console.log('[ipqueryEx()]');
    data.forEach(function (item) {
        console.log(item);
    });
    client.close();
});






