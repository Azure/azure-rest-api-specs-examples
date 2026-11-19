package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/CreateOrUpdateArcDeploymentWithGpu.json
func ExampleArcDeploymentsClient_BeginCreateOrUpdate_createOrUpdateArcDeploymentWithGpu() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArcDeploymentsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "accountName", "phi-3-gpu-arc", armcognitiveservices.ArcDeployment{
		Properties: &armcognitiveservices.ArcDeploymentProperties{
			Model: &armcognitiveservices.ArcDeploymentModel{
				AssetID: to.Ptr("azureml://registries/azureml/models/Phi-3-mini-128k-instruct-cuda-gpu/versions/2"),
			},
			ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
			Replicas:    to.Ptr[int32](2),
			Resources: &armcognitiveservices.ArcDeploymentKubernetesResources{
				Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
					Gpu: to.Ptr[int32](5),
				},
			},
			NodeSelector: map[string]*string{
				"agentpool": to.Ptr("a100"),
			},
		},
		SKU: &armcognitiveservices.ArcDeploymentSKU{
			Name: to.Ptr(armcognitiveservices.ArcDeploymentSKUNameArc),
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
	// res = armcognitiveservices.ArcDeploymentsClientCreateOrUpdateResponse{
	// 	ArcDeployment: armcognitiveservices.ArcDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments/phi-3-gpu-arc"),
	// 		Name: to.Ptr("phi-3-gpu-arc"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/arcDeployments"),
	// 		Properties: &armcognitiveservices.ArcDeploymentProperties{
	// 			Model: &armcognitiveservices.ArcDeploymentModel{
	// 				Format: to.Ptr("OpenAI"),
	// 				Name: to.Ptr("Phi-3-mini-128k-instruct-cuda-gpu"),
	// 				Version: to.Ptr("2"),
	// 				AssetID: to.Ptr("azureml://registries/azureml/models/Phi-3-mini-128k-instruct-cuda-gpu/versions/2"),
	// 			},
	// 			ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
	// 			Runtime: to.Ptr(armcognitiveservices.ArcDeploymentRuntimeOnnx),
	// 			Compute: to.Ptr(armcognitiveservices.ArcDeploymentComputeTypeGpu),
	// 			Replicas: to.Ptr[int32](2),
	// 			Resources: &armcognitiveservices.ArcDeploymentKubernetesResources{
	// 				Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
	// 					Gpu: to.Ptr[int32](5),
	// 				},
	// 			},
	// 			NodeSelector: map[string]*string{
	// 				"agentpool": to.Ptr("a100"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
	// 			RaiPolicyName: to.Ptr("Microsoft.DefaultV2"),
	// 		},
	// 		SKU: &armcognitiveservices.ArcDeploymentSKU{
	// 			Name: to.Ptr(armcognitiveservices.ArcDeploymentSKUNameArc),
	// 		},
	// 	},
	// }
}
