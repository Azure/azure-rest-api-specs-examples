package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionImagesClient_Get_virtualMachineExtensionImageGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineExtensionImagesClient().Get(ctx, "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineExtensionImagesClientGetResponse{
	// 	VirtualMachineExtensionImage: armcompute.VirtualMachineExtensionImage{
	// 		Properties: &armcompute.VirtualMachineExtensionImageProperties{
	// 			OperatingSystem: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 			ComputeRole: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 			HandlerSchema: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			VMScaleSetEnabled: to.Ptr(true),
	// 			SupportsMultipleExtensions: to.Ptr(true),
	// 		},
	// 		ID: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key9885": to.Ptr("aaaaaaaaa"),
	// 		},
	// 	},
	// }
}
