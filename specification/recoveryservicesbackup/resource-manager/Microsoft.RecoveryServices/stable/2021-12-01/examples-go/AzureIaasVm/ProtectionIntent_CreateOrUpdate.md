Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.3.1/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
func ExampleProtectionIntentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewProtectionIntentClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<intent-object-name>",
		armrecoveryservicesbackup.ProtectionIntentResource{
			Properties: &armrecoveryservicesbackup.AzureResourceProtectionIntent{
				PolicyID:                 to.StringPtr("<policy-id>"),
				ProtectionIntentItemType: armrecoveryservicesbackup.ProtectionIntentItemType("AzureResourceItem").ToPtr(),
				SourceResourceID:         to.StringPtr("<source-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProtectionIntentClientCreateOrUpdateResult)
}
```
