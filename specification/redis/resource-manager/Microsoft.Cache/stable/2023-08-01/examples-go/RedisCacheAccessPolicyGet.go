package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/20312e2b31df58f0ea7560e87062d62aa92f0a14/specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheAccessPolicyGet.json
func ExampleAccessPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPolicyClient().Get(ctx, "rg1", "cache1", "accessPolicy1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CacheAccessPolicy = armredis.CacheAccessPolicy{
	// 	Name: to.Ptr("accessPolicy1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/accessPolicies/accessPolicy1"),
	// 	Properties: &armredis.CacheAccessPolicyProperties{
	// 		Type: to.Ptr(armredis.AccessPolicyTypeCustom),
	// 		Permissions: to.Ptr("+get +hget"),
	// 		ProvisioningState: to.Ptr(armredis.AccessPolicyProvisioningStateSucceeded),
	// 	},
	// }
}
