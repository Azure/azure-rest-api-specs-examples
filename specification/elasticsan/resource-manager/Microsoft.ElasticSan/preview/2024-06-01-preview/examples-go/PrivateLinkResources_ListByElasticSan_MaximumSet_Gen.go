package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/PrivateLinkResources_ListByElasticSan_MaximumSet_Gen.json
func ExamplePrivateLinkResourcesClient_ListByElasticSan_privateLinkResourcesListByElasticSanMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.PrivateLinkResourceListResult = armelasticsan.PrivateLinkResourceListResult{
	// 	Value: []*armelasticsan.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("ggbyptukibs"),
	// 			Type: to.Ptr("qrfsowhtanlj"),
	// 			ID: to.Ptr("iidc"),
	// 			SystemData: &armelasticsan.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
	// 				CreatedBy: to.Ptr("bgurjvijz"),
	// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("uvexylihjrtinzkeluohusnaxatfqh"),
	// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			},
	// 			Properties: &armelasticsan.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("mbouakfumvbeqnevmgxpk"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("wujfilzifgumbvxbdhazmzf")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("hzxhgoqxxiaf")},
	// 					},
	// 			}},
	// 		}
}
