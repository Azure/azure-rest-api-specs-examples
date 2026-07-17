package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ChargesListForDepartmentFilterByStartEndDate.json
func ExampleChargesClient_List_chargesListByDepartmentLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChargesClient().List(ctx, "providers/Microsoft.Billing/BillingAccounts/1234/departments/42425", &armconsumption.ChargesClientListOptions{
		Filter: to.Ptr("usageStart eq '2018-04-01' AND usageEnd eq '2018-05-30'")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.ChargesClientListResponse{
	// 	ChargesListResult: armconsumption.ChargesListResult{
	// 		Value: []armconsumption.ChargeSummaryClassification{
	// 			&armconsumption.LegacyChargeSummary{
	// 				Name: to.Ptr("chargeSummaryId1"),
	// 				Type: to.Ptr("Microsoft.Consumption/charges"),
	// 				ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234/departments/42425/providers/Microsoft.Consumption/charges/chargeSummaryId1"),
	// 				Kind: to.Ptr(armconsumption.ChargeSummaryKindLegacy),
	// 				Properties: &armconsumption.LegacyChargeSummaryProperties{
	// 					AzureCharges: to.Ptr[float64](5000),
	// 					MarketplaceCharges: to.Ptr[float64](100),
	// 					BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234/providers/Microsoft.Billing/billingPeriods/201804"),
	// 					ChargesBilledSeparately: to.Ptr[float64](60.9),
	// 					Currency: to.Ptr("USD"),
	// 					UsageEnd: to.Ptr("2018-04-30"),
	// 					UsageStart: to.Ptr("2018-04-01"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
