package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TaskRunsCreate.json
func ExampleTaskRunsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTaskRunsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myRun", armcontainerregistry.TaskRun{
		Properties: &armcontainerregistry.TaskRunProperties{
			ForceUpdateTag: to.Ptr("test"),
			RunRequest: &armcontainerregistry.EncodedTaskRunRequest{
				Type:                 to.Ptr("EncodedTaskRunRequest"),
				Credentials:          &armcontainerregistry.Credentials{},
				EncodedTaskContent:   to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
				EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
					OS:           to.Ptr(armcontainerregistry.OSLinux),
				},
				Values: []*armcontainerregistry.SetValue{},
			},
		},
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
	// res.TaskRun = armcontainerregistry.TaskRun{
	// 	Name: to.Ptr("myrun"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
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
	// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd5"),
	// 			Properties: &armcontainerregistry.RunProperties{
	// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 					CPU: to.Ptr[int32](2),
	// 				},
	// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.926Z"); return t}()),
	// 				IsArchiveEnabled: to.Ptr(true),
	// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-06T17:23:21.000Z"); return t}()),
	// 				Platform: &armcontainerregistry.PlatformProperties{
	// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 					OS: to.Ptr(armcontainerregistry.OSLinux),
	// 				},
	// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 				RunID: to.Ptr("yd5"),
	// 				RunType: to.Ptr(armcontainerregistry.RunTypeQuickRun),
	// 				Status: to.Ptr(armcontainerregistry.RunStatusQueued),
	// 			},
	// 		},
	// 	},
	// }
}
