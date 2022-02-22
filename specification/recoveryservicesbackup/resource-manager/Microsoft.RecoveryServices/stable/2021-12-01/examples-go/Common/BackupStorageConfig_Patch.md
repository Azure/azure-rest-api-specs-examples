Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.3.1/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/BackupStorageConfig_Patch.json
func ExampleBackupResourceStorageConfigsNonCRRClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupResourceStorageConfigsNonCRRClient("<subscription-id>", cred, nil)
	_, err = client.Patch(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.BackupResourceConfigResource{
			Properties: &armrecoveryservicesbackup.BackupResourceConfig{
				StorageType:      armrecoveryservicesbackup.StorageType("LocallyRedundant").ToPtr(),
				StorageTypeState: armrecoveryservicesbackup.StorageTypeState("Unlocked").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
