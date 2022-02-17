Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.3.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureStorage/ProtectionContainers_Register.json
func ExampleProtectionContainersClient_Register() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionContainersClient("<subscription-id>", cred, nil)
	res, err := client.Register(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		armrecoveryservicesbackup.ProtectionContainerResource{
			Properties: &armrecoveryservicesbackup.AzureStorageContainer{
				BackupManagementType:      armrecoveryservicesbackup.BackupManagementType("AzureStorage").ToPtr(),
				ContainerType:             armrecoveryservicesbackup.ContainerType("StorageContainer").ToPtr(),
				FriendlyName:              to.StringPtr("<friendly-name>"),
				AcquireStorageAccountLock: armrecoveryservicesbackup.AcquireStorageAccountLock("Acquire").ToPtr(),
				SourceResourceID:          to.StringPtr("<source-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProtectionContainersClientRegisterResult)
}
```
