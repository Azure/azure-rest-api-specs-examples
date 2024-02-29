package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseDatabasesCreateWithGeoReplication.json
func ExampleDatabasesClient_BeginCreate_redisEnterpriseDatabasesCreateWithActiveGeoReplication() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreate(ctx, "rg1", "cache1", "default", armredisenterprise.Database{
		Properties: &armredisenterprise.DatabaseProperties{
			ClientProtocol:   to.Ptr(armredisenterprise.ProtocolEncrypted),
			ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
			EvictionPolicy:   to.Ptr(armredisenterprise.EvictionPolicyNoEviction),
			GeoReplication: &armredisenterprise.DatabasePropertiesGeoReplication{
				GroupNickname: to.Ptr("groupName"),
				LinkedDatabases: []*armredisenterprise.LinkedDatabase{
					{
						ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
					},
					{
						ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
					}},
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
	// res.Database = armredisenterprise.Database{
	// 	Name: to.Ptr("cache1/default"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise/databases"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
	// 	Properties: &armredisenterprise.DatabaseProperties{
	// 		ClientProtocol: to.Ptr(armredisenterprise.ProtocolEncrypted),
	// 		ClusteringPolicy: to.Ptr(armredisenterprise.ClusteringPolicyEnterpriseCluster),
	// 		EvictionPolicy: to.Ptr(armredisenterprise.EvictionPolicyNoEviction),
	// 		GeoReplication: &armredisenterprise.DatabasePropertiesGeoReplication{
	// 			GroupNickname: to.Ptr("groupName"),
	// 			LinkedDatabases: []*armredisenterprise.LinkedDatabase{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
	// 					State: to.Ptr(armredisenterprise.LinkStateLinking),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
	// 					State: to.Ptr(armredisenterprise.LinkStateLinking),
	// 			}},
	// 		},
	// 		Port: to.Ptr[int32](10000),
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateUpdating),
	// 	},
	// }
}
