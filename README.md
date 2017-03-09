# Simple Key-Value Storage
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master)
[![wercker status](https://app.wercker.com/status/717adbfffa215daf21462bfa273a5a16/s/master "wercker status")](https://app.wercker.com/project/byKey/717adbfffa215daf21462bfa273a5a16)
[![Release](https://img.shields.io/badge/release-0.1.0-brightgreen.svg)](https://github.com/itslab-kyushu/simple-kvs/releases/tag/v0.1.0)
[![Dockerhub](https://img.shields.io/badge/dockerhub-itslabq%2Fsimple--kvs-blue.svg)](https://hub.docker.com/r/itslabq/simple-kvs/)
[![MicroBadger](https://images.microbadger.com/badges/image/itslabq/simple-kvs.svg)](https://microbadger.com/images/itslabq/simple-kvs)

This software is a simple key-value storage implemented by
[Go](https://golang.org/) using [gRPC](http://www.grpc.io/).
This software is for comparing performance of secret sharing based key-value
storages.


## Installation
Compiled binaries of them are available on
[Github](https://github.com/itslab-kyushu/simple-kvs/releases).
After downloading a binary to your environment,
decompress and put it in a directory included in `$PATH`.

If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```sh
$ brew tap itslab-kyushu/simple-kvs
$ brew install simple-kvs
```

You can also compile by yourself.
First, you need to download the code base:

```sh
$ git clone https://github.com/itslab-kyushu/simple-kvs $GOPATH/src/itslab-kyushu/simple-kvs
```

Then, build client command `kvs`:

```sh
$ cd $GOPATH/src/itslab-kyushu/simple-kvs/client
$ go get -d -t -v .
$ go build -o kvs
```

and build server command `kvs-server`:

```sh
$ cd $GOPATH/src/itslab-kyushu/simple-kvs/server
$ go get -d -t -v .
$ go build -o kvs-server
```

To build both commands, [Go](https://golang.org/) > 1.7.4 is required.

## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
