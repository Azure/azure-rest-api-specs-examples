package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ListOperationsPartner.json
func ExampleAPISClient_NewListOperationsPartnerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorderpartner.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPISClient().NewListOperationsPartnerPager(nil)
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
		// page.OperationListResult = armedgeorderpartner.OperationListResult{
		// 	Value: []*armedgeorderpartner.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrderPartner/operations/read"),
		// 			Display: &armedgeorderpartner.OperationDisplay{
		// 				Description: to.Ptr("List or get the Operations"),
		// 				Operation: to.Ptr("List or Get Operations"),
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armedgeorderpartner.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrderPartner/searchInventories/action"),
		// 			Display: &armedgeorderpartner.OperationDisplay{
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_Microsoft.EdgeOrderPartner"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armedgeorderpartner.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrderPartner/locations/productFamilies/inventories/manageLink/action"),
		// 			Display: &armedgeorderpartner.OperationDisplay{
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_inventories"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armedgeorderpartner.OriginUser),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EdgeOrderPartner/locations/productFamilies/inventories/manageInventoryMetadata/action"),
		// 			Display: &armedgeorderpartner.OperationDisplay{
		// 				Provider: to.Ptr("Edge Ordering"),
		// 				Resource: to.Ptr("ArmApiRes_inventories"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armedgeorderpartner.OriginUser),
		// 	}},
		// }
	}
}
