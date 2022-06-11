```go
package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetwork/ManagedNetworksPut.json
func ExampleManagedNetworksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewManagedNetworksClient("subscriptionA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myManagedNetwork",
		armmanagednetwork.ManagedNetwork{
			Location: to.Ptr("eastus"),
			Tags:     map[string]*string{},
			Properties: &armmanagednetwork.Properties{
				Scope: &armmanagednetwork.Scope{
					ManagementGroups: []*armmanagednetwork.ResourceID{
						{
							ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
						},
						{
							ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0002-0000-0000-000000000000"),
						}},
					Subnets: []*armmanagednetwork.ResourceID{
						{
							ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetA"),
						},
						{
							ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetB"),
						}},
					Subscriptions: []*armmanagednetwork.ResourceID{
						{
							ID: to.Ptr("subscriptionA"),
						},
						{
							ID: to.Ptr("subscriptionB"),
						}},
					VirtualNetworks: []*armmanagednetwork.ResourceID{
						{
							ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
						},
						{
							ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagednetwork%2Farmmanagednetwork%2Fv0.1.0/sdk/resourcemanager/managednetwork/armmanagednetwork/README.md) on how to add the SDK to your project and authenticate.
