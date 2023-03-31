package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_PreRestore.json
func ExampleCloudEndpointsClient_BeginPreRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudEndpointsClient().BeginPreRestore(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleCloudEndpoint_1", armstoragesync.PreRestoreRequest{
		AzureFileShareURI: to.Ptr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
		RestoreFileSpec: []*armstoragesync.RestoreFileSpec{
			{
				Path:  to.Ptr("text1.txt"),
				Isdir: to.Ptr(false),
			},
			{
				Path:  to.Ptr("MyDir"),
				Isdir: to.Ptr(true),
			},
			{
				Path:  to.Ptr("MyDir/SubDir"),
				Isdir: to.Ptr(false),
			},
			{
				Path:  to.Ptr("MyDir/SubDir/File1.pdf"),
				Isdir: to.Ptr(false),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
