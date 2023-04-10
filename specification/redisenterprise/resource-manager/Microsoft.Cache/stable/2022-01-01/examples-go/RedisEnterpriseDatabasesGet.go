package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesGet.json
func ExampleDatabasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().Get(ctx, "rg1", "cache1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Database = armredisenterprise.Database{
	// 	Name: to.Ptr("cache1/default"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
	// 	Properties: &armredisenterprise.DatabaseProperties{
	// 		ClientProtocol: to.Ptr(armredisenterprise.ProtocolEncrypted),
	// 		ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyOSSCluster),
	// 		EvictionPolicy: to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
	// 		Modules: []*armredisenterprise.Module{
	// 			{
	// 				Name: to.Ptr("RediSearch"),
	// 				Args: to.Ptr(""),
	// 				Version: to.Ptr("1.0.0"),
	// 		}},
	// 		Persistence: &armredisenterprise.Persistence{
	// 			RdbEnabled: to.Ptr(true),
	// 			RdbFrequency: to.Ptr(armredisenterprise.RdbFrequencyTwelveH),
	// 		},
	// 		Port: to.Ptr[int32](10000),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 	},
	// }
}
