package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armtestbase.OperationListResult{
		// 	Value: []*armtestbase.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.TestBase/testBaseAccounts/read"),
		// 			Display: &armtestbase.OperationDisplay{
		// 				Description: to.Ptr("View the properties of a Test Base Account."),
		// 				Operation: to.Ptr("View Test Base Account"),
		// 				Provider: to.Ptr("Microsoft Test Base"),
		// 				Resource: to.Ptr("Test Base Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TestBase/testBaseAccounts/write"),
		// 			Display: &armtestbase.OperationDisplay{
		// 				Description: to.Ptr("Create a new Test Base Account or update the properties of an existing Test Base Account."),
		// 				Operation: to.Ptr("Update Test Base Account"),
		// 				Provider: to.Ptr("Microsoft Test Base"),
		// 				Resource: to.Ptr("Test Base Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.TestBase/testBaseAccounts/delete"),
		// 			Display: &armtestbase.OperationDisplay{
		// 				Description: to.Ptr("Delete a Test Base Account."),
		// 				Operation: to.Ptr("Delete Test Base Account"),
		// 				Provider: to.Ptr("Microsoft Test Base"),
		// 				Resource: to.Ptr("Test Base Account"),
		// 			},
		// 	}},
		// }
	}
}
