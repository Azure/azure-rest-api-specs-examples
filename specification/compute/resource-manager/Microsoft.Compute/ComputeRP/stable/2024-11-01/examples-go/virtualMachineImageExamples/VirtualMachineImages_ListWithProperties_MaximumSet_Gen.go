package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/250861bb6a886b75255edfa0aa5ee2dd0d6e7a11/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineImageExamples/VirtualMachineImages_ListWithProperties_MaximumSet_Gen.json
func ExampleVirtualMachineImagesClient_NewListWithPropertiesPager_virtualMachineImagesListWithPropertiesMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineImagesClient().NewListWithPropertiesPager("eastus", "MicrosoftWindowsServer", "WindowsServer", "2022-datacenter-azure-edition", armcompute.ExpandProperties, &armcompute.VirtualMachineImagesClientListWithPropertiesOptions{Top: to.Ptr[int32](4),
		Orderby: to.Ptr("aa"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineImagesWithPropertiesListResult = armcompute.VirtualMachineImagesWithPropertiesListResult{
		// 	Value: []*armcompute.VirtualMachineImage{
		// 		{
		// 			ID: to.Ptr("aaaaaaaaaaa"),
		// 			Name: to.Ptr("aaaaaaaaa"),
		// 			ExtendedLocation: &armcompute.ExtendedLocation{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 				Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 			},
		// 			Location: to.Ptr("aaaaa"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.VirtualMachineImageProperties{
		// 				Architecture: to.Ptr(armcompute.ArchitectureTypesX64),
		// 				AutomaticOSUpgradeProperties: &armcompute.AutomaticOSUpgradeProperties{
		// 					AutomaticOSUpgradeSupported: to.Ptr(true),
		// 				},
		// 				DataDiskImages: []*armcompute.DataDiskImage{
		// 					{
		// 						Lun: to.Ptr[int32](17),
		// 				}},
		// 				Disallowed: &armcompute.DisallowedConfiguration{
		// 					VMDiskType: to.Ptr(armcompute.VMDiskTypesNone),
		// 				},
		// 				Features: []*armcompute.VirtualMachineImageFeature{
		// 					{
		// 						Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 						Value: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				}},
		// 				HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
		// 				ImageDeprecationStatus: &armcompute.ImageDeprecationStatus{
		// 					AlternativeOption: &armcompute.AlternativeOption{
		// 						Type: to.Ptr(armcompute.AlternativeTypeOffer),
		// 						Value: to.Ptr("aaaaaaa"),
		// 					},
		// 					ImageState: to.Ptr(armcompute.ImageStateScheduledForDeprecation),
		// 					ScheduledDeprecationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-13T00:00:00.000Z"); return t}()),
		// 				},
		// 				OSDiskImage: &armcompute.OSDiskImage{
		// 					OperatingSystem: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				},
		// 				Plan: &armcompute.PurchasePlan{
		// 					Name: to.Ptr("aaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaa"),
		// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
