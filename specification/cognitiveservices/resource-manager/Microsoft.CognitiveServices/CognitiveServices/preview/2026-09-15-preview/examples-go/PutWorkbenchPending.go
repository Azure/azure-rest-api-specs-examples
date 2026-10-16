package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/PutWorkbenchPending.json
func ExampleWorkbenchesClient_BeginCreateOrUpdate_putWorkbenchPending() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkbenchesClient().BeginCreateOrUpdate(ctx, "rgcognitiveservices", "myAccount", "myProject", "myWorkbench", armcognitiveservices.Workbench{
		Properties: &armcognitiveservices.WorkbenchProperties{
			TargetClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster"),
			ImageLink:       to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"),
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
	// res = armcognitiveservices.WorkbenchesClientCreateOrUpdateResponse{
	// 	Workbench: armcognitiveservices.Workbench{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject/workbenches/myWorkbench"),
	// 		Name: to.Ptr("myWorkbench"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/workbenches"),
	// 		Properties: &armcognitiveservices.WorkbenchProperties{
	// 			TargetClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster"),
	// 			ImageLink: to.Ptr("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.WorkbenchProvisioningStateCreating),
	// 			Status: to.Ptr(armcognitiveservices.WorkbenchStatusCreating),
	// 		},
	// 	},
	// }
}
