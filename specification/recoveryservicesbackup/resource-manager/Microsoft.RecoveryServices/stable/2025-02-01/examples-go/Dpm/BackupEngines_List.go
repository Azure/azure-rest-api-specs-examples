package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Dpm/BackupEngines_List.json
func ExampleBackupEnginesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupEnginesClient().NewListPager("testVault", "testRG", &armrecoveryservicesbackup.BackupEnginesClientListOptions{Filter: nil,
		SkipToken: nil,
	})
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
		// page.BackupEngineBaseResourceList = armrecoveryservicesbackup.BackupEngineBaseResourceList{
		// 	Value: []*armrecoveryservicesbackup.BackupEngineBaseResource{
		// 		{
		// 			Name: to.Ptr("testServer1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupEngines"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupEngines/testServer1"),
		// 			Properties: &armrecoveryservicesbackup.DpmBackupEngine{
		// 				AzureBackupAgentVersion: to.Ptr("2.0.9532.0"),
		// 				BackupEngineState: to.Ptr("Active"),
		// 				BackupEngineType: to.Ptr(armrecoveryservicesbackup.BackupEngineTypeDpmBackupEngine),
		// 				DpmVersion: to.Ptr("5.1.348.0"),
		// 				ExtendedInfo: &armrecoveryservicesbackup.BackupEngineExtendedInfo{
		// 					AvailableDiskSpace: to.Ptr[float64](50),
		// 					DiskCount: to.Ptr[int32](5),
		// 					ProtectedItemsCount: to.Ptr[int32](35),
		// 					ProtectedServersCount: to.Ptr[int32](21),
		// 					UsedDiskSpace: to.Ptr[float64](20),
		// 				},
		// 				FriendlyName: to.Ptr("testServer1"),
		// 				IsAzureBackupAgentUpgradeAvailable: to.Ptr(false),
		// 				IsDpmUpgradeAvailable: to.Ptr(false),
		// 				RegistrationStatus: to.Ptr("Registered"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testServer5"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupEngines"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupEngines/testServer5"),
		// 			Properties: &armrecoveryservicesbackup.DpmBackupEngine{
		// 				AzureBackupAgentVersion: to.Ptr("2.0.9530.0"),
		// 				BackupEngineState: to.Ptr("Active"),
		// 				BackupEngineType: to.Ptr(armrecoveryservicesbackup.BackupEngineTypeDpmBackupEngine),
		// 				DpmVersion: to.Ptr("5.1.348.0"),
		// 				ExtendedInfo: &armrecoveryservicesbackup.BackupEngineExtendedInfo{
		// 					AvailableDiskSpace: to.Ptr[float64](50),
		// 					DiskCount: to.Ptr[int32](5),
		// 					ProtectedItemsCount: to.Ptr[int32](35),
		// 					ProtectedServersCount: to.Ptr[int32](21),
		// 					UsedDiskSpace: to.Ptr[float64](20),
		// 				},
		// 				FriendlyName: to.Ptr("testServer5"),
		// 				IsAzureBackupAgentUpgradeAvailable: to.Ptr(false),
		// 				IsDpmUpgradeAvailable: to.Ptr(false),
		// 				RegistrationStatus: to.Ptr("Registered"),
		// 			},
		// 	}},
		// }
	}
}
