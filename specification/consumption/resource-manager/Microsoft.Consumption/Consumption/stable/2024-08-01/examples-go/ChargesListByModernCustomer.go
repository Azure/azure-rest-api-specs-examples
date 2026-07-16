package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ChargesListByModernCustomer.json
func ExampleChargesClient_List_chargesListByCustomerModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChargesClient().List(ctx, "providers/Microsoft.Billing/BillingAccounts/1234:56789/customers/67890", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.ChargesClientListResponse{
	// 	ChargesListResult: armconsumption.ChargesListResult{
	// 		Value: []armconsumption.ChargeSummaryClassification{
	// 			&armconsumption.ModernChargeSummary{
	// 				Name: to.Ptr("chargeSummaryId1"),
	// 				Type: to.Ptr("Microsoft.Consumption/charges"),
	// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/customers/67890/providers/Microsoft.Consumption/charges/chargeSummaryId1"),
	// 				Kind: to.Ptr(armconsumption.ChargeSummaryKindModern),
	// 				Properties: &armconsumption.ModernChargeSummaryProperties{
	// 					AzureCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 						Value: to.Ptr[float64](0),
	// 					},
	// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789"),
	// 					BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/providers/Microsoft.Billing/billingPeriods/201910"),
	// 					ChargesBilledSeparately: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 						Value: to.Ptr[float64](265.09),
	// 					},
	// 					IsInvoiced: to.Ptr(false),
	// 					MarketplaceCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 						Value: to.Ptr[float64](0),
	// 					},
	// 					UsageEnd: to.Ptr("2023-05-31"),
	// 					UsageStart: to.Ptr("2023-03-01"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
