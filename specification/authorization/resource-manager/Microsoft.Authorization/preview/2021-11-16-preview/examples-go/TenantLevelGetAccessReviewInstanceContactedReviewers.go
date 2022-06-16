package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-11-16-preview/examples/TenantLevelGetAccessReviewInstanceContactedReviewers.json
func ExampleTenantLevelAccessReviewInstanceContactedReviewersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewTenantLevelAccessReviewInstanceContactedReviewersClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("265785a7-a81f-4201-8a18-bb0db95982b7",
		"f25ed880-9c31-4101-bc57-825d8df3b58c",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
