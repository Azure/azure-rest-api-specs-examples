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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.json
func ExampleGenerateDetailedCostReportClient_BeginCreateOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewGenerateDetailedCostReportClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOperation(ctx,
		"<scope>",
		armcostmanagement.GenerateDetailedCostReportDefinition{
			BillingPeriod: to.Ptr("<billing-period>"),
			Metric:        to.Ptr(armcostmanagement.GenerateDetailedCostReportMetricTypeActualCost),
		},
		&armcostmanagement.GenerateDetailedCostReportClientBeginCreateOperationOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
