package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayConnectionUpdateTags.json
func ExampleVirtualNetworkGatewayConnectionsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewayConnectionsClient().BeginUpdateTags(ctx, "rg1", "test", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.VirtualNetworkGatewayConnection = armnetwork.VirtualNetworkGatewayConnection{
	// 	Name: to.Ptr("test"),
	// 	Type: to.Ptr("Microsoft.Network/connections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/connections/test"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.VirtualNetworkGatewayConnectionPropertiesFormat{
	// 		ConnectionStatus: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionStatusUnknown),
	// 		ConnectionType: to.Ptr(armnetwork.VirtualNetworkGatewayConnectionTypeIPsec),
	// 		EgressBytesTransferred: to.Ptr[int64](0),
	// 		EgressNatRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
	// 		}},
	// 		EnableBgp: to.Ptr(false),
	// 		GatewayCustomBgpIPAddresses: []*armnetwork.GatewayCustomBgpIPAddressIPConfiguration{
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.1"),
	// 				IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/default"),
	// 			},
	// 			{
	// 				CustomBgpIPAddress: to.Ptr("169.254.21.3"),
	// 				IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/ipConfigurations/ActiveActive"),
	// 		}},
	// 		IngressBytesTransferred: to.Ptr[int64](0),
	// 		IngressNatRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
	// 		}},
	// 		IPSecPolicies: []*armnetwork.IPSecPolicy{
	// 		},
	// 		LocalNetworkGateway2: &armnetwork.LocalNetworkGateway{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/lgw"),
	// 			Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		RoutingWeight: to.Ptr[int32](0),
	// 		TrafficSelectorPolicies: []*armnetwork.TrafficSelectorPolicy{
	// 		},
	// 		UsePolicyBasedTrafficSelectors: to.Ptr(false),
	// 		VirtualNetworkGateway1: &armnetwork.VirtualNetworkGateway{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw"),
	// 			Properties: &armnetwork.VirtualNetworkGatewayPropertiesFormat{
	// 			},
	// 		},
	// 	},
	// }
}
