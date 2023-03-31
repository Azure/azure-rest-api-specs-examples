package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1da7cbab8d4f554484dedb676ba7bdbdf6cdf78/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_ListByResourceGroup_MinimumSet_Gen.json
func ExampleElasticSansClient_NewListByResourceGroupPager_elasticSansListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticSansClient().NewListByResourceGroupPager("rgelasticsan", nil)
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
		// page.List = armelasticsan.List{
		// 	Value: []*armelasticsan.ElasticSan{
		// 		{
		// 			Properties: &armelasticsan.Properties{
		// 				BaseSizeTiB: to.Ptr[int64](26),
		// 				ExtendedCapacitySizeTiB: to.Ptr[int64](7),
		// 				SKU: &armelasticsan.SKU{
		// 					Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
