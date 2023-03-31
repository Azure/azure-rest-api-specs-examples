package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
func ExampleBudgetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBudgetsClient().Delete(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "TestBudget", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
