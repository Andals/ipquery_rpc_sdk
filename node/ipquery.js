var util = require('util');
var thrift = require('thrift');
var ipquery = require('./gen-nodejs/ipquery');
var ttypes = require('./gen-nodejs/ipquery_types');

var Client = function(host) {
    var pair = host.split(':');
    this.host = pair[0];
    this.port = pair[1];
    this.connect();
}

util.inherits(Client, process.EventEmitter);

Client.prototype.connect = function() {
    this.ready = false;
    this.queue = [];

    this.connection = thrift.createConnection(this.host, this.port, {
        transport: thrift.TFramedTransport,
        protocol: thrift.TBinaryProtocol        
    });

    this.connection.on('error', function(err) {
        this.emit('error', err);
    });

    this.client = thrift.createClient(ipquery, this.connection);

    var self = this;
    this.connection.on('connect', function(err) {
        if (err) {
            self.emit('error', err);
            return;
        }

        self.ready = true;
        self.dispatch();
    });
};

Client.prototype.ipquery = function(ip, tag, callback) {
    var args = Array.prototype.slice.call(arguments);
    if (!this.ready) {
        this.queue.push([arguments.callee, args]);
        return;
    }
    this.client.ipquery(ip, tag, callback);
};

Client.prototype.ipqueryEx = function(ipList, tag, callback) {
    var args = Array.prototype.slice.call(arguments);
    if (!this.ready) {
        this.queue.push([arguments.callee, args]);
        return;
    }
    this.client.ipqueryEx(ipList, tag, callback);
};

Client.prototype.ipqueryVersion = function() {
    var args = Array.prototype.slice.call(arguments);
    if (!this.ready) {
        this.queue.push([arguments.callee, args]);
        return;
    }

    var callback;
    if (typeof args[args.length - 1] === 'function') {
        callback = args.pop();
    }

    this.client.ipqueryVersion(callback);
};

Client.prototype.close = function() {
    this.connection.end();
};

Client.prototype.dispatch = function() {
    if (this.ready) {
        if (this.queue.length > 0) {
            var next = this.queue.shift();
            next[0].apply(this, next[1]);
            this.dispatch();
        }
    }
};

exports.Client = Client;


