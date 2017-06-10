---
title: Client application
menu: main
date: 2017-03-08
lastmod: 2017-03-08
weight: 15
description: >-
  Client application provides an interface to a client/server style data storage
  service. This application has four commands:
  get is downloading a file from a server,
  put is uploading a file to the server,
  delete is deleting a file from the server,
  and list is showing names of files stored in the server.
---
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/simple-kvs/releases/tag/v0.2.0)

## Summary
Client application provides an interface to a client/server style data storage
service.
This application has four commands:

* get: download a file from a server,
* put: upload a file to the server,
* delete: delete a file from the server,
* list: show names of files stored in the server,

and all commands connect a set of [data storage servers](../server).

To specify address information of those servers, all commands receives a
configuration file in YAML.
The configuration file has one root element `servers` and it takes a list of
server information, which is an object consisting of two element `address` and
`port`.
The following example defines one server in a same network:

```yaml
servers:
  - address: 192.168.0.1
    port: 13009
```

The default name of the configuration file is `kvs.yml` but you can set another
name via `--config` flag.

## Get command
```shell
$ kvs remote get --config kvs.yml --output result.dat <file name>
```

Get command download the given file name from the server
defined in the configuration file, and stores it as the given file name via
`--output` flag.

If `--config` flag is omitted, `kvs.yml` is used, and if `--output` flag is
omitted, `<file name>` is used.

To find available file names, use list command.


## Put command
```shell
$ kvs remote put --config kvs.yml <file>
```

Put command reads the given file and uploads it to the server defined in the
configuration file.

If `--config` flag is omitted, `kvs.yml` is used.

## Delete command
```shell
$ kvs remote delete --config kvs.yml <file name>
```

Delete command deletes all shares associated with the given file name from all
servers defined in the configuration file.

If `--config` flag is omitted, `kvs.yml` is used.

## List command
```shell
$ sss remote list --config kvs.yml
```

List command shows all file names stored in the servers.
If `--config` flag is omitted, `kvs.yml` is used.

## Installation
If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```shell
$ brew tap itslab-kyushu/simple-kvs
$ brew install simple-kvs
```

Compiled binaries for some platforms are available on
[Github](https://github.com/itslab-kyushu/simple-kvs/releases).
To use these binaries, after downloading a binary to your environment, decompress and put it in a directory included in your $PATH.

You can also compile the client application by yourself.
To compile it, you first download the code base:

```shell
$ git clone https://github.com/itslab-kyushu/simple-kvs $GOPATH/src/itslab-kyushu/simple-kvs
```

Then, build the client application `kvs`:

```shell
$ cd $GOPATH/src/itslab-kyushu/simple-kvs/client
$ go get -d -t -v .
$ go build -o kvs
```

To build the command, [Go](https://golang.org/) > 1.7.4 is required.
