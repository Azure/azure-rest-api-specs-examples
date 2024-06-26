package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ManagedNetwork/ManagedNetworksPatch.json
func ExampleManagedNetworksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedNetworksClient().BeginUpdate(ctx, "myResourceGroup", "myManagedNetwork", armmanagednetwork.Update{
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedNetwork = armmanagednetwork.ManagedNetwork{
	// 	Name: to.Ptr("myManagedNetwork"),
	// 	Type: to.Ptr("Microsoft.ManagedNetwork/managedNetworks"),
	// 	ID: to.Ptr("/subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmanagednetwork.Properties{
	// 		Etag: to.Ptr("sadf-asdf-asdf-asdf"),
	// 		ProvisioningState: to.Ptr(armmanagednetwork.ProvisioningStateSucceeded),
	// 		Connectivity: &armmanagednetwork.ConnectivityCollection{
	// 			Groups: []*armmanagednetwork.Group{
	// 			},
	// 			Peerings: []*armmanagednetwork.PeeringPolicy{
	// 			},
	// 		},
	// 		Scope: &armmanagednetwork.Scope{
	// 			ManagementGroups: []*armmanagednetwork.ResourceID{
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0002-0000-0000-000000000000"),
	// 			}},
	// 			Subnets: []*armmanagednetwork.ResourceID{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetA"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetB"),
	// 			}},
	// 			Subscriptions: []*armmanagednetwork.ResourceID{
	// 				{
	// 					ID: to.Ptr("subscriptionA"),
	// 				},
	// 				{
	// 					ID: to.Ptr("subscriptionB"),
	// 			}},
	// 			VirtualNetworks: []*armmanagednetwork.ResourceID{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
	// 			}},
	// 		},
	// 	},
	// }
}
