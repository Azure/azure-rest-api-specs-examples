package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_ListByResourceGroup_MaximumSet_Gen.json
func ExampleL2IsolationDomainsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewL2IsolationDomainsClient().NewListByResourceGroupPager("resourceGroupName", nil)
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
		// page.L2IsolationDomainsListResult = armmanagednetworkfabric.L2IsolationDomainsListResult{
		// 	Value: []*armmanagednetworkfabric.L2IsolationDomain{
		// 		{
		// 			Name: to.Ptr("wcpalyqmig"),
		// 			Type: to.Ptr("vvl"),
		// 			ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/example-l2domain"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-21T01:57:02.777Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T06:25:58.985Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserID"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armmanagednetworkfabric.L2IsolationDomainProperties{
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 				Mtu: to.Ptr[int32](1500),
		// 				NetworkFabricID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				VlanID: to.Ptr[int32](501),
		// 			},
		// 	}},
		// }
	}
}
