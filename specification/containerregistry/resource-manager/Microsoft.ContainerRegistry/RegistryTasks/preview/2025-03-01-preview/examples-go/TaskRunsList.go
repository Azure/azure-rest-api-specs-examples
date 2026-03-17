package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/TaskRunsList.json
func ExampleTaskRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTaskRunsClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerregistrytasks.TaskRunsClientListResponse{
		// 	TaskRunListResult: armcontainerregistrytasks.TaskRunListResult{
		// 		Value: []*armcontainerregistrytasks.TaskRun{
		// 			{
		// 				Name: to.Ptr("mytestrun"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/TaskRuns"),
		// 				ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/taskRuns/myRun"),
		// 				Properties: &armcontainerregistrytasks.TaskRunProperties{
		// 					ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
		// 					RunRequest: &armcontainerregistrytasks.EncodedTaskRunRequest{
		// 						Type: to.Ptr("EncodedTaskRunRequest"),
		// 						Credentials: &armcontainerregistrytasks.Credentials{
		// 						},
		// 						EncodedTaskContent: to.Ptr("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K"),
		// 						EncodedValuesContent: to.Ptr("Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg=="),
		// 						IsArchiveEnabled: to.Ptr(true),
		// 						Platform: &armcontainerregistrytasks.PlatformProperties{
		// 							Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
		// 							OS: to.Ptr(armcontainerregistrytasks.OSLinux),
		// 						},
		// 						Values: []*armcontainerregistrytasks.SetValue{
		// 						},
		// 					},
		// 					RunResult: &armcontainerregistrytasks.Run{
		// 						Name: to.Ptr("yd4"),
		// 						Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
		// 						ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/yd4"),
		// 						Properties: &armcontainerregistrytasks.RunProperties{
		// 							AgentConfiguration: &armcontainerregistrytasks.AgentProperties{
		// 								CPU: to.Ptr[int32](2),
		// 							},
		// 							CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:29.2278794+00:00"); return t}()),
		// 							FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:37.0349516+00:00"); return t}()),
		// 							IsArchiveEnabled: to.Ptr(true),
		// 							LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:37+00:00"); return t}()),
		// 							Platform: &armcontainerregistrytasks.PlatformProperties{
		// 								Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
		// 								OS: to.Ptr(armcontainerregistrytasks.OSLinux),
		// 							},
		// 							ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
		// 							RunID: to.Ptr("yd4"),
		// 							RunType: to.Ptr(armcontainerregistrytasks.RunTypeQuickRun),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T17:15:29.4589616+00:00"); return t}()),
		// 							Status: to.Ptr(armcontainerregistrytasks.RunStatusSucceeded),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
