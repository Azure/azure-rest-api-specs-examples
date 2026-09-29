package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/ValidationTests_ListBySubscription_MinimumSet_Gen.json
func ExampleValidationTestsClient_NewListBySubscriptionPager_validationTestsListBySubscriptionMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewValidationTestsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armplatformvalidation.ValidationTestsClientListBySubscriptionResponse{
		// 	ValidationTestListResult: armplatformvalidation.ValidationTestListResult{
		// 		Value: []*armplatformvalidation.ValidationTest{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation"),
		// 				Name: to.Ptr("linux-quality-validation"),
		// 				Type: to.Ptr("Microsoft.PlatformValidation/validationTests"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
