package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/listFeaturesOperations.json
func ExampleFeatureClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfeatures.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFeatureClient().NewListOperationsPager(nil)
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
		// page.OperationListResult = armfeatures.OperationListResult{
		// 	Value: []*armfeatures.Operation{
		// 		{
		// 			Name: to.Ptr("FeaturesOpeartion1"),
		// 			Display: &armfeatures.OperationDisplay{
		// 				Operation: to.Ptr("Read"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FeaturesOpeartion2"),
		// 			Display: &armfeatures.OperationDisplay{
		// 				Operation: to.Ptr("Write"),
		// 				Provider: to.Ptr("Microsoft.ResourceProvider"),
		// 				Resource: to.Ptr("Resource2"),
		// 			},
		// 	}},
		// }
	}
}
