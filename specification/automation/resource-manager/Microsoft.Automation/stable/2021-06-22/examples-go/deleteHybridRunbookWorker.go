package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteHybridRunbookWorker.json
func ExampleHybridRunbookWorkersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewHybridRunbookWorkersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"rg",
		"myAutomationAccount20",
		"myGroup",
		"c010ad12-ef14-4a2a-aa9e-ef22c4745ddd",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
