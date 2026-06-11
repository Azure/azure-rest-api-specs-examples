package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/AzureFirewallPutWithAfcConfiguration.json
func ExampleAzureFirewallsClient_BeginCreateOrUpdate_createAzureFirewallWithAfcControlPlane() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureFirewallsClient().BeginCreateOrUpdate(ctx, "rg1", "azurefirewall", armnetwork.AzureFirewall{
		Location: to.Ptr("West US"),
		Properties: &armnetwork.AzureFirewallPropertiesFormat{
			IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
				{
					Name: to.Ptr("azureFirewallIpConfiguration"),
					Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
						PublicIPAddress: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
						},
						Subnet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
						},
					},
				},
			},
			SKU: &armnetwork.AzureFirewallSKU{
				Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
				Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
			},
			ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Zones: []*string{},
	}, &armnetwork.AzureFirewallsClientBeginCreateOrUpdateOptions{
		CreateAfcControlPlane: to.Ptr(true)})
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
	// res = armnetwork.AzureFirewallsClientCreateOrUpdateResponse{
	// 	AzureFirewall: armnetwork.AzureFirewall{
	// 		Name: to.Ptr("azurefirewall"),
	// 		Type: to.Ptr("Microsoft.Network/azureFirewalls"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armnetwork.AzureFirewallPropertiesFormat{
	// 			IPConfigurations: []*armnetwork.AzureFirewallIPConfiguration{
	// 				{
	// 					Name: to.Ptr("azureFirewallIpConfiguration"),
	// 					Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall/ipConfigurations/azureFirewallIpConfiguration"),
	// 					Properties: &armnetwork.AzureFirewallIPConfigurationPropertiesFormat{
	// 						PrivateIPAddress: to.Ptr("10.0.0.0"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						PublicIPAddress: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 						},
	// 						Subnet: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			SKU: &armnetwork.AzureFirewallSKU{
	// 				Name: to.Ptr(armnetwork.AzureFirewallSKUNameAZFWVnet),
	// 				Tier: to.Ptr(armnetwork.AzureFirewallSKUTierStandard),
	// 			},
	// 			ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
	// 			AfcConfiguration: &armnetwork.AfcConfiguration{
	// 				ServiceEndpoint: to.Ptr("5e73bbe79102451d968d1cac9b5dbd41.afc.azure.com"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 		Zones: []*string{
	// 		},
	// 	},
	// }
}
