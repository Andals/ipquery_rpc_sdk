#ifndef _IPQUERY_CLIENT_H_
#define _IPQUERY_CLIENT_H_
#include "ipquery.h"
 
#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>

#include <string>
#include <vector>
using namespace std;
 
using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;
 
using namespace ipquery;

class IpqueryClient
{
public:
    IpqueryClient()
    {
        client_ = NULL;
    }
    ~IpqueryClient()
    {
        if (client_ != NULL) {
            delete client_;
            client_ = NULL;
        }
    }
    bool connect(const string& host, const int &port);
    bool connect(const string& path);
    bool close();

    bool ipquery(std::string& _return, const std::string& ip, const std::string& tag) const;
    bool ipqueryEx(std::vector<std::string>& _return, const std::vector<std::string> & ipList, const std::string& tag) const;
    bool ipqueryVersion(std::string& _return) const;
private:
    ipqueryClient* client_;
};

#endif
