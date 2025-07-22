package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/UpdateDeployment.json
func ExampleDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginUpdate(ctx, "resourceGroupName", "accountName", "deploymentName", armcognitiveservices.PatchResourceTagsAndSKU{
		SKU: &armcognitiveservices.SKU{
			Name:     to.Ptr("Standard"),
			Capacity: to.Ptr[int32](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Deployment = armcognitiveservices.Deployment{
	// 	Name: to.Ptr("deploymentName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/deployments/deploymentName"),
	// 	Properties: &armcognitiveservices.DeploymentProperties{
	// 		Model: &armcognitiveservices.DeploymentModel{
	// 			Name: to.Ptr("ada"),
	// 			Format: to.Ptr("OpenAI"),
	// 			Version: to.Ptr("1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcognitiveservices.DeploymentProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("Standard"),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// }
}
