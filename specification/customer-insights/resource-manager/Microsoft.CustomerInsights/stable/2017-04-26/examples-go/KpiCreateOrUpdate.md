```go
package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiCreateOrUpdate.json
func ExampleKpiClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewKpiClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"kpiTest45453647",
		armcustomerinsights.KpiResourceFormat{
			Properties: &armcustomerinsights.KpiDefinition{
				Description: map[string]*string{
					"en-us": to.Ptr("Kpi Description"),
				},
				Aliases: []*armcustomerinsights.KpiAlias{
					{
						AliasName:  to.Ptr("alias"),
						Expression: to.Ptr("Id+4"),
					}},
				CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
				DisplayName: map[string]*string{
					"en-us": to.Ptr("Kpi DisplayName"),
				},
				EntityType:     to.Ptr(armcustomerinsights.EntityTypesProfile),
				EntityTypeName: to.Ptr("testProfile2327128"),
				Expression:     to.Ptr("SavingAccountBalance"),
				Function:       to.Ptr(armcustomerinsights.KpiFunctionsSum),
				GroupBy: []*string{
					to.Ptr("SavingAccountBalance")},
				ThresHolds: &armcustomerinsights.KpiThresholds{
					IncreasingKpi: to.Ptr(true),
					LowerLimit:    to.Ptr[float64](5),
					UpperLimit:    to.Ptr[float64](50),
				},
				Unit: to.Ptr("unit"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomerinsights%2Farmcustomerinsights%2Fv1.0.0/sdk/resourcemanager/customerinsights/armcustomerinsights/README.md) on how to add the SDK to your project and authenticate.
