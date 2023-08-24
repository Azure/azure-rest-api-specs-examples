package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/LocalNetworkGatewayCreate.json
func ExampleLocalNetworkGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLocalNetworkGatewaysClient().BeginCreateOrUpdate(ctx, "rg1", "localgw", armnetwork.LocalNetworkGateway{
		Location: to.Ptr("Central US"),
		Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
			Fqdn:             to.Ptr("site1.contoso.com"),
			GatewayIPAddress: to.Ptr("11.12.13.14"),
			LocalNetworkAddressSpace: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("10.1.0.0/16")},
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
	// res.LocalNetworkGateway = armnetwork.LocalNetworkGateway{
	// 	Name: to.Ptr("localgw"),
	// 	Type: to.Ptr("Microsoft.Network/localNetworkGateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw"),
	// 	Location: to.Ptr("centralus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
	// 		GatewayIPAddress: to.Ptr("11.12.13.14"),
	// 		LocalNetworkAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.1.0.0/16")},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 	}
}
