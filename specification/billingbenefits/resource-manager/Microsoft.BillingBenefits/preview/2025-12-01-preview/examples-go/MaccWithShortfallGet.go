package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/MaccWithShortfallGet.json
func ExampleMaccsClient_Get_maccWithShortfallGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMaccsClient("10000000-0000-0000-0000-000000000000").Get(ctx, "resource_group_name_01", "macc_20230614", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.MaccsClientGetResponse{
	// 	Macc: armbillingbenefits.Macc{
	// 		Name: to.Ptr("macc_20230614"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/maccs"),
	// 		ID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.MaccModelProperties{
	// 			AllowContributors: to.Ptr(true),
	// 			AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeEnabled),
	// 			Commitment: &armbillingbenefits.Commitment{
	// 				Amount: to.Ptr[float64](20000),
	// 				CurrencyCode: to.Ptr("USD"),
	// 				Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 			},
	// 			DisplayName: to.Ptr("macc_20230614"),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-31T23:59:59Z"); return t}()),
	// 			EntityType: to.Ptr(armbillingbenefits.MaccEntityTypePrimary),
	// 			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
	// 			ProvisioningState: to.Ptr("Canceled"),
	// 			Shortfall: &armbillingbenefits.Shortfall{
	// 				BalanceVersion: to.Ptr[float32](2),
	// 				Charge: &armbillingbenefits.Commitment{
	// 					Amount: to.Ptr[float64](10000),
	// 					CurrencyCode: to.Ptr("USD"),
	// 					Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 				},
	// 				EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-30T23:59:59Z"); return t}()),
	// 				ProductCode: to.Ptr("000278da-0000-0160-330f-a0b98cdbbdc4"),
	// 				ResourceID: to.Ptr("/subscriptions/27507203-b094-420a-a716-b6fda046e1c2/resourceGroups/testcredit/providers/Microsoft.BillingBenefits/credits/taco0647052025"),
	// 				StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-01T00:00:00Z"); return t}()),
	// 				SystemID: to.Ptr("shortfall12345678"),
	// 			},
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.MaccStatusShortfallCharged),
	// 			SystemID: to.Ptr("13810867107109237"),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
