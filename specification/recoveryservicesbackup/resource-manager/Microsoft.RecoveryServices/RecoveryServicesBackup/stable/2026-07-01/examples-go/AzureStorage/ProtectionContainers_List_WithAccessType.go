package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-07-01/AzureStorage/ProtectionContainers_List_WithAccessType.json
func ExampleBackupProtectionContainersClient_NewListPager_listBackupProtectionContainersWithAccessType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupProtectionContainersClient().NewListPager("swaggertestvault", "SwaggerTestRg", &armrecoveryservicesbackup.BackupProtectionContainersClientListOptions{
		Filter: to.Ptr("backupManagementType eq 'AzureStorage'")})
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
		// page = armrecoveryservicesbackup.BackupProtectionContainersClientListResponse{
		// 	ProtectionContainerResourceList: armrecoveryservicesbackup.ProtectionContainerResourceList{
		// 		Value: []*armrecoveryservicesbackup.ProtectionContainerResource{
		// 			{
		// 				Name: to.Ptr("StorageContainer;Storage;SwaggerTestRg;swaggertestsakeybased"),
		// 				Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;SwaggerTestRg;swaggertestsakeybased"),
		// 				Properties: &armrecoveryservicesbackup.AzureStorageContainer{
		// 					BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 					ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 					FriendlyName: to.Ptr("swaggertestsakeybased"),
		// 					HealthStatus: to.Ptr("Healthy"),
		// 					ProtectedItemCount: to.Ptr[int64](2),
		// 					RegistrationStatus: to.Ptr("Registered"),
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsakeybased"),
		// 					AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
		// 					AccessType: to.Ptr(armrecoveryservicesbackup.AccessTypeKeyBased),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("StorageContainer;Storage;SwaggerTestRg;swaggertestsasami"),
		// 				Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;SwaggerTestRg;swaggertestsasami"),
		// 				Properties: &armrecoveryservicesbackup.AzureStorageContainer{
		// 					BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 					ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 					FriendlyName: to.Ptr("swaggertestsasami"),
		// 					HealthStatus: to.Ptr("Healthy"),
		// 					ProtectedItemCount: to.Ptr[int64](1),
		// 					RegistrationStatus: to.Ptr("Registered"),
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsasami"),
		// 					AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
		// 					AccessType: to.Ptr(armrecoveryservicesbackup.AccessTypeIdentityBased),
		// 					IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
		// 						IsSystemAssignedIdentity: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("StorageContainer;Storage;SwaggerTestRg;swaggertestsauami"),
		// 				Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/swaggertestvault/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;SwaggerTestRg;swaggertestsauami"),
		// 				Properties: &armrecoveryservicesbackup.AzureStorageContainer{
		// 					BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 					ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 					FriendlyName: to.Ptr("swaggertestsauami"),
		// 					HealthStatus: to.Ptr("Healthy"),
		// 					ProtectedItemCount: to.Ptr[int64](1),
		// 					RegistrationStatus: to.Ptr("Registered"),
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsauami"),
		// 					AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockNotAcquire),
		// 					AccessType: to.Ptr(armrecoveryservicesbackup.AccessTypeIdentityBased),
		// 					IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
		// 						IsSystemAssignedIdentity: to.Ptr(false),
		// 						ManagedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/swaggertestuami"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
