package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/AccessReviewInstanceAcceptRecommendations.json
func ExampleAccessReviewInstanceClient_AcceptRecommendations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewAccessReviewInstanceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.AcceptRecommendations(ctx,
		"488a6d0e-0a63-4946-86e3-1f5bbc934661",
		"d9b9e056-7004-470b-bf21-1635e98487da",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
