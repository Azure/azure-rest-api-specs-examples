package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetPrivateCloud.json
func ExamplePrivateCloudsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateCloudsClient().Get(ctx, "myPrivateCloud", "westus2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateCloud = armvmwarecloudsimple.PrivateCloud{
	// 	Name: to.Ptr("myPrivateCloud"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/privateClouds"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.PrivateCloudProperties{
	// 		Type: to.Ptr("vSphere"),
	// 		AvailabilityZoneID: to.Ptr("az1"),
	// 		AvailabilityZoneName: to.Ptr("Availability Zone 1"),
	// 		ClustersNumber: to.Ptr[int32](1),
	// 		CreatedBy: to.Ptr("john.doe@cloudsimple.com"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-22T09:55:36.627Z"); return t}()),
	// 		DNSServers: []*string{
	// 			to.Ptr("100.0.0.8"),
	// 			to.Ptr("100.0.0.9")},
	// 			Expires: to.Ptr("n/a"),
	// 			NsxType: to.Ptr("Advanced"),
	// 			PlacementGroupID: to.Ptr("n1"),
	// 			PlacementGroupName: to.Ptr("Placement Group 1"),
	// 			PrivateCloudID: to.Ptr("{pc-id}"),
	// 			ResourcePools: []*armvmwarecloudsimple.ResourcePool{
	// 				{
	// 					Name: to.Ptr("Workload"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/resourcePools"),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26"),
	// 					Location: to.Ptr("westus2"),
	// 					PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					Properties: &armvmwarecloudsimple.ResourcePoolProperties{
	// 						FullName: to.Ptr("myPrivateCloud/Datacenter/Cluster/Workload"),
	// 					},
	// 			}},
	// 			State: to.Ptr("operational"),
	// 			TotalCPUCores: to.Ptr[int32](28),
	// 			TotalNodes: to.Ptr[int32](1),
	// 			TotalRAM: to.Ptr[int32](256),
	// 			TotalStorage: to.Ptr[float32](5.625),
	// 			VSphereVersion: to.Ptr("6.7u1"),
	// 			VcenterFqdn: to.Ptr("vcsa-4-westus2.az.cloudsimple.io"),
	// 			VcenterRefid: to.Ptr("100.0.0.6"),
	// 			VirtualMachineTemplates: []*armvmwarecloudsimple.VirtualMachineTemplate{
	// 				{
	// 					Name: to.Ptr("centos-template"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachineTemplates"),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualMachineTemplateProperties{
	// 						Path: to.Ptr("Datacenter/Workload VMs"),
	// 						AmountOfRAM: to.Ptr[int32](4096),
	// 						Controllers: []*armvmwarecloudsimple.VirtualDiskController{
	// 							{
	// 								Name: to.Ptr("SCSI controller 0"),
	// 								Type: to.Ptr("SCSI"),
	// 								ID: to.Ptr("1000"),
	// 								SubType: to.Ptr("LSI_PARALEL"),
	// 						}},
	// 						Disks: []*armvmwarecloudsimple.VirtualDisk{
	// 							{
	// 								ControllerID: to.Ptr("1000"),
	// 								IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
	// 								TotalSize: to.Ptr[int32](10485760),
	// 								VirtualDiskID: to.Ptr("2000"),
	// 								VirtualDiskName: to.Ptr("Hard disk 1"),
	// 						}},
	// 						GuestOS: to.Ptr("Other (32-bit)"),
	// 						GuestOSType: to.Ptr("other"),
	// 						Nics: []*armvmwarecloudsimple.VirtualNic{
	// 							{
	// 								MacAddress: to.Ptr("00:50:56:a6:7e:93"),
	// 								Network: &armvmwarecloudsimple.VirtualNetwork{
	// 									ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
	// 								},
	// 								NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
	// 								PowerOnBoot: to.Ptr(true),
	// 								VirtualNicID: to.Ptr("4000"),
	// 								VirtualNicName: to.Ptr("Network adapter 1"),
	// 						}},
	// 						NumberOfCores: to.Ptr[int32](2),
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 						Vmwaretools: to.Ptr("10309"),
	// 					},
	// 			}},
	// 			VirtualNetworks: []*armvmwarecloudsimple.VirtualNetwork{
	// 				{
	// 					Name: to.Ptr("Datacenter/CS-Management"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(false),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Datacenter/CS-Rescue"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(true),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-20"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Datacenter/CS-VSAN"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(false),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-21"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Datacenter/CS-VMotion"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(false),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-22"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("Datacenter/net-01"),
	// 					Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
	// 					Assignable: to.Ptr(true),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-35"),
	// 					Location: to.Ptr("westus2"),
	// 					Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
	// 						PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 					},
	// 			}},
	// 		},
	// 	}
}
