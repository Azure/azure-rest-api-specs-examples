package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticBackups_ListByVault.json
func ExampleElasticBackupsClient_NewListByVaultPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticBackupsClient().NewListByVaultPager("myRG", "account1", "backupVault1", nil)
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
		// page = armnetapp.ElasticBackupsClientListByVaultResponse{
		// 	ElasticBackupListResult: armnetapp.ElasticBackupListResult{
		// 		Value: []*armnetapp.ElasticBackup{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticBackupVaults/backupVault1/elasticBackups/backup1"),
		// 				Name: to.Ptr("account1/backupVault1/backup1"),
		// 				Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticBackupVaults/elasticBackups"),
		// 				Properties: &armnetapp.ElasticBackupProperties{
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
		// 					SnapshotCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
		// 					CompletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					Size: to.Ptr[int64](10011),
		// 					Label: to.Ptr("myLabel"),
		// 					BackupType: to.Ptr(armnetapp.ElasticBackupTypeManual),
		// 					SnapshotUsage: to.Ptr(armnetapp.SnapshotUsage("Existing")),
		// 					ElasticSnapshotResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1/elasticSnapshots/snap1"),
		// 					ElasticVolumeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1"),
		// 					ElasticBackupPolicyResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticBackupPolicies/policy1"),
		// 					VolumeSize: to.Ptr(armnetapp.VolumeSizeLarge),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
