package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2025-04-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armdurabletask.OperationsClientListResponse{
		// 	OperationListResult: armdurabletask.OperationListResult{
		// 		Value: []*armdurabletask.Operation{
		// 			{
		// 				Display: &armdurabletask.OperationDisplay{
		// 					Description: to.Ptr("Create or Update Durable Task Scheduler"),
		// 					Operation: to.Ptr("Create or Update Durable Task Scheduler"),
		// 					Provider: to.Ptr("Durable Task Scheduler"),
		// 					Resource: to.Ptr("Scheduler"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.DurableTask/schedulers/write"),
		// 			},
		// 			{
		// 				Display: &armdurabletask.OperationDisplay{
		// 					Description: to.Ptr("Delete Durable Task Scheduler"),
		// 					Operation: to.Ptr("Delete Durable Task Scheduler"),
		// 					Provider: to.Ptr("Durable Task Scheduler"),
		// 					Resource: to.Ptr("Scheduler"),
		// 				},
		// 				IsDataAction: to.Ptr(false),
		// 				Name: to.Ptr("Microsoft.DurableTask/schedulers/delete"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/akpblld"),
		// 	},
		// }
	}
}
