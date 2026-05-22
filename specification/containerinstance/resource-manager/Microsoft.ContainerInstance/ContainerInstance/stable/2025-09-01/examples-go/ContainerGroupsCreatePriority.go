package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/ContainerGroupsCreatePriority.json
func ExampleContainerGroupsClient_BeginCreateOrUpdate_containerGroupsCreateWithPriority() {
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
		Location: to.Ptr("eastus"),
		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
			Containers: []*armcontainerinstance.Container{
				{
					Name: to.Ptr("test-container-001"),
					Properties: &armcontainerinstance.ContainerProperties{
						Command: []*string{
							to.Ptr("/bin/sh"),
							to.Ptr("-c"),
							to.Ptr("sleep 10"),
						},
						Image: to.Ptr("alpine:latest"),
						Resources: &armcontainerinstance.ResourceRequirements{
							Requests: &armcontainerinstance.ResourceRequests{
								CPU:        to.Ptr[float64](1),
								MemoryInGB: to.Ptr[float64](1),
							},
						},
					},
				},
			},
			OSType:        to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
			Priority:      to.Ptr(armcontainerinstance.ContainerGroupPrioritySpot),
			RestartPolicy: to.Ptr(armcontainerinstance.ContainerGroupRestartPolicyNever),
			SKU:           to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
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
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armcontainerinstance.ContainerGroupPropertiesProperties{
	// 			Containers: []*armcontainerinstance.Container{
	// 				{
	// 					Name: to.Ptr("test-container-001"),
	// 					Properties: &armcontainerinstance.ContainerProperties{
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh"),
	// 							to.Ptr("-c"),
	// 							to.Ptr("sleep 10"),
	// 						},
	// 						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 						},
	// 						Image: to.Ptr("alpine:latest"),
	// 						Ports: []*armcontainerinstance.ContainerPort{
	// 						},
	// 						Resources: &armcontainerinstance.ResourceRequirements{
	// 							Requests: &armcontainerinstance.ResourceRequests{
	// 								CPU: to.Ptr[float64](1),
	// 								MemoryInGB: to.Ptr[float64](1),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			InitContainers: []*armcontainerinstance.InitContainerDefinition{
	// 			},
	// 			InstanceView: &armcontainerinstance.ContainerGroupPropertiesInstanceView{
	// 				Events: []*armcontainerinstance.Event{
	// 				},
	// 				State: to.Ptr("Succeeded"),
	// 			},
	// 			IsCreatedFromStandbyPool: to.Ptr(false),
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			Priority: to.Ptr(armcontainerinstance.ContainerGroupPrioritySpot),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RestartPolicy: to.Ptr(armcontainerinstance.ContainerGroupRestartPolicyNever),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 		},
	// 	},
	// }
}
