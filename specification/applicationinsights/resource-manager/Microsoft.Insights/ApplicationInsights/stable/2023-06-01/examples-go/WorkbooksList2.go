package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: 2023-06-01/WorkbooksList2.json
func ExampleWorkbooksClient_NewListBySubscriptionPager_workbooksList2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("6b643656-33eb-422f-aee8-3ac145d124af", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkbooksClient().NewListBySubscriptionPager(armapplicationinsights.CategoryTypeWorkbook, nil)
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
		// page = armapplicationinsights.WorkbooksClientListBySubscriptionResponse{
		// 	WorkbooksListResult: armapplicationinsights.WorkbooksListResult{
		// 		Value: []*armapplicationinsights.Workbook{
		// 		},
		// 	},
		// }
	}
}
