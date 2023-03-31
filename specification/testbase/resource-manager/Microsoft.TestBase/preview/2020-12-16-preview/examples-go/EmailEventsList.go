package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/EmailEventsList.json
func ExampleEmailEventsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailEventsClient().NewListPager("contoso-rg", "contoso-testBaseAccount", nil)
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
		// page.EmailEventListResult = armtestbase.EmailEventListResult{
		// 	Value: []*armtestbase.EmailEventResource{
		// 		{
		// 			Name: to.Ptr("weekly-summary"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/emailEvents"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/emailEvents/weekly-summary"),
		// 			Properties: &armtestbase.EmailEventProperties{
		// 				DisplayName: to.Ptr("Weekly Summary"),
		// 				EventID: to.Ptr("weekly-summary"),
		// 				EventName: to.Ptr("weekly-summary"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily-summary"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/emailEvents"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/emailEvents/daily-summary"),
		// 			Properties: &armtestbase.EmailEventProperties{
		// 				DisplayName: to.Ptr("Daily Summary"),
		// 				EventID: to.Ptr("daily-summary"),
		// 				EventName: to.Ptr("daily-summary"),
		// 			},
		// 	}},
		// }
	}
}
