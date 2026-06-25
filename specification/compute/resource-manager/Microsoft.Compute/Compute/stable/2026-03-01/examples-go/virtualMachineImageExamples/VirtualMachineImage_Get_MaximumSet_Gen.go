package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_Get_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_Get_virtualMachineImageGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
	// res = armcompute.VirtualMachineImagesClientGetResponse{
	// 	VirtualMachineImage: armcompute.VirtualMachineImage{
	// 		Properties: &armcompute.VirtualMachineImageProperties{
	// 			Plan: &armcompute.PurchasePlan{
	// 				Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 				Name: to.Ptr("aaaaaaaaa"),
	// 				Product: to.Ptr("aaaaaaaaaaaaaa"),
	// 			},
	// 			OSDiskImage: &armcompute.OSDiskImage{
	// 				OperatingSystem: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 			DataDiskImages: []*armcompute.DataDiskImage{
	// 				{
	// 					Lun: to.Ptr[int32](17),
	// 				},
	// 			},
	// 			AutomaticOSUpgradeProperties: &armcompute.AutomaticOSUpgradeProperties{
	// 				AutomaticOSUpgradeSupported: to.Ptr(true),
	// 			},
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
	// 			Disallowed: &armcompute.DisallowedConfiguration{
	// 				VMDiskType: to.Ptr(armcompute.VMDiskTypesNone),
	// 			},
	// 			Features: []*armcompute.VirtualMachineImageFeature{
	// 				{
	// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 					Value: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 			ImageDeprecationStatus: &armcompute.ImageDeprecationStatus{
	// 				ImageState: to.Ptr(armcompute.ImageStateScheduledForDeprecation),
	// 				ScheduledDeprecationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T00:00:00+00:00"); return t}()),
	// 				AlternativeOption: &armcompute.AlternativeOption{
	// 					Type: to.Ptr(armcompute.AlternativeTypeOffer),
	// 					Value: to.Ptr("aaaaaaa"),
	// 				},
	// 			},
	// 		},
	// 		Name: to.Ptr("aaaaaaaaa"),
	// 		Location: to.Ptr("aaaaa"),
	// 		Tags: map[string]*string{
	// 			"key6817": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 	},
	// }
}
