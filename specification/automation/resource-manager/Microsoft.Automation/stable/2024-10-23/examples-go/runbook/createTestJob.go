package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/runbook/createTestJob.json
func ExampleTestJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTestJobClient().Create(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", armautomation.TestJobCreateParameters{
		Parameters: map[string]*string{
			"key01": to.Ptr("value01"),
			"key02": to.Ptr("value02"),
		},
		RunOn:              to.Ptr(""),
		RuntimeEnvironment: to.Ptr("runtimeEnvironmentName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
