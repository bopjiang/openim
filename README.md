# openIM

an open source IM (instance message) system


## Architecture
![openIM architecture](http://i1.piimg.com/4851/af516f10b57324f0.png "openIM architecture")

## Modules

### Connection/Access Layer
#### TCPConncetor
#### UDPConncetor
#### XMPPConncetor
#### MQTTConncetor
#### WSConncetor
support websocket client

#### HTTP2Conncetor

### Disctribute Layer
#### Router
* route logic message by message type
* maintain location database of all clients, is a distributed kv store cluster

### Logic Layer
#### Message
dispatch client message, and save message when client not online

#### PushProxy
* maintain remote push token(APNs, Google token)
* push message to remote push service (APNs, GCM)

### Data Layer
#### DBProxy
database proxy, support API to access database
#### Database
supported Database list:
* Redis
* MySQL
* PostgreSQL
* ElasticSearch
