#
# Dockerfile
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

#
# This dockerfile builds an image for sss server.
#
# Before building this image, kvs-server must be built in server.
#
FROM ubuntu:latest
MAINTAINER Junpei Kawamoto <kawamoto.junpei@gmail.com>

VOLUME /data
WORKDIR /root

EXPOSE 13009

ADD ./server/kvs-server /root/

CMD ["/root/kvs-server", "--root", "/data"]
