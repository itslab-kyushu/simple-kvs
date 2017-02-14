//
// commands.go
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
	"fmt"
	"os"

	"github.com/itslab-kyushu/simple-kvs/client/command"
	"github.com/urfave/cli"
)

// DefaultConfFile defines the default configuration file name.
const DefaultConfFile = "kvs.yml"

// GlobalFlags defines a set of global flags.
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "quiet",
		Usage: "not output logging infroamtion",
	},
}

// Commands defines a set of commands.
var Commands = []cli.Command{
	{
		Name:        "get",
		Usage:       "Download a file",
		Description: "Download a file from servers.",
		ArgsUsage:   "<file name>",
		Action:      command.CmdGet,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Usage: "Server configuration `FILE`.",
				Value: DefaultConfFile,
			},
			cli.StringFlag{
				Name:  "output",
				Usage: "Store the downloaded file to the `FILE`.",
			},
		},
	},
	{
		Name:        "put",
		Usage:       "Upload a file",
		Description: "Upload a given file to servers.",
		ArgsUsage:   "<file>",
		Action:      command.CmdPut,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Usage: "Server configuration `FILE`.",
				Value: DefaultConfFile,
			},
		},
	},
	{
		Name:        "delete",
		Usage:       "Delete a file from all servers",
		Description: "Delete a file from all servers.",
		ArgsUsage:   "<file name>",
		Action:      command.CmdDelete,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Usage: "Server configuration `FILE`.",
				Value: DefaultConfFile,
			},
		},
	},
	{
		Name:        "list",
		Usage:       "Get a list of files stored in servers",
		Description: "Receive a list of files stored in a random server.",
		ArgsUsage:   " ",
		Action:      command.CmdList,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Usage: "Server configuration `FILE`.",
				Value: DefaultConfFile,
			},
		},
	},
}

// CommandNotFound handles an error that the given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
