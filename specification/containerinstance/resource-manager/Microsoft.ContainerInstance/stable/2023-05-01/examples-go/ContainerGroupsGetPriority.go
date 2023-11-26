package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerGroupsGetPriority.json
func ExampleContainerGroupsClient_Get_containerGroupsGetWithPriority() {
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
	// 				Name: to.Ptr("test-container-001"),
	// 				Properties: &armcontainerinstance.ContainerProperties{
	// 					Command: []*string{
	// 						to.Ptr("/bin/sh"),
	// 						to.Ptr("-c"),
	// 						to.Ptr("sleep 10")},
	// 						EnvironmentVariables: []*armcontainerinstance.EnvironmentVariable{
	// 						},
	// 						Image: to.Ptr("alpine:latest"),
	// 						InstanceView: &armcontainerinstance.ContainerPropertiesInstanceView{
	// 							CurrentState: &armcontainerinstance.ContainerState{
	// 								DetailStatus: to.Ptr("Completed"),
	// 								ExitCode: to.Ptr[int32](0),
	// 								FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:33:55.565Z"); return t}()),
	// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:33:44.217Z"); return t}()),
	// 								State: to.Ptr("Terminated"),
	// 							},
	// 							Events: []*armcontainerinstance.Event{
	// 								{
	// 									Name: to.Ptr("Pulling"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](2),
	// 									FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:28:22.000Z"); return t}()),
	// 									LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:32:09.000Z"); return t}()),
	// 									Message: to.Ptr("pulling image \"alpine@sha256:1304f174557314a7ed9eddb4eab12fed12cb0cd9809e4c28f29af86979a3c870\""),
	// 								},
	// 								{
	// 									Name: to.Ptr("Pulled"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](2),
	// 									FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:32:08.000Z"); return t}()),
	// 									LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:32:38.000Z"); return t}()),
	// 									Message: to.Ptr("Successfully pulled image \"alpine@sha256:1304f174557314a7ed9eddb4eab12fed12cb0cd9809e4c28f29af86979a3c870\""),
	// 								},
	// 								{
	// 									Name: to.Ptr("Started"),
	// 									Type: to.Ptr("Normal"),
	// 									Count: to.Ptr[int32](1),
	// 									FirstTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:33:44.000Z"); return t}()),
	// 									LastTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T19:33:44.000Z"); return t}()),
	// 									Message: to.Ptr("Started container"),
	// 							}},
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
	// 			}},
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
	// 		Name: to.Ptr("demo1"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/containerGroups"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroups/demo1"),
	// 		Location: to.Ptr("eastus"),
	// 	}
}
