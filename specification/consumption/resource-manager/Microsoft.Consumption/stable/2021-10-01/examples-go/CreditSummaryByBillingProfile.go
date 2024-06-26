package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreditSummaryByBillingProfile.json
func ExampleCreditsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCreditsClient().Get(ctx, "1234:5678", "2468", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CreditSummary = armconsumption.CreditSummary{
	// 	Name: to.Ptr("balanceSummary1"),
	// 	Type: to.Ptr("Microsoft.Consumption/credits/balanceSummary"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/credits/balanceSummary1"),
	// 	Properties: &armconsumption.CreditSummaryProperties{
	// 		BalanceSummary: &armconsumption.CreditBalanceSummary{
	// 			CurrentBalance: &armconsumption.Amount{
	// 				Currency: to.Ptr("USD"),
	// 				Value: to.Ptr[float64](100),
	// 			},
	// 			EstimatedBalance: &armconsumption.Amount{
	// 				Currency: to.Ptr("USD"),
	// 				Value: to.Ptr[float64](600),
	// 			},
	// 		},
	// 		BillingCurrency: to.Ptr("USD"),
	// 		CreditCurrency: to.Ptr("USD"),
	// 		ExpiredCredit: &armconsumption.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float64](0),
	// 		},
	// 		PendingCreditAdjustments: &armconsumption.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float64](500),
	// 		},
	// 		PendingEligibleCharges: &armconsumption.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float64](0),
	// 		},
	// 		Reseller: &armconsumption.Reseller{
	// 			ResellerDescription: to.Ptr("Reseller information."),
	// 			ResellerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/billingProfiles/2468/providers/Microsoft.Consumption/reseller/reseller1"),
	// 		},
	// 	},
	// }
}
