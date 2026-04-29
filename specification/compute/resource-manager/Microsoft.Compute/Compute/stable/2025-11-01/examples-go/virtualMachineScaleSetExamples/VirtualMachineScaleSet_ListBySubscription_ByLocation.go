package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListBySubscription_ByLocation.json
func ExampleVirtualMachineScaleSetsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
		// page = armcompute.VirtualMachineScaleSetsClientListByLocationResponse{
		// 	VirtualMachineScaleSetListResult: armcompute.VirtualMachineScaleSetListResult{
		// 		Value: []*armcompute.VirtualMachineScaleSet{
		// 			{
		// 				Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"myTag1": to.Ptr("tagValue1"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_D2s_v3"),
		// 					Tier: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 					SinglePlacementGroup: to.Ptr(false),
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							EnableAutomaticOSUpgrade: to.Ptr(false),
		// 						},
		// 					},
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](30),
		// 							},
		// 							ImageReference: &armcompute.ImageReference{
		// 								Publisher: to.Ptr("azuredatabricks"),
		// 								Offer: to.Ptr("databricks"),
		// 								SKU: to.Ptr("databricksworker"),
		// 								Version: to.Ptr("3.15.2"),
		// 							},
		// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 							},
		// 						},
		// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 							ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 							AdminUsername: to.Ptr("admin"),
		// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 								DisablePasswordAuthentication: to.Ptr(false),
		// 							},
		// 						},
		// 						NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 							NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 								{
		// 									Name: to.Ptr("myNic"),
		// 									Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 										Primary: to.Ptr(true),
		// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 											{
		// 												Name: to.Ptr("myIPConfig"),
		// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 													Subnet: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 										NetworkSecurityGroup: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						SecurityProfile: &armcompute.SecurityProfile{
		// 							SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("succeeded"),
		// 					Overprovision: to.Ptr(false),
		// 					DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
		// 					PlatformFaultDomainCount: to.Ptr[int32](1),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}1"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"myTag1": to.Ptr("tagValue2"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_D2s_v3"),
		// 					Tier: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 					SinglePlacementGroup: to.Ptr(false),
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							EnableAutomaticOSUpgrade: to.Ptr(false),
		// 						},
		// 					},
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](30),
		// 							},
		// 							ImageReference: &armcompute.ImageReference{
		// 								Publisher: to.Ptr("azuredatabricks"),
		// 								Offer: to.Ptr("databricks"),
		// 								SKU: to.Ptr("databricksworker"),
		// 								Version: to.Ptr("3.15.2"),
		// 							},
		// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 							},
		// 						},
		// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 							ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 							AdminUsername: to.Ptr("admin"),
		// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 								DisablePasswordAuthentication: to.Ptr(false),
		// 							},
		// 						},
		// 						NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 							NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 								{
		// 									Name: to.Ptr("myNic1"),
		// 									Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 										Primary: to.Ptr(true),
		// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 											{
		// 												Name: to.Ptr("myIPConfig"),
		// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 													Subnet: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 										NetworkSecurityGroup: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						SecurityProfile: &armcompute.SecurityProfile{
		// 							SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("succeeded"),
		// 					Overprovision: to.Ptr(false),
		// 					DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(false),
		// 					PlatformFaultDomainCount: to.Ptr[int32](1),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
