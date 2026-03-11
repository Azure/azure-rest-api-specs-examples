package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
)

// Generated from example definition: 2025-08-01/ReplicationNetworkMappings_Delete.json
func ExampleReplicationNetworkMappingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("9112a37f-0f3e-46ec-9c00-060c6edca071", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationNetworkMappingsClient().BeginDelete(ctx, "srcBvte2a14C27", "srce2avaultbvtaC27", "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", "e2267b5c-2650-49bd-ab3f-d66aae694c06", "corpe2amap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
