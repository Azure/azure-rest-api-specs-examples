package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountUsagesList.json
func ExampleUsageClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsageClient().NewListPager("contoso-rg1", "contoso-testBaseAccount1", &armtestbase.UsageClientListOptions{Filter: nil})
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
		// page.AccountUsageDataList = armtestbase.AccountUsageDataList{
		// 	Value: []*armtestbase.AccountUsageData{
		// 		{
		// 			Name: &armtestbase.AccountUsageName{
		// 				LocalizedValue: to.Ptr("contoso-package1-usage"),
		// 				Value: to.Ptr("contoso-package1-usage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](7531),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/usages/contoso-package1-usage"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr("Seconds"),
		// 		},
		// 		{
		// 			Name: &armtestbase.AccountUsageName{
		// 				LocalizedValue: to.Ptr("contoso-package2-usage"),
		// 				Value: to.Ptr("contoso-package2-usage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3157),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/usages/contoso-package2-usage"),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr("Seconds"),
		// 	}},
		// }
	}
}
