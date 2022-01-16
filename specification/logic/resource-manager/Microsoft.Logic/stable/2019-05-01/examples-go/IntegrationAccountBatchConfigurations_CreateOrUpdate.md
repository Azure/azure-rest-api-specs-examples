Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_CreateOrUpdate.json
func ExampleIntegrationAccountBatchConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountBatchConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<batch-configuration-name>",
		armlogic.BatchConfiguration{
			Location: to.StringPtr("<location>"),
			Properties: &armlogic.BatchConfigurationProperties{
				BatchGroupName: to.StringPtr("<batch-group-name>"),
				ReleaseCriteria: &armlogic.BatchReleaseCriteria{
					BatchSize:    to.Int32Ptr(234567),
					MessageCount: to.Int32Ptr(10),
					Recurrence: &armlogic.WorkflowTriggerRecurrence{
						Frequency: armlogic.RecurrenceFrequency("Minute").ToPtr(),
						Interval:  to.Int32Ptr(1),
						StartTime: to.StringPtr("<start-time>"),
						TimeZone:  to.StringPtr("<time-zone>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationAccountBatchConfigurationsClientCreateOrUpdateResult)
}
```
