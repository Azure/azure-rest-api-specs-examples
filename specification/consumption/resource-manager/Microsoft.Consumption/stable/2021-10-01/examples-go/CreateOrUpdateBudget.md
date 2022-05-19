Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconsumption%2Farmconsumption%2Fv1.0.0/sdk/resourcemanager/consumption/armconsumption/README.md) on how to add the SDK to your project and authenticate.

```go
package armconsumption_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreateOrUpdateBudget.json
func ExampleBudgetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconsumption.NewBudgetsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"subscriptions/00000000-0000-0000-0000-000000000000",
		"TestBudget",
		armconsumption.Budget{
			ETag: to.Ptr("\"1d34d016a593709\""),
			Properties: &armconsumption.BudgetProperties{
				Amount:   to.Ptr[float64](100.65),
				Category: to.Ptr(armconsumption.CategoryTypeCost),
				Filter: &armconsumption.BudgetFilter{
					And: []*armconsumption.BudgetFilterProperties{
						{
							Dimensions: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("ResourceId"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("category"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("Dev"),
									to.Ptr("Prod")},
							},
						},
						{
							Tags: &armconsumption.BudgetComparisonExpression{
								Name:     to.Ptr("department"),
								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
								Values: []*string{
									to.Ptr("engineering"),
									to.Ptr("sales")},
							},
						}},
				},
				Notifications: map[string]*armconsumption.Notification{
					"Actual_GreaterThan_80_Percent": {
						ContactEmails: []*string{
							to.Ptr("johndoe@contoso.com"),
							to.Ptr("janesmith@contoso.com")},
						ContactGroups: []*string{
							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup")},
						ContactRoles: []*string{
							to.Ptr("Contributor"),
							to.Ptr("Reader")},
						Enabled:       to.Ptr(true),
						Locale:        to.Ptr(armconsumption.CultureCodeEnUs),
						Operator:      to.Ptr(armconsumption.OperatorTypeGreaterThan),
						Threshold:     to.Ptr[float64](80),
						ThresholdType: to.Ptr(armconsumption.ThresholdTypeActual),
					},
				},
				TimeGrain: to.Ptr(armconsumption.TimeGrainTypeMonthly),
				TimePeriod: &armconsumption.BudgetTimePeriod{
					EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T00:00:00Z"); return t }()),
					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00Z"); return t }()),
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
