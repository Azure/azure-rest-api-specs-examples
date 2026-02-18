package armmanagedops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedops/armmanagedops"
)

// Generated from example definition: 2025-07-28-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedops.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armmanagedops.OperationsClientListResponse{
		// 	OperationListResult: armmanagedops.OperationListResult{
		// 		Value: []*armmanagedops.Operation{
		// 			{
		// 				Name: to.Ptr("NameOfOperation"),
		// 				Display: &armmanagedops.OperationDisplay{
		// 					Provider: to.Ptr("myProvider"),
		// 					Resource: to.Ptr("myResource"),
		// 					Operation: to.Ptr("read"),
		// 					Description: to.Ptr("Read resources"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
