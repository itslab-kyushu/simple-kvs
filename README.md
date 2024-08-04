# Simple Key-Value Storage
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master)
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/simple-kvs/releases/tag/v0.2.0)
[![Dockerhub](https://img.shields.io/badge/dockerhub-itslabq%2Fsimple--kvs-blue.svg)](https://hub.docker.com/r/itslabq/simple-kvs/)

This software is a simple key-value storage implemented by
[Go](https://golang.org/) using [gRPC](http://www.grpc.io/).
This software has been made for comparing performance of secret sharing based
key-value storages in the following article:

* [Hiroaki Anada](https://sites.google.com/view/lab-hiroaki-anada/),
  [Junpei Kawamoto](https://www.jkawamoto.info),
  Chenyutao Ke,
  [Kirill Morozov](https://engineering.unt.edu/people/kirill-morozov.html), and
  [Kouichi Sakurai](http://itslab.inf.kyushu-u.ac.jp/~sakurai/),
  "[Cross-Group Secret Sharing Scheme for Secure Usage of Cloud Storage over Different Providers and Regions](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/article/10.1007%2Fs11227-017-2009-7),"
  [The Journal of Supercomputing](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/journal/11227), 2017.

Please consider to refer it, if you will publish any articles using this
software.

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
