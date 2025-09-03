package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/VirtualHardDisks_Upload.json
func ExampleVirtualHardDisksClient_BeginUpload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHardDisksClient().BeginUpload(ctx, "test-rg", "test-vhd", armazurestackhcivm.VirtualHardDiskUploadRequest{
		AzureManagedDiskUploadURL: to.Ptr("https://YourStorageAccountName.blob.core.windows.net/YourContainerName/YourVHDBlobName.vhd?<sas-token>"),
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
	// res = armazurestackhcivm.VirtualHardDisksClientUploadResponse{
	// 	VirtualHardDiskUploadResponse: &armazurestackhcivm.VirtualHardDiskUploadResponse{
	// 		VirtualHardDiskID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
	// 		UploadStatus: &armazurestackhcivm.VirtualHardDiskUploadStatus{
	// 			UploadedSizeInMB: to.Ptr[int64](10240),
	// 			ProgressPercentage: to.Ptr[int64](55),
	// 			Status: to.Ptr(armazurestackhcivm.StatusInProgress),
	// 		},
	// 	},
	// }
}
