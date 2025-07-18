package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armhybridcompute.OperationListResult{
		// 	Value: []*armhybridcompute.OperationValue{
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridCompute/operations/read"),
		// 			Display: &armhybridcompute.OperationValueDisplay{
		// 				Description: to.Ptr("Read all Operations for Azure Arc for Servers"),
		// 				Operation: to.Ptr("Read all Operations"),
		// 				Provider: to.Ptr("Microsoft HybridCompute"),
		// 				Resource: to.Ptr("Microsoft.HybridCompute Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.HybridCompute/register/action"),
		// 			Display: &armhybridcompute.OperationValueDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Microsoft.HybridCompute Resource Provider"),
		// 				Operation: to.Ptr("Register Subscription for Azure Arc for Servers"),
		// 				Provider: to.Ptr("Microsoft HybridCompute"),
		// 				Resource: to.Ptr("Microsoft.HybridCompute Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
