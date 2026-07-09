package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/ProtectionGroups_Backup_MaximumSet_Gen.json
func ExampleProtectionGroupsClient_Backup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionGroupsClient().Backup(ctx, "rgcommvault", "sample-cloudAccountName", "sample-protectionGroupName", armcommvaultcontentstore.BackupProtectionGroupRequest{
		VMList: []*armcommvaultcontentstore.VMListItem{
			{
				VMGUID: to.Ptr("40000000-aaaa-4bbb-8ccc-000000000000"),
			},
		},
		BackupOptions: &armcommvaultcontentstore.BackupOptions{
			BackupLevel:               to.Ptr(armcommvaultcontentstore.BackupLevelFull),
			JobDescription:            to.Ptr("Ad-hoc backup job"),
			BackupCopyImmediately:     to.Ptr(true),
			RunSnapShotBackup:         to.Ptr(true),
			NotifyUserOnJobCompletion: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.ProtectionGroupsClientBackupResponse{
	// 	BackupProtectionGroupResponse: armcommvaultcontentstore.BackupProtectionGroupResponse{
	// 		TaskID: to.Ptr[int32](9),
	// 		JobIDs: []*string{
	// 			to.Ptr("c"),
	// 		},
	// 	},
	// }
}
