package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdatabasewatcher.OperationsClientListResponse{
		// 	OperationListResult: armdatabasewatcher.OperationListResult{
		// 		Value: []*armdatabasewatcher.Operation{
		// 			{
		// 				Name: to.Ptr("snzrdvltunnrz"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armdatabasewatcher.OperationDisplay{
		// 					Provider: to.Ptr("dtfrqzamclscchmghtxn"),
		// 					Resource: to.Ptr("lvlhnsfnquorjuuutjxex"),
		// 					Operation: to.Ptr("vbgvamoxqwthpbdghxzaw"),
		// 					Description: to.Ptr("nvbtuwwjfehylzmoatd"),
		// 				},
		// 				Origin: to.Ptr(armdatabasewatcher.OriginUser),
		// 				ActionType: to.Ptr(armdatabasewatcher.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/awfba"),
		// 	},
		// }
	}
}
