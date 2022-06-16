package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/DeleteLab.json
func ExampleLabsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeducation.NewLabsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		"{invoiceSectionName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
