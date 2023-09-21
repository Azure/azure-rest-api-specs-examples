package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/VirtualClusterList.json
func ExampleVirtualClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualClustersClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualClusterListResult = armsql.VirtualClusterListResult{
		// 	Value: []*armsql.VirtualCluster{
		// 		{
		// 			Name: to.Ptr("VirtualClustercltest"),
		// 			Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 			ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/VirtualClustercltest"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.VirtualClusterProperties{
		// 				ChildResources: []*string{
		// 					to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl")},
		// 					SubnetID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest"),
		// 					Version: to.Ptr("1.0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vc2"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc2"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 					},
		// 					SubnetID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest2"),
		// 					Version: to.Ptr("2.0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vc1"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc1"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 						to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl"),
		// 						to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/mi")},
		// 						SubnetID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest"),
		// 						Version: to.Ptr("2.0"),
		// 					},
		// 			}},
		// 		}
	}
}
