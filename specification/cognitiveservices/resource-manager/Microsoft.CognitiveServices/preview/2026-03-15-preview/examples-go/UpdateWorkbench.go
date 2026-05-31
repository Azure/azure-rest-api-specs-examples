package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/UpdateWorkbench.json
func ExampleWorkbenchesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkbenchesClient().BeginUpdate(ctx, "rgcognitiveservices", "myAccount", "myProject", "myWorkbench", armcognitiveservices.Workbench{
		Properties: &armcognitiveservices.WorkbenchProperties{
			TargetClusterID:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster"),
			ImageLink:              to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:v2"),
			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
			DatasetID:              to.Ptr("dataset-67890"),
			SSHSettings: &armcognitiveservices.SSHSettings{
				SSHPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ..."),
				AdminEnabled: to.Ptr(true),
			},
		},
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
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
	// res = armcognitiveservices.WorkbenchesClientUpdateResponse{
	// 	Workbench: armcognitiveservices.Workbench{
	// 		Properties: &armcognitiveservices.WorkbenchProperties{
	// 			TargetClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster"),
	// 			ImageLink: to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:v2"),
	// 			IdleTimeBeforeShutdown: to.Ptr("PT1H"),
	// 			DatasetID: to.Ptr("dataset-67890"),
	// 			SSHSettings: &armcognitiveservices.SSHSettings{
	// 				SSHPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQ..."),
	// 				AdminEnabled: to.Ptr(true),
	// 			},
	// 			ConnectivityEndpoints: &armcognitiveservices.ConnectivityEndpoints{
	// 				PublicIPAddress: to.Ptr("20.42.51.100"),
	// 				SSHPort: to.Ptr[int32](50000),
	// 			},
	// 			WebEndpoint: to.Ptr("https://myworkbench.eastus.api.azureml.ms"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ComputeProvisioningStateAccepted),
	// 			Errors: []*armcognitiveservices.ErrorDetail{
	// 			},
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T12:00:00.000Z"); return t}()),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject/workbenches/myWorkbench"),
	// 		Name: to.Ptr("myWorkbench"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/workbenches"),
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedBy: to.Ptr("xxx@microsoft.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T12:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xxx@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T12:00:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
