package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/Dpm/BackupEngines_Get.json
func ExampleBackupEnginesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupEnginesClient().Get(ctx, "testVault", "testRG", "testServer", &armrecoveryservicesbackup.BackupEnginesClientGetOptions{Filter: nil,
		SkipToken: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupEngineBaseResource = armrecoveryservicesbackup.BackupEngineBaseResource{
	// 	Name: to.Ptr("testServer"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupEngines"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupEngines/testServer"),
	// 	Properties: &armrecoveryservicesbackup.DpmBackupEngine{
	// 		AzureBackupAgentVersion: to.Ptr("2.0.9532.0"),
	// 		BackupEngineState: to.Ptr("Active"),
	// 		BackupEngineType: to.Ptr(armrecoveryservicesbackup.BackupEngineTypeDpmBackupEngine),
	// 		DpmVersion: to.Ptr("5.1.348.0"),
	// 		ExtendedInfo: &armrecoveryservicesbackup.BackupEngineExtendedInfo{
	// 			AvailableDiskSpace: to.Ptr[float64](50),
	// 			DiskCount: to.Ptr[int32](5),
	// 			ProtectedItemsCount: to.Ptr[int32](35),
	// 			ProtectedServersCount: to.Ptr[int32](21),
	// 			UsedDiskSpace: to.Ptr[float64](20),
	// 		},
	// 		FriendlyName: to.Ptr("testServer"),
	// 		IsAzureBackupAgentUpgradeAvailable: to.Ptr(false),
	// 		IsDpmUpgradeAvailable: to.Ptr(false),
	// 		RegistrationStatus: to.Ptr("Registered"),
	// 	},
	// }
}
