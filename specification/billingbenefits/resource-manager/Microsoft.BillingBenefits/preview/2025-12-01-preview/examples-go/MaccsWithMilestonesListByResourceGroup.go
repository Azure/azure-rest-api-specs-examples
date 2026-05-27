package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/MaccsWithMilestonesListByResourceGroup.json
func ExampleMaccsClient_NewListByResourceGroupPager_maccsWithMilestonesListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMaccsClient("10000000-0000-0000-0000-000000000000").NewListByResourceGroupPager("resource_group_name_01", nil)
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
		// page = armbillingbenefits.MaccsClientListByResourceGroupResponse{
		// 	MaccList: armbillingbenefits.MaccList{
		// 		Value: []*armbillingbenefits.Macc{
		// 			{
		// 				Name: to.Ptr("macc_20230614"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/maccs"),
		// 				ID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armbillingbenefits.MaccModelProperties{
		// 					AllowContributors: to.Ptr(true),
		// 					AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeEnabled),
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
		// 					Commitment: &armbillingbenefits.Commitment{
		// 						Amount: to.Ptr[float64](20000),
		// 						CurrencyCode: to.Ptr("USD"),
		// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 					},
		// 					DisplayName: to.Ptr("macc 20230614"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2028-05-01T23:59:59Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.MaccEntityTypePrimary),
		// 					Milestones: []*armbillingbenefits.MaccMilestone{
		// 						{
		// 							AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeEnabled),
		// 							Commitment: &armbillingbenefits.Price{
		// 								Amount: to.Ptr[float64](10000),
		// 								CurrencyCode: to.Ptr("USD"),
		// 							},
		// 							EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-31T23:59:59Z"); return t}()),
		// 							MilestoneID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 							Status: to.Ptr(armbillingbenefits.MaccMilestoneStatusActive),
		// 						},
		// 						{
		// 							AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeDisabled),
		// 							Commitment: &armbillingbenefits.Price{
		// 								Amount: to.Ptr[float64](15000),
		// 								CurrencyCode: to.Ptr("USD"),
		// 							},
		// 							EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-05-31T23:59:59Z"); return t}()),
		// 							MilestoneID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 							Status: to.Ptr(armbillingbenefits.MaccMilestoneStatusScheduled),
		// 						},
		// 					},
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.MaccStatusActive),
		// 					SystemID: to.Ptr("13810867107109237"),
		// 				},
		// 				SystemData: &armbillingbenefits.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
