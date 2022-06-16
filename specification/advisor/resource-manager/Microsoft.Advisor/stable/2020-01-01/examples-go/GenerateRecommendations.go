package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GenerateRecommendations.json
func ExampleRecommendationsClient_Generate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewRecommendationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Generate(ctx,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
