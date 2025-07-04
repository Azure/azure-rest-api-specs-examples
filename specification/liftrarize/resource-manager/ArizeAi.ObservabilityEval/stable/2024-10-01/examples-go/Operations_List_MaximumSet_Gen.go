package armarizeaiobservabilityeval_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/arizeaiobservabilityeval/armarizeaiobservabilityeval"
)

// Generated from example definition: 2024-10-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armarizeaiobservabilityeval.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armarizeaiobservabilityeval.OperationsClientListResponse{
		// 	OperationListResult: armarizeaiobservabilityeval.OperationListResult{
		// 		Value: []*armarizeaiobservabilityeval.Operation{
		// 			{
		// 				Name: to.Ptr("abinn"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armarizeaiobservabilityeval.OperationDisplay{
		// 					Provider: to.Ptr("qinuigznqqbkipyomvvzom"),
		// 					Resource: to.Ptr("tifectcvxqkbwioajbpb"),
		// 					Operation: to.Ptr("afzzgqqbdy"),
		// 					Description: to.Ptr("xqtjxdhoglmtrh"),
		// 				},
		// 				Origin: to.Ptr(armarizeaiobservabilityeval.OriginUser),
		// 				ActionType: to.Ptr(armarizeaiobservabilityeval.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
