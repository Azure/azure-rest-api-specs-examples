package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dbd896bc9a795bcb3ec7db0a340b517fd3059620/specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armelastic.OperationListResult{
		// 	Value: []*armelastic.OperationResult{
		// 		{
		// 			Name: to.Ptr("Microsoft.Elastic/monitors/write"),
		// 			Display: &armelastic.OperationDisplay{
		// 				Description: to.Ptr("Write monitors resource"),
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Microsoft.Elastic"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 	}},
		// }
	}
}
