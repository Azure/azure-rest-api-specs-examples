package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_List_MaximumSet_Gen.json
func ExampleVirtualMachineInstancesClient_NewListPager_virtualMachineInstancesListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineInstancesClient().NewListPager("gtgclehcbsyave", nil)
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
		// page.VirtualMachineInstanceListResult = armscvmm.VirtualMachineInstanceListResult{
		// 	Value: []*armscvmm.VirtualMachineInstance{
		// 		{
		// 			Name: to.Ptr("uuqpsdoiyvedvqtrwop"),
		// 			Type: to.Ptr("zculorteltpvthtzgnpgdpoe"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
		// 			SystemData: &armscvmm.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
		// 				CreatedBy: to.Ptr("p"),
		// 				CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
		// 				LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 			},
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VirtualMachineInstanceProperties{
		// 				AvailabilitySets: []*armscvmm.AvailabilitySetListItem{
		// 					{
		// 						Name: to.Ptr("lwbhaseo"),
		// 						ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/availabilitySets/availabilitySetResourceName"),
		// 				}},
		// 				HardwareProfile: &armscvmm.HardwareProfile{
		// 					CPUCount: to.Ptr[int32](22),
		// 					DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
		// 					DynamicMemoryMaxMB: to.Ptr[int32](2),
		// 					DynamicMemoryMinMB: to.Ptr[int32](30),
		// 					IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
		// 					LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
		// 					MemoryMB: to.Ptr[int32](5),
		// 				},
		// 				InfrastructureProfile: &armscvmm.InfrastructureProfile{
		// 					BiosGUID: to.Ptr("xixivxifyql"),
		// 					CheckpointType: to.Ptr("jkbpzjxpeegackhsvikrnlnwqz"),
		// 					Checkpoints: []*armscvmm.Checkpoint{
		// 						{
		// 							Name: to.Ptr("keqn"),
		// 							Description: to.Ptr("qurzfrgyflrh"),
		// 							CheckpointID: to.Ptr("wsqmrje"),
		// 							ParentCheckpointID: to.Ptr("hqhhzikoxunuqguouw"),
		// 					}},
		// 					CloudID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/clouds/cloudResourceName"),
		// 					Generation: to.Ptr[int32](28),
		// 					InventoryItemID: to.Ptr("ihkkqmg"),
		// 					LastRestoredVMCheckpoint: &armscvmm.Checkpoint{
		// 						Name: to.Ptr("keqn"),
		// 						Description: to.Ptr("qurzfrgyflrh"),
		// 						CheckpointID: to.Ptr("wsqmrje"),
		// 						ParentCheckpointID: to.Ptr("hqhhzikoxunuqguouw"),
		// 					},
		// 					TemplateID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
		// 					UUID: to.Ptr("hrpw"),
		// 					VMName: to.Ptr("qovpayfydhcvfrhe"),
		// 					VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
		// 				},
		// 				NetworkProfile: &armscvmm.NetworkProfile{
		// 					NetworkInterfaces: []*armscvmm.NetworkInterface{
		// 						{
		// 							Name: to.Ptr("kvofzqulbjlbtt"),
		// 							DisplayName: to.Ptr("yoayfd"),
		// 							IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 							IPv4Addresses: []*string{
		// 								to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
		// 								IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 								IPv6Addresses: []*string{
		// 									to.Ptr("pk")},
		// 									MacAddress: to.Ptr("oaeqqegt"),
		// 									MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 									NetworkName: to.Ptr("lqbm"),
		// 									NicID: to.Ptr("roxpsvlo"),
		// 									VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
		// 							}},
		// 						},
		// 						OSProfile: &armscvmm.OsProfileForVMInstance{
		// 							ComputerName: to.Ptr("uuxpcxuxcufllc"),
		// 							OSSKU: to.Ptr("cxqnjxgkts"),
		// 							OSType: to.Ptr(armscvmm.OsTypeWindows),
		// 							OSVersion: to.Ptr("djt"),
		// 						},
		// 						PowerState: to.Ptr("dbqyxewvrbqcifpwfvxyllwyaffmvm"),
		// 						ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 						StorageProfile: &armscvmm.StorageProfile{
		// 							Disks: []*armscvmm.VirtualDisk{
		// 								{
		// 									Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
		// 									Bus: to.Ptr[int32](8),
		// 									BusType: to.Ptr("zu"),
		// 									CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
		// 									DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
		// 									DiskSizeGB: to.Ptr[int32](30),
		// 									DisplayName: to.Ptr("fgladknawlgjodo"),
		// 									Lun: to.Ptr[int32](10),
		// 									MaxDiskSizeGB: to.Ptr[int32](18),
		// 									StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
		// 										Name: to.Ptr("ceiyfrflu"),
		// 										ID: to.Ptr("o"),
		// 									},
		// 									TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
		// 									VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
		// 									VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
		// 									VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
		// 							}},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
