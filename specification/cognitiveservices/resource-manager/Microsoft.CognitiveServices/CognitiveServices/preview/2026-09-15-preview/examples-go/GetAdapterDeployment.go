package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/GetAdapterDeployment.json
func ExampleAdapterDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdapterDeploymentsClient().Get(ctx, "resourceGroupName", "accountName", "adapterDeploymentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AdapterDeploymentsClientGetResponse{
	// 	Etag: to.Ptr("\"0x8D1234567890ABC\""),
	// 	AdapterDeployment: armcognitiveservices.AdapterDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/adapterDeployments/adapterDeploymentName"),
	// 		Name: to.Ptr("adapterDeploymentName"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/adapterDeployments"),
	// 		Etag: to.Ptr("\"0x8D1234567890ABC\""),
	// 		Properties: &armcognitiveservices.AdapterDeploymentProperties{
	// 			SourceModelID: to.Ptr("azureai://accounts/accountName/projects/projectName/models/modelName/versions/1"),
	// 			TargetDeploymentName: to.Ptr("managedComputeDeploymentName"),
	// 			ActiveTargetDeploymentName: to.Ptr("managedComputeDeploymentName"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			LastOperation: &armcognitiveservices.AdapterDeploymentLastOperation{
	// 				Type: to.Ptr(armcognitiveservices.AdapterDeploymentOperationTypeCreate),
	// 				Status: to.Ptr(armcognitiveservices.AdapterDeploymentOperationStateSucceeded),
	// 				RequestedTargetDeploymentName: to.Ptr("managedComputeDeploymentName"),
	// 				StartedAt: to.Ptr(time.Date(2026, time.August, 31, 18, 30, 0, 0, time.UTC)),
	// 				CompletedAt: to.Ptr(time.Date(2026, time.August, 31, 18, 34, 0, 0, time.UTC)),
	// 			},
	// 		},
	// 	},
	// }
}
