package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkRackSkus_ListBySubscription_MaximumSet_Gen.json
func ExampleNetworkRackSKUsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkRackSKUsClient().NewListBySubscriptionPager(nil)
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
		// page.NetworkRackSKUsListResult = armmanagednetworkfabric.NetworkRackSKUsListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkRackSKU{
		// 		{
		// 			Name: to.Ptr("networkRackSkuName"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkrackskus"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/providers/Microsoft.ManagedNetworkFabric/networkRackSkus/networkRackSkuName"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-16T06:09:38.603Z"); return t}()),
		// 				CreatedBy: to.Ptr("d1bd24c7-b27f-477e-86dd-939e107873d7"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-16T06:19:40.603Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkRackSKUProperties{
		// 				MaximumServerCount: to.Ptr[int32](19),
		// 				MaximumStorageCount: to.Ptr[int32](18),
		// 				MaximumUplinks: to.Ptr[int32](12),
		// 				NetworkDevices: []*armmanagednetworkfabric.NetworkDeviceRoleProperties{
		// 					{
		// 						NetworkDeviceSKUName: to.Ptr("DefaultSku"),
		// 						RackSlot: to.Ptr[int32](8),
		// 						RoleType: to.Ptr(armmanagednetworkfabric.NetworkDeviceRackRoleTypeCE),
		// 				}},
		// 				RoleName: to.Ptr(armmanagednetworkfabric.NetworkRackRoleNameComputeRack),
		// 			},
		// 	}},
		// }
	}
}
