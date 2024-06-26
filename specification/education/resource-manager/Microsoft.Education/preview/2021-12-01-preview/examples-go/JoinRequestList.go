package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequestList.json
func ExampleJoinRequestsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJoinRequestsClient().NewListPager("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", &armeducation.JoinRequestsClientListOptions{IncludeDenied: to.Ptr(false)})
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
		// page.JoinRequestList = armeducation.JoinRequestList{
		// 	Value: []*armeducation.JoinRequestDetails{
		// 		{
		// 			Name: to.Ptr("{joinRequestName}"),
		// 			Type: to.Ptr("Microsoft.Education/JoinRequest"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}"),
		// 			Properties: &armeducation.JoinRequestProperties{
		// 				Email: to.Ptr("test@contoso.com"),
		// 				FirstName: to.Ptr("test"),
		// 				LastName: to.Ptr("user"),
		// 				Status: to.Ptr(armeducation.JoinRequestStatusPending),
		// 			},
		// 	}},
		// }
	}
}
