package roleAdmin

import (
	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/twistingmercury/godog/commands/utils"
)

func Roles() (data []datadog.Role, err error) {
	ctx := utils.NewContext()

	configuration := datadog.NewConfiguration()

	client := datadog.NewAPIClient(configuration)
	r, _, err := client.RolesApi.ListRoles(ctx)

	if err != nil {
		return
	}

	data = r.GetData()
	return
}
