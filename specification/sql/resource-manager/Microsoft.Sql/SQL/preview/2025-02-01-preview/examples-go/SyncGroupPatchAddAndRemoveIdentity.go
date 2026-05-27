package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/SyncGroupPatchAddAndRemoveIdentity.json
func ExampleSyncGroupsClient_BeginUpdate_updateASyncGroupWithNewUserAssignedIdentityAndRemoveOlderAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSyncGroupsClient().BeginUpdate(ctx, "syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", armsql.SyncGroup{
		Identity: &armsql.DataSyncParticipantIdentity{
			Type: to.Ptr(armsql.DataSyncParticipantIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armsql.DataSyncParticipantUserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi-2": {},
			},
		},
		Properties: &armsql.SyncGroupProperties{
			Interval:                 to.Ptr[int32](-1),
			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
			SyncDatabaseID:           to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
			UsePrivateLinkConnection: to.Ptr(false),
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
	// res = armsql.SyncGroupsClientUpdateResponse{
	// 	SyncGroup: armsql.SyncGroup{
	// 		Identity: &armsql.DataSyncParticipantIdentity{
	// 			Type: to.Ptr(armsql.DataSyncParticipantIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armsql.DataSyncParticipantUserAssignedIdentity{
	// 				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi-2": &armsql.DataSyncParticipantUserAssignedIdentity{
	// 					ClientID: to.Ptr("e09c8507-0000-0000-97e2-18c5beec59dc"),
	// 					PrincipalID: to.Ptr("e09c8507-0000-0000-97e2-18c5beec59dc"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armsql.SyncGroupProperties{
	// 			Interval: to.Ptr[int32](-1),
	// 			LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
	// 			ConflictResolutionPolicy: to.Ptr(armsql.SyncConflictResolutionPolicyHubWin),
	// 			SyncDatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328"),
	// 			SyncState: to.Ptr(armsql.SyncGroupStateNotReady),
	// 			UsePrivateLinkConnection: to.Ptr(false),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/syncgroupcrud-3521/providers/Microsoft.Sql/servers/syncgroupcrud-8475/databases/syncgroupcrud-4328/syncGroups/syncgroupcrud-3187"),
	// 		Name: to.Ptr("syncgroup"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/syncGroups"),
	// 	},
	// }
}
