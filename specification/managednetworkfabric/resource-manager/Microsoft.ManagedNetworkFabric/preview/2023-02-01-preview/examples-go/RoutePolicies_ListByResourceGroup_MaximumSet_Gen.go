package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/RoutePolicies_ListByResourceGroup_MaximumSet_Gen.json
func ExampleRoutePoliciesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoutePoliciesClient().NewListByResourceGroupPager("rgRoutePolicies", nil)
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
		// page.RoutePoliciesListResult = armmanagednetworkfabric.RoutePoliciesListResult{
		// 	Value: []*armmanagednetworkfabric.RoutePolicy{
		// 		{
		// 			Name: to.Ptr("routePolicyName"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/routePolicies"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T12:42:03.945Z"); return t}()),
		// 				CreatedBy: to.Ptr("User"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T12:42:03.946Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("email@address.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("EastUS"),
		// 			Tags: map[string]*string{
		// 				"key8254": to.Ptr(""),
		// 			},
		// 			Properties: &armmanagednetworkfabric.RoutePolicyProperties{
		// 				Annotation: to.Ptr("annotationValue"),
		// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				Statements: []*armmanagednetworkfabric.RoutePolicyStatementProperties{
		// 					{
		// 						Annotation: to.Ptr("annotationValue"),
		// 						Action: &armmanagednetworkfabric.StatementActionProperties{
		// 							ActionType: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
		// 							IPCommunityProperties: &armmanagednetworkfabric.ActionIPCommunityProperties{
		// 								Add: &armmanagednetworkfabric.IPCommunityIDList{
		// 									IPCommunityIDs: []*string{
		// 										to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName")},
		// 									},
		// 									Delete: &armmanagednetworkfabric.IPCommunityIDList{
		// 										IPCommunityIDs: []*string{
		// 											to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName")},
		// 										},
		// 										Set: &armmanagednetworkfabric.IPCommunityIDList{
		// 											IPCommunityIDs: []*string{
		// 												to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName")},
		// 											},
		// 										},
		// 										IPExtendedCommunityProperties: &armmanagednetworkfabric.ActionIPExtendedCommunityProperties{
		// 											Add: &armmanagednetworkfabric.IPExtendedCommunityIDList{
		// 												IPExtendedCommunityIDs: []*string{
		// 													to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName")},
		// 												},
		// 												Delete: &armmanagednetworkfabric.IPExtendedCommunityIDList{
		// 													IPExtendedCommunityIDs: []*string{
		// 														to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName")},
		// 													},
		// 													Set: &armmanagednetworkfabric.IPExtendedCommunityIDList{
		// 														IPExtendedCommunityIDs: []*string{
		// 															to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName")},
		// 														},
		// 													},
		// 													LocalPreference: to.Ptr[int64](20),
		// 												},
		// 												Condition: &armmanagednetworkfabric.StatementConditionProperties{
		// 													IPCommunityIDs: []*string{
		// 														to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName")},
		// 														IPExtendedCommunityIDs: []*string{
		// 															to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName")},
		// 															IPPrefixID: to.Ptr("subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
		// 														},
		// 														SequenceNumber: to.Ptr[int64](7),
		// 												}},
		// 											},
		// 									}},
		// 								}
	}
}
