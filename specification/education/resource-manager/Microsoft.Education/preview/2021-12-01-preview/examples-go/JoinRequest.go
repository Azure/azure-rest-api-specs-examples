package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequest.json
func ExampleJoinRequestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeducation.NewJoinRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		"{invoiceSectionName}",
		"{joinRequestName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
