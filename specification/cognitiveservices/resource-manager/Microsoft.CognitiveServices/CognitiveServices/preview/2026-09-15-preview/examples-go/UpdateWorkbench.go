package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/UpdateWorkbench.json
func ExampleWorkbenchesClient_Update_updateWorkbench() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkbenchesClient().Update(ctx, "rgcognitiveservices", "myAccount", "myProject", "myWorkbench", armcognitiveservices.WorkbenchUpdate{
		Properties: &armcognitiveservices.WorkbenchUpdateProperties{
			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.WorkbenchesClientUpdateResponse{
	// 	Workbench: armcognitiveservices.Workbench{
	// 		Properties: &armcognitiveservices.WorkbenchProperties{
	// 			TargetClusterID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/vc-rg/providers/Microsoft.MachineLearningServices/virtualClusters/test-vc"),
	// 			ImageLink: to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"),
	// 			InstanceType: to.Ptr("Singularity.ND12_H100_v5-n1"),
	// 			GpuCount: to.Ptr[int32](1),
	// 			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
	// 			DatasetID: to.Ptr("dataset-12345"),
	// 			SSHSettings: &armcognitiveservices.SSHSettings{
	// 				SSHPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ..."),
	// 				AdminEnabled: to.Ptr(true),
	// 			},
	// 			ConnectivityEndpoints: &armcognitiveservices.ConnectivityEndpoints{
	// 				PublicIPAddress: to.Ptr("20.42.51.100"),
	// 				SSHPort: to.Ptr[int32](50000),
	// 			},
	// 			WebEndpoint: to.Ptr("https://myworkbench.eastus.api.azureml.ms"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.WorkbenchProvisioningStateSucceeded),
	// 			Status: to.Ptr(armcognitiveservices.WorkbenchStatusRunning),
	// 			Errors: []*armcognitiveservices.ErrorDetail{
	// 			},
	// 			CreationTime: to.Ptr(time.Date(2026, time.March, 24, 12, 0, 0, 0, time.UTC)),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject/workbenches/myWorkbench"),
	// 		Name: to.Ptr("myWorkbench"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/workbenches"),
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedBy: to.Ptr("xxx@microsoft.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.March, 24, 12, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("xxx@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.March, 24, 12, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
