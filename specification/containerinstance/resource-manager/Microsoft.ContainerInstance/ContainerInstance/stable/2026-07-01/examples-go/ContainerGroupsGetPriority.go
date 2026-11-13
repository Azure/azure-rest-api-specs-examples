package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2026-07-01/ContainerGroupsGetPriority.json
func ExampleContainerGroupsClient_Get_containerGroupsGetWithPriority() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
	// res = armcontainerinstance.ContainerGroupsClientGetResponse{
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
	// 						InstanceView: &armcontainerinstance.ContainerPropertiesInstanceView{
	// 							CurrentState: &armcontainerinstance.ContainerState{
	// 								DetailStatus: to.Ptr("Completed"),
	// 								ExitCode: to.Ptr[int32](0),
	// 								FinishTime: to.Ptr(time.Date(2022, time.November, 11, 19, 33, 55, 565000000, time.UTC)),
	// 								StartTime: to.Ptr(time.Date(2022, time.November, 11, 19, 33, 44, 217000000, time.UTC)),
	// 								State: to.Ptr("Terminated"),
	// 							},
	// 							Events: []*armcontainerinstance.Event{
	// 								{
	// 									Name: to.Ptr("Pulling"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](2),
	// 									FirstTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 28, 22, 0, time.UTC)),
	// 									LastTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 32, 9, 0, time.UTC)),
	// 									Message: to.Ptr("pulling image \"alpine@sha256:1304f174557314a7ed9eddb4eab12fed12cb0cd9809e4c28f29af86979a3c870\""),
	// 								},
	// 								{
	// 									Name: to.Ptr("Pulled"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](2),
	// 									FirstTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 32, 8, 0, time.UTC)),
	// 									LastTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 32, 38, 0, time.UTC)),
	// 									Message: to.Ptr("Successfully pulled image \"alpine@sha256:1304f174557314a7ed9eddb4eab12fed12cb0cd9809e4c28f29af86979a3c870\""),
	// 								},
	// 								{
	// 									Name: to.Ptr("Started"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](1),
	// 									FirstTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 33, 44, 0, time.UTC)),
	// 									LastTimestamp: to.Ptr(time.Date(2022, time.November, 11, 19, 33, 44, 0, time.UTC)),
	// 									Message: to.Ptr("Started container"),
	// 								},
	// 							},
	// 							RestartCount: to.Ptr[int32](0),
	// 						},
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
	// 			OSType: to.Ptr(armcontainerinstance.OperatingSystemTypesLinux),
	// 			Priority: to.Ptr(armcontainerinstance.ContainerGroupPrioritySpot),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RestartPolicy: to.Ptr(armcontainerinstance.ContainerGroupRestartPolicyNever),
	// 			SKU: to.Ptr(armcontainerinstance.ContainerGroupSKUStandard),
	// 		},
	// 	},
	// }
}
