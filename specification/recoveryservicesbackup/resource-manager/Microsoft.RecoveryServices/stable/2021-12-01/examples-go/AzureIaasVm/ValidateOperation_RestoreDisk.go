package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk.json
func ExampleOperationClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Validate(ctx,
		"<vault-name>",
		"<resource-group-name>",
		&armrecoveryservicesbackup.ValidateIaasVMRestoreOperationRequest{
			ObjectType: to.Ptr("<object-type>"),
			RestoreRequest: &armrecoveryservicesbackup.IaasVMRestoreRequest{
				ObjectType:            to.Ptr("<object-type>"),
				CreateNewCloudService: to.Ptr(true),
				EncryptionDetails: &armrecoveryservicesbackup.EncryptionDetails{
					EncryptionEnabled: to.Ptr(false),
				},
				IdentityInfo: &armrecoveryservicesbackup.IdentityInfo{
					IsSystemAssignedIdentity:  to.Ptr(false),
					ManagedIdentityResourceID: to.Ptr("<managed-identity-resource-id>"),
				},
				OriginalStorageAccountOption: to.Ptr(false),
				RecoveryPointID:              to.Ptr("<recovery-point-id>"),
				RecoveryType:                 to.Ptr(armrecoveryservicesbackup.RecoveryTypeRestoreDisks),
				Region:                       to.Ptr("<region>"),
				SourceResourceID:             to.Ptr("<source-resource-id>"),
				StorageAccountID:             to.Ptr("<storage-account-id>"),
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
