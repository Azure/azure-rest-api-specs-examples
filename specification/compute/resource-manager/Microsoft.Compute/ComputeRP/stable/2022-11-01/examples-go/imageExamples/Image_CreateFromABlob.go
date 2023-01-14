package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/imageExamples/Image_CreateFromABlob.json
func ExampleImagesClient_BeginCreateOrUpdate_createAVirtualMachineImageFromABlob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewImagesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myImage", armcompute.Image{
		Location: to.Ptr("West US"),
		Properties: &armcompute.ImageProperties{
			StorageProfile: &armcompute.ImageStorageProfile{
				OSDisk: &armcompute.ImageOSDisk{
					BlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
					OSType:  to.Ptr(armcompute.OperatingSystemTypesLinux),
				},
				ZoneResilient: to.Ptr(true),
			},
		},
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
