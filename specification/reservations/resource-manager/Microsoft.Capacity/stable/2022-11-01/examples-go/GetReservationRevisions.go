package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationRevisions.json
func ExampleReservationClient_NewListRevisionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationClient().NewListRevisionsPager("276e7ae4-84d0-4da6-ab4b-d6b94f3557da", "6ef59113-3482-40da-8d79-787f823e34bc", nil)
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
		// page.ReservationList = armreservations.ReservationList{
		// 	Value: []*armreservations.ReservationResponse{
		// 		{
		// 			Name: to.Ptr("4"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations/revisions"),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc/revisions/4"),
		// 			Etag: to.Ptr[int32](4),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armreservations.Properties{
		// 				AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
		// 				BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 				DisplayName: to.Ptr("cabri_test"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T23:57:48.189Z"); return t}()),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T23:57:48.189Z"); return t}()),
		// 				InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 				LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T23:57:54.376Z"); return t}()),
		// 				MergeProperties: &armreservations.ReservationMergeProperties{
		// 					MergeSources: []*string{
		// 						to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 						to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 					},
		// 					ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 					Quantity: to.Ptr[int32](3),
		// 					ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 					SKUDescription: to.Ptr("D1 v2"),
		// 				},
		// 				SKU: &armreservations.SKUName{
		// 					Name: to.Ptr("Standard_DS1_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("3"),
		// 				Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations/revisions"),
		// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc/revisions/3"),
		// 				Etag: to.Ptr[int32](3),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armreservations.Properties{
		// 					AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 						DisplayName: to.Ptr("Azure subscription 1"),
		// 						SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 					},
		// 					AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 					BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 					DisplayName: to.Ptr("cabri_test"),
		// 					EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 					ExtendedStatusInfo: &armreservations.ExtendedStatusInfo{
		// 						Message: to.Ptr("An operation is in progress on your reservation. Please wait for operation to complete before taking further action"),
		// 						StatusCode: to.Ptr(armreservations.ReservationStatusCodePending),
		// 					},
		// 					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 					LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T23:57:47.488Z"); return t}()),
		// 					MergeProperties: &armreservations.ReservationMergeProperties{
		// 						MergeSources: []*string{
		// 							to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 							to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 						},
		// 						ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 						Quantity: to.Ptr[int32](3),
		// 						ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 						SKUDescription: to.Ptr("D1 v2"),
		// 					},
		// 					SKU: &armreservations.SKUName{
		// 						Name: to.Ptr("Standard_DS1_v2"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("2"),
		// 					Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations/revisions"),
		// 					ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc/revisions/2"),
		// 					Etag: to.Ptr[int32](2),
		// 					Location: to.Ptr("eastus"),
		// 					Properties: &armreservations.Properties{
		// 						AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 							DisplayName: to.Ptr("Azure subscription 1"),
		// 							SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 						},
		// 						AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 						BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 						DisplayName: to.Ptr("cabri_test"),
		// 						EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 						ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 						ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 						InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 						LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 						MergeProperties: &armreservations.ReservationMergeProperties{
		// 							MergeSources: []*string{
		// 								to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 								to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 							},
		// 							ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 							Quantity: to.Ptr[int32](3),
		// 							ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 							SKUDescription: to.Ptr("D1 v2"),
		// 							SwapProperties: &armreservations.ReservationSwapProperties{
		// 								SwapDestination: to.Ptr("/providers/microsoft.capacity/reservationOrders/afadf486-3432-4254-b297-4db8ef055f38/reservations/317efb41-b3ef-3706-8447-c2c045ab3ef5"),
		// 							},
		// 						},
		// 						SKU: &armreservations.SKUName{
		// 							Name: to.Ptr("Standard_DS1_v2"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("1"),
		// 						Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations/revisions"),
		// 						ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc/revisions/1"),
		// 						Etag: to.Ptr[int32](1),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armreservations.Properties{
		// 							AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 								DisplayName: to.Ptr("Azure subscription 1"),
		// 								SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 							},
		// 							AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 							DisplayName: to.Ptr("cabri_test"),
		// 							EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:27.331Z"); return t}()),
		// 							ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 							ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:27.331Z"); return t}()),
		// 							InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 							LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:27.331Z"); return t}()),
		// 							MergeProperties: &armreservations.ReservationMergeProperties{
		// 								MergeSources: []*string{
		// 									to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 									to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 								},
		// 								ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 								Quantity: to.Ptr[int32](3),
		// 								ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 								SKUDescription: to.Ptr("D1 v2"),
		// 							},
		// 							SKU: &armreservations.SKUName{
		// 								Name: to.Ptr("Standard_DS1_v2"),
		// 							},
		// 					}},
		// 				}
	}
}
