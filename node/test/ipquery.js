var ipquery = require('../ipquery.js');
var should = require('should');
var client = new ipquery.Client('localhost:9090');
describe('ipqueryClient', function() {
    describe('- ipqueryVersion()', function() {
        it('should return the ipquery version 1.0', function (done) {
            client.ipqueryVersion(function(err, data) {
                data.should.equal('1.0')
                done()
            });
        });
    });

    describe('- ipquery(ip, tag)', function() {
        it('should return "辽宁\t大连\t大连\t101070201联通 when call the function ipquery("61.189.127.109", "prov_city")', function (done) {
            client.ipquery("61.189.127.109", "prov_city", function(err, data) {
                data.should.equal('辽宁\t大连\t大连\t101070201联通')
                done()
            });
        });
    });    

    describe('- ipqueryEx(ipList, tag)', function() {
        it('should return ["天津\t天津\t西青\t101030500联通","河北\t唐山\t唐山\t101090501联通","广西\t南宁\t南宁\t101300101电信"] when call the function ipquery(["60.25.71.203", "110.231.68.79", "113.12.144.23"], "prov_city")', function (done) {
            var ipList = [
                "60.25.71.203",
                "110.231.68.79",
                "113.12.144.23",    
            ];            
            client.ipqueryEx(ipList, "prov_city", function(err, data) {
                data = JSON.stringify(data);
                data.should.equal('["天津\\t天津\\t西青\\t101030500联通","河北\\t唐山\\t唐山\\t101090501联通","广西\\t南宁\\t南宁\\t101300101电信"]');
                done()
            });
        });
    });      
});
