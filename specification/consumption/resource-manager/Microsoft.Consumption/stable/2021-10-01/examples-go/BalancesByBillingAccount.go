package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BalancesByBillingAccount.json
func ExampleBalancesClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBalancesClient().GetByBillingAccount(ctx, "123456", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Balance = armconsumption.Balance{
	// 	Name: to.Ptr("balanceId1"),
	// 	Type: to.Ptr("Microsoft.Consumption/balances"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/providers/Microsoft.Billing/billingPeriods/201702/providers/Microsoft.Consumption/balances/balanceId1"),
	// 	Properties: &armconsumption.BalanceProperties{
	// 		AdjustmentDetails: []*armconsumption.BalancePropertiesAdjustmentDetailsItem{
	// 			{
	// 				Name: to.Ptr("Promo Credit"),
	// 				Value: to.Ptr[float64](1.1),
	// 			},
	// 			{
	// 				Name: to.Ptr("SIE Credit"),
	// 				Value: to.Ptr[float64](1),
	// 		}},
	// 		Adjustments: to.Ptr[float64](0),
	// 		AzureMarketplaceServiceCharges: to.Ptr[float64](609.82),
	// 		BeginningBalance: to.Ptr[float64](3396469.19),
	// 		BillingFrequency: to.Ptr(armconsumption.BillingFrequencyMonth),
	// 		ChargesBilledSeparately: to.Ptr[float64](0),
	// 		Currency: to.Ptr("USD  "),
	// 		EndingBalance: to.Ptr[float64](2922371.02),
	// 		NewPurchases: to.Ptr[float64](0),
	// 		NewPurchasesDetails: []*armconsumption.BalancePropertiesNewPurchasesDetailsItem{
	// 			{
	// 				Name: to.Ptr("Promo Purchase"),
	// 				Value: to.Ptr[float64](1),
	// 		}},
	// 		PriceHidden: to.Ptr(false),
	// 		ServiceOverage: to.Ptr[float64](0),
	// 		TotalOverage: to.Ptr[float64](0),
	// 		TotalUsage: to.Ptr[float64](474098.17),
	// 		Utilized: to.Ptr[float64](474098.17),
	// 	},
	// }
}
