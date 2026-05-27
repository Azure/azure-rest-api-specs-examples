package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/SyncMemberPatchWithIdentity.json
func ExampleSyncMembersClient_BeginUpdate_updateAnExistingSyncMemberWithUserAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncMembersClient().BeginUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879", armsql.SyncMember{
		Identity: &armsql.DataSyncParticipantIdentity{
			Type: to.Ptr(armsql.DataSyncParticipantIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armsql.DataSyncParticipantUserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi": {},
			},
		},
		Properties: &armsql.SyncMemberProperties{
			DatabaseType:                      to.Ptr(armsql.SyncMemberDbTypeAzureSQLDatabase),
			ServerName:                        to.Ptr("syncgroupcrud-3379.database.windows.net"),
			DatabaseName:                      to.Ptr("syncgroupcrud-7421"),
			SyncDirection:                     to.Ptr(armsql.SyncDirectionBidirectional),
			UsePrivateLinkConnection:          to.Ptr(true),
			SyncMemberAzureDatabaseResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.SyncMembersClientUpdateResponse{
	// 	SyncMember: armsql.SyncMember{
	// 		Identity: &armsql.DataSyncParticipantIdentity{
	// 			Type: to.Ptr(armsql.DataSyncParticipantIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armsql.DataSyncParticipantUserAssignedIdentity{
	// 				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi": &armsql.DataSyncParticipantUserAssignedIdentity{
	// 					ClientID: to.Ptr("e09c8507-0000-0000-97e2-18c5beec59dc"),
	// 					PrincipalID: to.Ptr("0c29d9b7-0ae2-4014-96ea-faf8e0cf2bc7"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armsql.SyncMemberProperties{
	// 			DatabaseType: to.Ptr(armsql.SyncMemberDbTypeAzureSQLDatabase),
	// 			ServerName: to.Ptr("syncgroupcrud-3379.database.windows.net"),
	// 			DatabaseName: to.Ptr("syncgroupcrud-7421"),
	// 			SyncDirection: to.Ptr(armsql.SyncDirectionBidirectional),
	// 			SyncState: to.Ptr(armsql.SyncMemberStateUnProvisioned),
	// 			UsePrivateLinkConnection: to.Ptr(true),
	// 			PrivateEndpointName: to.Ptr("PE_67FDBBD6-B2D8-4014-9CC6-C68ABBCFD481_syncmembercrud-4879"),
	// 			SyncMemberAzureDatabaseResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-65440/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187/syncMembers/syncmembercrud-4879"),
	// 		Name: to.Ptr("syncmembercrud-4879"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups/syncMembers"),
	// 	},
	// }
}
