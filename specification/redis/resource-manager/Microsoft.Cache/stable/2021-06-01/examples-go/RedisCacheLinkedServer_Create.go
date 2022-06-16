package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheLinkedServer_Create.json
func ExampleLinkedServerClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredis.NewLinkedServerClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cache1",
		"cache2",
		armredis.LinkedServerCreateParameters{
			Properties: &armredis.LinkedServerCreateProperties{
				LinkedRedisCacheID:       to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache2"),
				LinkedRedisCacheLocation: to.Ptr("West US"),
				ServerRole:               to.Ptr(armredis.ReplicationRoleSecondary),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
