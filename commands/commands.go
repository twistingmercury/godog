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
const role = "--role"

type Command struct {
	Moniker string
	Help    string
	Args    []string
}

// NewCommandArg parses and validates command arguments.
func NewCommandArg(expected []string, args ...string) (ca map[string]string, err error) {
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

// the reason for not using flag or spf13/pflag is that the commands themselves require different
// arguments themselves.  This seemed like the easiest way to do that at the time.
var (
	cmdAdd       = Command{"--add-user", "Add a new user to Datadog.", []string{email, name, role}}
	cmdDel       = Command{"--del-user", "Disable an existing user in Datadog.", []string{email}}
	cmdFindUser  = Command{"--find-user", "Find a specific user by email or name.", []string{email}}
	cmdListUsers = Command{"--list-users", "List all of the currently active and pending Datadog users.", nil}
	cmdListRoles = Command{"--list-roles", "Lists all of the available roles supported by Datadog.", nil}
	cmdHelp      = Command{"--help", "Shows command line help.", nil}
)

// Execute reads in the command line arguments and invokes the desired Datadog API.
func Execute() (err error) {
	cmd := os.Args[1]

	switch cmd {
	case cmdAdd.Moniker:
		err = execCreateUser()
	case cmdDel.Moniker:
		err = execDisableUser()
	case cmdListUsers.Moniker:
		err = execListAllUsers()
	case cmdFindUser.Moniker:
		err = execFindUser()
	case cmdListRoles.Moniker:
		err = execListAllRoles()
	case cmdHelp.Moniker:
		Help()
	default:
		err = fmt.Errorf("command '%s' is not valid", cmd)
	}
	return
}

func execCreateUser() (err error) {
	if len(os.Args[2:]) != 3 {
		return fmt.Errorf("command '%s' requires 3 args: %s", cmdAdd.Moniker, strings.Join(cmdAdd.Args, "\n"))
	}
	args, err := NewCommandArg(cmdAdd.Args, os.Args[2:]...)
	if err != nil {
		return
	}
	user, err := userAdmin.CreateUser(args[email], args[name], args[role])
	if err != nil {
		return
	}
	data, err := yaml.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	return
}

func execDisableUser() (err error) {
	if len(os.Args[2:]) != 1 {
		return fmt.Errorf("command '%s' requires one argument: %s", cmdDel.Moniker, name)
	}
	args, err := NewCommandArg(cmdDel.Args, os.Args[2:]...)
	if err != nil {
		return
	}
	ok, user, e := userAdmin.DisableUser(args[email])
	switch {
	case e != nil:
		err = e
	case !ok:
		email := strings.Split(os.Args[2], "=")[2]
		fmt.Printf("user '%s' could not be found\n", email)
	default:
		data, e := yaml.Marshal(user)
		if e != nil {
			err = e
			break
		}
		fmt.Println(string(data))
	}
	return
}

func execFindUser() (err error) {
	if len(os.Args[2:]) != 1 {
		err = fmt.Errorf("command '%s' requires one argument: %s", cmdDel.Moniker, name)
		return
	}

	args, err := NewCommandArg(cmdDel.Args, os.Args[2:]...)
	if err != nil {
		return
	}

	ok, user, err := userAdmin.FindUser(args[email])
	switch {
	case err != nil:
		return
	case !ok:
		email := strings.Split(os.Args[2], "=")[2]
		fmt.Printf("user '%s' could not be found\n", email)
	default:
		data, e := yaml.Marshal(user)
		if e != nil {
			err = e
			break
		}
		fmt.Println(string(data))
	}
	return
}

func execListAllUsers() (err error) {
	users, err := userAdmin.ListAllUsers()
	if err != nil {
		return
	}
	data, _ := yaml.Marshal(users)
	fmt.Println(string(data))
	return
}

func execListAllRoles() (err error) {
	roles, err := roleAdmin.Roles()
	if err != nil {
		return
	}
	data, _ := yaml.Marshal(roles)
	fmt.Println(string(data))
	return
}

//Help displays the help information for Godog.
func Help() {
	Logo()
	cmds := []Command{
		cmdAdd,
		cmdDel,
		cmdListUsers,
		cmdFindUser,
		// cmdListRoles, //--> let's keep this a secret!
		cmdHelp,
	}
	t := &table.Table{}
	t.SetStyle(table.StyleColoredDark)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Command", "Description", "Required Args"})
	for _, c := range cmds {
		t.AppendRows([]table.Row{{c.Moniker, c.Help, ""}})
		if len(c.Args) > 0 {
			for _, a := range c.Args {
				t.AppendRows([]table.Row{{"", "", a}})
			}
		}
	}
	t.Render()
	os.Exit(0)
}
