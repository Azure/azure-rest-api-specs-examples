package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/createJob.json
func ExampleJobClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewJobClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx,
		"mygroup",
		"ContoseAutomationAccount",
		"foo",
		armautomation.JobCreateParameters{
			Properties: &armautomation.JobCreateProperties{
				Parameters: map[string]*string{
					"key01": to.Ptr("value01"),
					"key02": to.Ptr("value02"),
				},
				RunOn: to.Ptr(""),
				Runbook: &armautomation.RunbookAssociationProperty{
					Name: to.Ptr("TestRunbook"),
				},
			},
		},
		&armautomation.JobClientCreateOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
