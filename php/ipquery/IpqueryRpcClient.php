<?php

namespace ipquery;

define("THRIFT_LIB_DIR", dirname(__FILE__) . "/lib/");
define("IPQUERY_LIB_DIR", dirname(__FILE__));

require_once THRIFT_LIB_DIR . 'Thrift/ClassLoader/ThriftClassLoader.php';

require_once IPQUERY_LIB_DIR . '/Types.php';
require_once IPQUERY_LIB_DIR . '/ipquery.php';

use Thrift\ClassLoader\ThriftClassLoader;

$GEN_DIR = realpath(dirname(__FILE__));


$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift', THRIFT_LIB_DIR);
$loader->register();

use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\THttpClient;
use Thrift\Transport\TFramedTransport;
use Thrift\Exception\TException;

class IpqueryRpcClient
{
    private $client = null;
    private function init($socket)
    {
        $transport = new TFramedTransport($socket, 1024, 1024);
        $protocol = new TBinaryProtocol($transport);
        $this->client = new \ipqueryClient($protocol);
        $transport->open();
    }

    public function tcpConnect($host, $port)
    {
        $socket = new TSocket($host, $port);
        $this->init($socket);
    }

    public function ipquery($ip, $tag)
    {
        if ($this->client === null) {
            return false;
        }
        return $this->client->ipquery($ip, $tag);
    }

    public function ipqueryEx($ipList, $tag)
    {
        if ($this->client === null) {
            return false;
        }
        return $this->client->ipqueryEx($ipList, $tag);
    }

    public function ipqueryVersion()
    {
        if ($this->client === null) {
            return false;
        }
        return $this->client->ipqueryVersion();
    }
}
