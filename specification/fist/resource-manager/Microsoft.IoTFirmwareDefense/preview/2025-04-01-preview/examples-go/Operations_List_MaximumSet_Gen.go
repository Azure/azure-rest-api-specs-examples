package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armiotfirmwaredefense.OperationsClientListResponse{
		// 	OperationListResult: armiotfirmwaredefense.OperationListResult{
		// 		Value: []*armiotfirmwaredefense.Operation{
		// 			{
		// 				Name: to.Ptr("wdbuw"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armiotfirmwaredefense.OperationDisplay{
		// 					Provider: to.Ptr("avfdrxtfudhdoyhlpwldgzq"),
		// 					Resource: to.Ptr("udukhbh"),
		// 					Operation: to.Ptr("ljpkuvkasdmwjxwibaeuhorzqy"),
		// 					Description: to.Ptr("hlsaojlacqrucpys"),
		// 				},
		// 				Origin: to.Ptr(armiotfirmwaredefense.OriginUser),
		// 				ActionType: to.Ptr(armiotfirmwaredefense.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/avs"),
		// 	},
		// }
	}
}
