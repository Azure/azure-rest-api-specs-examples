package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterpriseDatabasesCreate.json
func ExampleDatabasesClient_BeginCreate_redisEnterpriseDatabasesCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreate(ctx, "rg1", "cache1", "default", armredisenterprise.Database{
		Properties: &armredisenterprise.DatabaseCreateProperties{
			AccessKeysAuthentication: to.Ptr(armredisenterprise.AccessKeysAuthenticationEnabled),
			ClientProtocol:           to.Ptr(armredisenterprise.ProtocolEncrypted),
			ClusteringPolicy:         to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
			DeferUpgrade:             to.Ptr(armredisenterprise.DeferUpgradeSettingNotDeferred),
			EvictionPolicy:           to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
			Modules: []*armredisenterprise.Module{
				{
					Name: to.Ptr("RedisBloom"),
					Args: to.Ptr("ERROR_RATE 0.00 INITIAL_SIZE 400"),
				},
				{
					Name: to.Ptr("RedisTimeSeries"),
					Args: to.Ptr("RETENTION_POLICY 20"),
				},
				{
					Name: to.Ptr("RediSearch"),
				},
			},
			Persistence: &armredisenterprise.Persistence{
				AofEnabled:   to.Ptr(true),
				AofFrequency: to.Ptr(armredisenterprise.AofFrequencyOneS),
			},
			Port: to.Ptr[int32](10000),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredisenterprise.DatabasesClientCreateResponse{
	// 	Database: &armredisenterprise.Database{
	// 		Name: to.Ptr("cache1/default"),
	// 		Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases"),
	// 		ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
	// 		Properties: &armredisenterprise.DatabaseCreateProperties{
	// 			AccessKeysAuthentication: to.Ptr(armredisenterprise.AccessKeysAuthenticationEnabled),
	// 			ClientProtocol: to.Ptr(armredisenterprise.ProtocolEncrypted),
	// 			ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
	// 			DeferUpgrade: to.Ptr(armredisenterprise.DeferUpgradeSettingNotDeferred),
	// 			EvictionPolicy: to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
	// 			Modules: []*armredisenterprise.Module{
	// 				{
	// 					Name: to.Ptr("RedisBloom"),
	// 					Args: to.Ptr("ERROR_RATE 0.00 INITIAL_SIZE 400"),
	// 					Version: to.Ptr("1.0.0"),
	// 				},
	// 				{
	// 					Name: to.Ptr("RedisTimeSeries"),
	// 					Args: to.Ptr("RETENTION_POLICY 20"),
	// 					Version: to.Ptr("1.0.0"),
	// 				},
	// 				{
	// 					Name: to.Ptr("RediSearch"),
	// 					Args: to.Ptr(""),
	// 					Version: to.Ptr("1.0.0"),
	// 				},
	// 			},
	// 			Persistence: &armredisenterprise.Persistence{
	// 				AofEnabled: to.Ptr(true),
	// 				AofFrequency: to.Ptr(armredisenterprise.AofFrequencyOneS),
	// 			},
	// 			Port: to.Ptr[int32](10000),
	// 			ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateUpdating),
	// 			RedisVersion: to.Ptr("6.0"),
	// 			ResourceState: to.Ptr(armredisenterprise.ResourceStateUpdating),
	// 		},
	// 	},
	// }
}
