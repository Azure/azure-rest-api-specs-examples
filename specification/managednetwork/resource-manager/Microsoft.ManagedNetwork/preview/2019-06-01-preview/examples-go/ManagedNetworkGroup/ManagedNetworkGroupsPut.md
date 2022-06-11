```go
package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetworkGroup/ManagedNetworkGroupsPut.json
func ExampleGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewGroupsClient("subscriptionA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myManagedNetwork",
		"myManagedNetworkGroup1",
		armmanagednetwork.Group{
			Properties: &armmanagednetwork.GroupProperties{
				ManagementGroups: []*armmanagednetwork.ResourceID{},
				Subnets: []*armmanagednetwork.ResourceID{
					{
						ID: to.Ptr("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA/subnets/subnetA"),
					}},
				Subscriptions: []*armmanagednetwork.ResourceID{},
				VirtualNetworks: []*armmanagednetwork.ResourceID{
					{
						ID: to.Ptr("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
					},
					{
						ID: to.Ptr("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
					}},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagednetwork%2Farmmanagednetwork%2Fv0.1.0/sdk/resourcemanager/managednetwork/armmanagednetwork/README.md) on how to add the SDK to your project and authenticate.
