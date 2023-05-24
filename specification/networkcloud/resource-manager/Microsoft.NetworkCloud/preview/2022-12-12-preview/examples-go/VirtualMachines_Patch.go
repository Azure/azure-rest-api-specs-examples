package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/VirtualMachines_Patch.json
func ExampleVirtualMachinesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginUpdate(ctx, "resourceGroupName", "virtualMachineName", armnetworkcloud.VirtualMachinePatchParameters{
		Properties: &armnetworkcloud.VirtualMachinePatchProperties{
			VMImageRepositoryCredentials: &armnetworkcloud.ImageRepositoryCredentials{
				Password:    to.Ptr("{password}"),
				RegistryURL: to.Ptr("myacr.azurecr.io"),
				Username:    to.Ptr("myuser"),
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
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
	// res.VirtualMachine = armnetworkcloud.VirtualMachine{
	// 	Name: to.Ptr("virtualMachineName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.VirtualMachineProperties{
	// 		AdminUsername: to.Ptr("username"),
	// 		BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
	// 		BootMethod: to.Ptr(armnetworkcloud.VirtualMachineBootMethodUEFI),
	// 		CloudServicesNetworkAttachment: &armnetworkcloud.NetworkAttachment{
	// 			AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
	// 			IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
	// 		},
	// 		ClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 		CPUCores: to.Ptr[int64](2),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.VirtualMachineDetailedStatusAvailable),
	// 		DetailedStatusMessage: to.Ptr("VM is up and healthy"),
	// 		MemorySizeGB: to.Ptr[int64](8),
	// 		NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
	// 			{
	// 				AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
	// 				DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
	// 				IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
	// 				IPv4Address: to.Ptr("198.51.100.1"),
	// 				IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0000"),
	// 				MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
	// 				NetworkAttachmentName: to.Ptr("netAttachName01"),
	// 		}},
	// 		NetworkData: to.Ptr("bmV0d29ya0RhdGVTYW1wbGU="),
	// 		PlacementHints: []*armnetworkcloud.VirtualMachinePlacementHint{
	// 			{
	// 				HintType: to.Ptr(armnetworkcloud.VirtualMachinePlacementHintTypeAffinity),
	// 				ResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
	// 				SchedulingExecution: to.Ptr(armnetworkcloud.VirtualMachineSchedulingExecutionHard),
	// 				Scope: to.Ptr(armnetworkcloud.VirtualMachinePlacementHintPodAffinityScope("")),
	// 		}},
	// 		PowerState: to.Ptr(armnetworkcloud.VirtualMachinePowerStateOn),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.VirtualMachineProvisioningStateSucceeded),
	// 		SSHPublicKeys: []*armnetworkcloud.SSHPublicKey{
	// 			{
	// 				KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
	// 		}},
	// 		StorageProfile: &armnetworkcloud.StorageProfile{
	// 			OSDisk: &armnetworkcloud.OsDisk{
	// 				CreateOption: to.Ptr(armnetworkcloud.OsDiskCreateOptionEphemeral),
	// 				DeleteOption: to.Ptr(armnetworkcloud.OsDiskDeleteOptionDelete),
	// 				DiskSizeGB: to.Ptr[int64](120),
	// 			},
	// 			VolumeAttachments: []*string{
	// 				to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName")},
	// 			},
	// 			UserData: to.Ptr("dXNlckRhdGVTYW1wbGU="),
	// 			VMDeviceModel: to.Ptr(armnetworkcloud.VirtualMachineDeviceModelTypeT2),
	// 			VMImage: to.Ptr("myacr.azurecr.io/foobar:latest"),
	// 			VMImageRepositoryCredentials: &armnetworkcloud.ImageRepositoryCredentials{
	// 				RegistryURL: to.Ptr("myacr.azurecr.io"),
	// 				Username: to.Ptr("myuser"),
	// 			},
	// 			Volumes: []*string{
	// 				to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName")},
	// 			},
	// 		}
}
