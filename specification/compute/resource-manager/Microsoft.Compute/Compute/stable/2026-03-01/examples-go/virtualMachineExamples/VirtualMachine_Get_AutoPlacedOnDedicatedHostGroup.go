package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_Get_AutoPlacedOnDedicatedHostGroup.json
func ExampleVirtualMachinesClient_Get_getAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "myResourceGroup", "myVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientGetResponse{
	// 	VirtualMachine: armcompute.VirtualMachine{
	// 		Name: to.Ptr("myVM"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			VMID: to.Ptr("0f47b100-583c-48e3-a4c0-aefc2c9bbcc1"),
	// 			HostGroup: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/hostGroups/myHostGroup"),
	// 			},
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2016-Datacenter"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Name: to.Ptr("myOsDisk"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myOsDisk"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("myVM"),
	// 				AdminUsername: to.Ptr("admin"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(false),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{myNIC}"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	},
	// }
}
