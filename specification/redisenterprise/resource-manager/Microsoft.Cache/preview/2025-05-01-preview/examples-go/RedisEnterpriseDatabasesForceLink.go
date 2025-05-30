package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad9b489baef1d982f7641f6c47a00794c9a1a5be/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseDatabasesForceLink.json
func ExampleDatabasesClient_BeginForceLinkToReplicationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginForceLinkToReplicationGroup(ctx, "rg1", "cache1", "default", armredisenterprise.ForceLinkParameters{
		GeoReplication: &armredisenterprise.ForceLinkParametersGeoReplication{
			GroupNickname: to.Ptr("groupName"),
			LinkedDatabases: []*armredisenterprise.LinkedDatabase{
				{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
				},
				{
					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
