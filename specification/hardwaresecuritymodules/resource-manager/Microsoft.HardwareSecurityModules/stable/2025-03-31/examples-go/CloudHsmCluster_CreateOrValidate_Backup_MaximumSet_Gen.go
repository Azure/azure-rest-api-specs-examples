package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: 2025-03-31/CloudHsmCluster_CreateOrValidate_Backup_MaximumSet_Gen.json
func ExampleCloudHsmClustersClient_BeginBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudHsmClustersClient().BeginBackup(ctx, "rgcloudhsm", "chsm1", &armhardwaresecuritymodules.CloudHsmClustersClientBeginBackupOptions{
		BackupRequestProperties: &armhardwaresecuritymodules.BackupRequestProperties{
			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
			Token:                        to.Ptr("se=2018-02-01T00%3A00Z&spr=https&sv=2017-04-17&sr=b&sig=REDACTED"),
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
	// res = armhardwaresecuritymodules.CloudHsmClustersClientBackupResponse{
	// 	BackupResult: &armhardwaresecuritymodules.BackupResult{
	// 		Properties: &armhardwaresecuritymodules.BackupResultProperties{
	// 			AzureStorageBlobContainerURI: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/sasContainer"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			JobID: to.Ptr("572a45927fc240e1ac075de27371680b"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-12T12:00:00.0000000Z"); return t}()),
	// 			Status: to.Ptr(armhardwaresecuritymodules.BackupRestoreOperationStatusInProgress),
	// 			StatusDetails: to.Ptr("Backup operation is in progress"),
	// 		},
	// 	},
	// }
}
