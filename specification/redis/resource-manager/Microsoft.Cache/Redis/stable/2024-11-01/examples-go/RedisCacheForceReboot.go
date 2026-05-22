package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheForceReboot.json
func ExampleClient_ForceReboot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ForceReboot(ctx, "rg1", "cache1", armredis.RebootParameters{
		Ports: []*int32{
			to.Ptr[int32](13000),
			to.Ptr[int32](15001),
		},
		RebootType: to.Ptr(armredis.RebootTypeAllNodes),
		ShardID:    to.Ptr[int32](0),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredis.ClientForceRebootResponse{
	// 	ForceRebootResponse: &armredis.ForceRebootResponse{
	// 		Message: to.Ptr("reboot operation enqueued"),
	// 	},
	// }
}
