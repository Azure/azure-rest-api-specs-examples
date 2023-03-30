package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesForBillingPeriodByEnrollmentAccount.json
func ExampleChargesClient_List_changesForBillingPeriodByEnrollmentAccountLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChargesClient().List(ctx, "providers/Microsoft.Billing/BillingAccounts/1234/enrollmentAccounts/42425", &armconsumption.ChargesClientListOptions{StartDate: nil,
		EndDate: nil,
		Filter:  nil,
		Apply:   nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChargesListResult = armconsumption.ChargesListResult{
	// 	Value: []armconsumption.ChargeSummaryClassification{
	// 		&armconsumption.LegacyChargeSummary{
	// 			Name: to.Ptr("chargeSummaryId1"),
	// 			Type: to.Ptr("Microsoft.Consumption/charges"),
	// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234/enrollmentAccounts/42425/providers/Microsoft.Consumption/charges/chargeSummaryId1"),
	// 			Kind: to.Ptr(armconsumption.ChargeSummaryKindLegacy),
	// 			Properties: &armconsumption.LegacyChargeSummaryProperties{
	// 				AzureCharges: to.Ptr[float64](5000),
	// 				BillingPeriodID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/1234/providers/Microsoft.Billing/billingPeriods/201804"),
	// 				ChargesBilledSeparately: to.Ptr[float64](60.9),
	// 				Currency: to.Ptr("USD"),
	// 				MarketplaceCharges: to.Ptr[float64](100),
	// 				UsageEnd: to.Ptr("2018-04-30"),
	// 				UsageStart: to.Ptr("2018-04-01"),
	// 			},
	// 	}},
	// }
}
