package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/BackupInstanceOperations/GetRecoveryPoint.json
func ExampleRecoveryPointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().Get(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", "7fb2cddd-c5b3-44f6-a0d9-db3c4f9d5f25", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureBackupRecoveryPointResource = armdataprotection.AzureBackupRecoveryPointResource{
	// 	Name: to.Ptr("7fb2cddd-c5b3-44f6-a0d9-db3c4f9d5f25"),
	// 	Type: to.Ptr("microsoft.dataprotection/backupvaults/backupInstances/recoveryPoints"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/HelloTest/providers/Microsoft.DataProtection/backupVaults/HelloTestVault/backupInstances/653213d-c5b3-44f6-a0d9-db3c4f9d8e34/recoveryPoints/7fb2cddd-c5b3-44f6-a0d9-db3c4f9d5f25"),
	// 	Properties: &armdataprotection.AzureBackupDiscreteRecoveryPoint{
	// 		ObjectType: to.Ptr("AzureBackupDiscreteRecoveryPoint"),
	// 		FriendlyName: to.Ptr("panbha4"),
	// 		RecoveryPointDataStoresDetails: []*armdataprotection.RecoveryPointDataStoreDetails{
	// 			{
	// 				Type: to.Ptr("Snapshot"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00.000Z"); return t}()),
	// 				ID: to.Ptr("0ff03512-b333-4509-a6c7-12164c8b1dce"),
	// 				MetaData: to.Ptr("123456"),
	// 			},
	// 			{
	// 				Type: to.Ptr("BackupStorage"),
	// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00.000Z"); return t}()),
	// 				ID: to.Ptr("5d8cfd30-722e-4bab-85f6-4a9d01ffc6f1"),
	// 				MetaData: to.Ptr("123456"),
	// 		}},
	// 		RecoveryPointState: to.Ptr(armdataprotection.RecoveryPointCompletionStateCompleted),
	// 		RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T13:00:00.000Z"); return t}()),
	// 		RecoveryPointType: to.Ptr("Full"),
	// 	},
	// }
}
