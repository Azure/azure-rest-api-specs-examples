package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: 2023-06-01/WorkbookRevisionsList.json
func ExampleWorkbooksClient_NewRevisionsListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("6b643656-33eb-422f-aee8-3ac145d124af", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkbooksClient().NewRevisionsListPager("my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", nil)
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
		// page = armapplicationinsights.WorkbooksClientRevisionsListResponse{
		// 	WorkbooksListResult: armapplicationinsights.WorkbooksListResult{
		// 		Value: []*armapplicationinsights.Workbook{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 				Name: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 				Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 				Properties: &armapplicationinsights.WorkbookProperties{
		// 					DisplayName: to.Ptr("My Workbook 1"),
		// 					UserID: to.Ptr("userId"),
		// 					SourceID: to.Ptr("Azure Monitor"),
		// 					SerializedData: nil,
		// 					Version: to.Ptr("Notebook/1.0"),
		// 					Category: to.Ptr("workbook"),
		// 					TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-04T09:07:17.735638Z"); return t}()),
		// 					StorageURI: nil,
		// 					Description: to.Ptr("Sample workbook"),
		// 					Revision: to.Ptr("1e2f8435b98248febee70c64ac22e1ab"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group/providers/Microsoft.Insights/workbooks/deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 				Name: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
		// 				Type: to.Ptr("Microsoft.Insights/workbooks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"TagSample01": to.Ptr("sample01"),
		// 					"TagSample02": to.Ptr("sample02"),
		// 				},
		// 				Kind: to.Ptr(armapplicationinsights.WorkbookSharedTypeKindShared),
		// 				Properties: &armapplicationinsights.WorkbookProperties{
		// 					DisplayName: to.Ptr("My Workbook 2"),
		// 					UserID: to.Ptr("userId"),
		// 					SourceID: to.Ptr("Azure Monitor"),
		// 					SerializedData: nil,
		// 					Version: to.Ptr("Notebook/1.0"),
		// 					Category: to.Ptr("workbook"),
		// 					TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-04T09:07:17.735638Z"); return t}()),
		// 					StorageURI: nil,
		// 					Description: to.Ptr("Sample workbook"),
		// 					Revision: to.Ptr("1e2f8435b98248febee70c64ac22e1bb"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
