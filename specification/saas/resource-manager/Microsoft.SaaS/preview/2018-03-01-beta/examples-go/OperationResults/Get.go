package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/OperationResults/Get.json
func ExampleOperationClient_BeginGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsaas.NewOperationClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGet(ctx,
		"5f35cb4c-8065-45b3-9116-5ba335462e95",
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
