package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/SubnetCreateServiceEndpoint.json
func ExampleSubnetsClient_BeginCreateOrUpdate_createSubnetWithServiceEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubnetsClient().BeginCreateOrUpdate(ctx, "subnet-test", "vnetname", "subnet1", armnetwork.Subnet{
		Properties: &armnetwork.SubnetPropertiesFormat{
			AddressPrefix: to.Ptr("10.0.0.0/16"),
			ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
				{
					Service: to.Ptr("Microsoft.Storage"),
				}},
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
	// res.Subnet = armnetwork.Subnet{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/subnet-test/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnet1"),
	// 	Name: to.Ptr("subnet1"),
	// 	Properties: &armnetwork.SubnetPropertiesFormat{
	// 		AddressPrefix: to.Ptr("10.0.0.0/16"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
	// 			{
	// 				Locations: []*string{
	// 					to.Ptr("eastus2(stage)"),
	// 					to.Ptr("usnorth(stage)")},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Service: to.Ptr("Microsoft.Storage"),
	// 			}},
	// 		},
	// 	}
}
