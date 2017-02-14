#
# client/Makefile
#
# Copyright (c) 2017 Junpei Kawamoto
#
# This file is part of simple-kvs.
#
# simple-kvs is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# simple-kvs is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with simple-kvs.  If not, see <http://www.gnu.org/licenses/>.
#
VERSION = snapshot
.PHONY: build get-deps
default: build

build: get-deps
	goxc -d=../pkg -pv=$(VERSION) -os="linux,darwin,windows,freebsd,openbsd" -n=kvs

get-deps:
	go get -d -t -v .
