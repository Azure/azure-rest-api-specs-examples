package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupCreateOrUpdateStandbyPool.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate_containerGroupCreateOrUpdateWithStandbyPool() {
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
			ContainerGroupProfile: &armcontainerinstance.ContainerGroupProfileReferenceDefinition{
				ID:       to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/democgp"),
				Revision: to.Ptr[int32](1),
			},
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("demo1"),
					Properties: &armcontainerinstance.ContainerProperties{
						ConfigMap: &armcontainerinstance.ConfigMap{
							KeyValuePairs: map[string]*string{
								"Newkey": to.Ptr("value"),
							},
						},
					},
				}},
			StandbyPoolProfile: &armcontainerinstance.StandbyPoolProfileDefinition{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.StandbyPool/standbyContainerGroupPools/demopool"),
			},
		},
		Location: to.Ptr("west us"),
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
	// 		ContainerGroupProfile: &armcontainerinstance.ContainerGroupProfileReferenceDefinition{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/democgp"),
	// 			Revision: to.Ptr[int32](1),
	// 		},
	// 		Containers: []*armcontainerinstance.Container{
	// 			{
	// 				Name: to.Ptr("demo1"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 					},
	// 					ConfigMap: &armcontainerinstance.ConfigMap{
	// 						KeyValuePairs: map[string]*string{
	// 							"Newkey": to.Ptr("value"),
	// 						},
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
	// 		ImageRegistryCredentials: []*armcontainerinstance.ImageRegistryCredential{
	// 		},
	// 		InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 		},
	// 		InstanceView: &armcontainerinstance.ContainerGroupPropertiesInstanceView{
	// 			Events: []*armcontainerinstance.Event{
	// 			},
	// 			State: to.Ptr("Running"),
	// 		},
	// 		IsCreatedFromStandbyPool: to.Ptr(true),
	// 		OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 	},
	// 	Name: to.Ptr("demo1"),
	// 	Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 	Location: to.Ptr("eastus2"),
	// }
}
