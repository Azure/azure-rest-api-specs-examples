package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/Operations_List_MinimumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
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
		// 	},
		// }
	}
}
