package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/ContainerGroupEncryptionProperties.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate_containerGroupWithEncryptionProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerGroupsClient().BeginCreateOrUpdate(ctx, "demo", "demo1", armcontainerinstance.ContainerGroup{
		Identity: &armcontainerinstance.ContainerGroupIdentity{
			Type: to.Ptr(armcontainerinstance.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerinstance.UserAssignedIdentities{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity": {},
			},
		},
		Location: to.Ptr("eastus2"),
		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
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
					},
				},
			},
			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerinstance.ContainerGroupsClientCreateOrUpdateResponse{
	// 	ContainerGroup: armcontainerinstance.ContainerGroup{
	// 		Name: to.Ptr("demo1"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 		Identity: &armcontainerinstance.ContainerGroupIdentity{
	// 			Type: to.Ptr(armcontainerinstance.ResourceIdentityTypeUserAssigned),
	// 			TenantID: to.Ptr("<tenant id>"),
	// 			UserAssignedIdentities: map[string]*armcontainerinstance.UserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity": &armcontainerinstance.UserAssignedIdentities{
	// 					ClientID: to.Ptr("<client id>"),
	// 					PrincipalID: to.Ptr("<principal id>"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
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
	// 			EncryptionProperties: &armcontainerinstance.EncryptionProperties{
	// 				Identity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity"),
	// 				KeyName: to.Ptr("test-key"),
	// 				KeyVersion: to.Ptr("<key version>"),
	// 				VaultBaseURL: to.Ptr("https://testkeyvault.vault.azure.net/"),
	// 			},
	// 			ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 			},
	// 			InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 			},
	// 			InstanceView: &armcontainerinstance.ContainerGroupPropertiesInstanceView{
	// 				Events: []*armcontainerinstance.Event{
	// 				},
	// 				State: to.Ptr("Pending"),
	// 			},
	// 			IPAddress: &armcontainerinstance.IPAddress{
	// 				Type: to.Ptr(armcontainerinstance.ContainerGroupIPAddressTypePublic),
	// 				Ports: []*armcontainerinstance.Port{
	// 					{
	// 						Port: to.Ptr[int32](80),
	// 						Protocol: to.Ptr(armcontainerinstance.ContainerGroupNetworkProtocolTCP),
	// 					},
	// 				},
	// 			},
	// 			IsCreatedFromStandbyPool: to.Ptr(false),
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			ProvisioningState: to.Ptr("Pending"),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 		},
	// 	},
	// }
}
