package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateCache.json
func ExampleCacheClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewCacheClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"c1",
		armapimanagement.CacheContract{
			Properties: &armapimanagement.CacheContractProperties{
				Description:      to.Ptr("Redis cache instances in West India"),
				ConnectionString: to.Ptr("apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False"),
				ResourceID:       to.Ptr("https://management.azure.com/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1"),
				UseFromLocation:  to.Ptr("default"),
			},
		},
		&armapimanagement.CacheClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
