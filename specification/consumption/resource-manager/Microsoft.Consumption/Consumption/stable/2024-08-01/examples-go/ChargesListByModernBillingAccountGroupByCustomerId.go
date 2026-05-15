package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ChargesListByModernBillingAccountGroupByCustomerId.json
func ExampleChargesClient_List_chargesListByBillingAccountGroupByCustomerIdModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChargesClient().List(ctx, "providers/Microsoft.Billing/billingAccounts/1234:56789", &armconsumption.ChargesClientListOptions{
		Apply:     to.Ptr("groupby((properties/customerId))"),
		EndDate:   to.Ptr("2019-09-30"),
		StartDate: to.Ptr("2019-09-01")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.ChargesClientListResponse{
	// 	ChargesListResult: &armconsumption.ChargesListResult{
	// 		Value: []armconsumption.ChargeSummaryClassification{
	// 			&armconsumption.ModernChargeSummary{
	// 				Name: to.Ptr("chargeSummaryId1"),
	// 				Type: to.Ptr("Microsoft.Consumption/charges"),
	// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/customers/67890/providers/Microsoft.Consumption/charges/chargeSummaryId1"),
	// 				Kind: to.Ptr(armconsumption.ChargeSummaryKindModern),
	// 				Properties: &armconsumption.ModernChargeSummaryProperties{
	// 					AzureCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789"),
	// 					BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/providers/Microsoft.Billing/billingPeriods/201909"),
	// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/billingProfiles/42425"),
	// 					ChargesBilledSeparately: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/customers/67890"),
	// 					InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/invoiceSections/67890"),
	// 					IsInvoiced: to.Ptr(false),
	// 					MarketplaceCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					UsageEnd: to.Ptr("2019-09-30"),
	// 					UsageStart: to.Ptr("2019-09-01"),
	// 				},
	// 			},
	// 			&armconsumption.ModernChargeSummary{
	// 				Name: to.Ptr("chargeSummaryId2"),
	// 				Type: to.Ptr("Microsoft.Consumption/charges"),
	// 				ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234:56789/customers/123456/providers/Microsoft.Consumption/charges/chargeSummaryId2"),
	// 				Kind: to.Ptr(armconsumption.ChargeSummaryKindModern),
	// 				Properties: &armconsumption.ModernChargeSummaryProperties{
	// 					AzureCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789"),
	// 					BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/providers/Microsoft.Billing/billingPeriods/201909"),
	// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/billingProfiles/42425"),
	// 					ChargesBilledSeparately: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/customers/123456"),
	// 					InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/invoiceSections/67890"),
	// 					IsInvoiced: to.Ptr(false),
	// 					MarketplaceCharges: &armconsumption.Amount{
	// 						Currency: to.Ptr("USD"),
	// 					},
	// 					UsageEnd: to.Ptr("2019-09-30"),
	// 					UsageStart: to.Ptr("2019-09-01"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
