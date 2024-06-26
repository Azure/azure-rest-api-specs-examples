package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/proxy/ExposingAvailableOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationPage = armmixedreality.OperationPage{
		// 	Value: []*armmixedreality.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.MixedReality/register/action"),
		// 			Display: &armmixedreality.OperationDisplay{
		// 				Description: to.Ptr("Registers a subscription for the Mixed Reality resource provider."),
		// 				Operation: to.Ptr("Registers the Mixed Reality resource provider"),
		// 				Provider: to.Ptr("Microsoft.MixedReality"),
		// 				Resource: to.Ptr("Mixed Reality resource provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.MixedReality/SpatialAnchorsAccounts/delete"),
		// 			Display: &armmixedreality.OperationDisplay{
		// 				Description: to.Ptr("Deletes the resource for Microsoft.MixedReality/SpatialAnchorsAccounts"),
		// 				Operation: to.Ptr("Delete Spatial Anchors Accounts"),
		// 				Provider: to.Ptr("Microsoft.MixedReality"),
		// 				Resource: to.Ptr("SpatialAnchorsAccounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.MixedReality/SpatialAnchorsAccounts/read"),
		// 			Display: &armmixedreality.OperationDisplay{
		// 				Description: to.Ptr("Gets the resource for Microsoft.MixedReality/SpatialAnchorsAccounts"),
		// 				Operation: to.Ptr("Get Spatial Anchors Accounts"),
		// 				Provider: to.Ptr("Microsoft.MixedReality"),
		// 				Resource: to.Ptr("SpatialAnchorsAccounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.MixedReality/SpatialAnchorsAccounts/write"),
		// 			Display: &armmixedreality.OperationDisplay{
		// 				Description: to.Ptr("Updates the resource for Microsoft.MixedReality/SpatialAnchorsAccounts"),
		// 				Operation: to.Ptr("Update Spatial Anchors Accounts"),
		// 				Provider: to.Ptr("Microsoft.MixedReality"),
		// 				Resource: to.Ptr("SpatialAnchorsAccounts"),
		// 			},
		// 	}},
		// }
	}
}
