package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRestoreWithRehydration.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreWithRehydration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRestoreWithRehydrationRequest{
		ObjectType: to.Ptr("AzureBackupRestoreWithRehydrationRequest"),
		RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
			ObjectType:      to.Ptr("RestoreTargetInfo"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			RestoreLocation: to.Ptr("southeastasia"),
			DatasourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("testdb"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
				ResourceURI:      to.Ptr(""),
			},
			DatasourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("OssDB"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
				ResourceLocation: to.Ptr(""),
				ResourceName:     to.Ptr("viveksipgtest"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/servers"),
				ResourceURI:      to.Ptr(""),
			},
		},
		SourceDataStoreType:          to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
		SourceResourceID:             to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		RecoveryPointID:              to.Ptr("hardcodedRP"),
		RehydrationPriority:          to.Ptr(armdataprotection.RehydrationPriorityHigh),
		RehydrationRetentionDuration: to.Ptr("7D"),
	}, nil)
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
