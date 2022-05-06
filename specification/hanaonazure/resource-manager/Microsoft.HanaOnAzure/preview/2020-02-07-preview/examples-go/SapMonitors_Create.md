Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhanaonazure%2Farmhanaonazure%2Fv0.4.0/sdk/resourcemanager/hanaonazure/armhanaonazure/README.md) on how to add the SDK to your project and authenticate.

```go
package armhanaonazure_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Create.json
func ExampleSapMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		armhanaonazure.SapMonitor{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armhanaonazure.SapMonitorProperties{
				EnableCustomerAnalytics:        to.Ptr(true),
				LogAnalyticsWorkspaceArmID:     to.Ptr("<log-analytics-workspace-arm-id>"),
				LogAnalyticsWorkspaceID:        to.Ptr("<log-analytics-workspace-id>"),
				LogAnalyticsWorkspaceSharedKey: to.Ptr("<log-analytics-workspace-shared-key>"),
				MonitorSubnet:                  to.Ptr("<monitor-subnet>"),
			},
		},
		&armhanaonazure.SapMonitorsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
