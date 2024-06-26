package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
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
		// page.OperationListResult = armresourcegraph.OperationListResult{
		// 	Value: []*armresourcegraph.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceGraph/operations/read"),
		// 			Display: &armresourcegraph.OperationDisplay{
		// 				Description: to.Ptr("Gets the list of supported operations"),
		// 				Operation: to.Ptr("Get Operations"),
		// 				Provider: to.Ptr("Microsoft Resource Graph"),
		// 				Resource: to.Ptr("Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceGraph/resources/read"),
		// 			Display: &armresourcegraph.OperationDisplay{
		// 				Description: to.Ptr("Submits a query on resources within specified subscriptions, the specified management groups, or against all access granted in the tenant."),
		// 				Operation: to.Ptr("Query resources"),
		// 				Provider: to.Ptr("Microsoft Resource Graph"),
		// 				Resource: to.Ptr("Resources"),
		// 			},
		// 	}},
		// }
	}
}
