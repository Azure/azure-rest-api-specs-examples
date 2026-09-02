package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/UpdateArcDeployment.json
func ExampleArcDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArcDeploymentsClient().BeginUpdate(ctx, "resourceGroupName", "accountName", "phi-3-arc", armcognitiveservices.ArcDeploymentUpdate{
		Properties: &armcognitiveservices.ArcDeploymentUpdateProperties{
			Replicas: to.Ptr[int32](3),
			Resources: &armcognitiveservices.ArcDeploymentPatchKubernetesResources{
				Requests: &armcognitiveservices.ArcDeploymentPatchCPUMemoryResourceRequirements{
					CPU:    to.Ptr("500m"),
					Memory: to.Ptr("2Gi"),
				},
				Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
					CPU:    to.Ptr("4"),
					Memory: to.Ptr("16Gi"),
				},
			},
			NodeSelector: map[string]*string{
				"agentpool": to.Ptr("cpu"),
			},
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
	// res = armcognitiveservices.ArcDeploymentsClientUpdateResponse{
	// 	ArcDeployment: armcognitiveservices.ArcDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/arcDeployments/phi-3-arc"),
	// 		Name: to.Ptr("phi-3-arc"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/arcDeployments"),
	// 		Etag: to.Ptr("\"0x8D...\""),
	// 		Properties: &armcognitiveservices.ArcDeploymentProperties{
	// 			Model: &armcognitiveservices.ArcDeploymentModel{
	// 				Format: to.Ptr("OpenAI"),
	// 				Name: to.Ptr("phi-3-mini"),
	// 			},
	// 			ExtensionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator"),
	// 			Runtime: to.Ptr(armcognitiveservices.ArcDeploymentRuntimeOnnx),
	// 			Compute: to.Ptr(armcognitiveservices.ArcDeploymentComputeTypeCPU),
	// 			Replicas: to.Ptr[int32](3),
	// 			Resources: &armcognitiveservices.ArcDeploymentKubernetesResources{
	// 				Requests: &armcognitiveservices.ArcDeploymentCPUMemoryResourceRequirements{
	// 					CPU: to.Ptr("500m"),
	// 					Memory: to.Ptr("2Gi"),
	// 				},
	// 				Limits: &armcognitiveservices.ArcDeploymentResourceRequirements{
	// 					CPU: to.Ptr("4"),
	// 					Memory: to.Ptr("16Gi"),
	// 				},
	// 			},
	// 			NodeSelector: map[string]*string{
	// 				"agentpool": to.Ptr("cpu"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
	// 			RaiPolicyName: to.Ptr("Microsoft.DefaultV2"),
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
