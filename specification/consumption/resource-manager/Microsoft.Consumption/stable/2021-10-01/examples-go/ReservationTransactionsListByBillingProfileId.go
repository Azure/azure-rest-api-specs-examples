package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationTransactionsListByBillingProfileId.json
func ExampleReservationTransactionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationTransactionsClient().NewListByBillingProfilePager("fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30", "Z76D-SGAF-BG7-TGB", &armconsumption.ReservationTransactionsClientListByBillingProfileOptions{Filter: to.Ptr("properties/eventDate+ge+2020-05-20+AND+properties/eventDate+le+2020-05-30")})
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
		// page.ModernReservationTransactionsListResult = armconsumption.ModernReservationTransactionsListResult{
		// 	Value: []*armconsumption.ModernReservationTransaction{
		// 		{
		// 			Name: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationTransactions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB/providers/Microsoft.Consumption/reservationTransactions"),
		// 			Properties: &armconsumption.ModernReservationTransactionProperties{
		// 				Description: to.Ptr("Reserved VM Instance, Standard_B1ls, US East, 3 Years"),
		// 				Amount: to.Ptr[float64](1.44),
		// 				ArmSKUName: to.Ptr("Standard_B1ls"),
		// 				BillingFrequency: to.Ptr("Recurring"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB"),
		// 				BillingProfileName: to.Ptr("IT Department*"),
		// 				Currency: to.Ptr("USD"),
		// 				EventDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-25T21:21:38.000Z"); return t}()),
		// 				EventType: to.Ptr("Purchase"),
		// 				Invoice: to.Ptr("T000456437"),
		// 				InvoiceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/billingProfiles/Z76D-SGAF-BG7-TGB/invoices/T000456437"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30/invoiceSections/QBTB-EYAK-PJA-TGB"),
		// 				InvoiceSectionName: to.Ptr("IT Department"),
		// 				PurchasingSubscriptionGUID: to.Ptr("d924ad15-4a3d-4047-971d-c8b1b300a97b"),
		// 				PurchasingSubscriptionName: to.Ptr("contoso"),
		// 				Quantity: to.Ptr[float64](1),
		// 				Region: to.Ptr("eastus"),
		// 				ReservationOrderID: to.Ptr("a838a8c3-a408-49e1-ac90-42cb95bff9b2"),
		// 				ReservationOrderName: to.Ptr("VM_RI_03-25-2020_14-18"),
		// 				Term: to.Ptr("P3Y"),
		// 			},
		// 	}},
		// }
	}
}
