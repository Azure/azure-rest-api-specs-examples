package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armimpactreporting.OperationsClientListResponse{
		// 	OperationListResult: armimpactreporting.OperationListResult{
		// 		Value: []*armimpactreporting.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.Impact/workloadImpacts/write"),
		// 				Display: &armimpactreporting.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Impact"),
		// 					Resource: to.Ptr("workloadImpacts"),
		// 					Operation: to.Ptr("write"),
		// 					Description: to.Ptr("Write workloadImpact resources"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Impact/workloadImpacts/read"),
		// 				Display: &armimpactreporting.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Impact"),
		// 					Resource: to.Ptr("workloadimpacts"),
		// 					Operation: to.Ptr("read"),
		// 					Description: to.Ptr("Read workloadImpact resources"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Impact/impactCategories/read"),
		// 				Display: &armimpactreporting.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Impact"),
		// 					Resource: to.Ptr("impactCategories"),
		// 					Operation: to.Ptr("read"),
		// 					Description: to.Ptr("Read impactCategory resources"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
