package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/PriceSheet.json
func ExamplePriceSheetClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPriceSheetClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.PriceSheetClientGetResponse{
	// 	PriceSheetResult: &armconsumption.PriceSheetResult{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Consumption/pricesheets"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/billingPeriods/201702/providers/Microsoft.Consumption/pricesheets/default"),
	// 		Properties: &armconsumption.PriceSheetModel{
	// 			NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.consumption/pricesheets/default?api-version=2018-01-31&$skiptoken=AQAAAA%3D%3D"),
	// 			Pricesheets: []*armconsumption.PriceSheetProperties{
	// 				{
	// 					BillingPeriodID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Billing/billingPeriods/201702"),
	// 					CurrencyCode: to.Ptr("EUR"),
	// 					MeterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					OfferID: to.Ptr("OfferId 1"),
	// 					PartNumber: to.Ptr("XX-11110"),
	// 					SavingsPlan: &armconsumption.SavingsPlan{
	// 						Term: to.Ptr("P3Y"),
	// 					},
	// 					UnitOfMeasure: to.Ptr("100 Hours"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
