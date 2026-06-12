package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Endpoint/create.json
func ExampleEndpointClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointClient().BeginCreateOrUpdate(ctx, "test-rg", "aml-workspace-name", "Azure.OpenAI", armmachinelearning.EndpointResourcePropertiesBasicResource{
		Properties: &armmachinelearning.OpenAIEndpointResourceProperties{
			Name:                 to.Ptr("Azure.OpenAI"),
			AssociatedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveService/account/account-1"),
			EndpointType:         to.Ptr(armmachinelearning.EndpointTypeAzureOpenAI),
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
	// res = armmachinelearning.EndpointClientCreateOrUpdateResponse{
	// 	EndpointResourcePropertiesBasicResource: armmachinelearning.EndpointResourcePropertiesBasicResource{
	// 		Name: to.Ptr("Azure.OpenAI"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/endpoints"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/aml-workspace-name/endpoints/Azure.OpenAI"),
	// 		Properties: &armmachinelearning.OpenAIEndpointResourceProperties{
	// 			Name: to.Ptr("Azure.OpenAI"),
	// 			AssociatedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveService/account/account-1"),
	// 			EndpointType: to.Ptr(armmachinelearning.EndpointTypeAzureOpenAI),
	// 			EndpointURI: to.Ptr("https://www.contoso.com/"),
	// 			FailureReason: to.Ptr("some_string"),
	// 			ProvisioningState: to.Ptr(armmachinelearning.DefaultResourceProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			CreatedBy: to.Ptr("some_string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("some_string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
