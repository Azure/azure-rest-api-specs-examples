package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/27046dbff974e3901970aa53b29cec6d8ec1342a/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armelasticsan.OperationListResult{
		// 	Value: []*armelasticsan.Operation{
		// 		{
		// 			Name: to.Ptr("zgtqmglizaqlsx"),
		// 			ActionType: to.Ptr(armelasticsan.ActionTypeInternal),
		// 			Display: &armelasticsan.OperationDisplay{
		// 				Description: to.Ptr("pmkjqzjverubmslnrcadqur"),
		// 				Operation: to.Ptr("yumtqbnawcvunwda"),
		// 				Provider: to.Ptr("apbqaoiegbmipkbqdczsuvlak"),
		// 				Resource: to.Ptr("hjyuuxkj"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armelasticsan.OriginUser),
		// 	}},
		// }
	}
}
