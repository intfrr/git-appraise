/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package commands contains the assorted sub commands supported by the git-review tool.
package commands

import (
	"fmt"
	"log"
)

// Command represents the definition of a single command.
type Command struct {
	Usage     func(string)
	RunMethod func([]string)
}

// Run executes a command, given its arguments.
//
// The args parameter is all of the command line args that followed the
// subcommand.
func (cmd *Command) Run(args []string) {
	cmd.RunMethod(args)
}

func notImplemented(subcommand string) *Command {
	return &Command{
		Usage: func(arg0 string) {
			fmt.Printf("Subcommand \"%s\" is not yet implemented.\n", subcommand)
		},
		RunMethod: func(args []string) {
			log.Fatal("Not Implemented")
		},
	}
}

var (
	acceptCmd  = notImplemented("accept")
	commentCmd = notImplemented("comment")
	pullCmd    = notImplemented("pull")
	pushCmd    = notImplemented("push")
	showCmd    = notImplemented("show")
	submitCmd  = notImplemented("submit")
	syncCmd    = notImplemented("sync")
)

var CommandMap = map[string]*Command{
	"accept":  acceptCmd,
	"comment": commentCmd,
	"list":    listCmd,
	"pull":    pullCmd,
	"push":    pushCmd,
	"request": requestCmd,
	"show":    showCmd,
	"submit":  submitCmd,
	"sync":    syncCmd,
}