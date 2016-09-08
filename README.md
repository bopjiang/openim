# openIM

an open source IM (instance message) system.


## Architecture
![openIM architecture](http://i4.buimg.com/573902/4b6715129b0dc21e.png "openIM architecture")

## Modules

### Connection Layer

#### TCPConnector
openIM binary protocol client

#### UDPConnec/tor
openIM binary protocol client

#### XMPPConnector
standard XMPP client

#### MQTTConnector
standard MQTT client

#### WSConnector
websocket client using openIM binary protocol

#### HTTP2Connector
HTTP2 client using openIM binary protocol


### Disctribute Layer
#### Router
* route logic message by message type
* maintain location database of all clients, a distributed kv store cluster

#### Messenger
dispatch client message to each single user, and save the message when client is offline

#### PushProxy
* maintain remote push token(APNs, Google token)
* push message to remote push service (APNs, GCM)

#### GroupMan
handle group message distribution

#### XMPPGateway
implement XMPP Server2Server(s2s) protocol

### Data Layer
#### DBProxy
database proxy, provide API to access database

#### Database
database to be supported:
* Redis
* MySQL
* PostgreSQL
* ElasticSearch
