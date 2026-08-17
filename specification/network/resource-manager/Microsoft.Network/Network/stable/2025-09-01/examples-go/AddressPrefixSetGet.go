package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/AddressPrefixSetGet.json
func ExampleAddressPrefixSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAddressPrefixSetsClient().Get(ctx, "rg1", "test-asg", "test-prefix-set", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.AddressPrefixSetsClientGetResponse{
	// 	AddressPrefixSet: armnetwork.AddressPrefixSet{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationSecurityGroups/test-asg/addressPrefixSets/test-prefix-set"),
	// 		Name: to.Ptr("test-prefix-set"),
	// 		Type: to.Ptr("Microsoft.Network/applicationSecurityGroups/addressPrefixSets"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		Properties: &armnetwork.AddressPrefixSetPropertiesFormat{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/16"),
	// 				to.Ptr("192.168.1.0/24"),
	// 				to.Ptr("2001:db8::/32"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
