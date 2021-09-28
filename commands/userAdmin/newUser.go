package userAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/utils"
)

func NewUser(name, email string) (user *DatadogUser, err error) {
	ctx := utils.NewContext()
	attribs := *datadog.NewUserCreateAttributes(email)
	attribs.SetName(name)
	data := *datadog.NewUserCreateData(attribs, datadog.UsersType("users"))
	body := *datadog.NewUserCreateRequest(data)

	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, _, err := client.UsersApi.CreateUser(ctx, body)
	if err != nil {
		return
	}

	user = &DatadogUser{
		Name:   *r.Data.Attributes.Name,
		Email:  *r.Data.Attributes.Email,
		ID:     *r.Data.Id,
		Status: *r.Data.Attributes.Status,
	}

	return
}
