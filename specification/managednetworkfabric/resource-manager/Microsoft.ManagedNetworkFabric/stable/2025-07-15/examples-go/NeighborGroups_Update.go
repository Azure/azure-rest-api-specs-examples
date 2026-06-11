package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NeighborGroups_Update.json
func ExampleNeighborGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNeighborGroupsClient().BeginUpdate(ctx, "example-rg", "example-neighborGroup", armmanagednetworkfabric.NeighborGroupPatch{
		Tags: map[string]*string{
			"KeyId": to.Ptr("KeyValue"),
		},
		Identity: &armmanagednetworkfabric.ManagedServiceIdentityPatch{
			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
				"key8793": {},
			},
		},
		Properties: &armmanagednetworkfabric.NeighborGroupPatchProperties{
			Annotation: to.Ptr("Updating"),
			Destination: &armmanagednetworkfabric.NeighborGroupDestinationPatch{
				IPv4Addresses: []*string{
					to.Ptr("10.10.10.10"),
					to.Ptr("20.10.10.10"),
					to.Ptr("30.10.10.10"),
					to.Ptr("40.10.10.10"),
					to.Ptr("50.10.10.10"),
					to.Ptr("60.10.10.10"),
					to.Ptr("70.10.10.10"),
					to.Ptr("80.10.10.10"),
					to.Ptr("90.10.10.10"),
				},
				IPv6Addresses: []*string{
					to.Ptr("2F::/100"),
				},
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
	// res = armmanagednetworkfabric.NeighborGroupsClientUpdateResponse{
	// 	NeighborGroup: &armmanagednetworkfabric.NeighborGroup{
	// 		Properties: &armmanagednetworkfabric.NeighborGroupProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			Destination: &armmanagednetworkfabric.NeighborGroupDestination{
	// 				IPv4Addresses: []*string{
	// 					to.Ptr("10.10.10.10"),
	// 					to.Ptr("20.10.10.10"),
	// 					to.Ptr("30.10.10.10"),
	// 					to.Ptr("40.10.10.10"),
	// 					to.Ptr("50.10.10.10"),
	// 					to.Ptr("60.10.10.10"),
	// 					to.Ptr("70.10.10.10"),
	// 					to.Ptr("80.10.10.10"),
	// 					to.Ptr("90.10.10.10"),
	// 				},
	// 				IPv6Addresses: []*string{
	// 					to.Ptr("2F::/100"),
	// 				},
	// 			},
	// 			NetworkTapIDs: []*string{
	// 				to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-networkTap"),
	// 			},
	// 			NetworkTapRuleIDs: []*string{
	// 				to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-networkTapRule"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		},
	// 		Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 			Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
	// 				"key4876": &armmanagednetworkfabric.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 					ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
	// 		Name: to.Ptr("example-neighborGroup"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/neighborGroups"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
