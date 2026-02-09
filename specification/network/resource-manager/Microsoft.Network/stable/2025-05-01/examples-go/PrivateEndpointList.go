package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateEndpointList.json
func ExamplePrivateEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointsClient().NewListPager("rg1", nil)
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
		// page.PrivateEndpointListResult = armnetwork.PrivateEndpointListResult{
		// 	Value: []*armnetwork.PrivateEndpoint{
		// 		{
		// 			Name: to.Ptr("pe1"),
		// 			Type: to.Ptr("Microsoft.Network/privateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.PrivateEndpointProperties{
		// 				ApplicationSecurityGroups: []*armnetwork.ApplicationSecurityGroup{
		// 				},
		// 				CustomDNSConfigs: []*armnetwork.CustomDNSConfigPropertiesFormat{
		// 					{
		// 						Fqdn: to.Ptr("abc.cosmos.com"),
		// 						IPAddresses: []*string{
		// 							to.Ptr("192.168.0.4")},
		// 						},
		// 						{
		// 							Fqdn: to.Ptr("abc2.cosmos.com"),
		// 							IPAddresses: []*string{
		// 								to.Ptr("192.168.0.5")},
		// 						}},
		// 						CustomNetworkInterfaceName: to.Ptr("testPeNic"),
		// 						IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
		// 							{
		// 								Name: to.Ptr("pestaticconfig"),
		// 								Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
		// 									GroupID: to.Ptr("file"),
		// 									MemberName: to.Ptr("file"),
		// 									PrivateIPAddress: to.Ptr("192.168.0.10"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("pestaticconfig"),
		// 								Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
		// 									GroupID: to.Ptr("file"),
		// 									MemberName: to.Ptr("file2"),
		// 									PrivateIPAddress: to.Ptr("192.168.0.11"),
		// 								},
		// 						}},
		// 						ManualPrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
		// 						},
		// 						NetworkInterfaces: []*armnetwork.Interface{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/pe1.nic.abcd1234"),
		// 						}},
		// 						PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1/privateLinkServiceConnections/plsconnection1"),
		// 								Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
		// 									GroupIDs: []*string{
		// 										to.Ptr("groupIdFromResource")},
		// 										PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 											Description: to.Ptr("Auto-approved"),
		// 											ActionsRequired: to.Ptr("None"),
		// 											Status: to.Ptr("Approved"),
		// 										},
		// 										PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 										RequestMessage: to.Ptr("Please approve my connection for pe1."),
		// 									},
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Subnet: &armnetwork.Subnet{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("pe2"),
		// 						Type: to.Ptr("Microsoft.Network/privateEndpoints"),
		// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe2"),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armnetwork.PrivateEndpointProperties{
		// 							ApplicationSecurityGroups: []*armnetwork.ApplicationSecurityGroup{
		// 							},
		// 							CustomDNSConfigs: []*armnetwork.CustomDNSConfigPropertiesFormat{
		// 								{
		// 									Fqdn: to.Ptr("abc3.cosmos1.com"),
		// 									IPAddresses: []*string{
		// 										to.Ptr("192.168.0.6")},
		// 									},
		// 									{
		// 										Fqdn: to.Ptr("abc4.cosmos1.com"),
		// 										IPAddresses: []*string{
		// 											to.Ptr("192.168.0.7")},
		// 									}},
		// 									CustomNetworkInterfaceName: to.Ptr("testPeNic"),
		// 									IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
		// 										{
		// 											Name: to.Ptr("pestaticconfig3"),
		// 											Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
		// 												GroupID: to.Ptr("file"),
		// 												MemberName: to.Ptr("file"),
		// 												PrivateIPAddress: to.Ptr("192.168.0.8"),
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("pestaticconfig4"),
		// 											Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
		// 												GroupID: to.Ptr("file"),
		// 												MemberName: to.Ptr("file2"),
		// 												PrivateIPAddress: to.Ptr("192.168.0.9"),
		// 											},
		// 									}},
		// 									ManualPrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe2/privateLinkServiceConnections/plsconnection2"),
		// 											Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
		// 												GroupIDs: []*string{
		// 													to.Ptr("groupIdFromResource")},
		// 													PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 														Description: to.Ptr("Awaiting approval"),
		// 														ActionsRequired: to.Ptr("None"),
		// 														Status: to.Ptr("Pending"),
		// 													},
		// 													PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 													RequestMessage: to.Ptr("Please manually approve my connection for pe2."),
		// 												},
		// 										}},
		// 										NetworkInterfaces: []*armnetwork.Interface{
		// 											{
		// 												ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/pe2.nic.zyxw9876"),
		// 										}},
		// 										PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
		// 										},
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										Subnet: &armnetwork.Subnet{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet2"),
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}
