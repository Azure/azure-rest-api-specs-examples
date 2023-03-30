package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiListByHub.json
func ExampleKpiClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKpiClient().NewListByHubPager("TestHubRG", "sdkTestHub", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KpiListResult = armcustomerinsights.KpiListResult{
		// 	Value: []*armcustomerinsights.KpiResourceFormat{
		// 		{
		// 			Name: to.Ptr("sdkTestHub/kpiTest45453647"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/kpi"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/kpi/kpiTest45453647"),
		// 			Properties: &armcustomerinsights.KpiDefinition{
		// 				Description: map[string]*string{
		// 					"en-us": to.Ptr("Kpi Description"),
		// 				},
		// 				Aliases: []*armcustomerinsights.KpiAlias{
		// 					{
		// 						AliasName: to.Ptr("alias"),
		// 						Expression: to.Ptr("Id+4"),
		// 				}},
		// 				CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
		// 				DisplayName: map[string]*string{
		// 					"en-us": to.Ptr("Kpi DisplayName"),
		// 				},
		// 				EntityType: to.Ptr(armcustomerinsights.EntityTypesProfile),
		// 				EntityTypeName: to.Ptr("testProfile2327128"),
		// 				Expression: to.Ptr("SavingAccountBalance"),
		// 				Function: to.Ptr(armcustomerinsights.KpiFunctionsSum),
		// 				GroupBy: []*string{
		// 					to.Ptr("SavingAccountBalance")},
		// 					KpiName: to.Ptr("kpiTest45453647"),
		// 					ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 					TenantID: to.Ptr("sdktesthub"),
		// 					ThresHolds: &armcustomerinsights.KpiThresholds{
		// 						IncreasingKpi: to.Ptr(true),
		// 						LowerLimit: to.Ptr[float64](5),
		// 						UpperLimit: to.Ptr[float64](50),
		// 					},
		// 					Unit: to.Ptr("unit"),
		// 				},
		// 		}},
		// 	}
	}
}
