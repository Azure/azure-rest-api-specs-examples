package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/VirtualNetworkLinkList.json
func ExampleVirtualNetworkLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkLinksClient().NewListPager("resourceGroup1", "privatezone1.com", &armprivatedns.VirtualNetworkLinksClientListOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualNetworkLinkListResult = armprivatedns.VirtualNetworkLinkListResult{
		// 	Value: []*armprivatedns.VirtualNetworkLink{
		// 		{
		// 			Name: to.Ptr("virtualNetworkLink1"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/virtualNetworkLinks"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/virtualNetworkLinks/virtualNetworkLink1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.VirtualNetworkLinkProperties{
		// 				ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 				RegistrationEnabled: to.Ptr(false),
		// 				VirtualNetwork: &armprivatedns.SubResource{
		// 					ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
		// 				},
		// 				VirtualNetworkLinkState: to.Ptr(armprivatedns.VirtualNetworkLinkStateCompleted),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("virtualNetworkLink2"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/virtualNetworkLinks"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/virtualNetworkLinks/virtualNetworkLink2"),
		// 			Location: to.Ptr("global"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.VirtualNetworkLinkProperties{
		// 				ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 				RegistrationEnabled: to.Ptr(true),
		// 				VirtualNetwork: &armprivatedns.SubResource{
		// 					ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
		// 				},
		// 				VirtualNetworkLinkState: to.Ptr(armprivatedns.VirtualNetworkLinkStateInProgress),
		// 			},
		// 	}},
		// }
	}
}
