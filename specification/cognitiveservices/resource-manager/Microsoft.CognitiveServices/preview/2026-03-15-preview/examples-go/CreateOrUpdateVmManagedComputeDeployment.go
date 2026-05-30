package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/CreateOrUpdateVmManagedComputeDeployment.json
func ExampleManagedComputeDeploymentsClient_BeginCreateOrUpdate_createOrUpdateVMManagedComputeDeployment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedComputeDeploymentsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "accountName", "gpt-oss-120b-byoc", armcognitiveservices.ManagedComputeDeployment{
		Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
			Model:              to.Ptr("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4"),
			DeploymentTemplate: to.Ptr("projects/my-project/deploymentTemplates/gpt-oss-120b-vllm-tuned/versions/2"),
			ComputeID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/computes/my-h100-pool"),
			Priority:           to.Ptr("High"),
		},
		SKU: &armcognitiveservices.SKU{
			Name:     to.Ptr("VmManagedCompute"),
			Capacity: to.Ptr[int32](2),
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
	// res = armcognitiveservices.ManagedComputeDeploymentsClientCreateOrUpdateResponse{
	// 	ManagedComputeDeployment: armcognitiveservices.ManagedComputeDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/gpt-oss-120b-byoc"),
	// 		Name: to.Ptr("gpt-oss-120b-byoc"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedComputeDeployments"),
	// 		Etag: to.Ptr("\"0x8D...\""),
	// 		Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
	// 			Model: to.Ptr("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4"),
	// 			DeploymentTemplate: to.Ptr("projects/my-project/deploymentTemplates/gpt-oss-120b-vllm-tuned/versions/2"),
	// 			ComputeID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/computes/my-h100-pool"),
	// 			Priority: to.Ptr("High"),
	// 			VersionUpgradeOption: to.Ptr(armcognitiveservices.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			ProvisioningDetails: &armcognitiveservices.ManagedComputeDeploymentProvisioningDetails{
	// 				Message: to.Ptr("Deployment is healthy and serving traffic."),
	// 				LastOperationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-12T10:25:00Z"); return t}()),
	// 			},
	// 			Routes: &armcognitiveservices.ManagedComputeDeploymentRoutes{
	// 				ChatCompletionsScoringPath: to.Ptr("/managed-deployments/gpt-oss-120b-byoc/v1/chat/completions"),
	// 			},
	// 		},
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("VmManagedCompute"),
	// 			Capacity: to.Ptr[int32](2),
	// 		},
	// 	},
	// }
}
