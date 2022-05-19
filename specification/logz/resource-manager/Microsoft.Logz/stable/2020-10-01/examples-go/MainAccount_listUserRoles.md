Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogz%2Farmlogz%2Fv1.0.0/sdk/resourcemanager/logz/armlogz/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MainAccount_listUserRoles.json
func ExampleMonitorsClient_NewListUserRolesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogz.NewMonitorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListUserRolesPager("myResourceGroup",
		"myMonitor",
		&armlogz.MonitorsClientListUserRolesOptions{Body: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
