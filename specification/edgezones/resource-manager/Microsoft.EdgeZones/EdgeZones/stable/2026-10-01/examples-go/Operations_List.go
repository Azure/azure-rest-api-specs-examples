package armedgezones_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgezones/armedgezones"
)

// Generated from example definition: 2026-10-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgezones.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armedgezones.OperationsClientListResponse{
		// 	OperationListResult: armedgezones.OperationListResult{
		// 		Value: []*armedgezones.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.EdgeZones/extendedZones/read"),
		// 				Display: &armedgezones.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.EdgeZones"),
		// 					Resource: to.Ptr("extendedZones"),
		// 					Operation: to.Ptr("read"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
