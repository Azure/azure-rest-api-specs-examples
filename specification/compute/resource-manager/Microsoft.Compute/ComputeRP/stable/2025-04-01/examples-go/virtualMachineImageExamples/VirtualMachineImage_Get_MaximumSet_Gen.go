package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineImageExamples/VirtualMachineImage_Get_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_Get_virtualMachineImageGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesClient().Get(ctx, "aaaaaa", "aaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImage = armcompute.VirtualMachineImage{
	// 	ID: to.Ptr("aaaaaaaaaaa"),
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	ExtendedLocation: &armcompute.ExtendedLocation{
	// 		Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Location: to.Ptr("aaaaa"),
	// 	Tags: map[string]*string{
	// 		"key6817": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineImageProperties{
	// 		AutomaticOSUpgradeProperties: &armcompute.AutomaticOSUpgradeProperties{
	// 			AutomaticOSUpgradeSupported: to.Ptr(true),
	// 		},
	// 		DataDiskImages: []*armcompute.DataDiskImage{
	// 			{
	// 				Lun: to.Ptr[int32](17),
	// 		}},
	// 		Disallowed: &armcompute.DisallowedConfiguration{
	// 			VMDiskType: to.Ptr(armcompute.VMDiskTypesNone),
	// 		},
	// 		Features: []*armcompute.VirtualMachineImageFeature{
	// 			{
	// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 				Value: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		}},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
	// 		ImageDeprecationStatus: &armcompute.ImageDeprecationStatus{
	// 			AlternativeOption: &armcompute.AlternativeOption{
	// 				Type: to.Ptr(armcompute.AlternativeTypeOffer),
	// 				Value: to.Ptr("aaaaaaa"),
	// 			},
	// 			ImageState: to.Ptr(armcompute.ImageStateScheduledForDeprecation),
	// 			ScheduledDeprecationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T00:00:00.000Z"); return t}()),
	// 		},
	// 		OSDiskImage: &armcompute.OSDiskImage{
	// 			OperatingSystem: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		},
	// 		Plan: &armcompute.PurchasePlan{
	// 			Name: to.Ptr("aaaaaaaaa"),
	// 			Product: to.Ptr("aaaaaaaaaaaaaa"),
	// 			Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 	},
	// }
}
