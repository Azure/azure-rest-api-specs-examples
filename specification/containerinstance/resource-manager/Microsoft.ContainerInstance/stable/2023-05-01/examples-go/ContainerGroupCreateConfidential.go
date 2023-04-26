package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupCreateConfidential.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate_confidentialContainerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerGroupsClient().BeginCreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroup{
		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
			ConfidentialComputeProperties: &armcontainerinstance.ConfidentialComputeProperties{
				CcePolicy: to.Ptr("eyJhbGxvd19hbGwiOiB0cnVlLCAiY29udGFpbmVycyI6IHsibGVuZ3RoIjogMCwgImVsZW1lbnRzIjogbnVsbH19"),
			},
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("accdemo"),
					Properties: &armcontainerinstance.ContainerProperties{
						Command:              []*string{},
						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{},
						Image:                to.Ptr("confiimage"),
						Ports: []*armcontainerinstance.ContainerPort{
							{
								Port: to.Ptr[int32](8000),
							}},
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU:        to.Ptr[float64](1),
								MemoryInGB: to.Ptr[float64](1.5),
							},
						},
						SecurityContext: &armcontainerinstance.SecurityContextDefinition{
							Capabilities: &armcontainerinstance.SecurityContextCapabilitiesDefinition{
								Add: []*string{
									to.Ptr("CAP_NET_ADMIN")},
							},
							Privileged: to.Ptr(false),
						},
					},
				}},
			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
			IPAddress: &armcontainerinstance.IPAddress{
				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
				Ports: []*armcontainerinstance.Port{
					{
						Port:     to.Ptr[int32](8000),
						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
					}},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
			SKU:    to.Ptr(armcontainerinstance.ContainerGroupSKUConfidential),
		},
		Location: to.Ptr("westeurope"),
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
	// res.ContainerGroup = armcontainerinstance.ContainerGroup{
	// 	Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
	// 		ConfidentialComputeProperties: &armcontainerinstance.ConfidentialComputeProperties{
	// 			CcePolicy: to.Ptr("eyJhbGxvd19hbGwiOiB0cnVlLCAiY29udGFpbmVycyI6IHsibGVuZ3RoIjogMCwgImVsZW1lbnRzIjogbnVsbH19"),
	// 		},
	// 		Containers: []*armcontainerinstance.Container{
	// 			{
	// 				Name: to.Ptr("accdemo"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 					},
	// 					EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 					},
	// 					Image: to.Ptr("confiimage"),
	// 					Ports: []*armcontainerinstance.ContainerPort{
	// 						{
	// 							Port: to.Ptr[int32](8000),
	// 					}},
	// 					Resources: &armcontainerinstance.ResourceRequirements{
	// 						Requests: &armcontainerinstance.ResourceRequests{
	// 							CPU: to.Ptr[float64](1),
	// 							MemoryInGB: to.Ptr[float64](1.5),
	// 						},
	// 					},
	// 					SecurityContext: &armcontainerinstance.SecurityContextDefinition{
	// 						Capabilities: &armcontainerinstance.SecurityContextCapabilitiesDefinition{
	// 							Add: []*string{
	// 								to.Ptr("CAP_NET_ADMIN")},
	// 							},
	// 							Privileged: to.Ptr(false),
	// 						},
	// 					},
	// 			}},
	// 			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 			},
	// 			InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 			},
	// 			InstanceView: &armcontainerinstance.ContainerGroupPropertiesInstanceView{
	// 				Events: []*armcontainerinstance.Event{
	// 				},
	// 				State: to.Ptr("Running"),
	// 			},
	// 			IPAddress: &armcontainerinstance.IPAddress{
	// 				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 				Ports: []*armcontainerinstance.Port{
	// 					{
	// 						Port: to.Ptr[int32](8000),
	// 						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 				}},
	// 			},
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUConfidential),
	// 		},
	// 		Name: to.Ptr("demo1"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/test-rg/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 		Location: to.Ptr("westeurope"),
	// 	}
}
