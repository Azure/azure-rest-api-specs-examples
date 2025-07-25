package armagricultureplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agricultureplatform/armagricultureplatform"
)

// Generated from example definition: 2024-06-01-preview/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagricultureplatform.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armagricultureplatform.OperationsClientListResponse{
		// 	OperationListResult: armagricultureplatform.OperationListResult{
		// 		Value: []*armagricultureplatform.Operation{
		// 			{
		// 				Name: to.Ptr("on"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armagricultureplatform.OperationDisplay{
		// 					Provider: to.Ptr("yvkmdseicusbpedotoit"),
		// 					Resource: to.Ptr("gzqgh"),
		// 					Operation: to.Ptr("wdvqfctmtffessbckulhoswzv"),
		// 					Description: to.Ptr("xuvqlrnkwqsafmydwvywnpdxlsmnn"),
		// 				},
		// 				Origin: to.Ptr(armagricultureplatform.OriginUser),
		// 				ActionType: to.Ptr(armagricultureplatform.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
