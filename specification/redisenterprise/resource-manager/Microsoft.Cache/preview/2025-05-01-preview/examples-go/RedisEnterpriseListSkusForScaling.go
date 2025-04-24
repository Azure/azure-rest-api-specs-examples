package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad9b489baef1d982f7641f6c47a00794c9a1a5be/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseListSkusForScaling.json
func ExampleClient_ListSKUsForScaling() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListSKUsForScaling(ctx, "rg1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUDetailsList = armredisenterprise.SKUDetailsList{
	// 	SKUs: []*armredisenterprise.SKUDetails{
	// 		{
	// 			Name: to.Ptr("MemoryOptimized_M100"),
	// 			SizeInGB: to.Ptr[float32](120),
	// 		},
	// 		{
	// 			Name: to.Ptr("ComputeOptimized_X700"),
	// 			SizeInGB: to.Ptr[float32](720),
	// 		},
	// 		{
	// 			Name: to.Ptr("Balanced_B5"),
	// 			SizeInGB: to.Ptr[float32](6),
	// 	}},
	// }
}
