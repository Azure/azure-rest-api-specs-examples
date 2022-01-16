Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.2.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExportCreateOrUpdateByBillingAccount.json
func ExampleExportsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewExportsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<export-name>",
		armcostmanagement.Export{
			Properties: &armcostmanagement.ExportProperties{
				Format: armcostmanagement.FormatType("Csv").ToPtr(),
				Definition: &armcostmanagement.ExportDefinition{
					Type: armcostmanagement.ExportType("ActualCost").ToPtr(),
					DataSet: &armcostmanagement.ExportDataset{
						Configuration: &armcostmanagement.ExportDatasetConfiguration{
							Columns: []*string{
								to.StringPtr("Date"),
								to.StringPtr("MeterId"),
								to.StringPtr("ResourceId"),
								to.StringPtr("ResourceLocation"),
								to.StringPtr("Quantity")},
						},
						Granularity: armcostmanagement.GranularityType("Daily").ToPtr(),
					},
					Timeframe: armcostmanagement.TimeframeType("MonthToDate").ToPtr(),
				},
				DeliveryInfo: &armcostmanagement.ExportDeliveryInfo{
					Destination: &armcostmanagement.ExportDeliveryDestination{
						Container:      to.StringPtr("<container>"),
						ResourceID:     to.StringPtr("<resource-id>"),
						RootFolderPath: to.StringPtr("<root-folder-path>"),
					},
				},
				Schedule: &armcostmanagement.ExportSchedule{
					Recurrence: armcostmanagement.RecurrenceType("Weekly").ToPtr(),
					RecurrencePeriod: &armcostmanagement.ExportRecurrencePeriod{
						From: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T00:00:00Z"); return t }()),
						To:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-31T00:00:00Z"); return t }()),
					},
					Status: armcostmanagement.StatusType("Active").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExportsClientCreateOrUpdateResult)
}
```
