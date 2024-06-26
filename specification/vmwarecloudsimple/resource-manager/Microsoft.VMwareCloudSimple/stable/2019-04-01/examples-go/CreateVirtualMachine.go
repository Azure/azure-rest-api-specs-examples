package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateVirtualMachine.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "https://management.azure.com/", "myVirtualMachine", armvmwarecloudsimple.VirtualMachine{
		Location: to.Ptr("westus2"),
		Properties: &armvmwarecloudsimple.VirtualMachineProperties{
			AmountOfRAM: to.Ptr[int32](4096),
			Disks: []*armvmwarecloudsimple.VirtualDisk{
				{
					ControllerID:     to.Ptr("1000"),
					IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
					TotalSize:        to.Ptr[int32](10485760),
					VirtualDiskID:    to.Ptr("2000"),
				}},
			Nics: []*armvmwarecloudsimple.VirtualNic{
				{
					Network: &armvmwarecloudsimple.VirtualNetwork{
						ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
					},
					NicType:      to.Ptr(armvmwarecloudsimple.NICTypeE1000),
					PowerOnBoot:  to.Ptr(true),
					VirtualNicID: to.Ptr("4000"),
				}},
			NumberOfCores:  to.Ptr[int32](2),
			PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
			ResourcePool: &armvmwarecloudsimple.ResourcePool{
				ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26"),
			},
			TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
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
	// res.VirtualMachine = armvmwarecloudsimple.VirtualMachine{
	// 	Name: to.Ptr("myVirtualMachine"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/virtualMachines/myVirtualMachine"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.VirtualMachineProperties{
	// 		AmountOfRAM: to.Ptr[int32](4096),
	// 		Controllers: []*armvmwarecloudsimple.VirtualDiskController{
	// 		},
	// 		Disks: []*armvmwarecloudsimple.VirtualDisk{
	// 			{
	// 				ControllerID: to.Ptr("1000"),
	// 				IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
	// 				TotalSize: to.Ptr[int32](10485760),
	// 				VirtualDiskID: to.Ptr("2000"),
	// 				VirtualDiskName: to.Ptr("Hard disk 1"),
	// 		}},
	// 		GuestOS: to.Ptr("Other (32-bit)"),
	// 		GuestOSType: to.Ptr(armvmwarecloudsimple.GuestOSTypeOther),
	// 		Nics: []*armvmwarecloudsimple.VirtualNic{
	// 			{
	// 				Network: &armvmwarecloudsimple.VirtualNetwork{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
	// 				},
	// 				NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
	// 				PowerOnBoot: to.Ptr(true),
	// 				VirtualNicID: to.Ptr("4000"),
	// 				VirtualNicName: to.Ptr("Network adapter 1"),
	// 		}},
	// 		NumberOfCores: to.Ptr[int32](2),
	// 		PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourcePool: &armvmwarecloudsimple.ResourcePool{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26"),
	// 		},
	// 		TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
	// 		Vmwaretools: to.Ptr("0"),
	// 	},
	// }
}
