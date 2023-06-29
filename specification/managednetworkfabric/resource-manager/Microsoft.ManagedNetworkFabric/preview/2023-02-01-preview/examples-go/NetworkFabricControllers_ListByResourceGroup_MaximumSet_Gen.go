package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_ListByResourceGroup_MaximumSet_Gen.json
func ExampleNetworkFabricControllersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFabricControllersClient().NewListByResourceGroupPager("resourceGroupName", nil)
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
		// page.NetworkFabricControllersListResult = armmanagednetworkfabric.NetworkFabricControllersListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkFabricController{
		// 		{
		// 			Name: to.Ptr("NetworkFabricName"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkfabriccontrollers"),
		// 			ID: to.Ptr("/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/networkFabricControllerName"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-16T05:01:43.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@company.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-16T05:01:43.008Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armmanagednetworkfabric.NetworkFabricControllerProperties{
		// 				InfrastructureExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
		// 					{
		// 						ExpressRouteCircuitID: to.Ptr("/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
		// 				}},
		// 				WorkloadExpressRouteConnections: []*armmanagednetworkfabric.ExpressRouteConnectionInformation{
		// 					{
		// 						ExpressRouteCircuitID: to.Ptr("/subscriptions/xxxxx/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/expressRouteCircuitName"),
		// 				}},
		// 				InfrastructureServices: &armmanagednetworkfabric.InfrastructureServices{
		// 					IPv4AddressSpaces: []*string{
		// 						to.Ptr("172.253.0.0/19")},
		// 					},
		// 					ManagedResourceGroupConfiguration: &armmanagednetworkfabric.ManagedResourceGroupConfiguration{
		// 						Name: to.Ptr("managedResourceGroupName"),
		// 						Location: to.Ptr("eastus"),
		// 					},
		// 					NetworkFabricIDs: []*string{
		// 					},
		// 					WorkloadServices: &armmanagednetworkfabric.WorkloadServices{
		// 						IPv4AddressSpaces: []*string{
		// 							to.Ptr("172.253.28.0/22")},
		// 							IPv6AddressSpaces: []*string{
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
