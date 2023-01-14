package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/TriggerRestoreAsFiles.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreAsFiles() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx, "000pikumar", "PrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
		ObjectType: to.Ptr("AzureBackupRecoveryPointBasedRestoreRequest"),
		RestoreTargetInfo: &armdataprotection.RestoreFilesTargetInfo{
			ObjectType:      to.Ptr("RestoreFilesTargetInfo"),
			RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
			RestoreLocation: to.Ptr("southeastasia"),
			TargetDetails: &armdataprotection.TargetDetails{
				FilePrefix:                to.Ptr("restoredblob"),
				RestoreTargetLocationType: to.Ptr(armdataprotection.RestoreTargetLocationTypeAzureBlobs),
				URL:                       to.Ptr("https://teststorage.blob.core.windows.net/restoretest"),
			},
		},
		SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
		SourceResourceID:    to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		RecoveryPointID:     to.Ptr("hardcodedRP"),
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
