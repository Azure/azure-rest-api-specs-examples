package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armcommvaultcontentstore.OperationsClientListResponse{
		// 	OperationListResult: armcommvaultcontentstore.OperationListResult{
		// 		Value: []*armcommvaultcontentstore.Operation{
		// 			{
		// 				Name: to.Ptr("ilkmpavtyvddhqbv"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armcommvaultcontentstore.OperationDisplay{
		// 					Provider: to.Ptr("uzdcxxutyrgzuejnqscepoq"),
		// 					Resource: to.Ptr("gwcncdch"),
		// 					Operation: to.Ptr("fxbcyqkxxsqhrojjktogfo"),
		// 					Description: to.Ptr("tlbmxwddalehnkohdqc"),
		// 				},
		// 				Origin: to.Ptr(armcommvaultcontentstore.OriginUser),
		// 				ActionType: to.Ptr(armcommvaultcontentstore.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
