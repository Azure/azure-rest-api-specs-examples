```go
package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceVirtualNetworkPeeringCreateOrUpdate.json
func ExampleVNetPeeringClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabricks.NewVNetPeeringClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg",
		"myWorkspace",
		"vNetPeeringTest",
		armdatabricks.VirtualNetworkPeering{
			Properties: &armdatabricks.VirtualNetworkPeeringPropertiesFormat{
				AllowForwardedTraffic:     to.Ptr(false),
				AllowGatewayTransit:       to.Ptr(false),
				AllowVirtualNetworkAccess: to.Ptr(true),
				RemoteVirtualNetwork: &armdatabricks.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork{
					ID: to.Ptr("/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"),
				},
				UseRemoteGateways: to.Ptr(false),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabricks%2Farmdatabricks%2Fv0.6.0/sdk/resourcemanager/databricks/armdatabricks/README.md) on how to add the SDK to your project and authenticate.
