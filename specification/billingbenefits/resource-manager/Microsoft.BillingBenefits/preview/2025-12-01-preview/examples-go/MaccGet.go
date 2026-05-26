package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/MaccGet.json
func ExampleMaccsClient_Get_maccGet() {
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
	// 		ID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.MaccModelProperties{
	// 			AllowContributors: to.Ptr(true),
	// 			AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeEnabled),
	// 			BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
	// 			Commitment: &armbillingbenefits.Commitment{
	// 				Amount: to.Ptr[float64](20000),
	// 				CurrencyCode: to.Ptr("USD"),
	// 				Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 			},
	// 			DisplayName: to.Ptr("macc 20230614"),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00Z"); return t}()),
	// 			EntityType: to.Ptr(armbillingbenefits.MaccEntityTypePrimary),
	// 			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.MaccStatusActive),
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
