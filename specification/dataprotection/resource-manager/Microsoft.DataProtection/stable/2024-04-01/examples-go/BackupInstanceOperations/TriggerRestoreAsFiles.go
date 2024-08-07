package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/BackupInstanceOperations/TriggerRestoreAsFiles.json
func ExampleBackupInstancesClient_BeginTriggerRestore_triggerRestoreAsFiles() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginTriggerRestore(ctx, "000pikumar", "PrivatePreviewVault1", "testInstance1", &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
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
	}, &armdataprotection.BackupInstancesClientBeginTriggerRestoreOptions{XMSAuthorizationAuxiliary: nil})
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
