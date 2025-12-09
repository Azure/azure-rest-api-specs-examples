package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TaskRunsGetDetails.json
func ExampleTaskRunsClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTaskRunsClient().GetDetails(ctx, "myResourceGroup", "myRegistry", "myRun", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("myRun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 	Properties: &armcontainerregistry.TaskRunProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
	// 			Type: to.Ptr("EncodedTaskRunRequest"),
	// 			IsArchiveEnabled: to.Ptr(true),
	// 			Credentials: &armcontainerregistry.Credentials{
	// 			},
	// 			EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 			EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 			Platform: &armcontainerregistry.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistry.OSLinux),
	// 			},
	// 			Values: []*armcontainerregistry.SetValue{
	// 			},
	// 		},
	// 		RunResult: &armcontainerregistry.Run{
	// 			Name: to.Ptr("yd5"),
	// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 			ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.926Z"); return t}()),
	// 				FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.879Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd5"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:22.134Z"); return t}()),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 			},
	// 		},
	// 	},
	// }
}
