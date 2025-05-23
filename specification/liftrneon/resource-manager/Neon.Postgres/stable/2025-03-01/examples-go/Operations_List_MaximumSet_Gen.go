package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armneonpostgres.OperationsClientListResponse{
		// 	OperationListResult: armneonpostgres.OperationListResult{
		// 		Value: []*armneonpostgres.Operation{
		// 			{
		// 				Name: to.Ptr("zgpfeuj"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armneonpostgres.OperationDisplay{
		// 					Provider: to.Ptr("lottcde"),
		// 					Resource: to.Ptr("lchjffakidtthnuaa"),
		// 					Operation: to.Ptr("ipqvheuesyujwjqhnmg"),
		// 					Description: to.Ptr("gicugbypsgqayjlfyrxvnietwzx"),
		// 				},
		// 				Origin: to.Ptr(armneonpostgres.OriginUser),
		// 				ActionType: to.Ptr(armneonpostgres.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aveozf"),
		// 	},
		// }
	}
}
