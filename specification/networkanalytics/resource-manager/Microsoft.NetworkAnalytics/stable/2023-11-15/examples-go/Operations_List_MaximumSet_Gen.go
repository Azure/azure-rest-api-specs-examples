package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armnetworkanalytics.OperationListResult{
		// 	Value: []*armnetworkanalytics.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.NetworkAnalytics/resourceName/operationName"),
		// 			ActionType: to.Ptr(armnetworkanalytics.ActionTypeInternal),
		// 			Display: &armnetworkanalytics.OperationDisplay{
		// 				Description: to.Ptr("Description of the operation"),
		// 				Operation: to.Ptr("OperationName"),
		// 				Provider: to.Ptr("Microsoft.NetworkAnalytics"),
		// 				Resource: to.Ptr("ResourceName"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armnetworkanalytics.OriginUser),
		// 	}},
		// }
	}
}
