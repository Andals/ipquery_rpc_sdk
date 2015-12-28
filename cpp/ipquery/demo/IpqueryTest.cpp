#include "../IpqueryClient.h"

#include <iostream>
#include <vector>
using namespace std;

void test(const IpqueryClient &client)
{
    string result;
    client.ipquery(result, "223.103.15.1", "prov_city");
    cout<<"223.103.15.1 result: "<<result<<endl;

    vector<string> ips, resultVector;
    ips.push_back("223.103.15.1");
    ips.push_back("58.30.15.255");

    client.ipqueryEx(resultVector, ips, "prov_city");
    for (vector<string>::iterator iter = resultVector.begin(); iter != resultVector.end(); iter++) {
        cout<<*iter<<endl;
    }

    string version;
    client.ipqueryVersion(version);
    cout<<"ipquery version is: "<<version<<endl;

}

int main(int argc, char **argv)
{
    IpqueryClient tcpClient;
    tcpClient.connect("localhost", 9090);
    test(tcpClient);

    IpqueryClient domainSocketClient;
    domainSocketClient.connect("/home/haoyankai/devspace/ipquery_rpc/var/ipquery.sock");
    test(domainSocketClient);
}

