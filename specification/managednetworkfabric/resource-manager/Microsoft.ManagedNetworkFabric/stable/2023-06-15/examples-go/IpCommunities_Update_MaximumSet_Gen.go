package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_Update_MaximumSet_Gen.json
func ExampleIPCommunitiesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPCommunitiesClient().BeginUpdate(ctx, "example-rg", "example-ipcommunity", armmanagednetworkfabric.IPCommunityPatch{
		Properties: &armmanagednetworkfabric.IPCommunityPatchableProperties{
			IPCommunityRules: []*armmanagednetworkfabric.IPCommunityRule{
				{
					Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					CommunityMembers: []*string{
						to.Ptr("1:1")},
					SequenceNumber: to.Ptr[int64](4155123341),
					WellKnownCommunities: []*armmanagednetworkfabric.WellKnownCommunities{
						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesInternet)},
				}},
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
	// res.IPCommunity = armmanagednetworkfabric.IPCommunity{
	// 	Name: to.Ptr("example-ipcommunity"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipcommunities"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipcommunity"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:26:17.611Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:26:17.611Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPCommunityProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		IPCommunityRules: []*armmanagednetworkfabric.IPCommunityRule{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				CommunityMembers: []*string{
	// 					to.Ptr("1:1")},
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 					WellKnownCommunities: []*armmanagednetworkfabric.WellKnownCommunities{
	// 						to.Ptr(armmanagednetworkfabric.WellKnownCommunitiesInternet)},
	// 				}},
	// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 				ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 				ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 			},
	// 		}
}
