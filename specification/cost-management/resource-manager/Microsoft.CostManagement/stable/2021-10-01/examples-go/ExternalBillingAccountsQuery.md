Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.4.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountsQuery.json
func ExampleQueryClient_UsageByExternalCloudProviderType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewQueryClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UsageByExternalCloudProviderType(ctx,
		armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts,
		"<external-cloud-provider-id>",
		armcostmanagement.QueryDefinition{
			Type: to.Ptr(armcostmanagement.ExportTypeUsage),
			Dataset: &armcostmanagement.QueryDataset{
				Filter: &armcostmanagement.QueryFilter{
					And: []*armcostmanagement.QueryFilter{
						{
							Or: []*armcostmanagement.QueryFilter{
								{
									Dimensions: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("<name>"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("East US"),
											to.Ptr("West Europe")},
									},
								},
								{
									Tags: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("<name>"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("UAT"),
											to.Ptr("Prod")},
									},
								}},
						},
						{
							Dimensions: &armcostmanagement.QueryComparisonExpression{
								Name:     to.Ptr("<name>"),
								Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
								Values: []*string{
									to.Ptr("API")},
							},
						}},
				},
				Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
			},
			Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
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
