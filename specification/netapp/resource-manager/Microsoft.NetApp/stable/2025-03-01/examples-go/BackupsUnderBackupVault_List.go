package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupsUnderBackupVault_List.json
func ExampleBackupsClient_NewListByVaultPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByVaultPager("myRG", "account1", "backupVault1", &armnetapp.BackupsClientListByVaultOptions{Filter: nil})
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
		// page.BackupsList = armnetapp.BackupsList{
		// 	Value: []*armnetapp.Backup{
		// 		{
		// 			Name: to.Ptr("account1/backupVault1/backup1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupVaults/backups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1/backups/backup1"),
		// 			Properties: &armnetapp.BackupProperties{
		// 				BackupPolicyResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolicies/policy1"),
		// 				BackupType: to.Ptr(armnetapp.BackupTypeManual),
		// 				CompletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33.000Z"); return t}()),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33.000Z"); return t}()),
		// 				Label: to.Ptr("myLabel"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Size: to.Ptr[int64](10011),
		// 				SnapshotCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33.000Z"); return t}()),
		// 				SnapshotName: to.Ptr("backup1"),
		// 				VolumeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPool/pool1/volumes/volume1"),
		// 			},
		// 	}},
		// }
	}
}
