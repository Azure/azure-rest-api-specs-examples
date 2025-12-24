package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/IpamPools_GetPoolUsage.json
func ExampleIpamPoolsClient_GetPoolUsage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIpamPoolsClient().GetPoolUsage(ctx, "rg1", "TestNetworkManager", "TestPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PoolUsage = armnetwork.PoolUsage{
	// 	AddressPrefixes: []*string{
	// 		to.Ptr("10.0.0.0/8"),
	// 		to.Ptr("192.0.0.0/8")},
	// 		AllocatedAddressPrefixes: []*string{
	// 			to.Ptr("10.0.2.0/26"),
	// 			to.Ptr("10.0.3.0/26")},
	// 			AvailableAddressPrefixes: []*string{
	// 				to.Ptr("10.0.1.0/24"),
	// 				to.Ptr("10.0.2.64/26"),
	// 				to.Ptr("10.0.2.128/25"),
	// 				to.Ptr("10.0.3.64/26"),
	// 				to.Ptr("10.0.3.128/25"),
	// 				to.Ptr("10.0.4.0/22"),
	// 				to.Ptr("10.0.8.0/21"),
	// 				to.Ptr("10.0.16.0/20"),
	// 				to.Ptr("10.0.32.0/19"),
	// 				to.Ptr("10.0.64.0/18"),
	// 				to.Ptr("10.0.128.0/17"),
	// 				to.Ptr("10.1.0.0/16"),
	// 				to.Ptr("10.2.0.0/15"),
	// 				to.Ptr("10.4.0.0/14"),
	// 				to.Ptr("10.8.0.0/13"),
	// 				to.Ptr("10.16.0.0/12"),
	// 				to.Ptr("10.32.0.0/11"),
	// 				to.Ptr("10.64.0.0/10"),
	// 				to.Ptr("10.128.0.0/9"),
	// 				to.Ptr("192.0.0.0/8")},
	// 				ChildPools: []*armnetwork.ResourceBasics{
	// 					{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("10.0.2.0/26")},
	// 							ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool1"),
	// 						},
	// 						{
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("10.0.3.0/26")},
	// 								ResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool2"),
	// 						}},
	// 						NumberOfAllocatedIPAddresses: to.Ptr("128"),
	// 						NumberOfAvailableIPAddresses: to.Ptr("33554048"),
	// 						NumberOfReservedIPAddresses: to.Ptr("256"),
	// 						ReservedAddressPrefixes: []*string{
	// 							to.Ptr("10.0.0.0/24")},
	// 							TotalNumberOfIPAddresses: to.Ptr("33554432"),
	// 						}
}
