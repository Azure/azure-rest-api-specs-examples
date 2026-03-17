package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/TaskRunsGetDetails.json
func ExampleTaskRunsClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
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
	// res = armcontainerregistrytasks.TaskRunsClientGetDetailsResponse{
	// 	TaskRun: &armcontainerregistrytasks.TaskRun{
	// 		Name: to.Ptr("myRun"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 		ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
	// 		Properties: &armcontainerregistrytasks.TaskRunProperties{
	// 			ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
	// 			RunRequest: &armcontainerregistrytasks.EncodedTaskRunRequest{
	// 				Type: to.Ptr("EncodedTaskRunRequest"),
	// 				Credentials: &armcontainerregistrytasks.Credentials{
	// 				},
	// 				EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
	// 				EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				Platform: &armcontainerregistrytasks.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistrytasks.OSLinux),
	// 				},
	// 				Values: []*armcontainerregistrytasks.SetValue{
	// 				},
	// 			},
	// 			RunResult: &armcontainerregistrytasks.Run{
	// 				Name: to.Ptr("yd5"),
	// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 				ID: to.Ptr("/subscriptions/3647315e-0c5b-4ce4-8739-b071e144b2c9/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 				Properties: &armcontainerregistrytasks.RunProperties{
	// 					AgentConfiguration: &armcontainerregistrytasks.AgentProperties{
	// 						CPU: to.Ptr[int32](2),
	// 					},
	// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.9261521+00:00"); return t}()),
	// 					FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29.8792157+00:00"); return t}()),
	// 					IsArchiveEnabled: to.Ptr(true),
	// 					LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:29+00:00"); return t}()),
	// 					Platform: &armcontainerregistrytasks.PlatformProperties{
	// 						Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
	// 						OS: to.Ptr(armcontainerregistrytasks.OSLinux),
	// 					},
	// 					ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
	// 					RunID: to.Ptr("yd5"),
	// 					RunType: to.Ptr(armcontainerregistrytasks.RunTypeQuickRun),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:22.1348714+00:00"); return t}()),
	// 					Status: to.Ptr(armcontainerregistrytasks.RunStatusSucceeded),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
