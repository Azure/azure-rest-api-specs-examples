package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Update_MaximumSet_Gen.json
func ExampleNeighborGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNeighborGroupsClient().BeginUpdate(ctx, "example-rg", "example-neighborGroup", armmanagednetworkfabric.NeighborGroupPatch{
		Tags: map[string]*string{
			"key8107": to.Ptr("2345"),
		},
		Properties: &armmanagednetworkfabric.NeighborGroupPatchProperties{
			Annotation: to.Ptr("Updating"),
			Destination: &armmanagednetworkfabric.NeighborGroupDestination{
				IPv4Addresses: []*string{
					to.Ptr("10.10.10.10"),
					to.Ptr("20.10.10.10"),
					to.Ptr("30.10.10.10"),
					to.Ptr("40.10.10.10"),
					to.Ptr("50.10.10.10"),
					to.Ptr("60.10.10.10"),
					to.Ptr("70.10.10.10"),
					to.Ptr("80.10.10.10"),
					to.Ptr("90.10.10.10")},
				IPv6Addresses: []*string{
					to.Ptr("2F::/100")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NeighborGroup = armmanagednetworkfabric.NeighborGroup{
	// 	Name: to.Ptr("example-neighborGroup"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/neighborGroups"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-23T05:49:59.193Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-23T05:49:59.194Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key8107": to.Ptr("2345"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NeighborGroupProperties{
	// 		Annotation: to.Ptr("ivjs"),
	// 		Destination: &armmanagednetworkfabric.NeighborGroupDestination{
	// 			IPv4Addresses: []*string{
	// 				to.Ptr("10.10.10.10"),
	// 				to.Ptr("20.10.10.10"),
	// 				to.Ptr("30.10.10.10"),
	// 				to.Ptr("40.10.10.10"),
	// 				to.Ptr("50.10.10.10"),
	// 				to.Ptr("60.10.10.10"),
	// 				to.Ptr("70.10.10.10"),
	// 				to.Ptr("80.10.10.10"),
	// 				to.Ptr("90.10.10.10")},
	// 				IPv6Addresses: []*string{
	// 					to.Ptr("2F::/100")},
	// 				},
	// 				NetworkTapIDs: []*string{
	// 					to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-networkTap")},
	// 					NetworkTapRuleIDs: []*string{
	// 						to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-networkTapRule")},
	// 						ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 					},
	// 				}
}
