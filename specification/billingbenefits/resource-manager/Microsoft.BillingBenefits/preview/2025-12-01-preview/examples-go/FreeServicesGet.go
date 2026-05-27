package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/FreeServicesGet.json
func ExampleFreeServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFreeServicesClient("97ee05f2-07d5-494d-908c-081a197f4277").Get(ctx, "resource_group_name_01", "freeservice_20251001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.FreeServicesClientGetResponse{
	// 	FreeServices: armbillingbenefits.FreeServices{
	// 		ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/freeServices/freeservice_20251001"),
	// 		Name: to.Ptr("freeservice_20251001"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/freeServices"),
	// 		Properties: &armbillingbenefits.FreeServicesProperties{
	// 			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
	// 			Status: to.Ptr(armbillingbenefits.FreeServicesStatusActive),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T00:00:00Z"); return t}()),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-10-01T00:00:00Z"); return t}()),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 			BillingProfileResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 			CustomerResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/customers/40000000-0000-0000-0000-000000000000"),
	// 			SystemID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 		},
	// 		Location: to.Ptr("global"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-01T01:01:01.1075056Z"); return t}()),
	// 		},
	// 	},
	// }
}
