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

// the reason for not using flag or spf13/pflag is that the commands themselves require different
// arguments themselves.  This seemed like the easiest way to do that at the time.
var (
	cmdAdd       = Command{"--add-user", "Add a new user to Datadog.", []string{email, name, role}}
	cmdDel       = Command{"--del-user", "Disable an existing user in Datadog.", []string{email}}
	cmdListUsers = Command{"--list-users", "List all of the currently active and pending Datadog users.", nil}
	cmdFindUsers = Command{"--find-user", "Find a specific user by email or name.", nil}
	cmdListRoles = Command{"--list-roles", "Lists all of the available roles supported by Datadog.", nil}
	cmdHelp      = Command{"--help", "Shows command line help.", nil}
)

func Help() {
	Logo()
	cmds := []Command{
		cmdAdd,
		cmdDel,
		cmdListUsers,
		cmdFindUsers,
		// cmdListRoles, //--> let's keep this a secret!
		cmdHelp,
	}
	thelp := &table.Table{}
	thelp.SetStyle(table.StyleColoredDark)
	thelp.SetOutputMirror(os.Stdout)
	thelp.AppendHeader(table.Row{"Command", "Description", "Required Args"})
	for _, c := range cmds {
		thelp.AppendRows([]table.Row{{c.Moniker, c.Help, ""}})
		if len(c.Args) > 0 {
			for _, a := range c.Args {
				thelp.AppendRows([]table.Row{{"", "", a}})
			}
		}
	}
	thelp.Render()
	os.Exit(0)
}

func Execute() (err error) {
	cmd := os.Args[1]
	var data []byte
	switch cmd {
	case cmdAdd.Moniker:
		if len(os.Args[2:]) != 3 {
			return fmt.Errorf("command '%s' requires 3 args: %s", cmdAdd.Moniker, strings.Join(cmdAdd.Args, "\n"))
		}
		data, err = addUsrCmd(os.Args[2], os.Args[3], os.Args[4])
	case cmdDel.Moniker:
		if len(os.Args[2:]) != 1 {
			return fmt.Errorf("command '%s' requires one argument: %s", cmd, name)
		}
		data, err = delUsrCmd(os.Args[2])
	case cmdListUsers.Moniker:
		data, err = allUsersCmd()
	case cmdFindUsers.Moniker:
		if len(os.Args[2:]) != 1 {
			return fmt.Errorf("command '%s' requires one argument: %s", cmd, name)
		}
		data, err = findUsrCmd(os.Args[2])
	case cmdListRoles.Moniker:
		data, err = allRolesCmd()
	case cmdHelp.Moniker:
		Help()
	default:
		err = fmt.Errorf("command '%s' is not valid", cmd)
	}

	if err != nil {
		return
	}

	fmt.Println(string(data))

	return
}

func addUsrCmd(arg1, arg2, arg3 string) (user []byte, err error) {
	args, err := NewCommandArg(cmdAdd.Args, arg1, arg2, arg3)
	if err != nil {
		return
	}
	data, err := userAdmin.NewUser(args[email], args[name], args[role])
	if err != nil {
		return
	}
	return yaml.Marshal(data)
}

func delUsrCmd(arg1 string) (user []byte, err error) {
	args, err := NewCommandArg(cmdDel.Args, arg1)
	if err != nil {
		return
	}
	data, err := userAdmin.DisableUser(args[email])
	if err != nil {
		return
	}
	return yaml.Marshal(data)
}

func findUsrCmd(arg1 string) (user []byte, err error) {
	args, err := NewCommandArg(cmdFindUsers.Args, arg1)
	if err != nil {
		return
	}
	data, err := userAdmin.FindUserByEmail(args[email])
	if err != nil {
		return
	}
	return yaml.Marshal(data)
}

func allUsersCmd() (users []byte, err error) {
	println("fetching all users...")
	data, err := userAdmin.ListAllUsers()
	if err != nil {
		return
	}
	return yaml.Marshal(data)
}

func allRolesCmd() (roles []byte, err error) {
	data, err := roleAdmin.Roles()
	if err != nil {
		return
	}
	return yaml.Marshal(data)
}

type Command struct {
	Moniker string
	Help    string
	Args    []string
}

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
