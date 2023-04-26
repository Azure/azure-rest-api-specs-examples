package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupsListByResourceGroup.json
func ExampleContainerGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerGroupsClient().NewListByResourceGroupPager("demo", nil)
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
		// page.ContainerGroupListResult = armcontainerinstance.ContainerGroupListResult{
		// 	Value: []*armcontainerinstance.ContainerGroup{
		// 		{
		// 			Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
		// 				Containers: []*armcontainerinstance.Container{
		// 					{
		// 						Name: to.Ptr("demo1"),
		// 						Properties: &armcontainerinstance.ContainerProperties{
		// 							Command: []*string{
		// 							},
		// 							EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
		// 							},
		// 							Image: to.Ptr("nginx"),
		// 							Ports: []*armcontainerinstance.ContainerPort{
		// 								{
		// 									Port: to.Ptr[int32](80),
		// 							}},
		// 							Resources: &armcontainerinstance.ResourceRequirements{
		// 								Requests: &armcontainerinstance.ResourceRequests{
		// 									CPU: to.Ptr[float64](1),
		// 									MemoryInGB: to.Ptr[float64](1.5),
		// 								},
		// 							},
		// 							VolumeMounts: []*armcontainerinstance.VolumeMount{
		// 								{
		// 									Name: to.Ptr("volume1"),
		// 									MountPath: to.Ptr("/mnt/volume1"),
		// 									ReadOnly: to.Ptr(false),
		// 							}},
		// 						},
		// 				}},
		// 				ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
		// 					{
		// 						Server: to.Ptr("azcloudconsoleregistry.azurecr.io"),
		// 						Username: to.Ptr("azcloudconsoleregistry"),
		// 				}},
		// 				IPAddress: &armcontainerinstance.IPAddress{
		// 					Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
		// 					IP: to.Ptr("10.0.0.1"),
		// 					Ports: []*armcontainerinstance.Port{
		// 						{
		// 							Port: to.Ptr[int32](80),
		// 							Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
		// 					}},
		// 				},
		// 				OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Volumes: []*armcontainerinstance.Volume{
		// 					{
		// 						Name: to.Ptr("volume1"),
		// 						AzureFile: &armcontainerinstance.AzureFileVolume{
		// 							ReadOnly: to.Ptr(false),
		// 							ShareName: to.Ptr("share1"),
		// 							StorageAccountName: to.Ptr("storage1"),
		// 						},
		// 				}},
		// 			},
		// 			Name: to.Ptr("demo1"),
		// 			Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
		// 			Location: to.Ptr("WestUs"),
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 		}},
		// 	}
	}
}
