# openIM


[![Build Status](https://travis-ci.org/openim/openim.svg?branch=master)](https://travis-ci.org/openim/openim) [![Coverage Status](https://coveralls.io/repos/github/openim/openim/badge.svg?branch=master)](https://coveralls.io/github/openim/openim?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/openim/openim)](https://goreportcard.com/report/github.com/openim/openim)


an open source IM (instance message) system.
We are using slack to discuss projects. Find us at: https://openim.slack.com/

## Motivation
* XMPP is complicated, may be we could simplify it.
* Communication should be safe for everyone. Nobody could sniff the information content. The communication behavior (to whom you send message) should also be protected. 
* Instance Message is just a channel for communication, could be used for man-man, machine-to-machine communicating.
* Web should be open, not closed. IM is a part the Web.

That's why we need an open source IM project, we provides server side and client side implementation. So everybody could build a IM cluster even in your living-room.

The whole is an open, distributed, decentralized, XMPP-compatible, End-2-End encryption IM system.

## Develop Document

* https://github.com/openim/openim/wiki

* http://openim.readthedocs.io/   

Wiki first. May move to readthedocs later when API is stable


## Architecture
![openIM architecture](http://i4.buimg.com/573902/4b6715129b0dc21e.png "openIM architecture")

## Modules

### Connection Layer

#### TCPConnector
openIM binary protocol client

#### UDPConnector
openIM binary protocol client

#### XMPPConnector
standard XMPP client

#### MQTTConnector
standard MQTT client

#### WSConnector
websocket client using openIM binary protocol

#### HTTP2Connector
HTTP2 client using openIM binary protocol


### Distribution Layer
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


## Contributing
Contributions are welcomed and greatly appreciated. See [CONTRIBUTING.md](CONTRIBUTING.md)
for details
