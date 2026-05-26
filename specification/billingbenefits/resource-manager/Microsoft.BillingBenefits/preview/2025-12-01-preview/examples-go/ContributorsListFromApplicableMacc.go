package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ContributorsListFromApplicableMacc.json
func ExampleContributorsClient_NewListFromApplicableMaccPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContributorsClient("<subscriptionID>").NewListFromApplicableMaccPager("20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31", "13810867107109237", nil)
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
		// page = armbillingbenefits.ContributorsClientListFromApplicableMaccResponse{
		// 	ContributorList: armbillingbenefits.ContributorList{
		// 		Value: []*armbillingbenefits.Contributor{
		// 			{
		// 				Name: to.Ptr("contributor_20230614"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/applicableContributors"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}/providers/Microsoft.BillingBenefits/applicableMaccs/13810867107109237/providers/Microsoft.BillingBenefits/applicableContributors/contributor_20230614"),
		// 				Properties: &armbillingbenefits.MaccModelProperties{
		// 					BillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
		// 					Commitment: &armbillingbenefits.Commitment{
		// 						Amount: to.Ptr[float64](20000),
		// 						CurrencyCode: to.Ptr("USD"),
		// 						Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
		// 					},
		// 					DisplayName: to.Ptr("Contributor Macc 20230614"),
		// 					EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00Z"); return t}()),
		// 					EntityType: to.Ptr(armbillingbenefits.MaccEntityTypeContributor),
		// 					PrimaryBillingAccountResourceID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}"),
		// 					PrimaryResourceID: to.Ptr("/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
		// 					ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
		// 					Status: to.Ptr(armbillingbenefits.MaccStatusActive),
		// 					SystemID: to.Ptr("13810867107109237"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
