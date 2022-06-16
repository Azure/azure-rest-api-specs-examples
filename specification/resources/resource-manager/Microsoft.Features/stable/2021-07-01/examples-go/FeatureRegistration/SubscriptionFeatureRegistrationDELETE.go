package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationDELETE.json
func ExampleSubscriptionFeatureRegistrationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfeatures.NewSubscriptionFeatureRegistrationsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"subscriptionFeatureRegistrationGroupTestRG",
		"testFeature",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
