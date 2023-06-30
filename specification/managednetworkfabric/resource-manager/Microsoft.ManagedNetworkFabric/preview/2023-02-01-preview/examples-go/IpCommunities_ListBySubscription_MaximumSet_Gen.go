package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_ListBySubscription_MaximumSet_Gen.json
func ExampleIPCommunitiesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPCommunitiesClient().NewListBySubscriptionPager(nil)
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
		// page.IPCommunitiesListResult = armmanagednetworkfabric.IPCommunitiesListResult{
		// 	Value: []*armmanagednetworkfabric.IPCommunity{
		// 		{
		// 			Name: to.Ptr("example-ipCommunity"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/ipCommunities"),
		// 			ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T14:05:16.579Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@email.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T14:05:16.579Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserId"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("EastUS"),
		// 			Tags: map[string]*string{
		// 				"key2814": to.Ptr(""),
		// 			},
		// 			Properties: &armmanagednetworkfabric.IPCommunityProperties{
		// 				Annotation: to.Ptr("annotationValue"),
		// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
		// 				CommunityMembers: []*string{
		// 					to.Ptr("1234:5678")},
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 					WellKnownCommunities: []*armmanagednetworkfabric.WellKnownCommunities{
		// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesInternet),
		// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesLocalAS),
		// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesNoExport),
		// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesGShut)},
		// 					},
		// 			}},
		// 		}
	}
}
