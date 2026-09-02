package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/GetArcDeployment.json
func ExampleArcDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArcDeploymentsClient().Get(ctx, "resourceGroupName", "accountName", "qwen-template-arc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.ArcDeploymentsClientGetResponse{
	// 	ArcDeployment: armcognitiveservices.ArcDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments/qwen-template-arc"),
	// 		Name: to.Ptr("qwen-template-arc"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/arcDeployments"),
	// 		Etag: to.Ptr("\"0x8D...\""),
	// 		Properties: &armcognitiveservices.ArcDeploymentProperties{
	// 			Model: &armcognitiveservices.ArcDeploymentModel{
	// 				Format: to.Ptr("OpenAI"),
	// 				Name: to.Ptr("qwen3"),
	// 			},
	// 			ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
	// 			Runtime: to.Ptr(armcognitiveservices.ArcDeploymentRuntimeVllm),
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
	// 			DeploymentTemplate: to.Ptr("azureml://registries/azureml-openai-oss/deploymenttemplates/vllm-qwen--qwen3-5-0-8b/versions/1"),
	// 			VllmParameters: &armcognitiveservices.ArcDeploymentVllmParameters{
	// 				TensorParallelSize: to.Ptr[int32](2),
	// 				MaxModelLen: to.Ptr[int32](8192),
	// 				GpuMemoryUtilization: to.Ptr[float32](0.9),
	// 				EnforceEager: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			ProvisioningDetails: &armcognitiveservices.ArcDeploymentProvisioningDetails{
	// 				Message: to.Ptr("Deployment is healthy and serving traffic."),
	// 				LastOperationTimestamp: to.Ptr(time.Date(2026, time.July, 15, 14, 45, 0, 0, time.UTC)),
	// 			},
	// 			InferenceEndpoint: to.Ptr("https://edge-cluster.eastus.inference.ml.azure.com"),
	// 			Capabilities: map[string]*string{
	// 				"chatCompletion": to.Ptr("true"),
	// 			},
	// 		},
	// 		SKU: &armcognitiveservices.ArcDeploymentSKU{
	// 			Name: to.Ptr(armcognitiveservices.ArcDeploymentSKUNameArc),
	// 		},
	// 	},
	// }
}
