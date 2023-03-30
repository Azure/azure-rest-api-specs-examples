package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentList.json
func ExampleStudentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStudentsClient().NewListPager("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", &armeducation.StudentsClientListOptions{IncludeDeleted: nil})
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
		// page.StudentListResult = armeducation.StudentListResult{
		// 	Value: []*armeducation.StudentDetails{
		// 		{
		// 			Name: to.Ptr("{studentAlias}"),
		// 			Type: to.Ptr("Microsoft.Education/Students"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/students/{studentAlias}"),
		// 			Properties: &armeducation.StudentProperties{
		// 				Budget: &armeducation.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](100),
		// 				},
		// 				EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T21:43:54.161Z"); return t}()),
		// 				Email: to.Ptr("test@contoso.com"),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-09T21:43:54.161Z"); return t}()),
		// 				FirstName: to.Ptr("test"),
		// 				LastName: to.Ptr("user"),
		// 				Role: to.Ptr(armeducation.StudentRoleStudent),
		// 				Status: to.Ptr(armeducation.StudentLabStatusActive),
		// 				SubscriptionAlias: to.Ptr("000000000-0000-0000-0000-00000000000000"),
		// 				SubscriptionID: to.Ptr("000000000-0000-0000-0000-00000000000000"),
		// 				SubscriptionInviteLastSentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T21:43:54.161Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
