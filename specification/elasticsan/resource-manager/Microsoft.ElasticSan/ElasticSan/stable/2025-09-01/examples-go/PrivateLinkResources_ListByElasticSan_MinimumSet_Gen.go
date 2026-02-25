package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/PrivateLinkResources_ListByElasticSan_MinimumSet_Gen.json
func ExamplePrivateLinkResourcesClient_ListByElasticSan_privateLinkResourcesListByElasticSanMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByElasticSan(ctx, "resourcegroupname", "elasticsanname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armelasticsan.PrivateLinkResourcesClientListByElasticSanResponse{
	// 	PrivateLinkResourceListResult: &armelasticsan.PrivateLinkResourceListResult{
	// 		Value: []*armelasticsan.PrivateLinkResource{
	// 		},
	// 	},
	// }
}
