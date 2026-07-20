package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskExamples/Disk_Update_AddSupportsHibernation.json
func ExampleDisksClient_BeginUpdate_updateAManagedDiskToAddSupportsHibernation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginUpdate(ctx, "myResourceGroup", "myDisk", armcompute.DiskUpdate{
		Properties: &armcompute.DiskUpdateProperties{
			SupportsHibernation: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DisksClientUpdateResponse{
	// 	Disk: armcompute.Disk{
	// 		Name: to.Ptr("myDisk"),
	// 		Location: to.Ptr("westus"),
	// 		SKU: &armcompute.DiskSKU{
	// 			Name: to.Ptr(armcompute.DiskStorageAccountTypesStandardLRS),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		Properties: &armcompute.DiskProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			SupportsHibernation: to.Ptr(true),
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionImport),
	// 				SourceURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
	// 			},
	// 			DiskSizeGB: to.Ptr[int32](127),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
