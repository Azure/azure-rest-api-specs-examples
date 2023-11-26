package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/EventsGetByBillingAccountWithFilters.json
func ExampleEventsClient_NewListByBillingAccountPager_eventsGetByBillingAccountWithFilters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventsClient().NewListByBillingAccountPager("1234:5678", &armconsumption.EventsClientListByBillingAccountOptions{Filter: to.Ptr("lotid eq 'G202001083926600XXXXX' AND lotsource eq 'consumptioncommitment'")})
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
		// page.Events = armconsumption.Events{
		// 	Value: []*armconsumption.EventSummary{
		// 		{
		// 			Name: to.Ptr("eventId1"),
		// 			Type: to.Ptr("Microsoft.Consumption/events"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:56789/providers/Microsoft.Consumption/events/eventId1"),
		// 			Properties: &armconsumption.EventProperties{
		// 				Description: to.Ptr("MACC Canceled"),
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/X3TD-KVTT-BG7-TGB"),
		// 				CanceledCredit: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](200),
		// 				},
		// 				ClosedBalance: &armconsumption.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float64](500),
		// 				},
		// 				EventType: to.Ptr(armconsumption.EventType("CanceledCredit")),
		// 				InvoiceNumber: to.Ptr("3304"),
		// 				LotID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234:5678/Microsoft.Consumption/lots/G202001083926600XXXXX"),
		// 				LotSource: to.Ptr("ConsumptionCommitment"),
		// 				TransactionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
