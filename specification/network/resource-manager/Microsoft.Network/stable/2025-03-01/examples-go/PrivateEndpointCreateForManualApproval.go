package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateEndpointCreateForManualApproval.json
func ExamplePrivateEndpointsClient_BeginCreateOrUpdate_createPrivateEndpointWithManualApprovalConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointsClient().BeginCreateOrUpdate(ctx, "rg1", "testPe", armnetwork.PrivateEndpoint{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.PrivateEndpointProperties{
			CustomNetworkInterfaceName: to.Ptr("testPeNic"),
			IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
				{
					Name: to.Ptr("pestaticconfig"),
					Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
						GroupID:          to.Ptr("file"),
						MemberName:       to.Ptr("file"),
						PrivateIPAddress: to.Ptr("192.168.0.5"),
					},
				}},
			ManualPrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
				{
					Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
						GroupIDs: []*string{
							to.Ptr("groupIdFromResource")},
						PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
						RequestMessage:       to.Ptr("Please manually approve my connection."),
					},
				}},
			Subnet: &armnetwork.Subnet{
				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
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
	// res.PrivateEndpoint = armnetwork.PrivateEndpoint{
	// 	Name: to.Ptr("testPe"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PrivateEndpointProperties{
	// 		ApplicationSecurityGroups: []*armnetwork.ApplicationSecurityGroup{
	// 		},
	// 		CustomDNSConfigs: []*armnetwork.CustomDNSConfigPropertiesFormat{
	// 			{
	// 				Fqdn: to.Ptr("abc.cosmos.com"),
	// 				IPAddresses: []*string{
	// 					to.Ptr("192.168.0.4")},
	// 				},
	// 				{
	// 					Fqdn: to.Ptr("abc2.cosmos.com"),
	// 					IPAddresses: []*string{
	// 						to.Ptr("192.168.0.5")},
	// 				}},
	// 				CustomNetworkInterfaceName: to.Ptr("testPeNic"),
	// 				IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
	// 					{
	// 						Name: to.Ptr("pestaticconfig"),
	// 						Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
	// 							GroupID: to.Ptr("file"),
	// 							MemberName: to.Ptr("file"),
	// 							PrivateIPAddress: to.Ptr("192.168.0.5"),
	// 						},
	// 				}},
	// 				ManualPrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe/manualPrivateLinkServiceConnections/plsconnection1"),
	// 						Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
	// 							GroupIDs: []*string{
	// 								to.Ptr("groupIdFromResource")},
	// 								PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 									Description: to.Ptr("Awaiting approval"),
	// 									ActionsRequired: to.Ptr("None"),
	// 									Status: to.Ptr("Pending"),
	// 								},
	// 								PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 								RequestMessage: to.Ptr("Please manually approve my connection."),
	// 							},
	// 					}},
	// 					NetworkInterfaces: []*armnetwork.Interface{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPe.nic.abcd1234"),
	// 					}},
	// 					PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 					},
	// 				},
	// 			}
}
