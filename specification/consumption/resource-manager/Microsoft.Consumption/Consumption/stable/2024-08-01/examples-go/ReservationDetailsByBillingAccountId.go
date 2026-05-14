package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/ReservationDetailsByBillingAccountId.json
func ExampleReservationsDetailsClient_NewListPager_reservationDetailsByBillingAccountId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsDetailsClient().NewListPager("providers/Microsoft.Billing/billingAccounts/12345", &armconsumption.ReservationsDetailsClientListOptions{
		Filter: to.Ptr("properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05")})
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
		// page = armconsumption.ReservationsDetailsClientListResponse{
		// 	ReservationDetailsListResult: armconsumption.ReservationDetailsListResult{
		// 		Value: []*armconsumption.ReservationDetail{
		// 			{
		// 				Name: to.Ptr("reservationDetails_Id1"),
		// 				Type: to.Ptr("Microsoft.Consumption/reservationDetails"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/12345/providers/Microsoft.Consumption/reservationDetails/reservationDetails_Id1"),
		// 				Properties: &armconsumption.ReservationDetailProperties{
		// 					InstanceFlexibilityGroup: to.Ptr("DSv3 Series"),
		// 					InstanceFlexibilityRatio: to.Ptr("1"),
		// 					InstanceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sqlh1/providers/microsoft.compute/virtualmachines/sqlh1"),
		// 					ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					SKUName: to.Ptr("Standard_D2s_v3"),
		// 					UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-30T00:00:00-08:00"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
