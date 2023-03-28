package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterListByResourceGroup.json
func ExampleVirtualClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.VirtualClusterListResult = armsql.VirtualClusterListResult{
		// 	Value: []*armsql.VirtualCluster{
		// 		{
		// 			Name: to.Ptr("vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
		// 			Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 			ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2"),
		// 			Location: to.Ptr("onebox"),
		// 			Properties: &armsql.VirtualClusterProperties{
		// 				ChildResources: []*string{
		// 					to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance1"),
		// 					to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance2")},
		// 					Family: to.Ptr("Gen4"),
		// 					MaintenanceConfigurationID: to.Ptr("/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1"),
		// 					SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vc-subnet1-14b795bd-9c8f-46ec-adb8-2b8eff56ac16"),
		// 				Type: to.Ptr("Microsoft.Sql/virtualClusters"),
		// 				ID: to.Ptr("/subscriptions/20d7082a-0fc7-4468-82bd-542694d5042b/resourceGroups/testrg/providers/Microsoft.Sql/virtualClusters/vc-subnet2-14b795bd-9c8f-46ec-adb8-2b8eff56ac16"),
		// 				Location: to.Ptr("onebox"),
		// 				Properties: &armsql.VirtualClusterProperties{
		// 					ChildResources: []*string{
		// 						to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance3"),
		// 						to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance4")},
		// 						Family: to.Ptr("Gen4"),
		// 						MaintenanceConfigurationID: to.Ptr("/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1"),
		// 						SubnetID: to.Ptr("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/subnet2"),
		// 					},
		// 			}},
		// 		}
	}
}
