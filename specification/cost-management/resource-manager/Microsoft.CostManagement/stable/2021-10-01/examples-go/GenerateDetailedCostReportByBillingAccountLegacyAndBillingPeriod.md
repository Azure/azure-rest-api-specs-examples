Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.1.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.json
func ExampleGenerateDetailedCostReportClient_BeginCreateOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewGenerateDetailedCostReportClient(cred, nil)
	poller, err := client.BeginCreateOperation(ctx,
		"<scope>",
		armcostmanagement.GenerateDetailedCostReportDefinition{
			BillingPeriod: to.StringPtr("<billing-period>"),
			Metric:        armcostmanagement.GenerateDetailedCostReportMetricTypeActualCost.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GenerateDetailedCostReportOperationResult.ID: %s\n", *res.ID)
}
```
