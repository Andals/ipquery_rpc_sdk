#!/usr/local/thrift/bin/thrift --gen cpp
 
namespace cpp ipquery

service ipquery {
  string ipquery(1:string ip, 2:string tag)
  list<string> ipqueryEx(1:list<string> ipList, 2:string tag)
  string ipqueryVersion()
}
