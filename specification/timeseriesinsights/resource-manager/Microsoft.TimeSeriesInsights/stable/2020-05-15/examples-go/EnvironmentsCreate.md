```go
package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EnvironmentsCreate.json
func ExampleEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtimeseriesinsights.NewEnvironmentsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"env1",
		&armtimeseriesinsights.Gen1EnvironmentCreateOrUpdateParameters{
			Location: to.Ptr("West US"),
			Kind:     to.Ptr(armtimeseriesinsights.EnvironmentKindGen1),
			SKU: &armtimeseriesinsights.SKU{
				Name:     to.Ptr(armtimeseriesinsights.SKUNameS1),
				Capacity: to.Ptr[int32](1),
			},
			Properties: &armtimeseriesinsights.Gen1EnvironmentCreationProperties{
				DataRetentionTime: to.Ptr("P31D"),
				PartitionKeyProperties: []*armtimeseriesinsights.TimeSeriesIDProperty{
					{
						Name: to.Ptr("DeviceId1"),
						Type: to.Ptr(armtimeseriesinsights.PropertyTypeString),
					}},
			},
		},
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftimeseriesinsights%2Farmtimeseriesinsights%2Fv1.0.0/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights/README.md) on how to add the SDK to your project and authenticate.
