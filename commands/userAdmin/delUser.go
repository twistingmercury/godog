package userAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/utils"
)

func DisableUser(email string) (user *DatadogUser, err error) {
	user, err = FindUserByEmail(email)
	if err != nil {
		return
	}
	ctx := utils.NewContext()
	attribs := *datadog.NewUserUpdateAttributes()
	attribs.SetDisabled(true)
	data := *datadog.NewUserUpdateData(attribs, user.ID, datadog.UsersType("users"))
	body := *datadog.NewUserUpdateRequest(data)

	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, _, err := client.UsersApi.UpdateUser(ctx, user.ID, body)
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
