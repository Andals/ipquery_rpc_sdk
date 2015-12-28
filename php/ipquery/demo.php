<?php
require_once "IpqueryRpcClient.php";

$tcpClient = new \Ipquery\IpqueryRpcClient();
$tcpClient->tcpConnect("localhost", 9090);
test($tcpClient);

function test($client)
{
    $result = $client->ipquery("61.189.127.109", "prov_city");
    echo "61.189.127.109 result is: $result\n";

    $ipList = array(
        "60.25.71.203",
        "110.231.68.79",
        "113.12.144.23",
    );
    $resultList = $client->ipqueryEx($ipList, "prov_city");
    foreach ($resultList as $value) {
        echo "$value\n";
    }

    $version = $client->ipqueryVersion();
    echo "ipquery version is $version\n";
}
