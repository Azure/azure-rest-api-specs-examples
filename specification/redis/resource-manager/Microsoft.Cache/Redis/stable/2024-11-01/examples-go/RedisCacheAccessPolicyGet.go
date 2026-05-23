package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheAccessPolicyGet.json
func ExampleAccessPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
	// res = armredis.AccessPolicyClientGetResponse{
	// 	CacheAccessPolicy: &armredis.CacheAccessPolicy{
	// 		Name: to.Ptr("accessPolicy1"),
	// 		Type: to.Ptr("Microsoft.Cache/Redis/accessPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/accessPolicies/accessPolicy1"),
	// 		Properties: &armredis.CacheAccessPolicyProperties{
	// 			Type: to.Ptr(armredis.AccessPolicyTypeCustom),
	// 			Permissions: to.Ptr("+get +hget"),
	// 			ProvisioningState: to.Ptr(armredis.AccessPolicyProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
