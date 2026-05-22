package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns/v2"
)

// Generated from example definition: 2024-06-01/VirtualNetworkLinkList.json
func ExampleVirtualNetworkLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkLinksClient().NewListPager("resourceGroup1", "privatelink.contoso.com", nil)
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
		// page = armprivatedns.VirtualNetworkLinksClientListResponse{
		// 	VirtualNetworkLinkListResult: armprivatedns.VirtualNetworkLinkListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/virtualNetworkLinks?api-version=2024-06-01&$skipToken=skipToken"),
		// 		Value: []*armprivatedns.VirtualNetworkLink{
		// 			{
		// 				Name: to.Ptr("virtualNetworkLink1"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones/virtualNetworkLinks"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatelink.contoso.com/virtualNetworkLinks/virtualNetworkLink1"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armprivatedns.VirtualNetworkLinkProperties{
		// 					ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 					RegistrationEnabled: to.Ptr(false),
		// 					ResolutionPolicy: to.Ptr(armprivatedns.ResolutionPolicyNxDomainRedirect),
		// 					VirtualNetwork: &armprivatedns.SubResource{
		// 						ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
		// 					},
		// 					VirtualNetworkLinkState: to.Ptr(armprivatedns.VirtualNetworkLinkStateCompleted),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("virtualNetworkLink2"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones/virtualNetworkLinks"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatelink.contoso.com/virtualNetworkLinks/virtualNetworkLink2"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armprivatedns.VirtualNetworkLinkProperties{
		// 					ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 					RegistrationEnabled: to.Ptr(true),
		// 					ResolutionPolicy: to.Ptr(armprivatedns.ResolutionPolicyDefault),
		// 					VirtualNetwork: &armprivatedns.SubResource{
		// 						ID: to.Ptr("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
		// 					},
		// 					VirtualNetworkLinkState: to.Ptr(armprivatedns.VirtualNetworkLinkStateInProgress),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
