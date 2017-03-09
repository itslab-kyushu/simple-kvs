---
title: Server application
menu: main
date: 2017-03-08
lastmod: 2017-03-08
weight: 20
description: >-
  Server application provides a key-value datastore service based on gRPC.
---
[![Release](https://img.shields.io/badge/release-0.1.0-brightgreen.svg)](https://github.com/itslab-kyushu/simple-kvs/releases/tag/v0.1.0)
[![Dockerhub](https://img.shields.io/badge/dockerhub-itslabq%2Fsimple--kvs-blue.svg)](https://hub.docker.com/r/itslabq/simple-kvs/)
[![MicroBadger](https://images.microbadger.com/badges/image/itslabq/simple-kvs.svg)](https://microbadger.com/images/itslabq/simple-kvs)

## Summary
Server application provides a key-value datastore service based on [gRPC](http://www.grpc.io/).
You can find the service definition in [here](https://github.com/itslab-kyushu/simple-kvs/blob/master/kvs/kvs.proto).

[![gRPC](../img/grpc.png)](http://www.grpc.io/)

## Usage
The server application supports the following flags,

* `--port`: the port number the server will listen,
* `--root`: the document root path to store uploaded shares,
* `--no-compress`: if set, all shares will be stored without compression.

If those flags are omitted, default values are used.
Thus, you can start a server by just run `kvs-server`.

## Docker support
A docker image to run the server, `itslabq/simple-kvs`, is also available.

[![docker](../img/small_h-trans.png)](https://www.docker.com/)

Containers based on this image expose port 13009 and store uploaded shares into
`/data`. The following command runs a container with mapping local port 13009
and `$(pwd)/data` to the container's port 13009 and `/data`:

```shell
$ docker run -it -p 13009:13009 -v $(pwd)/data:/data itslabq/simple-kvs
```

## Installation
Compiled binaries of the server application `kvs-server` are available on
[Github](https://github.com/itslab-kyushu/simple-kvs/releases).
To use these binaries, after downloading a binary to your environment, decompress and put it in a directory included in your $PATH.

You can also compile by yourself.
First, you need to download the code base:

```shell
$ git clone https://github.com/itslab-kyushu/simple-kvs $GOPATH/src/itslab-kyushu/simple-kvs
```

Then, build server command `kvs-server`:

```shell
$ cd $GOPATH/src/itslab-kyushu/simple-kvs/server
$ go get -d -t -v .
$ go build -o kvs-server
```

To build it, [Go](https://golang.org/) > 1.7.4 is required.
