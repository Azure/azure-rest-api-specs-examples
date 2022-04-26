Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Freservations%2Farmreservations%2Fv0.4.0/sdk/resourcemanager/reservations/armreservations/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/PurchaseReservationOrder.json
func ExampleReservationOrderClient_BeginPurchase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationOrderClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPurchase(ctx,
		"<reservation-order-id>",
		armreservations.PurchaseRequest{
			Location: to.Ptr("<location>"),
			Properties: &armreservations.PurchaseRequestProperties{
				AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
				BillingPlan:      to.Ptr(armreservations.ReservationBillingPlanMonthly),
				BillingScopeID:   to.Ptr("<billing-scope-id>"),
				DisplayName:      to.Ptr("<display-name>"),
				Quantity:         to.Ptr[int32](1),
				Renew:            to.Ptr(false),
				ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
				},
				ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
				Term:                 to.Ptr(armreservations.ReservationTermP1Y),
			},
			SKU: &armreservations.SKUName{
				Name: to.Ptr("<name>"),
			},
		},
		&armreservations.ReservationOrderClientBeginPurchaseOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
