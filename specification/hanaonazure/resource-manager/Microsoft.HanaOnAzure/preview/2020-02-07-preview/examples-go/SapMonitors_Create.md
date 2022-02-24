Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhanaonazure%2Farmhanaonazure%2Fv0.2.1/sdk/resourcemanager/hanaonazure/armhanaonazure/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Create.json
func ExampleSapMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhanaonazure.NewSapMonitorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		armhanaonazure.SapMonitor{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			Properties: &armhanaonazure.SapMonitorProperties{
				EnableCustomerAnalytics:        to.BoolPtr(true),
				LogAnalyticsWorkspaceArmID:     to.StringPtr("<log-analytics-workspace-arm-id>"),
				LogAnalyticsWorkspaceID:        to.StringPtr("<log-analytics-workspace-id>"),
				LogAnalyticsWorkspaceSharedKey: to.StringPtr("<log-analytics-workspace-shared-key>"),
				MonitorSubnet:                  to.StringPtr("<monitor-subnet>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SapMonitorsClientCreateResult)
}
```
