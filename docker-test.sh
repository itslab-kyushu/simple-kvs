#!/bin/bash
#
# docker-test.sh
#
# Copyright (c) 2017 Junpei Kawamoto
#
# This file is part of sss.
#
# sss is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# sss is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with sss.  If not, see <http://www.gnu.org/licenses/>.
#

#
# Run docker based client/server tests.
#
docker run -d --name kvs-server -p 13009:13009 itslabq/simple-kvs

cat << EOS > kvs.yml
servers:
  - address: 127.0.0.1
    port: 13009
EOS

cd client
go build -o client
cd ../
./client/client put kvs.yml
./client/client get kvs.yml --output kvs2.yml

docker kill kvs-server
docker rm kvs-server

[[ -z $(diff kvs.yml kvs.yml) ]]
