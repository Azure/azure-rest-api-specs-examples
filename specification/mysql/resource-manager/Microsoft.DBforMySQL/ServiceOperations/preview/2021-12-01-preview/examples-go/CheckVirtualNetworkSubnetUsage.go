package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2021-12-01-preview/examples/CheckVirtualNetworkSubnetUsage.json
func ExampleCheckVirtualNetworkSubnetUsageClient_Execute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCheckVirtualNetworkSubnetUsageClient().Execute(ctx, "WestUS", armmysqlflexibleservers.VirtualNetworkSubnetUsageParameter{
		VirtualNetworkResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/testvnet"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkSubnetUsageResult = armmysqlflexibleservers.VirtualNetworkSubnetUsageResult{
	// 	DelegatedSubnetsUsage: []*armmysqlflexibleservers.DelegatedSubnetUsage{
	// 		{
	// 			SubnetName: to.Ptr("test-subnet-1"),
	// 			Usage: to.Ptr[int64](2),
	// 		},
	// 		{
	// 			SubnetName: to.Ptr("test-subnet-2"),
	// 			Usage: to.Ptr[int64](3),
	// 	}},
	// 	Location: to.Ptr("WestUS"),
	// 	SubscriptionID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff"),
	// }
}
