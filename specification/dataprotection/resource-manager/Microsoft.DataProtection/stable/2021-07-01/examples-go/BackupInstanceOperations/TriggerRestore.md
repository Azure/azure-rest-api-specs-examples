```go
package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/BackupInstanceOperations/TriggerRestore.json
func ExampleBackupInstancesClient_BeginTriggerRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginTriggerRestore(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<backup-instance-name>",
		&armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
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
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataprotection%2Farmdataprotection%2Fv0.1.0/sdk/resourcemanager/dataprotection/armdataprotection/README.md) on how to add the SDK to your project and authenticate.
