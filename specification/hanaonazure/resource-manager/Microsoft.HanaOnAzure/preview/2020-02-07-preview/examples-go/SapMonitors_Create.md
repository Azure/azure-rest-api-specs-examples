Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhanaonazure%2Farmhanaonazure%2Fv0.5.0/sdk/resourcemanager/hanaonazure/armhanaonazure/README.md) on how to add the SDK to your project and authenticate.

```go
package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Create.json
func ExampleSapMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhanaonazure.NewSapMonitorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"mySapMonitor",
		armhanaonazure.SapMonitor{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armhanaonazure.SapMonitorProperties{
				EnableCustomerAnalytics:        to.Ptr(true),
				LogAnalyticsWorkspaceArmID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
				LogAnalyticsWorkspaceID:        to.Ptr("00000000-0000-0000-0000-000000000000"),
				LogAnalyticsWorkspaceSharedKey: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000=="),
				MonitorSubnet:                  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
