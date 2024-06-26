package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listSolutionsOperations.json
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
		// 			Name: to.Ptr("SolutionsOperation1"),
		// 			Display: &armmanagedapplications.OperationDisplay{
		// 				Description: to.Ptr("Description of the operation"),
		// 				Operation: to.Ptr("Read"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource1"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("SolutionssOperation2"),
		// 			Display: &armmanagedapplications.OperationDisplay{
		// 				Description: to.Ptr("Description of the operation"),
		// 				Operation: to.Ptr("Write"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource2"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
