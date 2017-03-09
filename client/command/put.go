//
// client/command/remote/put.go
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

// putOpt defines option values for cmdPut.
type putOpt struct {
	Filename string
	Rename   string
	Config   *cfg.Config
	Log      io.Writer
}

// CmdPut prepares put command and run cmdPut.
func CmdPut(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}
	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}

	var output io.Writer
	if c.GlobalBool("quiet") {
		output = ioutil.Discard
	} else {
		output = os.Stderr
	}

	return cmdPut(&putOpt{
		Filename: c.Args().First(),
		Rename:   c.String("name"),
		Config:   conf,
		Log:      output,
	})

}

func cmdPut(opt *putOpt) (err error) {

	if opt.Config.NServers() == 0 {
		return fmt.Errorf("No server information is given.")
	}

	data, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Uploading a file")
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

	var name string
	if opt.Rename != "" {
		name = opt.Rename
	} else {
		name = opt.Filename
	}
	client := kvs.NewKvsClient(conn)
	_, err = client.Put(context.Background(), &kvs.Entry{
		Key: &kvs.Key{
			Name: name,
		},
		Value: &kvs.Value{
			Value: data,
		},
	})
	return

}
