package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfileCreateOrUpdate_EncryptionProperties.json
func ExampleContainerGroupProfilesClient_CreateOrUpdate_containerGroupProfileWithEncryptionProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerGroupProfilesClient().CreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroupProfile{
		Properties: &armcontainerinstance.ContainerGroupProfilePropertiesProperties{
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
							}},
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU:        to.Ptr[float64](1),
								MemoryInGB: to.Ptr[float64](1.5),
							},
						},
					},
				}},
			EncryptionProperties: &armcontainerinstance.EncryptionProperties{
				Identity:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity"),
				KeyName:      to.Ptr("test-key"),
				KeyVersion:   to.Ptr("<key version>"),
				VaultBaseURL: to.Ptr("https://testkeyvault.vault.azure.net"),
			},
			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{},
			IPAddress: &armcontainerinstance.IPAddress{
				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
				Ports: []*armcontainerinstance.Port{
					{
						Port:     to.Ptr[int32](80),
						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
					}},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
		},
		Location: to.Ptr("eastus2"),
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerGroupProfile = armcontainerinstance.ContainerGroupProfile{
	// 	Properties: &armcontainerinstance.ContainerGroupProfilePropertiesProperties{
	// 		Containers: []*armcontainerinstance.Container{
	// 			{
	// 				Name: to.Ptr("demo1"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 					},
	// 					EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 					},
	// 					Image: to.Ptr("nginx"),
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
	// 				},
	// 		}},
	// 		EncryptionProperties: &armcontainerinstance.EncryptionProperties{
	// 			Identity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity"),
	// 			KeyName: to.Ptr("test-key"),
	// 			KeyVersion: to.Ptr("<key version>"),
	// 			VaultBaseURL: to.Ptr("https://testkeyvault.vault.azure.net/"),
	// 		},
	// 		ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 		},
	// 		InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 		},
	// 		IPAddress: &armcontainerinstance.IPAddress{
	// 			Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 			Ports: []*armcontainerinstance.Port{
	// 				{
	// 					Port: to.Ptr[int32](80),
	// 					Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 			}},
	// 		},
	// 		OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 		Revision: to.Ptr[int32](0),
	// 		SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 	},
	// 	Name: to.Ptr("demo1"),
	// 	Type: to.Ptr("Microsoft.ContainerInstance/containerGroupProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ContainerInstance/containerGroupProfiles/demo1"),
	// 	Location: to.Ptr("eastus2"),
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
