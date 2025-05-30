package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/148c3b0b44f7789ced94859992493fafd0072f83/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureStorage/SoftDeletedContainers_List.json
func ExampleDeletedProtectionContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedProtectionContainersClient().NewListPager("testRg", "testVault", &armrecoveryservicesbackup.DeletedProtectionContainersClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureWorkload'")})
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
		// page.ProtectionContainerResourceList = armrecoveryservicesbackup.ProtectionContainerResourceList{
		// 	Value: []*armrecoveryservicesbackup.ProtectionContainerResource{
		// 		{
		// 			Name: to.Ptr("StorageContainer;Storage;testrg;suchandrtestsa125"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
		// 			ID: to.Ptr("/subscriptions/172424a4-d65f-421e-a8de-197d98aabeba/resourcegroups/testrg/providers/microsoft.recoveryservices/vaults/suchandr-test-vault-wcus/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;testrg;suchandrtestsa125"),
		// 			Properties: &armrecoveryservicesbackup.AzureStorageContainer{
		// 				BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 				ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 				FriendlyName: to.Ptr("suchandrtestsa125"),
		// 				HealthStatus: to.Ptr("Healthy"),
		// 				RegistrationStatus: to.Ptr("SoftDeleted"),
		// 				ProtectedItemCount: to.Ptr[int64](2),
		// 				SourceResourceID: to.Ptr("/subscriptions/172424a4-d65f-421e-a8de-197d98aabeba/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/suchandrtestsa125"),
		// 			},
		// 	}},
		// }
	}
}
