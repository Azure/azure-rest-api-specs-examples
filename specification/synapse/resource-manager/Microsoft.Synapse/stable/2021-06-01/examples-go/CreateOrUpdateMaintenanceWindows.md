Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.5.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateMaintenanceWindows.json
func ExampleSQLPoolMaintenanceWindowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolMaintenanceWindowsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"samplerg",
		"testworkspace",
		"testsp",
		"current",
		armsynapse.MaintenanceWindows{
			Properties: &armsynapse.MaintenanceWindowsProperties{
				TimeRanges: []*armsynapse.MaintenanceWindowTimeRange{
					{
						DayOfWeek: to.Ptr(armsynapse.DayOfWeekSaturday),
						Duration:  to.Ptr("PT60M"),
						StartTime: to.Ptr("00:00:00"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
