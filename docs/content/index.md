---
title: Simple Key-Value Datastore Service
type: homepage
menu:
  main:
    Name: Top
date: 2017-03-08
lastmod: 2017-03-08
weight: 1
description: >-
  This software provides both a Go library and command line tools implementing
  a threshold Secret Sharing scheme.
---
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/simple-kvs/tree/master)
[![wercker status](https://app.wercker.com/status/717adbfffa215daf21462bfa273a5a16/s/master "wercker status")](https://app.wercker.com/project/byKey/717adbfffa215daf21462bfa273a5a16)
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/simple-kvs/releases/tag/v0.2.0)

## Summary
This is a simple client/server application of a key-value datastore service.
It includes both client and server applications.
This software has been made for comparing performance with other key-value
datastore services using secret sharing schemes such as
[cgss](https://itslab-kyushu.github.io/cgss/) and
[sss](https://itslab-kyushu.github.io/sss/);
the results are described in the following article:

* [Hiroaki Anada](http://sun.ac.jp/prof/anada/),
  [Junpei Kawamoto](https://www.jkawamoto.info),
  Chenyutao Ke,
  [Kirill Morozov](http://www.is.c.titech.ac.jp/~morozov/), and
  [Kouichi Sakurai](http://itslab.inf.kyushu-u.ac.jp/~sakurai/),
  "[Cross-Group Secret Sharing Scheme for Secure Usage of Cloud Storage over Different Providers and Regions](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/article/10.1007%2Fs11227-017-2009-7),"
  [The Journal of Supercomputing](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/journal/11227), 2017.

Please consider to refer it, if you will publish any articles using this
software.

### Contents

* [Client application usage](client)
* [Server application usage](server)

## License
This software is released under The GNU General Public License Version 3,
see [license](./licenses/) for more detail.
