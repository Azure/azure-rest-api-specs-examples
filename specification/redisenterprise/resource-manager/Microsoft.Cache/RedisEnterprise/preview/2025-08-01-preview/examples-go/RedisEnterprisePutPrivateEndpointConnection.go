package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterprisePutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginPut(ctx, "rg1", "cache1", "pectest01", armredisenterprise.PrivateEndpointConnection{
		Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
			},
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
