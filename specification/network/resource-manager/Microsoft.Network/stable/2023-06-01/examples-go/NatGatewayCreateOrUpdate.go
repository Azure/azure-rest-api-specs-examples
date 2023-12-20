package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/NatGatewayCreateOrUpdate.json
func ExampleNatGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNatGatewaysClient().BeginCreateOrUpdate(ctx, "rg1", "test-natgateway", armnetwork.NatGateway{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.NatGatewayPropertiesFormat{
			PublicIPAddresses: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
				}},
			PublicIPPrefixes: []*armnetwork.SubResource{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
				}},
		},
		SKU: &armnetwork.NatGatewaySKU{
			Name: to.Ptr(armnetwork.NatGatewaySKUNameStandard),
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
	// res.NatGateway = armnetwork.NatGateway{
	// 	Name: to.Ptr("test-natGateway"),
	// 	Type: to.Ptr("Microsoft.Network/natGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/natGateways/test-natGateway"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.NatGatewayPropertiesFormat{
	// 		IdleTimeoutInMinutes: to.Ptr[int32](5),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddresses: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
	// 		}},
	// 		PublicIPPrefixes: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
	// 		}},
	// 		Subnets: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
	// 		}},
	// 	},
	// 	SKU: &armnetwork.NatGatewaySKU{
	// 		Name: to.Ptr(armnetwork.NatGatewaySKUNameStandard),
	// 	},
	// }
}
