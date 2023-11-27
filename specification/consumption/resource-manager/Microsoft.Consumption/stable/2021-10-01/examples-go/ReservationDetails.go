package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetails.json
func ExampleReservationsDetailsClient_NewListByReservationOrderPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsDetailsClient().NewListByReservationOrderPager("00000000-0000-0000-0000-000000000000", "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05", nil)
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
		// page.ReservationDetailsListResult = armconsumption.ReservationDetailsListResult{
		// 	Value: []*armconsumption.ReservationDetail{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000_20171129"),
		// 			Type: to.Ptr("Microsoft.Consumption/reservationDetails"),
		// 			ID: to.Ptr("providers/Microsoft.Capacity/reservationOrders/00000000-0000-0000-0000-000000000000/reservations/00000000-0000-0000-0000-000000000000/providers/Microsoft.Consumption/reservationDetails/20171129"),
		// 			Tags: map[string]*string{
		// 				"dev": to.Ptr("tools"),
		// 				"env": to.Ptr("newcrp"),
		// 			},
		// 			Properties: &armconsumption.ReservationDetailProperties{
		// 				InstanceFlexibilityGroup: to.Ptr("DSv2 Series"),
		// 				InstanceFlexibilityRatio: to.Ptr("0.25"),
		// 				InstanceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/wvn-sql/providers/microsoft.compute/virtualmachines/abc-sql2014sp33"),
		// 				ReservationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservationOrderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ReservedHours: to.Ptr[float64](24),
		// 				SKUName: to.Ptr("Standard_D2_v2"),
		// 				TotalReservedQuantity: to.Ptr[float64](1),
		// 				UsageDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-29T00:00:00.000Z"); return t}()),
		// 				UsedHours: to.Ptr[float64](24),
		// 			},
		// 	}},
		// }
	}
}
