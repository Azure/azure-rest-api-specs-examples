package armlargeinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstanceOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armlargeinstance.OperationListResult{
		// 	Value: []*armlargeinstance.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureLargeInstance/read"),
		// 			Display: &armlargeinstance.OperationDisplay{
		// 				Description: to.Ptr("Read any AzureLargeInstance"),
		// 				Operation: to.Ptr("Read AzureLargeInstance"),
		// 				Provider: to.Ptr("Microsoft Azure Large Instance"),
		// 				Resource: to.Ptr("AzureLargeInstance"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureLargeInstance"),
		// 			Display: &armlargeinstance.OperationDisplay{
		// 				Description: to.Ptr("Start any AzureLargeInstance"),
		// 				Operation: to.Ptr("Start AzureLargeInstance"),
		// 				Provider: to.Ptr("Microsoft Azure Large Instance"),
		// 				Resource: to.Ptr("AzureLargeInstance"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 	}},
		// }
	}
}
