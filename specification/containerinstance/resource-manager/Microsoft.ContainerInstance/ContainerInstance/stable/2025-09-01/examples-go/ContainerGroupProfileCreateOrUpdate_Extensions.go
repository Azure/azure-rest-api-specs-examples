package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/ContainerGroupProfileCreateOrUpdate_Extensions.json
func ExampleCGProfileClient_CreateOrUpdate_containerGroupProfileCreateWithExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCGProfileClient().CreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroupProfile{
		Location: to.Ptr("eastus2"),
		Properties: &armcontainerinstance.ContainerGroupProfileProperties{
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("demo1"),
					Properties: &armcontainerinstance.ContainerProperties{
						Command:              []*string{},
						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{},
						Image:                to.Ptr("nginx"),
						Ports: []*armcontainerinstance.ContainerPort{
							{
								Port: to.Ptr[int32](80),
							},
						},
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU:        to.Ptr[float64](1),
								MemoryInGB: to.Ptr[float64](1.5),
							},
						},
					},
				},
			},
			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
			IPAddress: &armcontainerinstance.IPAddress{
				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePrivate),
				Ports: []*armcontainerinstance.Port{
					{
						Port:     to.Ptr[int32](80),
						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
					},
				},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
			Extensions: []*armcontainerinstance.DeploymentExtensionSpec{
				{
					Name: to.Ptr("kube-proxy"),
					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
						ExtensionType: to.Ptr("kube-proxy"),
						ProtectedSettings: map[string]any{
							"kubeConfig": "<kubeconfig encoded string>",
						},
						Settings: map[string]any{
							"clusterCidr": "10.240.0.0/16",
							"kubeVersion": "v1.9.10",
						},
						Version: to.Ptr("1.0"),
					},
				},
				{
					Name: to.Ptr("vk-realtime-metrics"),
					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
						ExtensionType: to.Ptr("realtime-metrics"),
						Version:       to.Ptr("1.0"),
					},
				},
			},
		},
		Zones: []*string{
			to.Ptr("1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerinstance.CGProfileClientCreateOrUpdateResponse{
	// 	ContainerGroupProfile: armcontainerinstance.ContainerGroupProfile{
	// 		Name: to.Ptr("demo1"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/containerGroupProfiles"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/test-rg/providers/Microsoft.ContainerInstance/containerGroupProfiles/demo1"),
	// 		Location: to.Ptr("eastus2"),
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
	// 					},
	// 				},
	// 			},
	// 			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 			},
	// 			InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 			},
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			Revision: to.Ptr[int32](0),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 			Extensions: []*armcontainerinstance.DeploymentExtensionSpec{
	// 				{
	// 					Name: to.Ptr("kube-proxy"),
	// 					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
	// 						ExtensionType: to.Ptr("kube-proxy"),
	// 						Settings: map[string]any{
	// 							"clusterCidr": "10.240.0.0/16",
	// 							"kubeVersion": "v1.9.10",
	// 						},
	// 						Version: to.Ptr("1.0"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("vk-realtime-metrics"),
	// 					Properties: &armcontainerinstance.DeploymentExtensionSpecProperties{
	// 						ExtensionType: to.Ptr("realtime-metrics"),
	// 						Version: to.Ptr("1.0"),
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
