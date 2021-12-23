Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.1.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureStorage/ProtectionContainers_Register.json
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
			Properties: &armrecoveryservicesbackup.AzureVMAppContainerProtectionContainer{
				AzureWorkloadContainer: armrecoveryservicesbackup.AzureWorkloadContainer{
					ProtectionContainer: armrecoveryservicesbackup.ProtectionContainer{
						BackupManagementType: armrecoveryservicesbackup.BackupManagementTypeAzureWorkload.ToPtr(),
						ContainerType:        armrecoveryservicesbackup.ContainerTypeVMAppContainer.ToPtr(),
						FriendlyName:         to.StringPtr("<friendly-name>"),
					},
					SourceResourceID: to.StringPtr("<source-resource-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProtectionContainerResource.ID: %s\n", *res.ID)
}
```
