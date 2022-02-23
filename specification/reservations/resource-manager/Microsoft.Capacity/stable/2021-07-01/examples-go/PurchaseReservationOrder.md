Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Freservations%2Farmreservations%2Fv0.2.1/sdk/resourcemanager/reservations/armreservations/README.md) on how to add the SDK to your project and authenticate.

```go
package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/PurchaseReservationOrder.json
func ExampleReservationOrderClient_BeginPurchase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	poller, err := client.BeginPurchase(ctx,
		"<reservation-order-id>",
		armreservations.PurchaseRequest{
			Location: to.StringPtr("<location>"),
			Properties: &armreservations.PurchaseRequestProperties{
				AppliedScopeType: armreservations.AppliedScopeType("Shared").ToPtr(),
				BillingPlan:      armreservations.ReservationBillingPlan("Monthly").ToPtr(),
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
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReservationOrderClientPurchaseResult)
}
```
