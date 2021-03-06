package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureStorage/ProtectionContainers_Register.json
func ExampleProtectionContainersClient_Register() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewProtectionContainersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Register(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		armrecoveryservicesbackup.ProtectionContainerResource{
			Properties: &armrecoveryservicesbackup.AzureStorageContainer{
				BackupManagementType:      to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
				ContainerType:             to.Ptr(armrecoveryservicesbackup.ContainerTypeStorageContainer),
				FriendlyName:              to.Ptr("<friendly-name>"),
				AcquireStorageAccountLock: to.Ptr(armrecoveryservicesbackup.AcquireStorageAccountLockAcquire),
				SourceResourceID:          to.Ptr("<source-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
