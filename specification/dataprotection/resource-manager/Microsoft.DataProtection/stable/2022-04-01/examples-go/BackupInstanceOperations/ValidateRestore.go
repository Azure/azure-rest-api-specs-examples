package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2022-04-01/examples/BackupInstanceOperations/ValidateRestore.json
func ExampleBackupInstancesClient_BeginValidateForRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginValidateForRestore(ctx,
		"000pikumar",
		"PratikPrivatePreviewVault1",
		"testInstance1",
		armdataprotection.ValidateRestoreRequestObject{
			RestoreRequestObject: &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
				ObjectType: to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
				RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
					ObjectType:      to.Ptr("RestoreTargetInfo"),
					RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
					RestoreLocation: to.Ptr("southeastasia"),
					DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
						ObjectType: to.Ptr("SecretStoreBasedAuthCredentials"),
						SecretStoreResource: &armdataprotection.SecretStoreResource{
							SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
							URI:             to.Ptr("https://samplevault.vault.azure.net/secrets/credentials"),
						},
					},
					DatasourceInfo: &armdataprotection.Datasource{
						DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
						ObjectType:       to.Ptr("Datasource"),
						ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb"),
						ResourceLocation: to.Ptr(""),
						ResourceName:     to.Ptr("targetdb"),
						ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
						ResourceURI:      to.Ptr(""),
					},
					DatasourceSetInfo: &armdataprotection.DatasourceSet{
						DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
						ObjectType:       to.Ptr("DatasourceSet"),
						ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
						ResourceLocation: to.Ptr(""),
						ResourceName:     to.Ptr("viveksipgtest"),
						ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
						ResourceURI:      to.Ptr(""),
					},
				},
				SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
				SourceResourceID:    to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				RecoveryPointID:     to.Ptr("hardcodedRP"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
