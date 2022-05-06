Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.5.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_List.json
func ExampleAlertsClient_NewGetAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewGetAllPager(&armalertsmanagement.AlertsClientGetAllOptions{TargetResource: nil,
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
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
