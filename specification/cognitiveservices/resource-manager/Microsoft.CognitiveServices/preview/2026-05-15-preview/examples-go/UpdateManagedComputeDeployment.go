package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/UpdateManagedComputeDeployment.json
func ExampleManagedComputeDeploymentsClient_BeginUpdate_updateManagedComputeDeployment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedComputeDeploymentsClient().BeginUpdate(ctx, "resourceGroupName", "accountName", "gpt-oss-120b-gpu", armcognitiveservices.PatchResourceSKU{
		SKU: &armcognitiveservices.SKU{
			Name:     to.Ptr("GlobalManagedCompute"),
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
	// res = armcognitiveservices.ManagedComputeDeploymentsClientUpdateResponse{
	// 	ManagedComputeDeployment: armcognitiveservices.ManagedComputeDeployment{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/gpt-oss-120b-gpu"),
	// 		Name: to.Ptr("gpt-oss-120b-gpu"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedComputeDeployments"),
	// 		Etag: to.Ptr("\"0x8D...\""),
	// 		Properties: &armcognitiveservices.ManagedComputeDeploymentProperties{
	// 			Model: to.Ptr("azureml://registries/azureml-openai-oss/models/gpt-oss-120b/versions/4"),
	// 			DeploymentTemplate: to.Ptr("azureml://registries/azureml-openai-oss/deploymenttemplates/gpt-oss-120b-short-context/versions/1"),
	// 			AcceleratorType: to.Ptr("H100_80GB"),
	// 			AcceleratorsPerInstance: to.Ptr[int32](4),
	// 			TotalAccelerators: to.Ptr[int32](8),
	// 			VersionUpgradeOption: to.Ptr(armcognitiveservices.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
	// 			Capabilities: map[string]*string{
	// 				"assetsV2": to.Ptr("true"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 			ProvisioningDetails: &armcognitiveservices.ManagedComputeDeploymentProvisioningDetails{
	// 				Message: to.Ptr("Scale operation completed successfully."),
	// 				LastOperationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-23T15:15:00Z"); return t}()),
	// 			},
	// 			Routes: &armcognitiveservices.ManagedComputeDeploymentRoutes{
	// 				ChatCompletionsScoringPath: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/chat/completions"),
	// 				Swagger: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/swagger.json"),
	// 				MessagesAPIScoringPath: to.Ptr("/managedComputeDeployments/gpt-oss-120b-gpu/messages"),
	// 			},
	// 		},
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("GlobalManagedCompute"),
	// 			Capacity: to.Ptr[int32](2),
	// 		},
	// 	},
	// }
}
