package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/DeleteAccessReviewHistoryDefinition.json
func ExampleScopeAccessReviewHistoryDefinitionClient_DeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScopeAccessReviewHistoryDefinitionClient().DeleteByID(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
