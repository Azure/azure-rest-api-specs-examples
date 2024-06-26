package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
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
		// page.OperationListResult = armselfhelp.OperationListResult{
		// 	Value: []*armselfhelp.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Help/diagnostics/read"),
		// 			Display: &armselfhelp.OperationDisplay{
		// 				Description: to.Ptr("Created and Reads a diagnostic resource to troubleshoot an issue with a resource."),
		// 				Operation: to.Ptr("Create/Read a Diagnostic"),
		// 				Provider: to.Ptr("Microsoft.Diagnostics"),
		// 				Resource: to.Ptr("Diagnostics"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Help/discoverySolutions/read"),
		// 			Display: &armselfhelp.OperationDisplay{
		// 				Description: to.Ptr("Returns list of solutions based on ResourceType or ProblemClassficationId"),
		// 				Operation: to.Ptr("List of available solutions."),
		// 				Provider: to.Ptr("Microsoft.Help"),
		// 				Resource: to.Ptr("DiscoverySolutions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
