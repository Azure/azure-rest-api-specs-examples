package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_DeleteInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssessmentsMetadataClient().DeleteInSubscription(ctx, "ca039e75-a276-4175-aebc-bcd41e4b14b7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
