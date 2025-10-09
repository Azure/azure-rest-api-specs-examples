package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2025-11-01/RetentionPolicies_Delete_MaximumSet_Gen.json
func ExampleRetentionPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("194D3C1E-462F-4738-9025-092A628C06EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRetentionPoliciesClient().BeginDelete(ctx, "rgdurabletask", "testcheduler", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
