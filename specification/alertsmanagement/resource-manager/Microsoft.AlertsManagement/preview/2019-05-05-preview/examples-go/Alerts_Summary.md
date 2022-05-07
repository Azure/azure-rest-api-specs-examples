Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.5.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_Summary.json
func ExampleAlertsClient_GetSummary() {
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
	res, err := client.GetSummary(ctx,
		armalertsmanagement.AlertsSummaryGroupByFields("severity,alertState"),
		&armalertsmanagement.AlertsClientGetSummaryOptions{IncludeSmartGroupsCount: nil,
			TargetResource:      nil,
			TargetResourceType:  nil,
			TargetResourceGroup: nil,
			MonitorService:      nil,
			MonitorCondition:    nil,
			Severity:            nil,
			AlertState:          nil,
			AlertRule:           nil,
			TimeRange:           nil,
			CustomTimeRange:     nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
