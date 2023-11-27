package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupsGet_Failed.json
func ExampleContainerGroupsClient_Get_containerGroupsGetFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerGroupsClient().Get(ctx, "demo", "demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerGroup = armcontainerinstance.ContainerGroup{
	// 	Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
	// 		Containers: []*armcontainerinstance.Container{
	// 			{
	// 				Name: to.Ptr("demo1"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 					},
	// 					EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 					},
	// 					Image: to.Ptr("nginx"),
	// 					InstanceView: &armcontainerinstance.ContainerPropertiesInstanceView{
	// 						CurrentState: &armcontainerinstance.ContainerState{
	// 							DetailStatus: to.Ptr(""),
	// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-08T00:09:10.000Z"); return t}()),
	// 							State: to.Ptr("Waiting"),
	// 						},
	// 						Events: []*armcontainerinstance.Event{
	// 							{
	// 								Name: to.Ptr("Pulling"),
	// 								Type: to.Ptr("Normal"),
	// 								Count: to.Ptr[int32](1),
	// 								FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-08T00:09:01.000Z"); return t}()),
	// 								LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-08T00:09:01.000Z"); return t}()),
	// 								Message: to.Ptr("pulling image \"nginx\""),
	// 							},
	// 							{
	// 								Name: to.Ptr("Pulled"),
	// 								Type: to.Ptr("Normal"),
	// 								Count: to.Ptr[int32](1),
	// 								FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-08T00:09:10.000Z"); return t}()),
	// 								LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-08T00:09:10.000Z"); return t}()),
	// 								Message: to.Ptr("Successfully pulled image \"nginx\""),
	// 						}},
	// 						RestartCount: to.Ptr[int32](0),
	// 					},
	// 					Ports: []*armcontainerinstance.ContainerPort{
	// 						{
	// 							Port: to.Ptr[int32](80),
	// 					}},
	// 					Resources: &armcontainerinstance.ResourceRequirements{
	// 						Requests: &armcontainerinstance.ResourceRequests{
	// 							CPU: to.Ptr[float64](1),
	// 							MemoryInGB: to.Ptr[float64](1.5),
	// 						},
	// 					},
	// 					VolumeMounts: []*armcontainerinstance.VolumeMount{
	// 						{
	// 							Name: to.Ptr("volume1"),
	// 							MountPath: to.Ptr("/mnt/volume1"),
	// 							ReadOnly: to.Ptr(false),
	// 					}},
	// 				},
	// 		}},
	// 		ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 			{
	// 				Server: to.Ptr("azcloudconsoleregistry.azurecr.io"),
	// 				Username: to.Ptr("azcloudconsoleregistry"),
	// 		}},
	// 		InstanceView: &armcontainerinstance.ContainerGroupPropertiesInstanceView{
	// 			Events: []*armcontainerinstance.Event{
	// 				{
	// 					Name: to.Ptr("FailedMount"),
	// 					Type: to.Ptr("Normal"),
	// 					Count: to.Ptr[int32](1),
	// 					FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-08T00:19:10.000Z"); return t}()),
	// 					LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-08T00:19:10.000Z"); return t}()),
	// 					Message: to.Ptr("Output: mount error(2): Permission denied"),
	// 			}},
	// 			State: to.Ptr("Pending"),
	// 		},
	// 		IPAddress: &armcontainerinstance.IPAddress{
	// 			Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 			IP: to.Ptr("10.0.0.1"),
	// 			Ports: []*armcontainerinstance.Port{
	// 				{
	// 					Port: to.Ptr[int32](80),
	// 					Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 			}},
	// 		},
	// 		OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Volumes: []*armcontainerinstance.Volume{
	// 			{
	// 				Name: to.Ptr("volume1"),
	// 				AzureFile: &armcontainerinstance.AzureFileVolume{
	// 					ReadOnly: to.Ptr(false),
	// 					ShareName: to.Ptr("share1"),
	// 					StorageAccountName: to.Ptr("storage1"),
	// 				},
	// 		}},
	// 	},
	// 	Name: to.Ptr("demo1"),
	// 	Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 	Location: to.Ptr("WestUs"),
	// }
}
