package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureStorage/ProtectionContainers_Register.json
func ExampleProtectionContainersClient_Register() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionContainersClient().Register(ctx, "swaggertestvault", "SwaggerTestRg", "Azure", "StorageContainer;Storage;SwaggerTestRg;swaggertestsa", armrecoveryservicesbackup.ProtectionContainerResource{
		Properties: &armrecoveryservicesbackup.AzureStorageContainer{
			BackupManagementType:      to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
			ContainerType:             to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
			FriendlyName:              to.Ptr("swaggertestsa"),
			AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
			SourceResourceID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionContainerResource = armrecoveryservicesbackup.ProtectionContainerResource{
	// 	Name: to.Ptr("StorageContainer;Storage;SwaggerTestRg;swaggertestsa"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;SwaggerTestRg;swaggertestsa"),
	// 	Properties: &armrecoveryservicesbackup.AzureStorageContainer{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
	// 		ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
	// 		FriendlyName: to.Ptr("swaggertestsa"),
	// 		HealthStatus: to.Ptr("Healthy"),
	// 		RegistrationStatus: to.Ptr("Registered"),
	// 		AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
	// 		ProtectedItemCount: to.Ptr[int64](0),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa"),
	// 	},
	// }
}
