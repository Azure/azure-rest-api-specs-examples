package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/PutDeploymentWithSpeculativeDecoding.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_putDeploymentWithSpeculativeDecoding() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "accountName", "deploymentName", armcognitiveservices.Deployment{
		Properties: &armcognitiveservices.DeploymentProperties{
			DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
			Model: &armcognitiveservices.DeploymentModel{
				Format:  to.Ptr("Fireworks"),
				Name:    to.Ptr("FW-Qwen3-14B"),
				Version: to.Ptr("1"),
			},
			SpeculativeDecoding: &armcognitiveservices.DeploymentSpeculativeDecoding{
				DraftModel: &armcognitiveservices.DeploymentModel{
					Format:  to.Ptr("FireworksCustom"),
					Name:    to.Ptr("testDraftModel"),
					Version: to.Ptr("1"),
					Source:  to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/projects/projectName"),
				},
				DraftTokenCount: to.Ptr[int32](4),
			},
			ServiceTier: to.Ptr(armcognitiveservices.ServiceTierDefault),
		},
		SKU: &armcognitiveservices.SKU{
			Name:     to.Ptr("GlobalProvisionedManaged"),
			Capacity: to.Ptr[int32](80),
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
	// res = armcognitiveservices.DeploymentsClientCreateOrUpdateResponse{
	// 	Deployment: armcognitiveservices.Deployment{
	// 		Name: to.Ptr("deploymentName"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/deployments"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/deployments/deploymentName"),
	// 		Properties: &armcognitiveservices.DeploymentProperties{
	// 			DeploymentState: to.Ptr(armcognitiveservices.DeploymentStateRunning),
	// 			Model: &armcognitiveservices.DeploymentModel{
	// 				Format: to.Ptr("Fireworks"),
	// 				Name: to.Ptr("FW-Qwen3-14B"),
	// 				Version: to.Ptr("1"),
	// 			},
	// 			SpeculativeDecoding: &armcognitiveservices.DeploymentSpeculativeDecoding{
	// 				DraftModel: &armcognitiveservices.DeploymentModel{
	// 					Format: to.Ptr("FireworksCustom"),
	// 					Name: to.Ptr("testDraftModel"),
	// 					Version: to.Ptr("1"),
	// 					Source: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/projects/projectName"),
	// 				},
	// 				DraftTokenCount: to.Ptr[int32](4),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.DeploymentProvisioningStateSucceeded),
	// 			ServiceTier: to.Ptr(armcognitiveservices.ServiceTierDefault),
	// 		},
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("GlobalProvisionedManaged"),
	// 			Capacity: to.Ptr[int32](80),
	// 		},
	// 	},
	// }
}
