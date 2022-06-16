package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
func ExampleCostsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewCostsClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"targetCost",
		armdevtestlabs.LabCost{
			Properties: &armdevtestlabs.LabCostProperties{
				CurrencyCode:  to.Ptr("USD"),
				EndDateTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T23:59:59Z"); return t }()),
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00Z"); return t }()),
				TargetCost: &armdevtestlabs.TargetCostProperties{
					CostThresholds: []*armdevtestlabs.CostThresholdProperties{
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](25),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("00000000-0000-0000-0000-000000000001"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusEnabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](50),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusEnabled),
							ThresholdID:                  to.Ptr("00000000-0000-0000-0000-000000000002"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](75),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("00000000-0000-0000-0000-000000000003"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](100),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("00000000-0000-0000-0000-000000000004"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](125),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("00000000-0000-0000-0000-000000000005"),
						}},
					CycleEndDateTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T00:00:00.000Z"); return t }()),
					CycleStartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00.000Z"); return t }()),
					CycleType:          to.Ptr(armdevtestlabs.ReportingCycleTypeCalendarMonth),
					Status:             to.Ptr(armdevtestlabs.TargetCostStatusEnabled),
					Target:             to.Ptr[int32](100),
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
