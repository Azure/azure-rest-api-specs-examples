package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdellstorage.OperationsClientListResponse{
		// 	OperationListResult: armdellstorage.OperationListResult{
		// 		Value: []*armdellstorage.Operation{
		// 			{
		// 				Name: to.Ptr("rndmegxufrfrypvgt"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armdellstorage.OperationDisplay{
		// 					Provider: to.Ptr("bo"),
		// 					Resource: to.Ptr("mlvmbnno"),
		// 					Operation: to.Ptr("welhqbqmwwzytgtskhmxurcwhpsndg"),
		// 					Description: to.Ptr("gfcmuplcdhodralpafiiwxytg"),
		// 				},
		// 				Origin: to.Ptr(armdellstorage.OriginUser),
		// 				ActionType: to.Ptr(armdellstorage.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
