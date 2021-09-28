package commands

import (
	"fmt"
	"net/mail"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/twistingmercury/godog/commands/roleAdmin"
	"github.com/twistingmercury/godog/commands/userAdmin"
	"gopkg.in/yaml.v2"
)

const email = "--email"
const name = "--name"

var (
	cmdAdd       = command{"--add-user", []string{"-a"}, "Add a new user to Datadog.", []string{email, name}}
	cmdDel       = command{"--del-user", []string{"-d"}, "Disable an existing user in Datadog.", []string{email}}
	cmdListUsers = command{"--list-users", []string{"-l"}, "List all of the currently active and pending Datadog users.", nil}
	cmdFindUsers = command{"--find-user", []string{"-f"}, "Find a specific user by email or name.", nil}
	cmdListRoles = command{"--list-roles", []string{"-r"}, "Lists all of the available roles supported by Datadog.", nil}
	cmdHelp      = command{"--help", []string{"?", "-h"}, "Shows command line help.", nil}

	cmds = []command{
		cmdAdd,
		cmdDel,
		cmdListUsers,
		cmdFindUsers,
		cmdListRoles,
		cmdHelp,
	}
)

func help() {
	Logo()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Command", "Description"})
	for _, c := range cmds {
		alt := c.flag
		if c.alt != nil && len(c.alt) > 0 {
			alt = c.flag + " | " + strings.Join(c.alt, " | ")
		}
		t.AppendRows([]table.Row{
			{alt, c.help},
		})
	}
	t.Render()
	os.Exit(0)
}

func Execute() (err error) {
	cmd := os.Args[1]
	var r interface{}
	switch cmd {
	case cmdAdd.flag:
		if len(os.Args[2:]) != 2 {
			return fmt.Errorf("command '%s' requires two arguments: %s", cmdAdd.flag, strings.Join(cmdAdd.args, ","))
		}
		r, err = addUsrCmd(os.Args[2], os.Args[3])
	case cmdDel.flag:
		if len(os.Args[2:]) != 1 {
			return fmt.Errorf("command '%s' requires one argument: %s", cmd, name)
		}
		r, err = delUsrCmd(os.Args[2])
	case cmdListUsers.flag:
		r, err = allUsersCmd()
	case cmdFindUsers.flag:
		if len(os.Args[2:]) != 1 {
			return fmt.Errorf("command '%s' requires one argument: %s", cmd, name)
		}
		r, err = allUsersCmd()
	case cmdListRoles.flag:
		r, err = allRolesCmd()
	case cmdHelp.flag:
		help()
	default:
		err = fmt.Errorf("command '%s' is not valid", cmd)
	}

	if err != nil {
		return
	}

	yaml.Marshal(r)

	return
}

func addUsrCmd(arg1, arg2 string) (data interface{}, err error) {
	args, err := newCommandArg(cmdAdd.args, arg1, arg2)
	return userAdmin.NewUser(args[email], args[name])
}

func delUsrCmd(arg1 string) (data interface{}, err error) {
	args, err := newCommandArg(cmdDel.args, arg1)
	return userAdmin.DisableUser(args[email])
}

func findUsrCmd(arg1 string) (data interface{}, err error) {
	args, err := newCommandArg(cmdFindUsers.args, arg1)
	return userAdmin.FindUserByEmail(args[email])
}

func allUsersCmd() (data interface{}, err error) {
	return userAdmin.ListAllUsers()
}

func allRolesCmd() (data interface{}, err error) {
	return roleAdmin.Roles()
}

type command struct {
	flag string
	alt  []string
	help string
	args []string
}

func newCommandArg(expected []string, args ...string) (ca map[string]string, err error) {
	ca = make(map[string]string)
	for _, arg := range args {
		avp := strings.Split(arg, "=")
		if len(avp) != 2 {
			return nil, fmt.Errorf("command arg '%s' does not have a value - command args should be formatted like so: --[arg_name]=[value]", avp[0])
		}
		if avp[0] == email {
			_, err = mail.ParseAddress(avp[1])
			if err != nil {
				return
			}
		}
		ca[avp[0]] = avp[1]
	}
	return
}
