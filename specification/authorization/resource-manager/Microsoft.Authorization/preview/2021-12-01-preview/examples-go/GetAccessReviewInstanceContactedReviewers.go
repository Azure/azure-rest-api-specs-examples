package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewInstanceContactedReviewers.json
func ExampleScopeAccessReviewInstanceContactedReviewersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScopeAccessReviewInstanceContactedReviewersClient().NewListPager("subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "265785a7-a81f-4201-8a18-bb0db95982b7", "f25ed880-9c31-4101-bc57-825d8df3b58c", nil)
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
		// page.AccessReviewContactedReviewerListResult = armauthorization.AccessReviewContactedReviewerListResult{
		// 	Value: []*armauthorization.AccessReviewContactedReviewer{
		// 		{
		// 			Name: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		// 			Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions/instances/contactedReviewers"),
		// 			ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/instances/f8882fec-7d56-4e97-ad6d-5e3f4557971d/contactedReviewers/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		// 			Properties: &armauthorization.AccessReviewContactedReviewerProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
		// 				UserDisplayName: to.Ptr("Bob"),
		// 				UserPrincipalName: to.Ptr("bob@contoso.com"),
		// 			},
		// 	}},
		// }
	}
}
