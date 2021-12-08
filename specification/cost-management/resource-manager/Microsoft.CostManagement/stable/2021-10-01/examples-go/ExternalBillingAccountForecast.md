Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.1.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountForecast.json
func ExampleForecastClient_ExternalCloudProviderUsage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewForecastClient(cred, nil)
	res, err := client.ExternalCloudProviderUsage(ctx,
		armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts,
		"<external-cloud-provider-id>",
		armcostmanagement.ForecastDefinition{
			Type: armcostmanagement.ForecastTypeUsage.ToPtr(),
			Dataset: &armcostmanagement.ForecastDataset{
				Filter: &armcostmanagement.QueryFilter{
					And: []*armcostmanagement.QueryFilter{
						{
							Or: []*armcostmanagement.QueryFilter{
								{
									Dimension: &armcostmanagement.QueryComparisonExpression{
										Name:     to.StringPtr("<name>"),
										Operator: armcostmanagement.QueryOperatorTypeIn.ToPtr(),
										Values: []*string{
											to.StringPtr("East US"),
											to.StringPtr("West Europe")},
									},
								},
								{
									Tag: &armcostmanagement.QueryComparisonExpression{
										Name:     to.StringPtr("<name>"),
										Operator: armcostmanagement.QueryOperatorTypeIn.ToPtr(),
										Values: []*string{
											to.StringPtr("UAT"),
											to.StringPtr("Prod")},
									},
								}},
						},
						{
							Dimension: &armcostmanagement.QueryComparisonExpression{
								Name:     to.StringPtr("<name>"),
								Operator: armcostmanagement.QueryOperatorTypeIn.ToPtr(),
								Values: []*string{
									to.StringPtr("API")},
							},
						}},
				},
				Granularity: armcostmanagement.GranularityTypeDaily.ToPtr(),
			},
			Timeframe: armcostmanagement.ForecastTimeframeTypeMonthToDate.ToPtr(),
		},
		&armcostmanagement.ForecastExternalCloudProviderUsageOptions{Filter: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("QueryResult.ID: %s\n", *res.ID)
}
```
