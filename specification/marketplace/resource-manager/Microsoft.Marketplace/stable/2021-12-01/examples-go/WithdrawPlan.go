package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/WithdrawPlan.json
func ExamplePrivateStoreClient_WithdrawPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateStoreClient().WithdrawPlan(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "marketplacetestthirdparty.md-test-third-party-2", &armmarketplace.PrivateStoreClientWithdrawPlanOptions{Payload: &armmarketplace.WithdrawProperties{
		Properties: &armmarketplace.WithdrawDetails{
			PlanID:      to.Ptr("*"),
			PublisherID: to.Ptr("marketplacetestthirdparty"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
