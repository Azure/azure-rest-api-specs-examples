package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/PriceSheetExpand.json
func ExamplePriceSheetClient_GetByBillingPeriod_priceSheetExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPriceSheetClient().GetByBillingPeriod(ctx, "201801", &armconsumption.PriceSheetClientGetByBillingPeriodOptions{Expand: to.Ptr("meterDetails"),
		Skiptoken: nil,
		Top:       nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PriceSheetResult = armconsumption.PriceSheetResult{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Consumption/pricesheets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/billingPeriods/201702/providers/Microsoft.Consumption/pricesheets/default"),
	// 	Properties: &armconsumption.PriceSheetModel{
	// 		Pricesheets: []*armconsumption.PriceSheetProperties{
	// 			{
	// 				BillingPeriodID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/billingPeriods/201702"),
	// 				CurrencyCode: to.Ptr("EUR"),
	// 				IncludedQuantity: to.Ptr[float64](100),
	// 				MeterDetails: &armconsumption.MeterDetails{
	// 					MeterCategory: to.Ptr("Networking"),
	// 					MeterLocation: to.Ptr("Zone 2"),
	// 					MeterName: to.Ptr("Data Transfer Out (GB)"),
	// 					PretaxStandardRate: to.Ptr[float64](0.138),
	// 					TotalIncludedQuantity: to.Ptr[float64](0),
	// 					Unit: to.Ptr("GB"),
	// 				},
	// 				MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				OfferID: to.Ptr("OfferId 1"),
	// 				PartNumber: to.Ptr("XX-11110"),
	// 				UnitOfMeasure: to.Ptr("100 Hours"),
	// 				UnitPrice: to.Ptr[float64](0.00328),
	// 		}},
	// 	},
	// }
}
