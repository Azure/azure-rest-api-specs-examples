package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad9b489baef1d982f7641f6c47a00794c9a1a5be/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseDatabasesUpdate.json
func ExampleDatabasesClient_BeginUpdate_redisEnterpriseDatabasesUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginUpdate(ctx, "rg1", "cache1", "default", armredisenterprise.DatabaseUpdate{
		Properties: &armredisenterprise.DatabaseProperties{
			AccessKeysAuthentication: to.Ptr(armredisenterprise.AccessKeysAuthenticationEnabled),
			ClientProtocol:           to.Ptr(armredisenterprise.ProtocolEncrypted),
			EvictionPolicy:           to.Ptr(armredisenterprise.EvictionPolicyAllKeysLRU),
			Persistence: &armredisenterprise.Persistence{
				RdbEnabled:   to.Ptr(true),
				RdbFrequency: to.Ptr(armredisenterprise.RdbFrequencyTwelveH),
			},
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
	// res.Database = armredisenterprise.Database{
	// 	Name: to.Ptr("cache1/default"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases"),
	// 	ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
	// 	Properties: &armredisenterprise.DatabaseProperties{
	// 		AccessKeysAuthentication: to.Ptr(armredisenterprise.AccessKeysAuthenticationEnabled),
	// 		ClientProtocol: to.Ptr(armredisenterprise.ProtocolEncrypted),
	// 		ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyOSSCluster),
	// 		DeferUpgrade: to.Ptr(armredisenterprise.DeferUpgradeSettingNotDeferred),
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
	// 		RedisVersion: to.Ptr("6.0"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateUpdating),
	// 	},
	// }
}
