package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/CalculateExchange.json
func ExampleCalculateExchangeClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewCalculateExchangeClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPost(ctx,
		armreservations.CalculateExchangeRequest{
			Properties: &armreservations.CalculateExchangeRequestProperties{
				ReservationsToExchange: []*armreservations.ReservationToReturn{
					{
						Quantity:      to.Ptr[int32](1),
						ReservationID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6"),
					}},
				ReservationsToPurchase: []*armreservations.PurchaseRequest{
					{
						Location: to.Ptr("westus"),
						Properties: &armreservations.PurchaseRequestProperties{
							AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
							BillingPlan:      to.Ptr(armreservations.ReservationBillingPlanUpfront),
							BillingScopeID:   to.Ptr("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
							DisplayName:      to.Ptr("testDisplayName"),
							Quantity:         to.Ptr[int32](1),
							Renew:            to.Ptr(false),
							ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
								InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
							},
							ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
							Term:                 to.Ptr(armreservations.ReservationTermP1Y),
						},
						SKU: &armreservations.SKUName{
							Name: to.Ptr("Standard_B1ls"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
