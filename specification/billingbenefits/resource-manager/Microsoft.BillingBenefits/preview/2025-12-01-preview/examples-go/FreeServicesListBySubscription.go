package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/FreeServicesListBySubscription.json
func ExampleFreeServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFreeServicesClient("97ee05f2-07d5-494d-908c-081a197f4277").NewListBySubscriptionPager(nil)
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
		// page = armbillingbenefits.FreeServicesClientListBySubscriptionResponse{
		// 	FreeServicesList: armbillingbenefits.FreeServicesList{
		// 		Value: []*armbillingbenefits.FreeServices{
		// 			{
		// 				ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/providers/Microsoft.BillingBenefits/freeServices/freeservice_20251001"),
		// 				Name: to.Ptr("freeservice_20251001"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/freeServices"),
		// 				Properties: &armbillingbenefits.FreeServicesProperties{
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					Status: to.Ptr(armbillingbenefits.FreeServicesStatusActive),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 					CustomerResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/customers/40000000-0000-0000-0000-000000000000"),
		// 					SystemID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
		// 				},
		// 				Location: to.Ptr("global"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T01:01:01.1075056Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/providers/Microsoft.BillingBenefits/freeServices/freeservice_20250901"),
		// 				Name: to.Ptr("freeservice_20250901"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/freeServices"),
		// 				Properties: &armbillingbenefits.FreeServicesProperties{
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					Status: to.Ptr(armbillingbenefits.FreeServicesStatusActive),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 					CustomerResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/customers/40000000-0000-0000-0000-000000000000"),
		// 					SystemID: to.Ptr("b2c3d4e5-f6a7-8901-bcde-fa2345678901"),
		// 				},
		// 				Location: to.Ptr("global"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-01T01:01:01.1075056Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/providers/Microsoft.BillingBenefits/freeServices/freeservice_20250801"),
		// 				Name: to.Ptr("freeservice_20250801"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/freeServices"),
		// 				Properties: &armbillingbenefits.FreeServicesProperties{
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					Status: to.Ptr(armbillingbenefits.FreeServicesStatusActive),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 					CustomerResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/customers/40000000-0000-0000-0000-000000000000"),
		// 					SystemID: to.Ptr("c3d4e5f6-a7b8-9012-cdef-ab3456789012"),
		// 				},
		// 				Location: to.Ptr("global"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("development"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T01:01:01.1075056Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
