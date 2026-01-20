package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_Create_WithDiskControllerType.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVmWithDiskControllerType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineProperties{
			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
				BootDiagnostics: &armcompute.BootDiagnostics{
					Enabled:    to.Ptr(true),
					StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
				},
			},
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD4V3),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					}},
			},
			OSProfile: &armcompute.OSProfile{
				AdminPassword: to.Ptr("{your-password}"),
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
			},
			ScheduledEventsPolicy: &armcompute.ScheduledEventsPolicy{
				AllInstancesDown: &armcompute.AllInstancesDown{
					AutomaticallyApprove: to.Ptr(true),
				},
				ScheduledEventsAdditionalPublishingTargets: &armcompute.ScheduledEventsAdditionalPublishingTargets{
					EventGridAndResourceGraph: &armcompute.EventGridAndResourceGraph{
						Enable:                    to.Ptr(true),
						ScheduledEventsAPIVersion: to.Ptr("2020-07-01"),
					},
				},
				UserInitiatedReboot: &armcompute.UserInitiatedReboot{
					AutomaticallyApprove: to.Ptr(true),
				},
				UserInitiatedRedeploy: &armcompute.UserInitiatedRedeploy{
					AutomaticallyApprove: to.Ptr(true),
				},
			},
			StorageProfile: &armcompute.StorageProfile{
				DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
				ImageReference: &armcompute.ImageReference{
					Offer:     to.Ptr("WindowsServer"),
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					SKU:       to.Ptr("2016-Datacenter"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("myVMosdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
				},
			},
			UserData: to.Ptr("U29tZSBDdXN0b20gRGF0YQ=="),
		},
	}, &armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.VirtualMachine = armcompute.VirtualMachine{
	// 	Name: to.Ptr("myVM"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.VirtualMachineProperties{
	// 		DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 			BootDiagnostics: &armcompute.BootDiagnostics{
	// 				Enabled: to.Ptr(true),
	// 				StorageURI: to.Ptr("http://nsgdiagnostic.blob.core.windows.net"),
	// 			},
	// 		},
	// 		HardwareProfile: &armcompute.HardwareProfile{
	// 			VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD4V3),
	// 		},
	// 		NetworkProfile: &armcompute.NetworkProfile{
	// 			NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 					Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 						Primary: to.Ptr(true),
	// 					},
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.OSProfile{
	// 			AdminUsername: to.Ptr("{your-username}"),
	// 			ComputerName: to.Ptr("myVM"),
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 			WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 				EnableAutomaticUpdates: to.Ptr(true),
	// 				ProvisionVMAgent: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ScheduledEventsPolicy: &armcompute.ScheduledEventsPolicy{
	// 			AllInstancesDown: &armcompute.AllInstancesDown{
	// 				AutomaticallyApprove: to.Ptr(true),
	// 			},
	// 			ScheduledEventsAdditionalPublishingTargets: &armcompute.ScheduledEventsAdditionalPublishingTargets{
	// 				EventGridAndResourceGraph: &armcompute.EventGridAndResourceGraph{
	// 					Enable: to.Ptr(true),
	// 					ScheduledEventsAPIVersion: to.Ptr("2020-07-01"),
	// 				},
	// 			},
	// 			UserInitiatedReboot: &armcompute.UserInitiatedReboot{
	// 				AutomaticallyApprove: to.Ptr(true),
	// 			},
	// 			UserInitiatedRedeploy: &armcompute.UserInitiatedRedeploy{
	// 				AutomaticallyApprove: to.Ptr(true),
	// 			},
	// 		},
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 			},
	// 			DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 			ImageReference: &armcompute.ImageReference{
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 				SKU: to.Ptr("2016-Datacenter"),
	// 				Version: to.Ptr("latest"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("myVMosdisk"),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 		},
	// 		VMID: to.Ptr("676420ba-7a24-4bfe-80bd-9c841ee184fa"),
	// 	},
	// }
}
