package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: 2025-03-31/CloudHsmCluster_RequestOrValidate_Restore_MaximumSet_Gen_ValidateRestoreProperties.json
func ExampleCloudHsmClustersClient_BeginValidateRestoreProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginValidateRestoreProperties(ctx, "rgcloudhsm", "chsm1", &armhardwaresecuritymodules.CloudHsmClustersClientBeginValidateRestorePropertiesOptions{
		RestoreRequestProperties: &armhardwaresecuritymodules.RestoreRequestProperties{
			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
			BackupID:                     to.Ptr("backupId"),
		}})
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientValidateRestorePropertiesResponse{
	// 	RestoreResult: &armhardwaresecuritymodules.RestoreResult{
	// 		Properties: &armhardwaresecuritymodules.BackupRestoreBaseResultProperties{
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Restore operation is in progress"),
	// 		},
	// 	},
	// }
}
