package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/createDeployment.json
func ExampleConnectionClient_BeginCreateOrUpdateDeployment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionClient().BeginCreateOrUpdateDeployment(ctx, "resourceGroup-1", "testworkspace", "testConnection", "text-davinci-003", armmachinelearning.EndpointDeploymentResourcePropertiesBasicResource{
		Properties: &armmachinelearning.OpenAIEndpointDeploymentResourceProperties{
			Type: to.Ptr("Azure.OpenAI"),
			Model: &armmachinelearning.EndpointDeploymentModel{
				Name:    to.Ptr("text-davinci-003"),
				Format:  to.Ptr("OpenAI"),
				Version: to.Ptr("1"),
			},
			VersionUpgradeOption: to.Ptr(armmachinelearning.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
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
	// res = armmachinelearning.ConnectionClientCreateOrUpdateDeploymentResponse{
	// 	EndpointDeploymentResourcePropertiesBasicResource: armmachinelearning.EndpointDeploymentResourcePropertiesBasicResource{
	// 		Name: to.Ptr("text-davinci-003"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections/deployments"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/connections/testConnection/deployments/text-davinci-003"),
	// 		Properties: &armmachinelearning.OpenAIEndpointDeploymentResourceProperties{
	// 			Type: to.Ptr("Azure.OpenAI"),
	// 			Model: &armmachinelearning.EndpointDeploymentModel{
	// 				Name: to.Ptr("text-davinci-003"),
	// 				Format: to.Ptr("OpenAI"),
	// 				Version: to.Ptr("1"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmachinelearning.DefaultResourceProvisioningStateSucceeded),
	// 			RaiPolicyName: to.Ptr("Microsoft.Default"),
	// 			VersionUpgradeOption: to.Ptr(armmachinelearning.DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedBy: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 			LastModifiedBy: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
