package armcarbonoptimization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/carbonoptimization/armcarbonoptimization"
)

// Generated from example definition: 2025-04-01/listOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcarbonoptimization.NewClientFactory(cred, nil)
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
		// page = armcarbonoptimization.OperationsClientListResponse{
		// 	OperationListResult: armcarbonoptimization.OperationListResult{
		// 		Value: []*armcarbonoptimization.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.Carbon/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcarbonoptimization.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Carbon"),
		// 					Resource: to.Ptr("queryCarbonEmissionReport"),
		// 					Operation: to.Ptr("CarbonService_ListCarbonEmissionReports"),
		// 					Description: to.Ptr("Returns carbon emission reports."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Carbon/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armcarbonoptimization.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Carbon"),
		// 					Resource: to.Ptr("QueryCarbonEmissionDataAvailableDateRange"),
		// 					Operation: to.Ptr("CarbonService_QueryCarbonEmissionDataAvailableDateRange"),
		// 					Description: to.Ptr("Returns carbon emission data available date range."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
