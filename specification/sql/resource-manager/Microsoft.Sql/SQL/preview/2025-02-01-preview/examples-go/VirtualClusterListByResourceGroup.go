package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/VirtualClusterListByResourceGroup.json
func ExampleVirtualClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualClustersClient().NewListByResourceGroupPager("testrg", nil)
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
		// page = armsql.VirtualClustersClientListByResourceGroupResponse{
		// 	VirtualClusterListResult: armsql.VirtualClusterListResult{
		// 		Value: []*armsql.VirtualCluster{
		// 			{
		// 				Name: to.Ptr("VirtualClustercltest"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/VirtualClustercltest"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 						to.Ptr("/subscriptions/65dc2520-d3b9-4d11-b638-f3c41f962550/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl"),
		// 					},
		// 					SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest"),
		// 					Version: to.Ptr("1.0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vc2"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc2"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 					},
		// 					SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest2"),
		// 					Version: to.Ptr("2.0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vc1"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc1"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 						to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl"),
		// 						to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/mi"),
		// 					},
		// 					SubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/cltest"),
		// 					Version: to.Ptr("2.0"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tkey": to.Ptr("tvalue3"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
