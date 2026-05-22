package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: 2021-12-01-preview/LabListWithInvoiceSectionName.json
func ExampleLabsClient_NewListPager_labListWithInvoiceSectionName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLabsClient().NewListPager("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", &armeducation.LabsClientListOptions{
		IncludeBudget: to.Ptr(true)})
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
		// page = armeducation.LabsClientListResponse{
		// 	LabListResult: armeducation.LabListResult{
		// 		Value: []*armeducation.LabDetails{
		// 			{
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.Education/Labs"),
		// 				Properties: &armeducation.LabProperties{
		// 					DisplayName: to.Ptr("example lab"),
		// 					BudgetPerStudent: &armeducation.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float32](100),
		// 					},
		// 					Description: to.Ptr("example lab description"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-09T21:25:56.838Z"); return t}()),
		// 					EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T21:25:56.838Z"); return t}()),
		// 					Status: to.Ptr(armeducation.LabStatusActive),
		// 					TotalBudget: &armeducation.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float32](500),
		// 					},
		// 					TotalAllocatedBudget: &armeducation.Amount{
		// 						Currency: to.Ptr("USD"),
		// 						Value: to.Ptr[float32](500),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
