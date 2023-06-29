package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_ListBySubscription_MaximumSet_Gen.json
func ExampleL3IsolationDomainsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewL3IsolationDomainsClient().NewListBySubscriptionPager(nil)
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
		// page.L3IsolationDomainsListResult = armmanagednetworkfabric.L3IsolationDomainsListResult{
		// 	Value: []*armmanagednetworkfabric.L3IsolationDomain{
		// 		{
		// 			Name: to.Ptr("example-l3domain"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/l3isolationdomains"),
		// 			ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserId"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armmanagednetworkfabric.L3IsolationDomainProperties{
		// 				Description: to.Ptr("creating L3 isolation domain"),
		// 				AggregateRouteConfiguration: &armmanagednetworkfabric.AggregateRouteConfiguration{
		// 					IPv4Routes: []*armmanagednetworkfabric.AggregateRoute{
		// 						{
		// 							Prefix: to.Ptr("10.0.0.0/24"),
		// 					}},
		// 					IPv6Routes: []*armmanagednetworkfabric.AggregateRoute{
		// 						{
		// 							Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/29"),
		// 					}},
		// 				},
		// 				ConnectedSubnetRoutePolicy: &armmanagednetworkfabric.L3IsolationDomainPatchPropertiesConnectedSubnetRoutePolicy{
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 					ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 				},
		// 				RedistributeConnectedSubnets: to.Ptr(armmanagednetworkfabric.RedistributeConnectedSubnetsTrue),
		// 				RedistributeStaticRoutes: to.Ptr(armmanagednetworkfabric.RedistributeStaticRoutesFalse),
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 				DisabledOnResources: []*string{
		// 					to.Ptr("")},
		// 					NetworkFabricID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName"),
		// 					OptionBDisabledOnResources: []*string{
		// 						to.Ptr("")},
		// 						ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}
