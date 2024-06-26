package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsOperationsList_example.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdigitaltwins.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armdigitaltwins.OperationListResult{
		// 	Value: []*armdigitaltwins.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DigitalTwins/models/read"),
		// 			Display: &armdigitaltwins.OperationDisplay{
		// 				Description: to.Ptr("Read any Model"),
		// 				Operation: to.Ptr("Read Model"),
		// 				Provider: to.Ptr("Azure Digital Twins"),
		// 				Resource: to.Ptr("Model"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/write"),
		// 			Display: &armdigitaltwins.OperationDisplay{
		// 				Description: to.Ptr("Create or update a Digital Twins Service instance."),
		// 				Operation: to.Ptr("Create or update a Digital Twins Service instance."),
		// 				Provider: to.Ptr("Microsoft.DigitalTwins"),
		// 				Resource: to.Ptr("digitalTwinsInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/read"),
		// 			Display: &armdigitaltwins.OperationDisplay{
		// 				Description: to.Ptr("Get a Digital Twins Service instance."),
		// 				Operation: to.Ptr("Get a Digital Twins Service instance."),
		// 				Provider: to.Ptr("Microsoft.DigitalTwins"),
		// 				Resource: to.Ptr("digitalTwinsInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DigitalTwins/digitalTwinsInstances/delete"),
		// 			Display: &armdigitaltwins.OperationDisplay{
		// 				Description: to.Ptr("Delete a Digital Twins Service instance."),
		// 				Operation: to.Ptr("Delete a Digital Twins Service instance."),
		// 				Provider: to.Ptr("Microsoft.DigitalTwins"),
		// 				Resource: to.Ptr("digitalTwinsInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
