package roleAdmin

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

const (
	standardRoleId = "6b24b5d6-e15e-11e8-b9ec-bb5761ac6cc5"
	readOnlyRoleId = "4b015cf4-e164-11e8-af08-af11849ec75b"
	adminRoleId    = "60cf6c52-e140-11e8-bd15-f7250a8586e6"
)

func RolebyID(id string) (r UserRole, err error) {
	switch id {
	case standardRoleId:
		r = *standardRole
	case readOnlyRoleId:
		r = *readOnlyRole
	case adminRoleId:
		r = *adminRole
	default:
		err = fmt.Errorf("%v is not a valid user role id", id)
	}
	return
}

type UserRole struct {
	flagValue string
	id        string
	name      string
}

func (u *UserRole) FlagValue() string {
	return u.flagValue
}

func (u *UserRole) Id() string {
	return u.id
}

func (u *UserRole) Name() string {
	return u.name
}

var (
	standardRole = &UserRole{id: standardRoleId, name: "Datadog Standard Role", flagValue: "s"}
	readOnlyRole = &UserRole{id: readOnlyRoleId, name: "Datadog Read Only Role", flagValue: "r"}
	adminRole    = &UserRole{id: adminRoleId, name: "Datadog Admin Role", flagValue: "a"}
)

func GetRole(fv string) (r *UserRole, err error) {
	switch fv {
	case "s":
		r = standardRole
	case "r":
		r = readOnlyRole
	case "a":
		r = adminRole
	default:
		err = fmt.Errorf("%v is not a valid user role", fv)
	}
	return
}

func Roles() (data []datadog.Role, err error) {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()

	client := datadog.NewAPIClient(configuration)
	r, _, err := client.RolesApi.ListRoles(ctx)

	if err != nil {
		return
	}

	data = *r.Data
	return
}
