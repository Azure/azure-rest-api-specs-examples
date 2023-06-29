package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Get_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPExtendedCommunitiesClient().Get(ctx, "rgIpExtendedCommunityLists", "example_ipExtendedCommunity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPExtendedCommunity = armmanagednetworkfabric.IPExtendedCommunity{
	// 	Name: to.Ptr("example_ipExtendedCommunity"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example_ipExtendedCommunity"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T14:44:01.092Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T14:44:01.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("EastUs"),
	// 	Tags: map[string]*string{
	// 		"key5054": to.Ptr("key"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
	// 		Annotation: to.Ptr("annotationValue"),
	// 		Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		RouteTargets: []*string{
	// 			to.Ptr("1234:5678")},
	// 		},
	// 	}
}
