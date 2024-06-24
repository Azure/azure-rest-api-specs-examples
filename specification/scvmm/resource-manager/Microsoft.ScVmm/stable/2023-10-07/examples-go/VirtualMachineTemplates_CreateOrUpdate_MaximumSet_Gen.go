package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginCreateOrUpdate_virtualMachineTemplatesCreateOrUpdateMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginCreateOrUpdate(ctx, "rgscvmm", "6", armscvmm.VirtualMachineTemplate{
		Location: to.Ptr("ayxsyduviotylbojh"),
		Tags: map[string]*string{
			"key9494": to.Ptr("kkbmfpwhmvlobm"),
		},
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
			Type: to.Ptr("customLocation"),
		},
		Properties: &armscvmm.VirtualMachineTemplateProperties{
			DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
			InventoryItemID:      to.Ptr("qjrykoogccwlgkd"),
			IsCustomizable:       to.Ptr(armscvmm.IsCustomizableTrue),
			IsHighlyAvailable:    to.Ptr(armscvmm.IsHighlyAvailableTrue),
			LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
			OSType:               to.Ptr(armscvmm.OsTypeWindows),
			UUID:                 to.Ptr("12345678-1234-1234-1234-12345678abcd"),
			VmmServerID:          to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
		},
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
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
	// 	Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	Tags: map[string]*string{
	// 		"key9494": to.Ptr("kkbmfpwhmvlobm"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineTemplateProperties{
	// 		ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
	// 		CPUCount: to.Ptr[int32](23),
	// 		Disks: []*armscvmm.VirtualDisk{
	// 			{
	// 				Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
	// 				Bus: to.Ptr[int32](8),
	// 				BusType: to.Ptr("zu"),
	// 				CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
	// 				DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				DisplayName: to.Ptr("fgladknawlgjodo"),
	// 				Lun: to.Ptr[int32](10),
	// 				MaxDiskSizeGB: to.Ptr[int32](18),
	// 				StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
	// 					Name: to.Ptr("ceiyfrflu"),
	// 					ID: to.Ptr("o"),
	// 				},
	// 				TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
	// 				VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
	// 				VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
	// 				VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
	// 		}},
	// 		DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
	// 		DynamicMemoryMaxMB: to.Ptr[int32](21),
	// 		DynamicMemoryMinMB: to.Ptr[int32](21),
	// 		Generation: to.Ptr[int32](16),
	// 		InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
	// 		IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
	// 		IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
	// 		LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
	// 		MemoryMB: to.Ptr[int32](24),
	// 		NetworkInterfaces: []*armscvmm.NetworkInterface{
	// 			{
	// 				Name: to.Ptr("kvofzqulbjlbtt"),
	// 				DisplayName: to.Ptr("yoayfd"),
	// 				IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 				IPv4Addresses: []*string{
	// 					to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
	// 					IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 					IPv6Addresses: []*string{
	// 						to.Ptr("pk")},
	// 						MacAddress: to.Ptr("oaeqqegt"),
	// 						MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 						NetworkName: to.Ptr("lqbm"),
	// 						NicID: to.Ptr("roxpsvlo"),
	// 						VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
	// 				}},
	// 				OSName: to.Ptr("qcbolnbisklo"),
	// 				OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 				UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
	// 				VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
	// 			},
	// 		}
}
