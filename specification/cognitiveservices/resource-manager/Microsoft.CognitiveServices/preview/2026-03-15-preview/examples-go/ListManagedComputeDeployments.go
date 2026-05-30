package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/ListManagedComputeDeployments.json
func ExampleManagedComputeDeploymentsClient_NewListPager_listManagedComputeDeployments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedComputeDeploymentsClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.ManagedComputeDeploymentsClientListResponse{
		// 	ManagedComputeDeploymentListResult: armcognitiveservices.ManagedComputeDeploymentListResult{
		// 		Value: []*armcognitiveservices.ManagedComputeDeployment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/gpt-oss-120b-gpu"),
		// 				Name: to.Ptr("gpt-oss-120b-gpu"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedComputeDeployments"),
		// 				Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
		// 					Model: to.Ptr("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4"),
		// 					DeploymentTemplate: to.Ptr("azureml://registries/azureml-openai-oss/deploymenttemplates/gpt-oss-120b-short-context/versions/1"),
		// 					AcceleratorType: to.Ptr("H100_80GB"),
		// 					AcceleratorsPerInstance: to.Ptr[int32](4),
		// 					TotalAccelerators: to.Ptr[int32](4),
		// 					VersionUpgradeOption: to.Ptr(armcognitiveservices.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 					Routes: &armcognitiveservices.ManagedComputeDeploymentRoutes{
		// 						ChatCompletionsScoringPath: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/chat/completions"),
		// 						Swagger: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/swagger.json"),
		// 						MessagesAPIScoringPath: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/messages"),
		// 					},
		// 				},
		// 				SKU: &armcognitiveservices.SKU{
		// 					Name: to.Ptr("GlobalManagedCompute"),
		// 					Capacity: to.Ptr[int32](1),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/llama-3-gpu"),
		// 				Name: to.Ptr("llama-3-gpu"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedComputeDeployments"),
		// 				Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
		// 					Model: to.Ptr("azureml://registries/azureml-meta/models/Meta-Llama-3.1-405B-Instruct/versions/2"),
		// 					DeploymentTemplate: to.Ptr("azureml://registries/azureml-meta/deploymenttemplates/llama-3-405b-fp8/versions/1"),
		// 					AcceleratorType: to.Ptr("H100_80GB"),
		// 					AcceleratorsPerInstance: to.Ptr[int32](8),
		// 					TotalAccelerators: to.Ptr[int32](16),
		// 					VersionUpgradeOption: to.Ptr(armcognitiveservices.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 					Routes: &armcognitiveservices.ManagedComputeDeploymentRoutes{
		// 						ChatCompletionsScoringPath: to.Ptr("/managedComputeDeployments/llama-3-gpu/chat/completions"),
		// 						Swagger: to.Ptr("/managedComputeDeployments/llama-3-gpu/swagger.json"),
		// 					},
		// 				},
		// 				SKU: &armcognitiveservices.SKU{
		// 					Name: to.Ptr("GlobalManagedCompute"),
		// 					Capacity: to.Ptr[int32](2),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/gpt-oss-120b-byoc"),
		// 				Name: to.Ptr("gpt-oss-120b-byoc"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedComputeDeployments"),
		// 				Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
		// 					Model: to.Ptr("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4"),
		// 					DeploymentTemplate: to.Ptr("projects/my-project/deploymentTemplates/gpt-oss-120b-vllm-tuned/versions/2"),
		// 					ComputeID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/computes/my-h100-pool"),
		// 					Priority: to.Ptr("High"),
		// 					VersionUpgradeOption: to.Ptr(armcognitiveservices.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 					Routes: &armcognitiveservices.ManagedComputeDeploymentRoutes{
		// 						ChatCompletionsScoringPath: to.Ptr("/managed-deployments/gpt-oss-120b-byoc/v1/chat/completions"),
		// 					},
		// 				},
		// 				SKU: &armcognitiveservices.SKU{
		// 					Name: to.Ptr("VmManagedCompute"),
		// 					Capacity: to.Ptr[int32](2),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
