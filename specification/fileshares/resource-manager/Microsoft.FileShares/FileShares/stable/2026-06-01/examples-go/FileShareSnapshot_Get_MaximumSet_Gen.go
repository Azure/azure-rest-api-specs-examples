package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShareSnapshot_Get_MaximumSet_Gen.json
func ExampleFileShareSnapshotsClient_GetFileShareSnapshot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileShareSnapshotsClient().GetFileShareSnapshot(ctx, "rgfileshares", "fileshare", "testfilesharesnapshot", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armfileshares.FileShareSnapshotsClientGetFileShareSnapshotResponse{
	// 	FileShareSnapshot: armfileshares.FileShareSnapshot{
	// 		Properties: &armfileshares.FileShareSnapshotProperties{
	// 			SnapshotTime: to.Ptr("wgujspthevkfo"),
	// 			InitiatorID: to.Ptr("ghublxbeifhvhwokvhppbw"),
	// 			Metadata: map[string]*string{
	// 				"key9372": to.Ptr("jtc"),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/rg1Network/providers/Microsoft.FileShares/fileShares/testfileshare/fileShareSnapshots/testfilesharesnapshot"),
	// 		Name: to.Ptr("testfilesharesnapshot"),
	// 		Type: to.Ptr("Microsoft.FileShares/fileShares/fileShareSnapshots"),
	// 		SystemData: &armfileshares.SystemData{
	// 			CreatedBy: to.Ptr("moat"),
	// 			CreatedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("dxukcadzapywjcgnxsqchaa"),
	// 			LastModifiedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 		},
	// 	},
	// }
}
