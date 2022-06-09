```go
package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/PurchaseReservationOrder.json
func ExampleReservationOrderClient_BeginPurchase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationOrderClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPurchase(ctx,
		"a075419f-44cc-497f-b68a-14ee811d48b9",
		armreservations.PurchaseRequest{
			Location: to.Ptr("westus"),
			Properties: &armreservations.PurchaseRequestProperties{
				AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
				BillingPlan:      to.Ptr(armreservations.ReservationBillingPlanMonthly),
				BillingScopeID:   to.Ptr("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
				DisplayName:      to.Ptr("TestReservationOrder"),
				Quantity:         to.Ptr[int32](1),
				Renew:            to.Ptr(false),
				ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
				},
				ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
				Term:                 to.Ptr(armreservations.ReservationTermP1Y),
			},
			SKU: &armreservations.SKUName{
				Name: to.Ptr("standard_D1"),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Freservations%2Farmreservations%2Fv1.0.0/sdk/resourcemanager/reservations/armreservations/README.md) on how to add the SDK to your project and authenticate.
