Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.4.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportCreateOrUpdateByBillingAccount.json
func ExampleExportsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewExportsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<export-name>",
		armcostmanagement.Export{
			Properties: &armcostmanagement.ExportProperties{
				Format: to.Ptr(armcostmanagement.FormatTypeCSV),
				Definition: &armcostmanagement.ExportDefinition{
					Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
					DataSet: &armcostmanagement.ExportDataset{
						Configuration: &armcostmanagement.ExportDatasetConfiguration{
							Columns: []*string{
								to.Ptr("Date"),
								to.Ptr("MeterId"),
								to.Ptr("ResourceId"),
								to.Ptr("ResourceLocation"),
								to.Ptr("Quantity")},
						},
						Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
					},
					Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
				},
				DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
					Destination: &armcostmanagement.ExportDeliveryDestination{
						Container:      to.Ptr("<container>"),
						ResourceID:     to.Ptr("<resource-id>"),
						RootFolderPath: to.Ptr("<root-folder-path>"),
					},
				},
				Schedule: &armcostmanagement.ExportSchedule{
					Recurrence: to.Ptr(armcostmanagement.RecurrenceTypeWeekly),
					RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
						From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t }()),
						To:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-31T00:00:00Z"); return t }()),
					},
					Status: to.Ptr(armcostmanagement.StatusTypeActive),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
