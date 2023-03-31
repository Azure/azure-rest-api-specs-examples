package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListVirtualMachines.json
func ExampleVirtualMachinesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListBySubscriptionPager(&armvmwarecloudsimple.VirtualMachinesClientListBySubscriptionOptions{Filter: nil,
		Top:       nil,
		SkipToken: nil,
	})
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
		// page.VirtualMachineListResponse = armvmwarecloudsimple.VirtualMachineListResponse{
		// 	Value: []*armvmwarecloudsimple.VirtualMachine{
		// 		{
		// 			Name: to.Ptr("virtualMachine-1"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup-1/providers/Microsoft.VMwareCloudSimple/virtualMachines/virtualMachine-1"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualMachineProperties{
		// 				AmountOfRAM: to.Ptr[int32](4096),
		// 				Controllers: []*armvmwarecloudsimple.VirtualDiskController{
		// 					{
		// 						Name: to.Ptr("SCSI controller 0"),
		// 						Type: to.Ptr("SCSI"),
		// 						ID: to.Ptr("1000"),
		// 						SubType: to.Ptr("LSI_PARALEL"),
		// 				}},
		// 				Disks: []*armvmwarecloudsimple.VirtualDisk{
		// 					{
		// 						ControllerID: to.Ptr("1000"),
		// 						IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
		// 						TotalSize: to.Ptr[int32](10485760),
		// 						VirtualDiskID: to.Ptr("2000"),
		// 						VirtualDiskName: to.Ptr("Hard disk 1"),
		// 				}},
		// 				GuestOS: to.Ptr("Other (32-bit)"),
		// 				GuestOSType: to.Ptr(armvmwarecloudsimple.GuestOSTypeOther),
		// 				Nics: []*armvmwarecloudsimple.VirtualNic{
		// 					{
		// 						MacAddress: to.Ptr("00:50:56:a6:d0:e1"),
		// 						Network: &armvmwarecloudsimple.VirtualNetwork{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
		// 						},
		// 						NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
		// 						PowerOnBoot: to.Ptr(true),
		// 						VirtualNicID: to.Ptr("4000"),
		// 						VirtualNicName: to.Ptr("Network adapter 1"),
		// 				}},
		// 				NumberOfCores: to.Ptr[int32](2),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				ResourcePool: &armvmwarecloudsimple.ResourcePool{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/rp-test/resourcepools/resgroup-26"),
		// 				},
		// 				TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
		// 				VMID: to.Ptr("vm-100"),
		// 				Vmwaretools: to.Ptr("10309"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"inUse": to.Ptr("true"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("virtualMachine-2"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup-2/providers/Microsoft.VMwareCloudSimple/virtualMachines/virtualMachine-2"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualMachineProperties{
		// 				AmountOfRAM: to.Ptr[int32](4096),
		// 				Controllers: []*armvmwarecloudsimple.VirtualDiskController{
		// 					{
		// 						Name: to.Ptr("SCSI controller 0"),
		// 						Type: to.Ptr("SCSI"),
		// 						ID: to.Ptr("1000"),
		// 						SubType: to.Ptr("LSI_PARALEL"),
		// 				}},
		// 				Disks: []*armvmwarecloudsimple.VirtualDisk{
		// 					{
		// 						ControllerID: to.Ptr("1000"),
		// 						IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
		// 						TotalSize: to.Ptr[int32](10485760),
		// 						VirtualDiskID: to.Ptr("2000"),
		// 						VirtualDiskName: to.Ptr("Hard disk 1"),
		// 				}},
		// 				GuestOS: to.Ptr("Other (32-bit)"),
		// 				GuestOSType: to.Ptr(armvmwarecloudsimple.GuestOSTypeOther),
		// 				Nics: []*armvmwarecloudsimple.VirtualNic{
		// 					{
		// 						MacAddress: to.Ptr("00:50:56:a6:33:12"),
		// 						Network: &armvmwarecloudsimple.VirtualNetwork{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
		// 						},
		// 						NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
		// 						PowerOnBoot: to.Ptr(true),
		// 						VirtualNicID: to.Ptr("4000"),
		// 						VirtualNicName: to.Ptr("Network adapter 1"),
		// 				}},
		// 				NumberOfCores: to.Ptr[int32](2),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				ResourcePool: &armvmwarecloudsimple.ResourcePool{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/rp-test/resourcepools/resgroup-26"),
		// 				},
		// 				TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
		// 				VMID: to.Ptr("vm-101"),
		// 				Vmwaretools: to.Ptr("10309"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"inUse": to.Ptr("true"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("virtualMachine-3"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup-1/providers/Microsoft.VMwareCloudSimple/virtualMachines/virtualMachine-3"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualMachineProperties{
		// 				AmountOfRAM: to.Ptr[int32](4096),
		// 				Controllers: []*armvmwarecloudsimple.VirtualDiskController{
		// 					{
		// 						Name: to.Ptr("SCSI controller 0"),
		// 						Type: to.Ptr("SCSI"),
		// 						ID: to.Ptr("1000"),
		// 						SubType: to.Ptr("LSI_PARALEL"),
		// 				}},
		// 				Disks: []*armvmwarecloudsimple.VirtualDisk{
		// 					{
		// 						ControllerID: to.Ptr("1000"),
		// 						IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
		// 						TotalSize: to.Ptr[int32](10485760),
		// 						VirtualDiskID: to.Ptr("2000"),
		// 						VirtualDiskName: to.Ptr("Hard disk 1"),
		// 				}},
		// 				GuestOS: to.Ptr("Other (32-bit)"),
		// 				GuestOSType: to.Ptr(armvmwarecloudsimple.GuestOSTypeOther),
		// 				Nics: []*armvmwarecloudsimple.VirtualNic{
		// 					{
		// 						MacAddress: to.Ptr("00:50:56:a6:63:f3"),
		// 						Network: &armvmwarecloudsimple.VirtualNetwork{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
		// 						},
		// 						NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
		// 						PowerOnBoot: to.Ptr(true),
		// 						VirtualNicID: to.Ptr("4000"),
		// 						VirtualNicName: to.Ptr("Network adapter 1"),
		// 				}},
		// 				NumberOfCores: to.Ptr[int32](2),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				ResourcePool: &armvmwarecloudsimple.ResourcePool{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/rp-test/resourcepools/resgroup-26"),
		// 				},
		// 				TemplateID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
		// 				VMID: to.Ptr("vm-102"),
		// 				Vmwaretools: to.Ptr("10309"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"inUse": to.Ptr("true"),
		// 			},
		// 	}},
		// }
	}
}
