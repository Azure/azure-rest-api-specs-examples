package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_Get_virtualMachineImagesEdgeZoneGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().Get(ctx, "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
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
