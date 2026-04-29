package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithEncryptionIdentity.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVMWithEncryptionIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Identity: &armcompute.VirtualMachineIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": {},
			},
		},
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
			},
			SecurityProfile: &armcompute.SecurityProfile{
				EncryptionIdentity: &armcompute.EncryptionIdentity{
					UserAssignedIdentityResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
				},
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					Offer:     to.Ptr("WindowsServer"),
					SKU:       to.Ptr("2019-Datacenter"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadOnly),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardSSDLRS),
					},
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Name:         to.Ptr("myVMosdisk"),
				},
			},
			OSProfile: &armcompute.OSProfile{
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
				AdminPassword: to.Ptr("{your-password}"),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					},
				},
			},
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
	// res = armcompute.VirtualMachinesClientCreateOrUpdateResponse{
	// 	VirtualMachine: &armcompute.VirtualMachine{
	// 		Name: to.Ptr("myVM"),
	// 		Identity: &armcompute.VirtualMachineIdentity{
	// 			Type: to.Ptr(armcompute.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 				"/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": &armcompute.UserAssignedIdentitiesValue{
	// 				},
	// 			},
	// 		},
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				ComputerName: to.Ptr("myVM"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 				},
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 						Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 							Primary: to.Ptr(true),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2019-Datacenter"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadOnly),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardSSDLRS),
	// 					},
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Name: to.Ptr("myVMosdisk"),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				EncryptionIdentity: &armcompute.EncryptionIdentity{
	// 					UserAssignedIdentityResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),
	// 				},
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			VMID: to.Ptr("5c0d55a7-c407-4ed6-bf7d-ddb810267c85"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2SV3),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Location: to.Ptr("westus"),
	// 	},
	// }
}
