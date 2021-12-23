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

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk.json
func ExampleOperationClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewOperationClient("<subscription-id>", cred, nil)
	_, err = client.Validate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		&armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
			ValidateRestoreOperationRequest: armrecoveryservicesbackup.ValidateRestoreOperationRequest{
				ValidateOperationRequest: armrecoveryservicesbackup.ValidateOperationRequest{
					ObjectType: to.StringPtr("<object-type>"),
				},
				RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
					RestoreRequest: armrecoveryservicesbackup.RestoreRequest{
						ObjectType: to.StringPtr("<object-type>"),
					},
					CreateNewCloudService: to.BoolPtr(true),
					EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
						EncryptionEnabled: to.BoolPtr(false),
					},
					IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
						IsSystemAssignedIdentity:  to.BoolPtr(false),
						ManagedIdentityResourceID: to.StringPtr("<managed-identity-resource-id>"),
					},
					OriginalStorageAccountOption: to.BoolPtr(false),
					RecoveryPointID:              to.StringPtr("<recovery-point-id>"),
					RecoveryType:                 armrecoveryservicesbackup.RecoveryTypeRestoreDisks.ToPtr(),
					Region:                       to.StringPtr("<region>"),
					SourceResourceID:             to.StringPtr("<source-resource-id>"),
					StorageAccountID:             to.StringPtr("<storage-account-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
