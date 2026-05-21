package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/ContainerGroupProfilesGetPriority.json
func ExampleCGProfileClient_Get_containerGroupProfilesGetWithPriority() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCGProfileClient().Get(ctx, "demo", "demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerinstance.CGProfileClientGetResponse{
	// 	ContainerGroupProfile: armcontainerinstance.ContainerGroupProfile{
	// 		Name: to.Ptr("demo1"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/containerGroupProfiles"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/demo1"),
	// 		Location: to.Ptr("WestUs"),
	// 		Properties: &armcontainerinstance.ContainerGroupProfileProperties{
	// 			Containers: []*armcontainerinstance.Container{
	// 				{
	// 					Name: to.Ptr("demo1"),
	// 					Properties: &armcontainerinstance.ContainerProperties{
	// 						Command: []*string{
	// 						},
	// 						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 						},
	// 						Image: to.Ptr("nginx"),
	// 						Ports: []*armcontainerinstance.ContainerPort{
	// 							{
	// 								Port: to.Ptr[int32](80),
	// 							},
	// 						},
	// 						Resources: &armcontainerinstance.ResourceRequirements{
	// 							Requests: &armcontainerinstance.ResourceRequests{
	// 								CPU: to.Ptr[float64](1),
	// 								MemoryInGB: to.Ptr[float64](1.5),
	// 							},
	// 						},
	// 						VolumeMounts: []*armcontainerinstance.VolumeMount{
	// 							{
	// 								Name: to.Ptr("volume1"),
	// 								MountPath: to.Ptr("/mnt/volume1"),
	// 								ReadOnly: to.Ptr(false),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 				{
	// 					Server: to.Ptr("azcloudconsoleregistry.azurecr.io"),
	// 					Username: to.Ptr("azcloudconsoleregistry"),
	// 				},
	// 			},
	// 			IPAddress: &armcontainerinstance.IPAddress{
	// 				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 				IP: to.Ptr("10.0.0.1"),
	// 				Ports: []*armcontainerinstance.Port{
	// 					{
	// 						Port: to.Ptr[int32](80),
	// 						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 					},
	// 				},
	// 			},
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			Priority: to.Ptr(armcontainerinstance.ContainerGroupPrioritySpot),
	// 			Revision: to.Ptr[int32](1),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 			Volumes: []*armcontainerinstance.Volume{
	// 				{
	// 					Name: to.Ptr("volume1"),
	// 					AzureFile: &armcontainerinstance.AzureFileVolume{
	// 						ReadOnly: to.Ptr(false),
	// 						ShareName: to.Ptr("share1"),
	// 						StorageAccountName: to.Ptr("storage1"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
