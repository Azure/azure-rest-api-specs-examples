package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/ElasticSans_Update_MinimumSet_Gen.json
func ExampleElasticSansClient_BeginUpdate_elasticSansUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticSansClient().BeginUpdate(ctx, "resourcegroupname", "elasticsanname", armelasticsan.Update{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ElasticSan = armelasticsan.ElasticSan{
	// 	Location: to.Ptr("eupkimcjwiphitotibv"),
	// 	Properties: &armelasticsan.Properties{
	// 		BaseSizeTiB: to.Ptr[int64](1),
	// 		ExtendedCapacitySizeTiB: to.Ptr[int64](3),
	// 		SKU: &armelasticsan.SKU{
	// 			Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
	// 		},
	// 	},
	// }
}
