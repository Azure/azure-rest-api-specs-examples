Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/createOrUpdateDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-uri>",
		"<name>",
		armmonitor.DiagnosticSettingsResource{
			Properties: &armmonitor.DiagnosticSettings{
				EventHubAuthorizationRuleID: to.StringPtr("<event-hub-authorization-rule-id>"),
				EventHubName:                to.StringPtr("<event-hub-name>"),
				LogAnalyticsDestinationType: to.StringPtr("<log-analytics-destination-type>"),
				Logs: []*armmonitor.LogSettings{
					{
						Category: to.StringPtr("<category>"),
						Enabled:  to.BoolPtr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Int32Ptr(0),
							Enabled: to.BoolPtr(false),
						},
					}},
				Metrics: []*armmonitor.MetricSettings{
					{
						Category: to.StringPtr("<category>"),
						Enabled:  to.BoolPtr(true),
						RetentionPolicy: &armmonitor.RetentionPolicy{
							Days:    to.Int32Ptr(0),
							Enabled: to.BoolPtr(false),
						},
					}},
				StorageAccountID: to.StringPtr("<storage-account-id>"),
				WorkspaceID:      to.StringPtr("<workspace-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientCreateOrUpdateResult)
}
```
