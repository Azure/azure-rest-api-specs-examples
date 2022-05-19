Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv1.0.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_CreateOrUpdate.json
func ExampleIntegrationAccountBatchConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationAccountBatchConfigurationsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testResourceGroup",
		"testIntegrationAccount",
		"testBatchConfiguration",
		armlogic.BatchConfiguration{
			Location: to.Ptr("westus"),
			Properties: &armlogic.BatchConfigurationProperties{
				BatchGroupName: to.Ptr("DEFAULT"),
				ReleaseCriteria: &armlogic.BatchReleaseCriteria{
					BatchSize:    to.Ptr[int32](234567),
					MessageCount: to.Ptr[int32](10),
					Recurrence: &armlogic.WorkflowTriggerRecurrence{
						Frequency: to.Ptr(armlogic.RecurrenceFrequencyMinute),
						Interval:  to.Ptr[int32](1),
						StartTime: to.Ptr("2017-03-24T11:43:00"),
						TimeZone:  to.Ptr("India Standard Time"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
