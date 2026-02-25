package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/Skus_List_MaximumSet_Gen.json
func ExampleSKUsClient_NewListPager_skusListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(&armelasticsan.SKUsClientListOptions{
		Filter: to.Ptr("obwwdrkq")})
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
		// page = armelasticsan.SKUsClientListResponse{
		// 	SKUInformationList: armelasticsan.SKUInformationList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionid/providers/Microsoft.ElasticSan/skus?api-version=2024-07-01-preview&%24skiptoken=def456uvw890"),
		// 		Value: []*armelasticsan.SKUInformation{
		// 			{
		// 				Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 				Capabilities: []*armelasticsan.SKUCapability{
		// 					{
		// 						Name: to.Ptr("qkht"),
		// 						Value: to.Ptr("eoayvlyzyjjziecxymlpk"),
		// 					},
		// 				},
		// 				LocationInfo: []*armelasticsan.SKULocationInfo{
		// 					{
		// 						Location: to.Ptr("ngycrsoihxdfctigejlf"),
		// 						Zones: []*string{
		// 							to.Ptr("1"),
		// 						},
		// 					},
		// 				},
		// 				Locations: []*string{
		// 					to.Ptr("ceszpxwqyifrxobqykplm"),
		// 				},
		// 				ResourceType: to.Ptr("tlqickysdtjahoanstgancifxfu"),
		// 				Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 			},
		// 		},
		// 	},
		// }
	}
}
