package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/OpenAI_CreateOrUpdate.json
func ExampleOpenAIClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOpenAIClient().CreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "default", &armelastic.OpenAIClientCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OpenAIIntegrationRPModel = armelastic.OpenAIIntegrationRPModel{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Elastic/monitors/openAIIntegration"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor/openAIIntegration/default"),
	// 	Properties: &armelastic.OpenAIIntegrationProperties{
	// 		LastRefreshAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-12T09:28:50.957Z"); return t}()),
	// 		OpenAIResourceEndpoint: to.Ptr("https://myOpenAI.openai.azure.com/"),
	// 		OpenAIResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myOpenAI"),
	// 	},
	// }
}
