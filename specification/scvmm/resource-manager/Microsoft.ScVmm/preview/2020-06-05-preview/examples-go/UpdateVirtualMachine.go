package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/UpdateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginUpdate(ctx, "testrg", "DemoVM", armscvmm.VirtualMachineUpdate{
		Properties: &armscvmm.VirtualMachineUpdateProperties{
			HardwareProfile: &armscvmm.HardwareProfileUpdate{
				CPUCount: to.Ptr[int32](4),
				MemoryMB: to.Ptr[int32](4096),
			},
			NetworkProfile: &armscvmm.NetworkProfileUpdate{
				NetworkInterfaces: []*armscvmm.NetworkInterfacesUpdate{
					{
						Name:            to.Ptr("test"),
						IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
						IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
						MacAddressType:  to.Ptr(armscvmm.AllocationMethodStatic),
					}},
			},
			StorageProfile: &armscvmm.StorageProfileUpdate{
				Disks: []*armscvmm.VirtualDiskUpdate{
					{
						Name:       to.Ptr("test"),
						DiskSizeGB: to.Ptr[int32](10),
					}},
			},
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.VirtualMachine = armscvmm.VirtualMachine{
	// 	Name: to.Ptr("DemoVM"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VirtualMachines"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachines/DemoVM"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armscvmm.VirtualMachineProperties{
	// 		CloudID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
	// 		Generation: to.Ptr[int32](0),
	// 		HardwareProfile: &armscvmm.HardwareProfile{
	// 			CPUCount: to.Ptr[int32](4),
	// 			DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledFalse),
	// 			DynamicMemoryMaxMB: to.Ptr[int32](0),
	// 			DynamicMemoryMinMB: to.Ptr[int32](0),
	// 			IsHighlyAvailable: to.Ptr("string"),
	// 			MemoryMB: to.Ptr[int32](4096),
	// 		},
	// 		NetworkProfile: &armscvmm.NetworkProfile{
	// 			NetworkInterfaces: []*armscvmm.NetworkInterfaces{
	// 				{
	// 					Name: to.Ptr("test"),
	// 					DisplayName: to.Ptr("string"),
	// 					IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 					IPv4Addresses: []*string{
	// 						to.Ptr("string")},
	// 						IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 						IPv6Addresses: []*string{
	// 							to.Ptr("string")},
	// 							MacAddress: to.Ptr("string"),
	// 							MacAddressType: to.Ptr(armscvmm.AllocationMethodStatic),
	// 							NetworkName: to.Ptr("string"),
	// 							NicID: to.Ptr("string"),
	// 							VirtualNetworkID: to.Ptr("string"),
	// 					}},
	// 				},
	// 				OSProfile: &armscvmm.OsProfile{
	// 					ComputerName: to.Ptr("DemoVM"),
	// 					OSName: to.Ptr("string"),
	// 					OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 				},
	// 				PowerState: to.Ptr("string"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				StorageProfile: &armscvmm.StorageProfile{
	// 					Disks: []*armscvmm.VirtualDisk{
	// 						{
	// 							Name: to.Ptr("test"),
	// 							Bus: to.Ptr[int32](0),
	// 							BusType: to.Ptr("string"),
	// 							CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskFalse),
	// 							DiskID: to.Ptr("string"),
	// 							DiskSizeGB: to.Ptr[int32](10),
	// 							DisplayName: to.Ptr("string"),
	// 							Lun: to.Ptr[int32](0),
	// 							MaxDiskSizeGB: to.Ptr[int32](10),
	// 							StorageQoSPolicy: &armscvmm.StorageQoSPolicyDetails{
	// 								Name: to.Ptr("string"),
	// 								ID: to.Ptr("string"),
	// 							},
	// 							TemplateDiskID: to.Ptr("string"),
	// 							VhdFormatType: to.Ptr("string"),
	// 							VhdType: to.Ptr("string"),
	// 							VolumeType: to.Ptr("string"),
	// 					}},
	// 				},
	// 				TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
	// 				UUID: to.Ptr("string"),
	// 				VMName: to.Ptr("string"),
	// 				VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 			},
	// 			Tags: map[string]*string{
	// 				"tag1": to.Ptr("value1"),
	// 				"tag2": to.Ptr("value2"),
	// 			},
	// 		}
}
