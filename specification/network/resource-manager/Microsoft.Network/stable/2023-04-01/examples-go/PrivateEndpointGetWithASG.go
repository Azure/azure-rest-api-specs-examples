package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/PrivateEndpointGetWithASG.json
func ExamplePrivateEndpointsClient_Get_getPrivateEndpointWithApplicationSecurityGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointsClient().Get(ctx, "rg1", "testPe", &armnetwork.PrivateEndpointsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpoint = armnetwork.PrivateEndpoint{
	// 	Name: to.Ptr("testPe"),
	// 	Type: to.Ptr("Microsoft.Network/privateEndpoints"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PrivateEndpointProperties{
	// 		ApplicationSecurityGroups: []*armnetwork.ApplicationSecurityGroup{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1"),
	// 		}},
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
	// 				ManualPrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
	// 				},
	// 				NetworkInterfaces: []*armnetwork.Interface{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPe.nic.abcd1234"),
	// 				}},
	// 				PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
	// 					{
	// 						Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
	// 							GroupIDs: []*string{
	// 								to.Ptr("groupIdFromResource")},
	// 								PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 									Description: to.Ptr("Auto-approved"),
	// 									ActionsRequired: to.Ptr("None"),
	// 									Status: to.Ptr("Approved"),
	// 								},
	// 								PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 								RequestMessage: to.Ptr("Please approve my connection."),
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 					},
	// 				},
	// 			}
}
