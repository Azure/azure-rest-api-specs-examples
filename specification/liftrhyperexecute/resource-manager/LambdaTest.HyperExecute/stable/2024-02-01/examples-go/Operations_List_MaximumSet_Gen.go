package armlambdatesthyperexecute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/lambdatesthyperexecute/armlambdatesthyperexecute"
)

// Generated from example definition: 2024-02-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGenGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlambdatesthyperexecute.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armlambdatesthyperexecute.OperationsClientListResponse{
		// 	OperationListResult: armlambdatesthyperexecute.OperationListResult{
		// 		Value: []*armlambdatesthyperexecute.Operation{
		// 			{
		// 				Name: to.Ptr("pqjblljnopojofemsawnpfetwqdakn"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armlambdatesthyperexecute.OperationDisplay{
		// 					Provider: to.Ptr("pyqznxwjmapxcfr"),
		// 					Resource: to.Ptr("xlrbrkdnznqfgdzwbubj"),
		// 					Operation: to.Ptr("fujfkctgncoiygfunwfczhvf"),
		// 					Description: to.Ptr("fyqutlzqriswplfce"),
		// 				},
		// 				Origin: to.Ptr(armlambdatesthyperexecute.OriginUser),
		// 				ActionType: to.Ptr(armlambdatesthyperexecute.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
