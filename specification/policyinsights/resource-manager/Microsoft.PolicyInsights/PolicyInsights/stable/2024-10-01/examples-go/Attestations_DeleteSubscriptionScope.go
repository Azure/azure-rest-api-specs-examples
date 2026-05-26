package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: 2024-10-01/Attestations_DeleteSubscriptionScope.json
func ExampleAttestationsClient_DeleteAtSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("35ee058e-5fa0-414c-8145-3ebb8d09b6e2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttestationsClient().DeleteAtSubscription(ctx, "790996e6-9871-4b1f-9cd9-ec42cd6ced1e", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicyinsights.AttestationsClientDeleteAtSubscriptionResponse{
	// }
}
