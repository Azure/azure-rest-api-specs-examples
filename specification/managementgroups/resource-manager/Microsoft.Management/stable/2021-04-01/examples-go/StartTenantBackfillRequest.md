Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagementgroups%2Farmmanagementgroups%2Fv0.3.0/sdk/resourcemanager/managementgroups/armmanagementgroups/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/StartTenantBackfillRequest.json
func ExampleAPIClient_StartTenantBackfill() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagementgroups.NewAPIClient(cred, nil)
	res, err := client.StartTenantBackfill(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.APIClientStartTenantBackfillResult)
}
```
