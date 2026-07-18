package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBDatabaseAccountFailoverPriorityChange.json
func ExampleDatabaseAccountsClient_BeginFailoverPriorityChange() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseAccountsClient().BeginFailoverPriorityChange(ctx, "rg1", "ddb1-failover", armcosmos.FailoverPolicies{
		FailoverPolicies: []*armcosmos.FailoverPolicy{
			{
				LocationName:     to.Ptr("eastus"),
				FailoverPriority: to.Ptr[int32](0),
			},
			{
				LocationName:     to.Ptr("westus"),
				FailoverPriority: to.Ptr[int32](1),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
