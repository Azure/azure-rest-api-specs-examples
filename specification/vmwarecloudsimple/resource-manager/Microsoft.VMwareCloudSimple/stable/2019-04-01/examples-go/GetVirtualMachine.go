package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachine.json
func ExampleVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "myResourceGroup", "myVirtualMachine", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachine = armvmwarecloudsimple.VirtualMachine{
	// 	Name: to.Ptr("myVirtualMachine"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/virtualMachines/myVirtualMachine"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.VirtualMachineProperties{
	// 		AmountOfRAM: to.Ptr[int32](4096),
	// 		Controllers: []*armvmwarecloudsimple.VirtualDiskController{
	// 			{
	// 				Name: to.Ptr("SCSI controller 0"),
	// 				Type: to.Ptr("SCSI"),
	// 				ID: to.Ptr("1000"),
	// 				SubType: to.Ptr("LSI_PARALEL"),
	// 		}},
	// 		Disks: []*armvmwarecloudsimple.VirtualDisk{
	// 			{
	// 				ControllerID: to.Ptr("1000"),
	// 				IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
	// 				TotalSize: to.Ptr[int32](10485760),
	// 				VirtualDiskID: to.Ptr("2000"),
	// 				VirtualDiskName: to.Ptr("Hard disk 1"),
	// 		}},
	// 		Folder: to.Ptr("Datacenter/Workload VMs"),
	// 		GuestOS: to.Ptr("Other (32-bit)"),
	// 		GuestOSType: to.Ptr(armvmwarecloudsimple.GuestOSTypeOther),
	// 		Nics: []*armvmwarecloudsimple.VirtualNic{
	// 			{
	// 				MacAddress: to.Ptr("00:50:56:a6:d0:e1"),
	// 				Network: &armvmwarecloudsimple.VirtualNetwork{
	// 					Name: to.Ptr("Datacenter/CS-Management"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(false),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualnetworks/dvportgroup-19"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
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
	// 			Name: to.Ptr("Workload"),
	// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/resourcePools"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcepools/resgroup-26"),
	// 			Location: to.Ptr("westus2"),
	// 			PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 			Properties: &armvmwarecloudsimple.ResourcePoolProperties{
	// 				FullName: to.Ptr("Datacenter/Cluster/Workload"),
	// 			},
	// 		},
	// 		Status: to.Ptr(armvmwarecloudsimple.VirtualMachineStatusRunning),
	// 		TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualmachinetemplates/vm-34"),
	// 		VMID: to.Ptr("vm-100"),
	// 		Vmwaretools: to.Ptr("10309"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"inUse": to.Ptr("true"),
	// 	},
	// }
}
