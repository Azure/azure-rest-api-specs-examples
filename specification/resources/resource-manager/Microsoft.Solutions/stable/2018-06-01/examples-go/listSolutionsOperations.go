package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listSolutionsOperations.json
func ExampleApplicationClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationClient().NewListOperationsPager(nil)
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
		// page.OperationListResult = armmanagedapplications.OperationListResult{
		// 	Value: []*armmanagedapplications.Operation{
		// 		{
		// 			Name: to.Ptr("SolutionsOpeartion1"),
		// 			Display: &armmanagedapplications.OperationDisplay{
		// 				Operation: to.Ptr("Read"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SolutionssOpeartion2"),
		// 			Display: &armmanagedapplications.OperationDisplay{
		// 				Operation: to.Ptr("Write"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource2"),
		// 			},
		// 	}},
		// }
	}
}
