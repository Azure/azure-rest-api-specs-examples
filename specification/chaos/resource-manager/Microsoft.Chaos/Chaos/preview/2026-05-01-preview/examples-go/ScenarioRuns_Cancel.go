package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/ScenarioRuns_Cancel.json
func ExampleScenarioRunsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScenarioRunsClient().Cancel(ctx, "exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012", "abcd1234-5678-9012-3456-789012345678", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
