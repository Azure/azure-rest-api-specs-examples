package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns/v2"
)

// Generated from example definition: 2024-06-01/VirtualNetworkLinkPut.json
func ExampleVirtualNetworkLinksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkLinksClient().BeginCreateOrUpdate(ctx, "resourceGroup1", "privatezone1.com", "virtualNetworkLink1", armprivatedns.VirtualNetworkLink{
		Location: to.Ptr("Global"),
		Properties: &armprivatedns.VirtualNetworkLinkProperties{
			RegistrationEnabled: to.Ptr(false),
			VirtualNetwork: &armprivatedns.SubResource{
				ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprivatedns.VirtualNetworkLinksClientCreateOrUpdateResponse{
	// 	VirtualNetworkLink: armprivatedns.VirtualNetworkLink{
	// 		Name: to.Ptr("virtualNetworkLink1"),
	// 		Type: to.Ptr("Microsoft.Network/privateDnsZones/virtualNetworkLinks"),
	// 		Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/virtualNetworkLinks/virtualNetworkLink1"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armprivatedns.VirtualNetworkLinkProperties{
	// 			ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
	// 			RegistrationEnabled: to.Ptr(false),
	// 			VirtualNetwork: &armprivatedns.SubResource{
	// 				ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
	// 			},
	// 			VirtualNetworkLinkState: to.Ptr(armprivatedns.VirtualNetworkLinkStateCompleted),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
