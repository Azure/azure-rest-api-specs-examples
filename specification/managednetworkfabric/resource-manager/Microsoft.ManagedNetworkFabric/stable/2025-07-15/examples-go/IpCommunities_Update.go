package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/IpCommunities_Update.json
func ExampleIPCommunitiesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPCommunitiesClient().BeginUpdate(ctx, "example-rg", "example-ipcommunity", armmanagednetworkfabric.IPCommunityPatch{
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.IPCommunityPatchableProperties{
			IPCommunityRules: []*armmanagednetworkfabric.IPCommunityRule{
				{
					Action:         to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					SequenceNumber: to.Ptr[int64](4155123341),
					WellKnownCommunities: []*armmanagednetworkfabric.WellKnownCommunities{
						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesInternet),
					},
					CommunityMembers: []*string{
						to.Ptr("1:1"),
					},
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
	// res = armmanagednetworkfabric.IPCommunitiesClientUpdateResponse{
	// 	IPCommunity: &armmanagednetworkfabric.IPCommunity{
	// 		Properties: &armmanagednetworkfabric.IPCommunityProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			IPCommunityRules: []*armmanagednetworkfabric.IPCommunityRule{
	// 				{
	// 					Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 					WellKnownCommunities: []*armmanagednetworkfabric.WellKnownCommunities{
	// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesInternet),
	// 					},
	// 					CommunityMembers: []*string{
	// 						to.Ptr("1:1"),
	// 					},
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipcommunity"),
	// 		Name: to.Ptr("example-ipcommunity"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/ipcommunities"),
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
