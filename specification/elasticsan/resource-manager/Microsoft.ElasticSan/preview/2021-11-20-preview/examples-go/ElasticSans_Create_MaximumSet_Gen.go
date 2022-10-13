package armelasticsans_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_Create_MaximumSet_Gen.json
func ExampleClient_BeginCreate_elasticSansCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armelasticsans.NewClient("aaaaaaaaaaaaaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rgelasticsan", "ti7q-k952-1qB3J_5", armelasticsans.ElasticSan{
		Tags: map[string]*string{
			"key896": to.Ptr("aaaaaaaaaaaaaaaaaa"),
		},
		Location: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		Properties: &armelasticsans.ElasticSanProperties{
			AvailabilityZones: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
			BaseSizeTiB:             to.Ptr[int64](26),
			ExtendedCapacitySizeTiB: to.Ptr[int64](7),
			SKU: &armelasticsans.SKU{
				Name: to.Ptr(armelasticsans.SKUNamePremiumLRS),
				Tier: to.Ptr(armelasticsans.SKUTierPremium),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
