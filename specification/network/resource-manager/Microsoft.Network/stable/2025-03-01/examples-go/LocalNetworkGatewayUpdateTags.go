package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LocalNetworkGatewayUpdateTags.json
func ExampleLocalNetworkGatewaysClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalNetworkGatewaysClient().UpdateTags(ctx, "rg1", "lgw", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocalNetworkGateway = armnetwork.LocalNetworkGateway{
	// 	Name: to.Ptr("lgw"),
	// 	Type: to.Ptr("Microsoft.Network/localNetworkGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/lgw"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 		GatewayIPAddress: to.Ptr("12.0.0.1"),
	// 		LocalNetworkAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}
