package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationsFromOrder.json
func ExampleReservationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationClient().NewListPager("276e7ae4-84d0-4da6-ab4b-d6b94f3557da", nil)
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
		// 			Name: to.Ptr("bcae77cd-3119-4766-919f-b50d36c75c7a"),
		// 			Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 			ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/bcae77cd-3119-4766-919f-b50d36c75c7a"),
		// 			Etag: to.Ptr[int32](10),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armreservations.Properties{
		// 				AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("Azure subscription 1"),
		// 					SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T01:00:30.925Z"); return t}()),
		// 				BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 				BillingScopeID: to.Ptr("/subscriptions/19376483-64b8-49e4-a931-d5248828720a"),
		// 				DisplayName: to.Ptr("cabri_test"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T01:00:30.925Z"); return t}()),
		// 				ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T01:00:30.925Z"); return t}()),
		// 				ExtendedStatusInfo: &armreservations.ExtendedStatusInfo{
		// 					Message: to.Ptr("This reservation was split and is no longer active."),
		// 					StatusCode: to.Ptr(armreservations.ReservationStatusCodeSplit),
		// 				},
		// 				InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 				LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T01:00:36.319Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 				Quantity: to.Ptr[int32](3),
		// 				Renew: to.Ptr(false),
		// 				ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 				SKUDescription: to.Ptr("D1 v2"),
		// 				SplitProperties: &armreservations.ReservationSplitProperties{
		// 					SplitDestinations: []*string{
		// 						to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 						to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 					},
		// 				},
		// 				SKU: &armreservations.SKUName{
		// 					Name: to.Ptr("Standard_DS1_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("00238563-7312-4c20-a134-8c030bf938a7"),
		// 				Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 				Etag: to.Ptr[int32](5),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armreservations.Properties{
		// 					AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 						DisplayName: to.Ptr("Azure subscription 1"),
		// 						SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 					},
		// 					AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 					BillingScopeID: to.Ptr("/subscriptions/19376483-64b8-49e4-a931-d5248828720a"),
		// 					DisplayName: to.Ptr("cabri_test"),
		// 					EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 					ExtendedStatusInfo: &armreservations.ExtendedStatusInfo{
		// 						Message: to.Ptr("This reservation was merged and is no longer active."),
		// 						StatusCode: to.Ptr(armreservations.ReservationStatusCodeMerged),
		// 					},
		// 					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 					LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					MergeProperties: &armreservations.ReservationMergeProperties{
		// 						MergeDestination: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc"),
		// 					},
		// 					ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 					Quantity: to.Ptr[int32](1),
		// 					Renew: to.Ptr(false),
		// 					ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 					SKUDescription: to.Ptr("D1 v2"),
		// 					SplitProperties: &armreservations.ReservationSplitProperties{
		// 						SplitSource: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/bcae77cd-3119-4766-919f-b50d36c75c7a"),
		// 					},
		// 				},
		// 				SKU: &armreservations.SKUName{
		// 					Name: to.Ptr("Standard_DS1_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de"),
		// 				Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de"),
		// 				Etag: to.Ptr[int32](4),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armreservations.Properties{
		// 					AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 						DisplayName: to.Ptr("Azure subscription 1"),
		// 						SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 					},
		// 					AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 					BillingScopeID: to.Ptr("/subscriptions/19376483-64b8-49e4-a931-d5248828720a"),
		// 					DisplayName: to.Ptr("cabri_test"),
		// 					EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 					ExtendedStatusInfo: &armreservations.ExtendedStatusInfo{
		// 						Message: to.Ptr("This reservation was merged and is no longer active."),
		// 						StatusCode: to.Ptr(armreservations.ReservationStatusCodeMerged),
		// 					},
		// 					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 					LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					MergeProperties: &armreservations.ReservationMergeProperties{
		// 						MergeDestination: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc"),
		// 					},
		// 					ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 					Quantity: to.Ptr[int32](2),
		// 					Renew: to.Ptr(false),
		// 					ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 					SKUDescription: to.Ptr("D1 v2"),
		// 					SplitProperties: &armreservations.ReservationSplitProperties{
		// 						SplitSource: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/bcae77cd-3119-4766-919f-b50d36c75c7a"),
		// 					},
		// 				},
		// 				SKU: &armreservations.SKUName{
		// 					Name: to.Ptr("Standard_DS1_v2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("6ef59113-3482-40da-8d79-787f823e34bc"),
		// 				Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 				ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/6ef59113-3482-40da-8d79-787f823e34bc"),
		// 				Etag: to.Ptr[int32](2),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armreservations.Properties{
		// 					AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 						DisplayName: to.Ptr("Azure subscription 1"),
		// 						SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 					},
		// 					AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 					BillingScopeID: to.Ptr("/subscriptions/19376483-64b8-49e4-a931-d5248828720a"),
		// 					DisplayName: to.Ptr("cabri_test"),
		// 					EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 					InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 					LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 					MergeProperties: &armreservations.ReservationMergeProperties{
		// 						MergeSources: []*string{
		// 							to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/00238563-7312-4c20-a134-8c030bf938a7"),
		// 							to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/e0e4b4f5-77ea-4984-9ee4-6bf9850ee6de")},
		// 						},
		// 						ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 						Quantity: to.Ptr[int32](3),
		// 						Renew: to.Ptr(false),
		// 						ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 						SKUDescription: to.Ptr("D1 v2"),
		// 					},
		// 					SKU: &armreservations.SKUName{
		// 						Name: to.Ptr("Standard_DS1_v2"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("5se63214-3382-40da-ed79-787f823e34bc"),
		// 					Type: to.Ptr("Microsoft.Capacity/reservationOrders/reservations"),
		// 					ID: to.Ptr("/providers/microsoft.capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/5se63214-3382-40da-ed79-787f823e34bc"),
		// 					Etag: to.Ptr[int32](2),
		// 					Location: to.Ptr("eastus"),
		// 					Properties: &armreservations.Properties{
		// 						AppliedScopeProperties: &armreservations.AppliedScopeProperties{
		// 							DisplayName: to.Ptr("Azure subscription 1"),
		// 							SubscriptionID: to.Ptr("/subscriptions/98df3792-7962-4f18-8be2-d5576f122de3"),
		// 						},
		// 						AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
		// 						BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:40:32.763Z"); return t}()),
		// 						BillingPlan: to.Ptr(armreservations.ReservationBillingPlanMonthly),
		// 						BillingScopeID: to.Ptr("/subscriptions/19376483-64b8-49e4-a931-d5248828720a"),
		// 						DisplayName: to.Ptr("cabri_test"),
		// 						EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 						ExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2018-09-22"); return t}()),
		// 						ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-22T22:46:32.763Z"); return t}()),
		// 						InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
		// 						LastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-22T22:46:32.763Z"); return t}()),
		// 						ProvisioningState: to.Ptr(armreservations.ProvisioningStateSucceeded),
		// 						Quantity: to.Ptr[int32](3),
		// 						Renew: to.Ptr(false),
		// 						ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
		// 						SKUDescription: to.Ptr("D1 v2"),
		// 						SwapProperties: &armreservations.ReservationSwapProperties{
		// 							SwapDestination: to.Ptr("/providers/microsoft.capacity/reservationOrders/63027ae3-389a-4716-ad77-04f42ce93c06/reservations/4jf50113-3382-40da-ed79-787f823e34bc"),
		// 						},
		// 					},
		// 					SKU: &armreservations.SKUName{
		// 						Name: to.Ptr("Standard_D1_v2"),
		// 					},
		// 			}},
		// 		}
	}
}
