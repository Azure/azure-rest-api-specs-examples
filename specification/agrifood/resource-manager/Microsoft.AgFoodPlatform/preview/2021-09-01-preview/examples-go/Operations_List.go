package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a7af6049f4b4743ef3b649f3852bcc7bd9a43ee0/specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagrifood.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armagrifood.OperationListResult{
		// 	Value: []*armagrifood.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/read"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Gets or Lists existing AgFoodPlatform FarmBeats resource(s)."),
		// 				Operation: to.Ptr("Get or List AgFoodPlatform FarmBeats resource(s)."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform FarmBeats"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/write"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates AgFoodPlatform FarmBeats."),
		// 				Operation: to.Ptr("Create or Update AgFoodPlatform FarmBeats."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform FarmBeats"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/delete"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing AgFoodPlatform FarmBeats resource."),
		// 				Operation: to.Ptr("Delete AgFoodPlatform FarmBeats resource."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform FarmBeats"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/locations/checkNameAvailability/action"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Checks that resource name is valid and is not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/operations/read"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("List all operations in Microsoft AgFoodPlatform resource provider."),
		// 				Operation: to.Ptr("List all operations."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("List all operations in Microsoft AgFoodPlatform resource provider."),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/extensions/read"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Gets or Lists existing AgFoodPlatform Extensions resource(s)."),
		// 				Operation: to.Ptr("Get or List AgFoodPlatform Extensions resource(s)."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform Extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/extensions/write"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates AgFoodPlatform Extensions."),
		// 				Operation: to.Ptr("Create or Update AgFoodPlatform Extensions."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform Extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/extensions/delete"),
		// 			Display: &armagrifood.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing AgFoodPlatform Extensions resource."),
		// 				Operation: to.Ptr("Delete AgFoodPlatform Extensions resource."),
		// 				Provider: to.Ptr("Microsoft AgFoodPlatform"),
		// 				Resource: to.Ptr("AgFoodPlatform Extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
