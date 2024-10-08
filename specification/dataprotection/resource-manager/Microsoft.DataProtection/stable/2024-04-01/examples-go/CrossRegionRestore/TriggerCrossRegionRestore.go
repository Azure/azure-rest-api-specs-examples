package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/CrossRegionRestore/TriggerCrossRegionRestore.json
func ExampleBackupInstancesClient_BeginTriggerCrossRegionRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginTriggerCrossRegionRestore(ctx, "000pikumar", "EastAsia", armdataprotection.CrossRegionRestoreRequestObject{
		CrossRegionRestoreDetails: &armdataprotection.CrossRegionRestoreDetails{
			SourceBackupInstanceID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/BackupInstances/harshitbi1"),
			SourceRegion:           to.Ptr("east asia"),
		},
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationJobExtendedInfo = armdataprotection.OperationJobExtendedInfo{
	// 	ObjectType: to.Ptr("OperationJobExtendedInfo"),
	// 	JobID: to.Ptr("c60cb49-63e8-4b21-b9bd-26277b3fdfae"),
	// }
}
