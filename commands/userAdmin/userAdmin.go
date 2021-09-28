package userAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/roleAdmin"
	"gopkg.in/yaml.v2"
)

// DatadogUser is a truncated datadog.User object.
type DatadogUser struct {
	Name   string
	Email  string
	Id     string
	Status string
	Roles  []string
}

// String 'stringifies' the instance.
func (d *DatadogUser) String() string {
	yml, _ := yaml.Marshal(d)
	return string(yml)
}

func NewUser(du datadog.User) (ddu DatadogUser) {
	ddu = DatadogUser{
		Name:   *du.Attributes.Name,
		Email:  *du.Attributes.Email,
		Id:     *du.Id,
		Status: *du.Attributes.Status,
	}

	if du.HasRelationships() {
		ddu.Roles = make([]string, 0)
		for _, rel := range *du.Relationships.Roles.Data {
			ur, err := roleAdmin.RolebyID(*rel.Id)
			if err != nil {
				continue
			}
			ddu.Roles = append(ddu.Roles, ur.Name())
		}
	}

	return
}
