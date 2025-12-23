package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/BastionHostDeveloperPut.json
func ExampleBastionHostsClient_BeginCreateOrUpdate_createDeveloperBastionHost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginCreateOrUpdate(ctx, "rg2", "bastionhostdeveloper", armnetwork.BastionHost{
		Properties: &armnetwork.BastionHostPropertiesFormat{
			IPConfigurations: []*armnetwork.BastionHostIPConfiguration{},
			NetworkACLs: &armnetwork.BastionHostPropertiesFormatNetworkACLs{
				IPRules: []*armnetwork.IPRule{
					{
						AddressPrefix: to.Ptr("1.1.1.1/16"),
					}},
			},
			VirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
		},
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
	// res.BastionHost = armnetwork.BastionHost{
	// 	Name: to.Ptr("bastionhostdeveloper"),
	// 	Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/bastionHosts/bastionhostdeveloper'"),
	// 	Location: to.Ptr("West US"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.BastionHostPropertiesFormat{
	// 		DNSName: to.Ptr("omnibrain.uswest.bastionglobal.azure.com"),
	// 		IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 		},
	// 		NetworkACLs: &armnetwork.BastionHostPropertiesFormatNetworkACLs{
	// 			IPRules: []*armnetwork.IPRule{
	// 				{
	// 					AddressPrefix: to.Ptr("1.1.1.1/16"),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualNetwork: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 		},
	// 	},
	// 	SKU: &armnetwork.SKU{
	// 		Name: to.Ptr(armnetwork.BastionHostSKUNameDeveloper),
	// 	},
	// }
}
