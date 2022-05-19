Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.6.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createJobSchedule.json
func ExampleJobScheduleClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewJobScheduleClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx,
		"rg",
		"ContoseAutomationAccount",
		"0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc",
		armautomation.JobScheduleCreateParameters{
			Properties: &armautomation.JobScheduleCreateProperties{
				Parameters: map[string]*string{
					"jobscheduletag01": to.Ptr("jobschedulevalue01"),
					"jobscheduletag02": to.Ptr("jobschedulevalue02"),
				},
				Runbook: &armautomation.RunbookAssociationProperty{
					Name: to.Ptr("TestRunbook"),
				},
				Schedule: &armautomation.ScheduleAssociationProperty{
					Name: to.Ptr("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
