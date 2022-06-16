package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/ValidateRestore.json
func ExampleBackupInstancesClient_BeginValidateForRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginValidateForRestore(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-instance-name>",
		armdataprotection.ValidateRestoreRequestObject{
			RestoreRequestObject: &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
				AzureBackupRestoreRequest: armdataprotection.AzureBackupRestoreRequest{
					ObjectType: to.StringPtr("<object-type>"),
					RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
						RestoreTargetInfoBase: armdataprotection.RestoreTargetInfoBase{
							ObjectType:      to.StringPtr("<object-type>"),
							RecoveryOption:  armdataprotection.RecoveryOptionFailIfExists.ToPtr(),
							RestoreLocation: to.StringPtr("<restore-location>"),
						},
						DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
							AuthCredentials: armdataprotection.AuthCredentials{
								ObjectType: to.StringPtr("<object-type>"),
							},
							SecretStoreResource: &armdataprotection.SecretStoreResource{
								SecretStoreType: armdataprotection.SecretStoreTypeAzureKeyVault.ToPtr(),
								URI:             to.StringPtr("<uri>"),
							},
						},
						DatasourceInfo: &armdataprotection.Datasource{
							DatasourceType:   to.StringPtr("<datasource-type>"),
							ObjectType:       to.StringPtr("<object-type>"),
							ResourceID:       to.StringPtr("<resource-id>"),
							ResourceLocation: to.StringPtr("<resource-location>"),
							ResourceName:     to.StringPtr("<resource-name>"),
							ResourceType:     to.StringPtr("<resource-type>"),
							ResourceURI:      to.StringPtr("<resource-uri>"),
						},
						DatasourceSetInfo: &armdataprotection.DatasourceSet{
							DatasourceType:   to.StringPtr("<datasource-type>"),
							ObjectType:       to.StringPtr("<object-type>"),
							ResourceID:       to.StringPtr("<resource-id>"),
							ResourceLocation: to.StringPtr("<resource-location>"),
							ResourceName:     to.StringPtr("<resource-name>"),
							ResourceType:     to.StringPtr("<resource-type>"),
							ResourceURI:      to.StringPtr("<resource-uri>"),
						},
					},
					SourceDataStoreType: armdataprotection.SourceDataStoreTypeVaultStore.ToPtr(),
				},
				RecoveryPointID: to.StringPtr("<recovery-point-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
