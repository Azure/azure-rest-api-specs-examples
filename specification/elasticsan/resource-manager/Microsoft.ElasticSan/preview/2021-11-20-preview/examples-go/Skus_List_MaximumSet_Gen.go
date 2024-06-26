package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1da7cbab8d4f554484dedb676ba7bdbdf6cdf78/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Skus_List_MaximumSet_Gen.json
func ExampleSKUsClient_NewListPager_skusListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(&armelasticsan.SKUsClientListOptions{Filter: to.Ptr("aaaa")})
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
		// page.SKUInformationList = armelasticsan.SKUInformationList{
		// 	Value: []*armelasticsan.SKUInformation{
		// 		{
		// 			Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 			Capabilities: []*armelasticsan.SKUCapability{
		// 				{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					Value: to.Ptr("aaaaaaaa"),
		// 			}},
		// 			LocationInfo: []*armelasticsan.SKULocationInfo{
		// 				{
		// 					Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					Zones: []*string{
		// 						to.Ptr("aaa")},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("aaaaaaaaa")},
		// 					ResourceType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 			}},
		// 		}
	}
}
