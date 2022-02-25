Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
func ExampleCostsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewCostsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.LabCost{
			Properties: &armdevtestlabs.LabCostProperties{
				CurrencyCode:  to.StringPtr("<currency-code>"),
				EndDateTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T23:59:59Z"); return t }()),
				StartDateTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00Z"); return t }()),
				TargetCost: &armdevtestlabs.TargetCostProperties{
					CostThresholds: []*armdevtestlabs.CostThresholdProperties{
						{
							DisplayOnChart: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Float64Ptr(25),
							},
							SendNotificationWhenExceeded: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							ThresholdID:                  to.StringPtr("<threshold-id>"),
						},
						{
							DisplayOnChart: armdevtestlabs.CostThresholdStatus("Enabled").ToPtr(),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Float64Ptr(50),
							},
							SendNotificationWhenExceeded: armdevtestlabs.CostThresholdStatus("Enabled").ToPtr(),
							ThresholdID:                  to.StringPtr("<threshold-id>"),
						},
						{
							DisplayOnChart: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Float64Ptr(75),
							},
							SendNotificationWhenExceeded: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							ThresholdID:                  to.StringPtr("<threshold-id>"),
						},
						{
							DisplayOnChart: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Float64Ptr(100),
							},
							SendNotificationWhenExceeded: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							ThresholdID:                  to.StringPtr("<threshold-id>"),
						},
						{
							DisplayOnChart: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Float64Ptr(125),
							},
							SendNotificationWhenExceeded: armdevtestlabs.CostThresholdStatus("Disabled").ToPtr(),
							ThresholdID:                  to.StringPtr("<threshold-id>"),
						}},
					CycleEndDateTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T00:00:00.000Z"); return t }()),
					CycleStartDateTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00.000Z"); return t }()),
					CycleType:          armdevtestlabs.ReportingCycleType("CalendarMonth").ToPtr(),
					Status:             armdevtestlabs.TargetCostStatus("Enabled").ToPtr(),
					Target:             to.Int32Ptr(100),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CostsClientCreateOrUpdateResult)
}
```
