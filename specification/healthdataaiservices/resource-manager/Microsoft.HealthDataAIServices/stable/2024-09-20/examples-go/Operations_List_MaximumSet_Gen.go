package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: 2024-09-20/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGeneratedByMaximumSetRuleStable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armhealthdataaiservices.OperationsClientListResponse{
		// 	OperationListResult: armhealthdataaiservices.OperationListResult{
		// 		Value: []*armhealthdataaiservices.Operation{
		// 			{
		// 				Name: to.Ptr("vlozcqymdxttexvmhouwzob"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armhealthdataaiservices.OperationDisplay{
		// 					Provider: to.Ptr("hwy"),
		// 					Resource: to.Ptr("yxabgnzjshmqldqthxonpam"),
		// 					Operation: to.Ptr("quwaawjasjgpqhskxoxzx"),
		// 					Description: to.Ptr("ayqiodducsbwvzcgno"),
		// 				},
		// 				Origin: to.Ptr(armhealthdataaiservices.OriginUser),
		// 				ActionType: to.Ptr(armhealthdataaiservices.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
