package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasGet.json
func ExampleReservationOrderAliasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationOrderAliasClient().Get(ctx, "reservationOrderAlias123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationOrderAliasResponse = armbillingbenefits.ReservationOrderAliasResponse{
	// 	Name: to.Ptr("reservationOrderAlias123"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/reservationOrderAliases"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/reservationOrderAliases/reservationOrderAlias123"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armbillingbenefits.ReservationOrderAliasResponseProperties{
	// 		AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
	// 			ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		DisplayName: to.Ptr("ReservationOrder_2022-06-02"),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		Quantity: to.Ptr[int32](5),
	// 		Renew: to.Ptr(true),
	// 		ReservationOrderID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/30000000-0000-0000-0000-000000000000"),
	// 		ReservedResourceProperties: &armbillingbenefits.ReservationOrderAliasResponsePropertiesReservedResourceProperties{
	// 			InstanceFlexibility: to.Ptr(armbillingbenefits.InstanceFlexibilityOn),
	// 		},
	// 		ReservedResourceType: to.Ptr(armbillingbenefits.ReservedResourceTypeVirtualMachines),
	// 		Term: to.Ptr(armbillingbenefits.TermP1Y),
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Standard_M64s_v2"),
	// 	},
	// }
}
