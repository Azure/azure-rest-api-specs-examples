package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2024-02-01-preview/examples/DeletedWorkbooksList.json
func ExampleDeletedWorkbooksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedWorkbooksClient().NewListBySubscriptionPager(&armapplicationinsights.DeletedWorkbooksClientListBySubscriptionOptions{Category: nil,
		Tags: []string{},
	})
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
		// page.DeletedWorkbooksListResult = armapplicationinsights.DeletedWorkbooksListResult{
		// 	Value: []*armapplicationinsights.DeletedWorkbook{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999"),
		// 			Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/55555555-6666-7777-8888-999999999999"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 			Properties: &armapplicationinsights.DeletedWorkbookProperties{
		// 				Description: to.Ptr("Sample workbook1"),
		// 				Category: to.Ptr("workbook"),
		// 				DisplayName: to.Ptr("Workbook 1"),
		// 				Revision: to.Ptr("12345678901234567890123456789010"),
		// 				SourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/MyApp1"),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T09:07:00.123Z"); return t}()),
		// 				UserID: to.Ptr("userId"),
		// 				Version: to.Ptr("Notebook/1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999998"),
		// 			Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/my-resource-group2/providers/Microsoft.Insights/workbooks/55555555-6666-7777-8888-999999999998"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"TagSample01": to.Ptr("sample01"),
		// 				"TagSample02": to.Ptr("sample02"),
		// 			},
		// 			Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 			Properties: &armapplicationinsights.DeletedWorkbookProperties{
		// 				Description: to.Ptr("Sample workbook2"),
		// 				Category: to.Ptr("workbook"),
		// 				DisplayName: to.Ptr("Workbook 2"),
		// 				Revision: to.Ptr("12345678901234567890123456789011"),
		// 				SourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/my-resource-group2/providers/Microsoft.Web/sites/MyApp2"),
		// 				TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-04T09:07:01.123Z"); return t}()),
		// 				UserID: to.Ptr("userId"),
		// 				Version: to.Ptr("Notebook/1.0"),
		// 			},
		// 	}},
		// }
	}
}
