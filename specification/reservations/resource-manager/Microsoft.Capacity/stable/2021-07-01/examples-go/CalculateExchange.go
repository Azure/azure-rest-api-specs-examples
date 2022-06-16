package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/CalculateExchange.json
func ExampleCalculateExchangeClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewCalculateExchangeClient(cred, nil)
	poller, err := client.BeginPost(ctx,
		armreservations.CalculateExchangeRequest{
			Properties: &armreservations.CalculateExchangeRequestProperties{
				ReservationsToExchange: []*armreservations.ReservationToReturn{
					{
						Quantity:      to.Int32Ptr(1),
						ReservationID: to.StringPtr("<reservation-id>"),
					}},
				ReservationsToPurchase: []*armreservations.PurchaseRequest{
					{
						Location: to.StringPtr("<location>"),
						Properties: &armreservations.PurchaseRequestProperties{
							AppliedScopeType: armreservations.AppliedScopeType("Shared").ToPtr(),
							BillingPlan:      armreservations.ReservationBillingPlan("Upfront").ToPtr(),
							BillingScopeID:   to.StringPtr("<billing-scope-id>"),
							DisplayName:      to.StringPtr("<display-name>"),
							Quantity:         to.Int32Ptr(1),
							Renew:            to.BoolPtr(false),
							ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
								InstanceFlexibility: armreservations.InstanceFlexibility("On").ToPtr(),
							},
							ReservedResourceType: armreservations.ReservedResourceType("VirtualMachines").ToPtr(),
							Term:                 armreservations.ReservationTerm("P1Y").ToPtr(),
						},
						SKU: &armreservations.SKUName{
							Name: to.StringPtr("<name>"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CalculateExchangeClientPostResult)
}
