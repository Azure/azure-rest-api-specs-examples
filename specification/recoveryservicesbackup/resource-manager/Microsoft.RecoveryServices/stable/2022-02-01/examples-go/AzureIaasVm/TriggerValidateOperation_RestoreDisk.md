Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/TriggerValidateOperation_RestoreDisk.json
func ExampleValidateOperationClient_BeginTrigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewValidateOperationClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTrigger(ctx,
		"testVault",
		"testRG",
		&armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
			ObjectType: to.Ptr("ValidateIaasVMRestoreOperationRequest"),
			RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
				ObjectType:            to.Ptr("IaasVMRestoreRequest"),
				CreateNewCloudService: to.Ptr(true),
				EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
					EncryptionEnabled: to.Ptr(false),
				},
				IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
					IsSystemAssignedIdentity:  to.Ptr(false),
					ManagedIdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"),
				},
				OriginalStorageAccountOption: to.Ptr(false),
				RecoveryPointID:              to.Ptr("348916168024334"),
				RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
				Region:                       to.Ptr("southeastasia"),
				SourceResourceID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
				StorageAccountID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
