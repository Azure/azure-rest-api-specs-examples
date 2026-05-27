package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateCache.json
func ExampleCacheClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCacheClient().CreateOrUpdate(ctx, "rg1", "apimService1", "c1", armapimanagement.CacheContract{
		Properties: &armapimanagement.CacheContractProperties{
			Description:      to.Ptr("Redis cache instances in West India"),
			ConnectionString: to.Ptr("apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False"),
			ResourceID:       to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1"),
			UseFromLocation:  to.Ptr("default"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.CacheClientCreateOrUpdateResponse{
	// 	CacheContract: armapimanagement.CacheContract{
	// 		Name: to.Ptr("c1"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/caches"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/caches/c1"),
	// 		Properties: &armapimanagement.CacheContractProperties{
	// 			Description: to.Ptr("Redis cache instances in West India"),
	// 			ConnectionString: to.Ptr("{{5f7fbca77a891a2200f3db38}}"),
	// 			ResourceID: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1"),
	// 			UseFromLocation: to.Ptr("default"),
	// 		},
	// 	},
	// }
}
