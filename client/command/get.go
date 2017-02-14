//
// client/command/remote/get.go
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

package command

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"google.golang.org/grpc"

	"github.com/itslab-kyushu/simple-kvs/kvs"
	"github.com/itslab-kyushu/sss/cfg"
	"github.com/urfave/cli"
)

type getOpt struct {
	Config     *cfg.Config
	Name       string
	OutputFile string
	Log        io.Writer
}

// CmdGet prepares get command and run cmdGet.
func CmdGet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}

	output := c.String("output")
	if output == "" {
		output = c.Args().First()
	}

	var log io.Writer
	if c.GlobalBool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	return cmdGet(&getOpt{
		Config:     conf,
		Name:       c.Args().First(),
		OutputFile: output,
		Log:        log,
	})

}

func cmdGet(opt *getOpt) (err error) {

	if opt.Config.NServers() == 0 {
		return fmt.Errorf("No server information is given.")
	}

	fmt.Fprintln(opt.Log, "Downloading a file")
	server := opt.Config.Servers[0]

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", server.Address, server.Port),
		grpc.WithInsecure(),
		grpc.WithCompressor(grpc.NewGZIPCompressor()),
		grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
	)
	if err != nil {
		return
	}
	defer conn.Close()

	client := kvs.NewKvsClient(conn)
	value, err := client.Get(context.Background(), &kvs.Key{
		Name: opt.Name,
	})
	if err != nil {
		return
	}
	return ioutil.WriteFile(opt.OutputFile, value.Value, 0644)

}
