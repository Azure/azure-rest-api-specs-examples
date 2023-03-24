package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/PublicIpAddressList.json
func ExamplePublicIPAddressesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicIPAddressesClient().NewListPager("rg1", nil)
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
		// page.PublicIPAddressListResult = armnetwork.PublicIPAddressListResult{
		// 	Value: []*armnetwork.PublicIPAddress{
		// 		{
		// 			Name: to.Ptr("testDNS-ip"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPAddresses"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/testDNS-ip"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
		// 				DdosSettings: &armnetwork.DdosSettings{
		// 					ProtectionMode: to.Ptr(armnetwork.DdosSettingsProtectionModeVirtualNetworkInherited),
		// 				},
		// 				IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 				IPConfiguration: &armnetwork.IPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testDNS649/ipConfigurations/ipconfig1"),
		// 				},
		// 				IPTags: []*armnetwork.IPTag{
		// 					{
		// 						IPTagType: to.Ptr("FirstPartyUsage"),
		// 						Tag: to.Ptr("SQL"),
		// 					},
		// 					{
		// 						IPTagType: to.Ptr("FirstPartyUsage"),
		// 						Tag: to.Ptr("Storage"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ip03"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPAddresses"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/ip03"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
		// 				DdosSettings: &armnetwork.DdosSettings{
		// 					ProtectionMode: to.Ptr(armnetwork.DdosSettingsProtectionModeVirtualNetworkInherited),
		// 				},
		// 				DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
		// 					DomainNameLabel: to.Ptr("testlbl"),
		// 					Fqdn: to.Ptr("testlbl.westus.cloudapp.azure.com"),
		// 				},
		// 				IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 				IPAddress: to.Ptr("40.85.154.247"),
		// 				IPConfiguration: &armnetwork.IPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/testLb/frontendIPConfigurations/LoadBalancerFrontEnd"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 			},
		// 	}},
		// }
	}
}
