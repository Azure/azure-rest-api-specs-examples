package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheForceReboot.json
func ExampleClient_ForceReboot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredis.NewClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ForceReboot(ctx,
		"rg1",
		"cache1",
		armredis.RebootParameters{
			Ports: []*int32{
				to.Ptr[int32](13000),
				to.Ptr[int32](15001)},
			RebootType: to.Ptr(armredis.RebootTypeAllNodes),
			ShardID:    to.Ptr[int32](0),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
