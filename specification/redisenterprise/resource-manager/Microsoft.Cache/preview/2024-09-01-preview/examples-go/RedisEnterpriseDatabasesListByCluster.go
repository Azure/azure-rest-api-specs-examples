package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a287afb3721dee0d88f11502ec123470bc52a28/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseDatabasesListByCluster.json
func ExampleDatabasesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByClusterPager("rg1", "cache1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatabaseList = armredisenterprise.DatabaseList{
		// 	Value: []*armredisenterprise.Database{
		// 		{
		// 			Name: to.Ptr("cache1/default"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases"),
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
		// 			Properties: &armredisenterprise.DatabaseProperties{
		// 				ClientProtocol: to.Ptr(armredisenterprise.ProtocolEncrypted),
		// 				ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyOSSCluster),
		// 				DeferUpgrade: to.Ptr(armredisenterprise.DeferUpgradeSettingNotDeferred),
		// 				EvictionPolicy: to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
		// 				Modules: []*armredisenterprise.Module{
		// 					{
		// 						Name: to.Ptr("RediSearch"),
		// 						Args: to.Ptr(""),
		// 						Version: to.Ptr("1.0.0"),
		// 				}},
		// 				Persistence: &armredisenterprise.Persistence{
		// 					RdbEnabled: to.Ptr(true),
		// 					RdbFrequency: to.Ptr(armredisenterprise.RdbFrequencyTwelveH),
		// 				},
		// 				Port: to.Ptr[int32](10000),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				RedisVersion: to.Ptr("6.0"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 			},
		// 	}},
		// }
	}
}
