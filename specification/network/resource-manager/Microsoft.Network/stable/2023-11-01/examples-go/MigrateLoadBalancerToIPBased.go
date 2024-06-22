package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/MigrateLoadBalancerToIPBased.json
func ExampleLoadBalancersClient_MigrateToIPBased() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().MigrateToIPBased(ctx, "rg1", "lb1", &armnetwork.LoadBalancersClientMigrateToIPBasedOptions{Parameters: &armnetwork.MigrateLoadBalancerToIPBasedRequest{
		Pools: []*string{
			to.Ptr("pool1"),
			to.Ptr("pool2")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigratedPools = armnetwork.MigratedPools{
	// 	MigratedPools: []*string{
	// 		to.Ptr("pool1"),
	// 		to.Ptr("pool2")},
	// 	}
}
