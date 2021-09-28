package userAdmin

import (
	"context"
	"fmt"
	"log"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/roleAdmin"
)

// CreateUser invokes the DD api to create a new user account.
func CreateUser(email, name, role string) (user DatadogUser, err error) {
	defer func() {
		if rcr := recover(); rcr != nil {
			log.Fatalln(rcr)
		}
	}()

	userRole, err := roleAdmin.GetRole(role)
	if err != nil {
		return
	}

	fmt.Printf("new user info: email: %s; name: %s; role: %s\n", email, name, userRole.Name())

	ctx := datadog.NewDefaultContext(context.Background())
	rId := userRole.Id()
	attribs := *datadog.NewUserCreateAttributes(email)
	attribs.SetName(name)

	rd := *datadog.NewRelationshipToRoleDataWithDefaults()
	rd.Id = &rId
	rtr := *datadog.NewRelationshipToRoles()
	rtr.SetData([]datadog.RelationshipToRoleData{rd})
	rel := *datadog.NewUserRelationships()
	rel.SetRoles(rtr)

	createUserData := *datadog.NewUserCreateData(attribs, datadog.UsersType("users"))
	createUserData.SetRelationships(rel)
	body := *datadog.NewUserCreateRequest(createUserData)

	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, raw, err := client.UsersApi.CreateUser(ctx, body)
	if err != nil {
		err = fmt.Errorf("respose: %v", raw)
		return
	}

	user = NewUser(*r.Data)
	return
}
