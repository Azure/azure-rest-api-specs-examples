Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.3.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_List.json
func ExampleAlertsClient_GetAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armalertsmanagement.NewAlertsClient("<subscription-id>", cred, nil)
	pager := client.GetAll(&armalertsmanagement.AlertsClientGetAllOptions{TargetResource: nil,
		TargetResourceType:  nil,
		TargetResourceGroup: nil,
		MonitorService:      nil,
		MonitorCondition:    nil,
		Severity:            nil,
		AlertState:          nil,
		AlertRule:           nil,
		SmartGroupID:        nil,
		IncludeContext:      nil,
		IncludeEgressConfig: nil,
		PageCount:           nil,
		SortBy:              nil,
		SortOrder:           nil,
		Select:              nil,
		TimeRange:           nil,
		CustomTimeRange:     nil,
	})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
