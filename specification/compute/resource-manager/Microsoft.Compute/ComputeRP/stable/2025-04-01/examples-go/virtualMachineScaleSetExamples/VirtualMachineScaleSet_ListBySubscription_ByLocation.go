package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListBySubscription_ByLocation.json
func ExampleVirtualMachineScaleSetsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetsClient().NewListByLocationPager("eastus", nil)
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
		// page.VirtualMachineScaleSetListResult = armcompute.VirtualMachineScaleSetListResult{
		// 	Value: []*armcompute.VirtualMachineScaleSet{
		// 		{
		// 			Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 				DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
		// 				Overprovision: to.Ptr(false),
		// 				PlatformFaultDomainCount: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SinglePlacementGroup: to.Ptr(false),
		// 				UpgradePolicy: &armcompute.UpgradePolicy{
		// 					AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 						EnableAutomaticOSUpgrade: to.Ptr(false),
		// 					},
		// 					Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 				},
		// 				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 							{
		// 								Name: to.Ptr("myNic"),
		// 								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 										{
		// 											Name: to.Ptr("myIPConfig"),
		// 											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 												Primary: to.Ptr(true),
		// 												Subnet: &armcompute.APIEntityReference{
		// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 												},
		// 											},
		// 									}},
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 									},
		// 									Primary: to.Ptr(true),
		// 								},
		// 						}},
		// 					},
		// 					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 						AdminUsername: to.Ptr("admin"),
		// 						ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(false),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 						DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 						},
		// 						ImageReference: &armcompute.ImageReference{
		// 							Offer: to.Ptr("databricks"),
		// 							Publisher: to.Ptr("azuredatabricks"),
		// 							SKU: to.Ptr("databricksworker"),
		// 							Version: to.Ptr("3.15.2"),
		// 						},
		// 						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 							Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							DiskSizeGB: to.Ptr[int32](30),
		// 							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_D2s_v3"),
		// 				Capacity: to.Ptr[int64](4),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"myTag1": to.Ptr("tagValue2"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 				DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
		// 				Overprovision: to.Ptr(false),
		// 				PlatformFaultDomainCount: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SinglePlacementGroup: to.Ptr(false),
		// 				UpgradePolicy: &armcompute.UpgradePolicy{
		// 					AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 						EnableAutomaticOSUpgrade: to.Ptr(false),
		// 					},
		// 					Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 				},
		// 				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 							{
		// 								Name: to.Ptr("myNic1"),
		// 								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 										{
		// 											Name: to.Ptr("myIPConfig"),
		// 											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 												Primary: to.Ptr(true),
		// 												Subnet: &armcompute.APIEntityReference{
		// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 												},
		// 											},
		// 									}},
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 									},
		// 									Primary: to.Ptr(true),
		// 								},
		// 						}},
		// 					},
		// 					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 						AdminUsername: to.Ptr("admin"),
		// 						ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(false),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 						DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 						},
		// 						ImageReference: &armcompute.ImageReference{
		// 							Offer: to.Ptr("databricks"),
		// 							Publisher: to.Ptr("azuredatabricks"),
		// 							SKU: to.Ptr("databricksworker"),
		// 							Version: to.Ptr("3.15.2"),
		// 						},
		// 						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 							Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							DiskSizeGB: to.Ptr[int32](30),
		// 							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_D2s_v3"),
		// 				Capacity: to.Ptr[int64](4),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
