#include "IpqueryClient.h"  // As an example

bool IpqueryClient::connect(const string& host, const int &port)
{
    if (client_ != NULL) {
        delete client_;
    }
    boost::shared_ptr<TSocket> socket(new TSocket(host, port));
    boost::shared_ptr<TTransport> transport(new TFramedTransport(socket));
    boost::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));
    client_ = new ipqueryClient(protocol);
    transport->open();
    return true;
}

bool IpqueryClient::connect(const string& path)
{
    if (client_ != NULL) {
        delete client_;
    }

    boost::shared_ptr<TSocket> socket(new TSocket(path));
    boost::shared_ptr<TTransport> transport(new TFramedTransport(socket));
    boost::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));

    client_ = new ipqueryClient(protocol);
    transport->open();
    return true;
}

bool IpqueryClient::ipquery(std::string& _return, const std::string& ip, const std::string& tag) const
{
    if (client_ == NULL) {
        return false;
    }
    client_->ipquery(_return, ip, tag);
    return true;
}

bool IpqueryClient::ipqueryEx(std::vector<std::string>& _return, const std::vector<std::string>& ipList, const std::string& tag) const
{
    if (client_ == NULL) {
        return false;
    }
    client_->ipqueryEx(_return, ipList, tag);
    return true;
}

bool IpqueryClient::ipqueryVersion(std::string& _return) const
{
    if (client_ == NULL) {
        return false;
    }
    client_->ipqueryVersion(_return);
    return true;
}
