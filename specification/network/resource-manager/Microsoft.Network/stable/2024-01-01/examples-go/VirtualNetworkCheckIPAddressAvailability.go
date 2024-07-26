package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VirtualNetworkCheckIPAddressAvailability.json
func ExampleVirtualNetworksClient_CheckIPAddressAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().CheckIPAddressAvailability(ctx, "rg1", "test-vnet", "10.0.1.4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPAddressAvailabilityResult = armnetwork.IPAddressAvailabilityResult{
	// 	Available: to.Ptr(false),
	// 	AvailableIPAddresses: []*string{
	// 		to.Ptr("10.0.1.5"),
	// 		to.Ptr("10.0.1.6"),
	// 		to.Ptr("10.0.1.7"),
	// 		to.Ptr("10.0.1.8"),
	// 		to.Ptr("10.0.1.9")},
	// 	}
}
