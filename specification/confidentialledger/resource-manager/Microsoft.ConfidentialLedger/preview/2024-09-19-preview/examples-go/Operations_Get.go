package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/40c12a2dd76da721d480d55a44e6ec666045d508/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/Operations_Get.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfidentialledger.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ResourceProviderOperationList = armconfidentialledger.ResourceProviderOperationList{
		// 	Value: []*armconfidentialledger.ResourceProviderOperationDefinition{
		// 		{
		// 			Name: to.Ptr("Microsoft.ConfidentialLedger/ledgers/read"),
		// 			Display: &armconfidentialledger.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("View the properties of a Confidential Ledger"),
		// 				Operation: to.Ptr("Get Confidential Ledger properties"),
		// 				Resource: to.Ptr("ledger"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
