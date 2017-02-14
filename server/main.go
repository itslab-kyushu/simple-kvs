//
// server/main.go
//
// Copyright (c) 2017 Junpei Kawamoto
//
// This file is part of simple-kvs.
//
// simple-kvs is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// simple-kvs is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with simple-kvs.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"os"
	"runtime"

	"github.com/itslab-kyushu/simple-kvs/server/command"
	"github.com/urfave/cli"
)

func main() {

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = "Start a key-value store server"
	app.UsageText = "kvs-server [global options]"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Action = command.CmdRun

	app.EnableBashCompletion = true
	app.Copyright = `simple-kvs  Copyright (C) 2017  Junpei Kawamoto
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.`

	app.Run(os.Args)
}
