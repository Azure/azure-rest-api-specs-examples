package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/ListArcDeployments.json
func ExampleArcDeploymentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArcDeploymentsClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.ArcDeploymentsClientListResponse{
		// 	ArcDeploymentListResult: armcognitiveservices.ArcDeploymentListResult{
		// 		Value: []*armcognitiveservices.ArcDeployment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments/phi-3-arc"),
		// 				Name: to.Ptr("phi-3-arc"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/arcDeployments"),
		// 				Properties: &armcognitiveservices.ArcDeploymentProperties{
		// 					Model: &armcognitiveservices.ArcDeploymentModel{
		// 						Format: to.Ptr("OpenAI"),
		// 						Name: to.Ptr("Phi-3-mini-128k-instruct-generic-cpu"),
		// 						Version: to.Ptr("3"),
		// 						AssetID: to.Ptr("azureml://registries/azureml/models/Phi-3-mini-128k-instruct-generic-cpu/versions/3"),
		// 					},
		// 					ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
		// 					Runtime: to.Ptr(armcognitiveservices.ArcDeploymentRuntimeOnnx),
		// 					Compute: to.Ptr(armcognitiveservices.ArcDeploymentComputeTypeCPU),
		// 					Replicas: to.Ptr[int32](2),
		// 					Resources: &armcognitiveservices.ArcDeploymentKubernetesResources{
		// 						Requests: &armcognitiveservices.ArcDeploymentCPUMemoryResourceRequirements{
		// 							CPU: to.Ptr("8"),
		// 							Memory: to.Ptr("16Gi"),
		// 						},
		// 						Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
		// 							CPU: to.Ptr("8"),
		// 							Memory: to.Ptr("16Gi"),
		// 						},
		// 					},
		// 					NodeSelector: map[string]*string{
		// 						"agentpool": to.Ptr("cpu"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 					DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
		// 					RaiPolicyName: to.Ptr("Microsoft.DefaultV2"),
		// 					Capabilities: map[string]*string{
		// 						"chatCompletion": to.Ptr("true"),
		// 					},
		// 				},
		// 				SKU: &armcognitiveservices.ArcDeploymentSKU{
		// 					Name: to.Ptr(armcognitiveservices.ArcDeploymentSKUNameArc),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments/phi-3-gpu-arc"),
		// 				Name: to.Ptr("phi-3-gpu-arc"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/arcDeployments"),
		// 				Properties: &armcognitiveservices.ArcDeploymentProperties{
		// 					Model: &armcognitiveservices.ArcDeploymentModel{
		// 						Format: to.Ptr("OpenAI"),
		// 						Name: to.Ptr("Phi-3-mini-128k-instruct-cuda-gpu"),
		// 						Version: to.Ptr("2"),
		// 						AssetID: to.Ptr("azureml://registries/azureml/models/Phi-3-mini-128k-instruct-cuda-gpu/versions/2"),
		// 					},
		// 					ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
		// 					Runtime: to.Ptr(armcognitiveservices.ArcDeploymentRuntimeOnnx),
		// 					Compute: to.Ptr(armcognitiveservices.ArcDeploymentComputeTypeGpu),
		// 					Replicas: to.Ptr[int32](2),
		// 					Resources: &armcognitiveservices.ArcDeploymentKubernetesResources{
		// 						Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
		// 							Gpu: to.Ptr[int32](5),
		// 						},
		// 					},
		// 					NodeSelector: map[string]*string{
		// 						"agentpool": to.Ptr("a100"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 					DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
		// 					RaiPolicyName: to.Ptr("Microsoft.DefaultV2"),
		// 					Capabilities: map[string]*string{
		// 						"chatCompletion": to.Ptr("true"),
		// 					},
		// 				},
		// 				SKU: &armcognitiveservices.ArcDeploymentSKU{
		// 					Name: to.Ptr(armcognitiveservices.ArcDeploymentSKUNameArc),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments?api-version=2026-11-15-preview&$skipToken=next"),
		// 	},
		// }
	}
}
